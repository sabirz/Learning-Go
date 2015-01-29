package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int

//var help string

func main() {
	flag.IntVar(&port, "port", 3000, "The port to run the server on")

	help := flag.Int("help", 5, "test test test")
	flag.Parse()

	fmt.Printf("Serving files on localhost: %v\n", port)
	fmt.Printf("taking requests on: %v\n", *help)

	err := ServeStatic(port)
	if err != nil {
		log.Fatalln(err)
	}
}

func ServeStatic(port int) error {
	host := fmt.Sprintf("localhost:%v", port)
	return http.ListenAndServe(host, http.FileServer(http.Dir(".")))

}
