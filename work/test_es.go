package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"os"
	"time"
)

type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

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
	fmt.Println("client start:", client)
	if err != nil {
		fmt.Println(err)
	}

	indexName := "twitter"
	ctx := context.Background()

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	if exists {
		fmt.Println("index exists")
	} else {
		// Create a new index.
		createIndex, err := client.CreateIndex(indexName).Do(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
		if !createIndex.Acknowledged {
			fmt.Println("create index not acknowledged")
		}
	}

	// Index a tweet (using JSON serialization)
	tweet1 := Tweet{User: "olivere", Message: "Take Five", Retweets: 0}
	_, err = client.Index().
		Index(indexName). // es index name
		Type("tweet").
		Id("1").
		BodyJson(tweet1).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Search with a term query
	searchResult, err := client.Search().
		Index(indexName). // search in index t.index
		Query(elastic.NewTermQuery("user", "olivere")). // specify the query
		//Sort("user", true).                             // sort by "user" field, ascending
		//From(0).Size(10).                               // take documents 0-9
		Pretty(true). // pretty print request and response JSON
		Do(ctx) // execute
	if err != nil {
		fmt.Printf("Search failed: %v\n", err)
		return
	}
	// Number of hits
	if searchResult.Hits.TotalHits.Value > 0 {
		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var tweet Tweet
			err := json.Unmarshal(hit.Source, &tweet)
			if err != nil {
				// Deserialization failed
				fmt.Printf("Deserialize failed: %v\n", err)
				continue
			}
			fmt.Printf("tweet: %v\n", tweet)
		}
	} else {
		// No hits
		fmt.Print("Found no tweets\n")
	}
}
