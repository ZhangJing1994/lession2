package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	userInfo "lession2/week4/api/user/v1"
	"lession2/week4/internal/pkg/grpc"
	"lession2/week4/internal/service"

	"golang.org/x/sync/errgroup"
)

const (
	address = ":8100"
)

func main() {
	// init service api
	us := InitUserUsecase()
	service := service.NewUserService(us)

	// register grpc service
	s := grpc.NewServer(address)
	userInfo.RegisterUserServer(s, service)

	// context
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	// start grpc server
	g.Go(func() error {
		return s.Start(ctx)
	})

	// catch signals
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-sigs:
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	// wait stop
	if err := g.Wait(); err != nil {
		log.Printf("error: %v", err)
	}
}
