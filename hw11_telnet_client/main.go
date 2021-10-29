package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "")
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)
	client := NewTelnetClient(address, *timeout, os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		log.Fatalln("failed to connect:", err)
	}

	defer client.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer cancel()
		client.Send()
	}()

	go func() {
		defer cancel()
		client.Receive()
	}()

	defer cancel()
	<-ctx.Done()
}
