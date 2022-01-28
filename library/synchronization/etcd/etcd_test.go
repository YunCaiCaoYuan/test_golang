package etcd_test

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"testing"
	"time"
)

// fixme 没通！！！

var endpoints = []string{"localhost:12379", "localhost:22379", "localhost:32379"}
var dialTimeout = 5 * time.Second

func Test_ExampleMaintenance_status(t *testing.T) {
	for _, ep := range endpoints {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   []string{ep},
			DialTimeout: dialTimeout,
		})
		if err != nil {
			log.Fatal(err)
		}
		defer cli.Close()

		resp, err := cli.Status(context.Background(), ep)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("endpoint: %s / Leader: %v\n", ep, resp.Header.MemberId == resp.Leader)
	}
	// endpoint: localhost:2379 / Leader: false
	// endpoint: localhost:22379 / Leader: false
	// endpoint: localhost:32379 / Leader: true
}
