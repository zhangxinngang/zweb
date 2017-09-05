package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime/pprof"
	"zweb/handler"
	"zweb/route"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	route.AddRoute("GET", "/main", handler.GetBlogHandler)
	route.AddRoute("POST", "/main", handler.PostBlogHandler)
	route.AddRoute("GET", "/hello", handler.GetHelloHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
