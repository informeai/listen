package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	usage := "LISTEN:\nExpose file in port local\nUSAGE:\nlisten --file [file] -p [port]"
	if len(os.Args) != 5 {
		fmt.Println(usage)
		os.Exit(0)
	}
	var file string
	var port string
	flag.StringVar(&file, "file", "", "file for expose")
	flag.StringVar(&port, "p", "", "port for listen")
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slcBytes, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		w.Write(slcBytes)
	})
	fmt.Printf("listen execute in port: %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
