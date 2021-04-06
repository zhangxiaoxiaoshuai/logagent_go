package tailfile

import (
	"code.oldboy.com/studygolang/logagent/kafka"
	"context"
	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

// tail相关

type tailTask struct {
	path   string
	topic  string
	tObj   *tail.Tail
	ctx    context.Context
	cancel context.CancelFunc
}

func newTailTask(path, topic string) *tailTask {
	ctx, cancel := context.WithCancel(context.Background())
	tt := &tailTask{
		path:   path,
		topic:  topic,
		ctx:    ctx,
		cancel: cancel,
	}
	return tt
}

func (t *tailTask) Init() (err error) {
	cfg := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	t.tObj, err = tail.TailFile(t.path, cfg)
	return
}

func (t *tailTask) run() {
	// 读取日志,发往kafka
	logrus.Infof("collect for path:%s is running...", t.path)
	for {
		select {
		case <-t.ctx.Done(): // 只要调用t.cancel() 就会收到信号
			logrus.Infof("path:%s is stopping...", t.path)
			return
			// 循环读数据
		case line, ok := <-t.tObj.Lines: // chan tail.Line
			if !ok {
				logrus.Warn("tail file close reopen, path:%s\n", t.path)
				time.Sleep(time.Second) // 读取出错等一秒
				continue
			}
			// 如果是空行就略过
			//fmt.Printf("%#v\n", line.Text)
			if len(strings.Trim(line.Text, "\r")) == 0 {
				logrus.Info("出现空行拉,直接跳过...")
				continue
			}
			// 利用通道将同步的代码改为异步的
			// 把读出来的一行日志包装秤kafka里面的msg类型
			msg := &sarama.ProducerMessage{}
			msg.Topic = t.topic // 每个tailObj自己的topic
			msg.Value = sarama.StringEncoder(line.Text)
			// 丢到通道中
			kafka.ToMsgChan(msg)
		}
	}
}
