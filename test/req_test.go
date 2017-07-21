package test

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestLoopReq(t *testing.T) {
	for {
		time.Sleep(time.Second)
		go req()
	}
}

func req() {
	for i := 1; i < 10000; i++ {
		cli := http.Client{}
		//_, err := cli.Get("http://127.0.0.1:10301/ads")
		_, err := cli.Get("http://127.0.0.1:8000/main?id=1")
		if err != nil {
			fmt.Println(err, i)
		}
	}
}
