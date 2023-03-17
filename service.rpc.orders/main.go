package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-pg/pg/v10"
	"google.golang.org/grpc"

	"github.com/yunuseskisan/touch/service.rpc.orders/internal/db/migrations"
	"github.com/yunuseskisan/touch/service.rpc.orders/internal/handler"
	ordersproto "github.com/yunuseskisan/touch/service.rpc.orders/proto"
)

func main() {
	opt, err := pg.ParseURL("postgresql://touch:touch@postgres:5432/touch?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	conn := pg.Connect(opt)
	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	_, _, err = migrations.Collection.Run(conn, "init")
	if err != nil {
		log.Fatal(err)
	}

	oldVersion, toVersion, err := migrations.Collection.Run(conn, "up")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully migrated database. From Version %d To Version %d", oldVersion, toVersion)

	mux := grpc.NewServer()
	ordersproto.RegisterOrdersServiceServer(mux, handler.New(conn))

	var wg sync.WaitGroup
	wg.Add(1)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer wg.Done()
		if err := mux.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		for {
			<-sigint
			mux.GracefulStop()
		}
	}()

	wg.Wait()
}
