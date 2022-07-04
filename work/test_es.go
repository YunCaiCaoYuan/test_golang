package main

import (
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Obtain a client for an Elasticsearch cluster of two nodes,
	// running on 10.0.1.1 and 10.0.1.2. Do not run the sniffer.
	// Set the healthcheck interval to 10s. When requests fail,
	// retry 5 times. Print error messages to os.Stderr and informational
	// messages to os.Stdout.
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetRetrier(NewCustomRetrier()),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetHeaders(http.Header{
			"X-Caller-Id": []string{"..."},
		}),
	)
	fmt.Println("client start", client)
	if err != nil {
		fmt.Println(err)
	}
}
