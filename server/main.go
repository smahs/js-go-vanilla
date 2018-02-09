package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	// Version or build identifier; set it during building.
	Version     string
	versionFlag = flag.Bool("version", false, "Print build version and exit")
)

func init() {
	flag.Parse()
	if Version == "" {
		Version = "v0"
	}
}

func formatVersion() string {
	return fmt.Sprintf("Server Version: %s", Version)
}

func printVersion() {
	fmt.Println(formatVersion())
	os.Exit(0)
}

func serveVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(formatVersion()))
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}

	data, err := Asset(path[1:len(path)])
	if err != nil {
		w.WriteHeader(404)
		return
	}

	// Set the mime type for css
	if strings.HasSuffix(path, "css") {
		w.Header().Set("Content-Type", "text/css")
	}

	w.Write(data)
}

func main() {
	if *versionFlag {
		printVersion()
	}

	http.HandleFunc("/version", serveVersion)
	http.HandleFunc("/", serveStatic)

	http.ListenAndServe(":9000", nil)
}
