package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

const timeoutOpenAPIQuit     = 2 * time.Minute

func main(){
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 一个服务挂了全挂
	g, ctx := errgroup.WithContext(ctx)

	e := gin.Default()

	e.GET("/api/v1/hello",Hello)

	g.Go(func() (err error) {
		defer func() {
			err = WrapPanicErr(err, recover())

		}()

		err = ServerRun(ctx, ":8080", e, timeoutOpenAPIQuit)
		if err != nil {
			err = fmt.Errorf("failed to run open api serv: %w", err)
		}
		return
	})

	if err := g.Wait(); err != nil {
		logrus.Error(err)
	}

	logrus.Info("all services are down")

}