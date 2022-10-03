package ms

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func (s *Service) Server() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := s.Echo.Start(":" + s.Config.Port); err != nil && err != http.ErrServerClosed {
			s.Echo.Logger.Fatalf("listen: %\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Echo.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}
