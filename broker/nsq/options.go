package nsq

import (
	"context"

	"github.com/divisionone/go-micro/broker"
)

type contextKeyT string

var (
	concurrentHandlerKey = contextKeyT("github.com/divisionone/go-plugins/broker/nsq/concurrentHandlers")
)

func ConcurrentHandlers(n int) broker.SubscribeOption {
	return func(o *broker.SubscribeOptions) {
		o.Context = context.WithValue(o.Context, concurrentHandlerKey, n)
	}
}
