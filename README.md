# grpc-google-tutorial
this is just a  learning project to use gRPC


here I used gRPC so,  it's necessary to install the protobuf compiler,  and was installed on Ubuntu

to install, open terminal, and type the following commands:

>sudo apt update
>sudo install protobuf-compiler

to delete protobuf if it's necessary to use:

>sudo apt-get remove protobuf-compiler

Run the following command to install the Go protocol buffers plugin:

>go get google.golang.org/protobuf/cmd/protoc-gen-go

compile .proto file in go syntax using:

> "protoc -I= $SRC_DIR --go_out= $DST_DIR $SRC_DIR/addressbook.proto"

where, $SRC_DIR is your source directory, $DST_DIR destination code, where your code will be generated, and the path where the '.proto' is.

then to execute add_person, navigate to the path were add_person.go file is, and then run:

>go run add_person.go <file_name_addresses>

file_name_addresses is a file where is being saved the directory of persons.

for test I'm using "github.com/stretchr/testify/assert" for assertions in tests you can add with

> go get "github.com/stretchr/testify/assert"