package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"registry-factory/lib"
	"syscall"
)

func main() {
	ctx := context.Background()
	cfg := lib.ServerConfig{
		Port:        7878,
		DockerdHost: "10.160.162.129",
	}

	s := lib.NewProxyServer(cfg)
	go func() {
		s.Start()
	}()

	log.Printf("Server is listening at %s:%d...\n", cfg.Host, cfg.Port)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, os.Kill)

	select {
	case <-ctx.Done():
		log.Println("ctx done!")
	case <-sig:
		log.Println("Gracefully shutting down the server...")
		if err := s.Stop(ctx); err != nil {
			log.Printf("Failed to shutdown server with error: %s\n", err)
		}
	}

	log.Println("Server is shutdown")
}
