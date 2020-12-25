package gnet

import (
	"golang.org/x/net/context"
	"log"
	"time"
)

type RetryCallGet struct {
	Next CallGetter
	Count  int  //重试次数
	Interval int //重试间隔
}

func (tcg *RetryCallGet) CallGet(ctx context.Context, key string, client interface{}) ( []byte, error) {
	var err error
	var value []byte
	for i:=0; i< tcg.Count; i++ {
		value, err = tcg.Next.CallGet(ctx, key, client)
		log.Printf("Retry number %v|error=%v", i+1, err)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(tcg.Interval)*time.Millisecond)
	}
	return value, err
}