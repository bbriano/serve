// Serve serves a directory over HTTP.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var port = flag.Int("port", 9000, "port to listen on")

func main() {
	flag.Parse()
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: serve [options] directory")
		os.Exit(1)
	}
	dir := os.Args[1]
	stat, err := os.Stat(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	if !stat.IsDir() {
		fmt.Fprintf(os.Stderr, "Error: %s is not a directory\n", dir)
		os.Exit(1)
	}
	addr := "0.0.0.0:" + strconv.Itoa(*port)
	handler := http.FileServer(http.Dir(dir))
	fmt.Println("Serving on http://" + addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
