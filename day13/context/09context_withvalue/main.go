package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// context.WithValue

type TraceCode string
type UserID string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE") // 造一个TraceCode类型的结果 类似于n:=int32(100)

	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
	useridKey := UserID("USER_ID")
	userid, ok := ctx.Value(useridKey).(int64) // 在子goroutine中获取userID
	if !ok {
		fmt.Println("invalid user id")
	}
	log.Printf("%s worker func...", traceCode)
	log.Printf("userid:%d", userid)
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	ctx = context.WithValue(ctx, UserID("USER_ID"), int64(21341313213213132))
	log.Printf("%s main 函数", "12512312234")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
