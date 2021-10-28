package main

import (
	"encoding/json"
	"fmt"
	"strings"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/binding/spec"
	"github.com/cloudevents/sdk-go/v2/types"
)

const (
	prefix = "ce-"
)

var (
	specs = spec.WithPrefix(prefix)
)

func main() {
	event := cloudevents.NewEvent()
	event2 := cloudevents.NewEvent()
	/*event.SetID("example-uuid-32943bac6fea")
	event.SetSource("example/uri")
	event.SetType("example.type")
	event.SetData(cloudevents.ApplicationJSON, map[string]string{"blah": "0.3"})
	event.SetExtension("test", "testing")
	event.SetSubject("testsub")
	event.SetDataContentType("testdatacontt")
	event.SetDataSchema("testdatasch")
	event.SetSpecVersion("0.3")
	fmt.Println(event.SpecVersion())
	fmt.Println("---------------------------------")
	fmt.Println()*/

	data := `{
		"specversion": "1.0",
		"type": "example.type",
		"source": "example/uri",
		"id": "1234",
		"extra": "hey im extra",
		"extra2": "hey im extra"
	}`
	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		fmt.Println(err, data)
	}

	data2 := `{
		"specversion": "1.0",
		"type": "example.type",
		"id": "1234",
		"source": "example/uri",
		"extra2": "hey im extra",
		"extra": "hey im extra"
	}`
	err = json.Unmarshal([]byte(data2), &event2)
	if err != nil {
		fmt.Println(err, data)
	}
	fmt.Println(event.String() == event2.String())
	fmt.Println("---------------------------------")
	fmt.Println()

	body := map[string]interface{}{
		"specversion":         "0.3",
		"type":                "test",
		"extra":               0.3,
		"source_erASD123123 ": "jiji",
		"id":                  2,
	}
	event = cloudevents.NewEvent()
	err = setEventAttributes(&event, "ce-Test-esd", "ce-ce-testsd", "body")
	if err != nil {
		fmt.Println(err)
	}

	for key, val := range event.FieldErrors {
		fmt.Printf("%s:: %s\n", key, val)
	}

	fmt.Println(event)
	fmt.Println("---------------------------------")
	fmt.Println()
	fmt.Println(specs.Version(body["specversion"].(string)).Attribute("ce-specversion"))
}

func setBinaryMessageProperties(event *cloudevents.Event, headers map[string]interface{}) error {
	for ceKey, val := range headers {
		lowerKey := strings.ToLower(ceKey)
		if strings.HasPrefix(lowerKey, prefix) {
			key := strings.TrimPrefix(lowerKey, prefix)
			err := setEventAttributes(event, key, lowerKey, val)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func setStructuredMessageProperties(event *cloudevents.Event, body map[string]interface{}) error {
	for key, val := range body {
		ceKey := strings.ToLower(prefix + key)
		err := setEventAttributes(event, key, ceKey, val)
		if err != nil {
			return err
		}
	}

	return nil
}

func setEventAttributes(event *cloudevents.Event, key, ceKey string, val interface{}) error {
	attr := specs.Version(event.SpecVersion()).Attribute(ceKey)
	// convert everything to string after this point
	s, err := types.Format(val)
	if err != nil {
		return err
	}

	if attr != nil {
		err = specs.Version(event.SpecVersion()).SetAttribute(event.Context, ceKey, s)
	} else {
		event.SetExtension(key, s)
		err = event.FieldErrors["extension:"+key]
	}

	return err
}
