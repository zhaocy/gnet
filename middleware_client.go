package gnet

import "golang.org/x/net/context"

type CallGetter interface {
	CallGet(ctx context.Context, key string, client interface{}) ( []byte, error)
}

type CallGetMiddleware struct {
	Next CallGetter
}

func (cg *CallGetMiddleware) CallGet(ctx context.Context, key string, client interface{}) ( []byte, error) {
	return cg.Next.CallGet(ctx, key, client)
}

