package gnet

import (
	"golang.org/x/net/context"
	"time"
)

//调用超时
type TimeoutCallGet struct {
	Next CallGetter
	Timeout int  //毫秒
}

func (tcg *TimeoutCallGet) CallGet(ctx context.Context, key string, client interface{}) ( []byte, error) {
	var cancelFunc context.CancelFunc
	var ch = make(chan bool)
	var err error
	var value []byte
	ctx, cancelFunc= context.WithTimeout(ctx, time.Duration(tcg.Timeout)*time.Millisecond)
	go func () {
		value, err = tcg.Next.CallGet(ctx, key, client)
		ch<- true
	} ()
	select {
	case <-ctx.Done():
		LogDebug("ctx timeout")
		//ctx timeout, call cancelFunc to cancel all the sub-processes in the calling chain
		cancelFunc()
		err = ctx.Err()
	case <-ch:
		LogDebug("call finished normally")
	}
	return value, err
}