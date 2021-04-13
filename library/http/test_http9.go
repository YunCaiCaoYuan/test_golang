package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	//"net/url"
	"time"
)

func main() {
	//proxy := func(_ *http.Request) (*url.URL, error) {
	//	return url.Parse("http://127.0.0.1:1080")
	//}
	noRedirect := func(req *http.Request, via []*http.Request) error {
		return errors.New("No Redirect")
	}
	transport := &http.Transport{
		//Proxy: proxy,
		IdleConnTimeout: time.Second * 2048,
		ResponseHeaderTimeout: time.Second * 5,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}
	client := &http.Client{
		Transport:     transport,
		CheckRedirect: noRedirect,
	}

	// create cancelCtx
	ctx, cancel := context.WithCancel(context.TODO())
	// create timer and set the timeout operation
	timer := time.AfterFunc(5*time.Second, func() {
		fmt.Println("now:", time.Now())
		cancel()
	})
	req, err := http.NewRequest("GET", "http://vt1.doubanio.com/201706150839/fd34356de8a03bb99869924284b0ef46/view/movie/F/302160401.flv", nil)
	if err != nil {
		panic(err)
	}
	// set the request context
	req = req.WithContext(ctx)
	fmt.Println("request create")

	fmt.Println(time.Now())
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("request send")
	var length int64
	for {
		//timer.Reset(2 * time.Second)
		// Try instead:
		timer.Reset(50 * time.Millisecond)
		num, err := io.CopyN(ioutil.Discard, resp.Body, 2048*4)
		length += num
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println("received length:", length)
	}

	fmt.Println("jobs done length:", length)
}
