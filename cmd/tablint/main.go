package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aerth/tabs/htmlparse"
)

func main() {
	flag.Parse()
	filenames := flag.Args()
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatalln(err)
		}
		b, err := htmlparse.Parse(f)
		if err != nil {
			log.Fatalln(err)
		}
		filename = strings.TrimSuffix(filename, ".html")
		fmt.Fprintf(os.Stdout, "<h1>%s</h1>\n\n", filename)
		os.Stdout.Write(b)
		fmt.Fprintf(os.Stdout, "\n\n\n\n")
	}
}
