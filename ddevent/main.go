package main

import (
	"github.com/DataDog/datadog-go/statsd"
	"log"
)

func main() {
	client, err := statsd.New("127.0.0.1:8125",
		statsd.WithNamespace("flubber."),               // prefix every metric with the app name
		statsd.WithTags([]string{"region:us-east-1a"}), // send the EC2 availability zone as a tag with every metric
		// add more options here...
	)
	err = client.Event(&statsd.Event{Title: "TESTING DD GO SDK", AlertType: "success", Tags: []string{"env:prod", "source:jenkins","test:123"},Text:"some text"})
	if err != nil {
		log.Fatal(err)
	}
}
