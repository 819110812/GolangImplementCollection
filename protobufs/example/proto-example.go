package main

import (
	test "GolangImplementation/protobufs"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {

	// Output:
	// Hello, world.
	println("Hello, world.")
	person := doSomething()
	WriteToFile("person.bin", person)

	// readFromFile()
	persons := &test.Person{}
	ReadFromFile("person.bin", persons)

	fmt.Println(toJson(person))
}

func WriteToFile(fname string, protoMessage proto.Message) error {
	outFile, err := proto.Marshal(protoMessage)
	if err != nil {
		fmt.Println("Error marshalling proto message: ", err)
		return err
	}
	if err := ioutil.WriteFile(fname, outFile, 0644); err != nil {
		log.Fatalln("Error writing file: ", err)
		return err
	}
	return nil

}

func ReadFromFile(fname string, pd proto.Message) error {
	out, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	if err := proto.Unmarshal(out, pd); err != nil {
		log.Fatalln("error: ", err)
	}
	fmt.Println(string(out))
	return nil
}

func toJson(pd proto.Message) string {
	marshal := jsonpb.Marshaler{}
	out, err := marshal.MarshalToString(pd)
	if err != nil {
		log.Fatalln(err)
	}

	return out
}

func doSomething() *test.Person {
	person := test.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*test.Person_PhoneNumber{
			{Number: "555-4321", Type: test.Person_HOME},
		},
	}

	person.Name = "Jane Doe1"
	fmt.Printf(string(person.GetId()))
	return &person
}
