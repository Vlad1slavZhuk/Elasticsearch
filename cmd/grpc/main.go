package main

import (
	api "Elasticsearch/api/protoc"
	"Elasticsearch/internal/pkg/storage/elastic"
	s "Elasticsearch/internal/pkg/storage/grpc"
	mem "Elasticsearch/internal/pkg/storage/in-memory"
	"Elasticsearch/internal/pkg/storage/mongo"
	"Elasticsearch/internal/pkg/storage/postgres"
	"Elasticsearch/internal/pkg/storage/redis"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

var port string

func init() {
	if port = os.Getenv("GRPC_PORT"); port == "" {
		log.Fatal("Set GRPC_PORT!")
	}
	port = ":" + port
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	ls, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	typeStorage := os.Getenv("TYPE_STORAGE")
	switch typeStorage {
	case "", "memory":
		log.Printf("gRPC server: launching on port %v with in-memory storage\n", ":8001")
		api.RegisterServiceProtobufServer(srv, s.NewStorageGrpcServer(mem.NewStorage()))
	case "redis":
		log.Printf("gRPC server: launching on port %v with redis storage\n", ":8002")
		api.RegisterServiceProtobufServer(srv, s.NewStorageGrpcServer(redis.NewRedisStr()))
	case "postgres":
		log.Printf("GRPC server: launching on port %v with postgres storage\n", ":8003")
		api.RegisterServiceProtobufServer(srv, s.NewStorageGrpcServer(postgres.NewPostgresStorage()))
	case "mongo":
		log.Printf("GRPC server: launching on port %v with mongodb storage\n", ":8004")
		api.RegisterServiceProtobufServer(srv, s.NewStorageGrpcServer(mongo.NewMongoStorage()))
	case "elasticsearch":
		log.Printf("GRPC server: launching on port %v with elasticsearch storage\n", ":8005")
		api.RegisterServiceProtobufServer(srv, s.NewStorageGrpcServer(elastic.NewStorageElasticSearch()))
	default:
		log.Fatal("Set valid STORAGE_SERVICE!")
	}

	go func() {
		log.Fatal(srv.Serve(ls))
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)
	defer close(c)
	srv.GracefulStop()
	log.Printf("gRPC server is shutdown...")
}
