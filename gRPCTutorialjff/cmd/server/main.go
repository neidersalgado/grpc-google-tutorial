package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"

	"github.com/neidersalgado/grpc-google-tutorial/gRPCTutorialjff/todo"
)

func main() {
	
}

type length int64

const (
	sizeOfLength = 1
	dbPath       = "mydb.pb"
)

var endianness = binary.LittleEndian

func add(text string) error {
	task := &todo.Task{
		Text: text,
		Done: false,
	}

	b, err := proto.Marshal(task)
	if err != nil {
		return fmt.Errorf("could not encode task: %v", err)
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("could not open %s: %v", dbPath, err)
	}

	if err := binary.Write(f, endianness, length(len(b))); err != nil {
		return fmt.Errorf("could not encode length of message: %v", err)
	}
	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("could not write task to file: %v", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close file %s: %v", dbPath, err)
	}
	return nil
}

func list() error {
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return fmt.Errorf("could not read %s: %v", dbPath, err)
	}

	for {
		if len(b) == 0 {
			return nil
		} else if len(b) < sizeOfLength {
			return fmt.Errorf("remaining odd %d bytes, what to do?", len(b))
		}

		var l length
		if err := binary.Read(bytes.NewReader(b[:sizeOfLength]), endianness, &l); err != nil {
			return fmt.Errorf("could not decode message length: %v", err)
		}
		b = b[sizeOfLength:]
		//TODO improve the printing of file to avoid print while tru, and EOF
		var task todo.Task
		if err := proto.Unmarshal(b[:l], &task); err != nil {
			return fmt.Errorf("could not read task: %v", err)
		}
		b = b[l:]

		if task.Done {
			fmt.Printf("ok")
		} else {
			fmt.Printf("not ok")
		}
		fmt.Printf(" %s\n", task.Text)
	}
}
