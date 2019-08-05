package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	dataservice "dataservice/proto/dataservice"
)

type Dataservice struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Dataservice) Call(ctx context.Context, req *dataservice.Request, rsp *dataservice.Response) error {
	log.Log("Received Dataservice.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Dataservice) Stream(ctx context.Context, req *dataservice.StreamingRequest, stream dataservice.Dataservice_StreamStream) error {
	log.Logf("Received Dataservice.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&dataservice.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Dataservice) PingPong(ctx context.Context, stream dataservice.Dataservice_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&dataservice.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
