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

	"github.com/yunuseskisan/touch/service.rpc.assets/internal/db/migrations"
	"github.com/yunuseskisan/touch/service.rpc.assets/internal/handler"
	assetsproto "github.com/yunuseskisan/touch/service.rpc.assets/proto"
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

	ordersConn, err := grpc.Dial("service.rpc.orders:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	ordersClient := ordersproto.NewOrdersServiceClient(ordersConn)

	assetsproto.RegisterAssetsServiceServer(mux, handler.New(conn, ordersClient))

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
