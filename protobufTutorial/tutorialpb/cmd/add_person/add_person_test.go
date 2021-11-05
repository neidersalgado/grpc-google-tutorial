package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/neidersalgado/grpc-google-tutorial/tutorialpb"
)

func TestPromptForAddress_WhenValidAddres_ThenReturnAddress(t *testing.T) {
	in := `12345
    Fake Name
    fake@mail.void
    123-234-345
    home
    234-345-456
    mobile
    234-345-456
    work
    234-345-456
    unknown
    `
	expectedOut := getValidExpectedPersonAddress()

	got, _ := promptForAddress(strings.NewReader(in))

	assert.Equal(t, expectedOut, got)
}

func getValidExpectedPersonAddress() *pb.Person {
	return &pb.Person{
		Name:   "Fake Name",
		Id:     12345,
		Email:  "fake@mail.void",
		Phones: getValidPhones(),
	}
}

func getValidPhones() []*pb.Person_PhoneNumber {
	return []*pb.Person_PhoneNumber{
		{Number: "123-234-345", Type: pb.Person_HOME},
		{Number: "234-345-456", Type: pb.Person_MOBILE},
		{Number: "234-345-456", Type: pb.Person_WORK},
		{Number: "234-345-456", Type: 0},
	}
}
