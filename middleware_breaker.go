package gnet

import (
	"github.com/sony/gobreaker"
	"golang.org/x/net/context"
)

var cb *gobreaker.CircuitBreaker

type CircuitBreakerCallGet struct {
	Next CallGetter
}

func init() {
	var st gobreaker.Settings
	st.Name = "CircuitBreakerCallGet"
	st.MaxRequests = 2
	st.Timeout = 10
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 2 && failureRatio >= 0.6
	}
	cb = gobreaker.NewCircuitBreaker(st)
}

func (tcg *CircuitBreakerCallGet) CallGet(ctx context.Context, key string, client interface{}) ( []byte, error) {
	var err error
	var value []byte
	var serviceUp bool
	LogDebug("state:%v", cb.State().String())
	_, err = cb.Execute(func() (interface{}, error) {
		value, err = tcg.Next.CallGet(ctx, key, client)
		if err != nil {
			return nil, err
		}
		serviceUp = true
		return value, nil
	})
	if !serviceUp {
		//return a default value here. You can also run a downgrade function here
		LogError("circuit breaker return error:%v\n", err)
		return []byte{0}, nil
	}
	return value, nil
}