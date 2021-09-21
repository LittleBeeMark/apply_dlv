package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ServerRun server run
func ServerRun(ctx context.Context, addr string, e *gin.Engine, quitTimeout time.Duration) (err error) {

	defer func() {
		if err != nil {
			if gin.IsDebugging() {
				fmt.Fprintf(os.Stderr, "[GIN-debug] [ERROR] %v\n", err)
			}
		}
	}()

	s := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	format := "Listening and serving HTTP on %s\n"

	if gin.IsDebugging() {
		fmt.Fprintf(os.Stdout, "[GIN-debug] "+format, addr)
	}

	go func() {
		err = s.ListenAndServe()
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), quitTimeout)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		logrus.Errorf("failed to http server shutdown: %s", err.Error())
	}

	return
}


// WrapPanicErr wrap panic to error
func WrapPanicErr(err error, panicErr interface{}) error {
	if panicErr == nil {
		return err
	}

	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	newErr := fmt.Errorf("unexpected error: %v\nstack: %s", panicErr, buf[:n])

	if err == nil {
		return newErr
	}

	return fmt.Errorf("panic error: %s\nreturn error: %w", newErr, err)
}
