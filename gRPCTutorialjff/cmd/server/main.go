package main

import (
	"bytes"
	context "context"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"

	"github.com/neidersalgado/grpc-google-tutorial/gRPCTutorialjff/todo"
)

func main() {
	srv := grpc.NewServer()
	var tasks T
	todo.RegisterTasksServer(srv, tasks)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}
	log.Fatal(srv.Serve(l))
}