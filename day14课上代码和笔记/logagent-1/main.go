package main

import (
	"code.oldboy.com/studygolang/logagent/kafka"
	"code.oldboy.com/studygolang/logagent/tailfile"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

// 日志收集的客户端
// 类似的开源项目还有filebeat
// 收集指定目录下的日志文件,发送到kafka中

// 现在的技能包:
// 往kafka发数据
// 使用tail读日志文件

// 整个logaent的配置结构体
type Config struct {
	KafkaConfig `ini:"kafka"`
	CollectConfig `ini:"collect"`
}
type KafkaConfig struct {
	Address string `ini:"address"`
	ChanSize int64 `ini:"chan_size"`
}

type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`
}

// run 真正的业务逻辑
func run ()(err error){
	// logfile --> TailObj --> log --> Client --> kafka
	for {
		// 循环读数据
		line, ok := <-tailfile.TailObj.Lines // chan tail.Line
		if !ok {
			logrus.Warn("tail file close reopen, filename:%s\n", tailfile.TailObj.Filename)
			time.Sleep(time.Second) // 读取出错等一秒
			continue
		}
		// 如果是空行就略过
		//fmt.Printf("%#v\n", line.Text)
		if len(strings.Trim(line.Text,"\r")) == 0 {
			logrus.Info("出现空行拉,直接跳过...")
			continue
		}
		// 利用通道将同步的代码改为异步的
		// 把读出来的一行日志包装秤kafka里面的msg类型
		msg := &sarama.ProducerMessage{}
		msg.Topic = "web_log"
		msg.Value = sarama.StringEncoder(line.Text)
		// 丢到通道中
		kafka.ToMsgChan(msg)
	}
}
func main(){
	var configObj = new(Config)
	// 0. 读配置文件 `go-ini`
	err := ini.MapTo(configObj, "./conf/config.ini")
	if err != nil {
		logrus.Error("load config failed,err:%v", err)
		return
	}
	fmt.Printf("%#v\n", configObj)
	// 1. 初始化连接kafka(做好准备工作)
	err = kafka.Init([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)
	if err != nil {
		logrus.Error("init kafka failed, err:%v", err)
		return
	}
	logrus.Info("init kafka success!")
	// 2. 根据配置中的日志路径初始化tail
	err = tailfile.Init(configObj.CollectConfig.LogFilePath)
	if err != nil {
		logrus.Error("init tailfile failed, err:%v", err)
		return
	}
	logrus.Info("init tailfile success!")
	// 3. 把日志通过sarama发往kafka
	err = run()
	if err != nil {
		logrus.Error("run failed ,err :%v", err)
		return
	}
}
