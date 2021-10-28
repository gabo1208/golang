package main

import (
	"encoding/json"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/binding/spec"
)

const (
	prefix = "ce-"
)

var (
	specs = spec.WithPrefix(prefix)
)

func main() {
	event := cloudevents.NewEvent()
	event.SetID("example-uuid-32943bac6fea")
	event.SetSource("example/uri")
	event.SetType("example.type")
	event.SetData(cloudevents.ApplicationJSON, map[string]string{"specversion": "0.3"})
	event.SetExtension("test", "testing")
	event.SetSubject("testsub")
	event.SetDataContentType("testdatacontt")
	event.SetDataSchema("testdatasch")
	//fmt.Println(event)
	//fmt.Println("---------------------------------")
	//fmt.Println()*/

	data := `{
		"specversion": "1.0",
		"type": "example.type",
		"source": "example/uri",
		"id": 1234,
		"extra": "hey im extra"
	}`
	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		fmt.Println(err, data)
	}
	fmt.Println(event)
	fmt.Println("---------------------------------")
	fmt.Println()

	err = specs.Version(event.SpecVersion()).SetAttribute(event.Context, "ce-time", "2021-10-27T19:55:17.3224399Z")
	if err != nil {
		fmt.Println(err, data)
	}

	fmt.Println(event)
	fmt.Println("---------------------------------")
	fmt.Println()
	fmt.Println(specs.Version(event.SpecVersion()).Attribute("ce-time"))
}
