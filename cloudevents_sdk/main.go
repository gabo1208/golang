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
	event.SetID("example-uuid-32943bac6fea")
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
	fmt.Println()

	data := `{
		"specversion": "1.0",
		"type": "example.type",
		"source": "example/uri",
		"id": "1234",
		"extra": "hey im extra",
		"extra2": "hey im extra",
		"datacontenttype": "test"
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
	fmt.Println(event)
	fmt.Println("---------------------------------")
	fmt.Println()

	body := map[string]interface{}{
		"specversion": "1.0",
		"type":        "test",
		"extra":       0.3,
		"source ":     "jiji",
		"id":          2,
		"data":        "testdata",
	}
	event = cloudevents.NewEvent()
	event.SetDataContentType("test")
	fmt.Println(event)

	event = cloudevents.NewEvent()
	err = setEventAttributes(&event, "datacontenttype", "ce-datacontenttype", body)
	if err != nil {
		fmt.Println(err)
	}

	for key, val := range specs.Version(event.SpecVersion()).Attributes() {
		fmt.Printf("%d:: %s\n", key, val)
	}

	fmt.Println(event)
	fmt.Println("---------------------------------")
	fmt.Println()
	fmt.Println(specs.Version(event.SpecVersion()).Attributes()[0])
}

func isBinary(headers map[string]interface{}) bool {
	return headers["ce-specversion"] != nil || headers["Ce-Specversion"] != nil
}

func isStructured(body map[string]interface{}) bool {
	return body["specversion"] != nil || body["Specversion"] != nil
}

func setBinaryMessageProperties(event *cloudevents.Event, headers map[string]interface{}) error {
	for ceKey, val := range headers {
		lowerKey := strings.ToLower(ceKey)

		if lowerKey == "content-type" {
			continue
		}

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
	var dataVal interface{}

	for key, val := range body {
		if strings.ToLower(key) == "content-type" {
			continue
		}

		// To prevent setting the data before the contenttype, the setData call
		// is left after the cycle is done
		if strings.ToLower(key) == "data" {
			dataVal = val
			continue
		}

		ceKey := strings.ToLower(prefix + key)
		err := setEventAttributes(event, key, ceKey, val)
		if err != nil {
			return err
		}
	}

	// Here it sets the data with its correct type
	err := event.SetData(event.DataContentType(), dataVal)
	if err != nil {
		return err
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
		// this is how the extension error keys are stored on the event.FieldErrors
		err = event.FieldErrors["extension:"+key]
	}

	return err
}
