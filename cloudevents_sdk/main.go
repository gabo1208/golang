package main

import (
	"encoding/json"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func main() {
	event := cloudevents.NewEvent()
	event.SetID("example-uuid-32943bac6fea")
	event.SetSource("example/uri")
	event.SetType("example.type")
	event.SetData(cloudevents.ApplicationJSON, map[string]string{"hello": "world"})
	event.SetExtension("test", "testing")
	fmt.Println(event.Extensions())

	bytes, err := json.Marshal(event)
	fmt.Println(err)
	fmt.Println(event)

	event = cloudevents.NewEvent()
	err = json.Unmarshal(bytes, &event)
	fmt.Println(err)
	fmt.Println(event)
	fmt.Println(event.Extensions())
}
