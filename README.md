# grpc-google-tutorial
this is just a  learning project to use gRPC


here I used gRPC so,  it's necessary to install  protobuf compiler,  the OS over  I worked is Ubuntu.

to install, open terminal and type following commands:

>sudo apt update
>sudo install protobuf-compiler

to delete protobuf if it's necessary use:

>sudo apt-get remove protobuf-compiler


Run the following command to install the Go protocol buffers plugin:

>go get google.golang.org/protobuf/cmd/protoc-gen-go

compile proto file in go syntax using:

> "protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto"

where, $SRC_DIR is your source directory, $DST_DIR destination code, where your code w0ill be generated, and the path where the '-proto' is.

then to execute add_person, navitate to the path were add_person.go file is, an the run:

>go run add_person.go <file_name_addreses>

file_name_addresses is a file were  is being saving the directory of persons