package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	dataservice "github.com/dexinq/dataservice/proto/dataservice"
)

type Dataservice struct{}

func (e *Dataservice) Handle(ctx context.Context, msg *dataservice.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *dataservice.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
