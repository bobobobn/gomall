package service

import (
	"context"
	example "demo/hello_demo/kitex_gen/hello/example"
	"testing"
)

func TestHelloMethod_Run(t *testing.T) {
	ctx := context.Background()
	s := NewHelloMethodService(ctx)
	// init req and assert value

	request := &example.HelloReq{}
	resp, err := s.Run(request)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
