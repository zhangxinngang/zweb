package test

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestLoopReq(t *testing.T) {
	//for {
	//time.Sleep(time.Second)
	start := time.Now().UnixNano()
	failed := 0
	for i := 1; i < 3000; i++ {
		go req()
		//fmt.Println(string(data))
	}
	fmt.Println(failed, "failed")
	end := time.Now().UnixNano()
	fmt.Println("last ", end-start, "seconds")
	//}
	time.Sleep(time.Second * 10)
}

func req() {
	cli := http.Client{}
	//_, err := cli.Get("http://127.0.0.1:10301/ads")
	_, err := cli.Get("http://127.0.0.1:8000/main?id=1")
	if err != nil {
		fmt.Println(err)
	}
	// data, _ := ioutil.ReadAll(resp.Body)
	// if string(data) == "" {
	// 	fmt.Println("failed")
	// }
}
