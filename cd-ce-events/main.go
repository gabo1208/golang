package main

import (
	"fmt"
	"log"

	cdevents "github.com/cdevents/sdk-go/pkg/api"
)

func main() {
	event, err := cdevents.NewPipelineRunQueuedEvent()
	if err != nil {
		log.Fatalf("could not create a cdevent, %v", err)
	}

	event.SetId("1")
	err = event.SetCustomData("text", []byte("hello"))
	if err != nil {
		log.Fatalf("could not set cdevent custom data, %v", err)
	}
	// Set the required context fields
	event.SetSubjectId("myPipelineRun1")
	event.SetSource("my/first/cdevent/program")

	// Set the required subject fields
	event.SetSubjectPipelineName("myPipeline")
	event.SetSubjectUrl("https://example.com/myPipeline")
	fmt.Println(event)

	ce, err := cdevents.AsCloudEvent(event)
	if err != nil {
		log.Fatalf("could not create a cloudevent, %v", err)
	}
	fmt.Println(ce)
}
