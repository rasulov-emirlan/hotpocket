package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	cfg.Directory = dir + "/"
	wtchr, err := StartWatching(cfg.Directory, cfg.Exceptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		if err := cfg.ExecCommands(ctx); err != nil && err != context.Canceled {
			panic(err)
		}
	}()

	log.Println("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	for {
		select {
		case <-wtchr.Events:
			cancel()
			time.Sleep(time.Second * 5)
			go func() {
				if err := cfg.ExecCommands(ctx); err != nil {
					panic(err)
				}
			}()
			log.Println("Reloaded")
		case <-quit:
			cancel()
			log.Println("Gracefully ended hotpocket session")
			return
		default:
			time.Sleep(time.Second * 5)
			continue
		}
	}
}

func (c *Config) ExecCommands(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, c.Command, c.Arguments...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
