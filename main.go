package main

import (
	"fmt"
	"log"
	"net/http"
	"flag"
)

var (
	configFile  = flag.String("config", "", "YAML Configuration File Path")
	listenPort     = flag.String("port", "9595", "Listen Port")
	defaultAddress = flag.String("default", "127.0.0.1:8080", "Default Server Address")
)

func usage() {
	fmt.Println("Univility 0.1")
	fmt.Println("usage: univility [options]")
	fmt.Println("")

	fmt.Println("Options:")
	fmt.Println("  --config   <>         YAML Configuration File Path")
	fmt.Println("  --port     <:9595>    Listen Port")
	fmt.Println("")

	fmt.Println("Examples:")
	fmt.Println("  To serve on port 9598 and use config file at /tmp/unility.yml")
	fmt.Println("  $ univility --config /tmp/unility.yml ---port 9598")
	fmt.Println("")
}


func main() {
	flag.Usage = usage
	flag.Parse()

	http.HandleFunc("/", sayhelloName)       // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("[FATAL] ListenAndServe: ", err)
	}
}
