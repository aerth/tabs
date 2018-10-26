package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aerth/tabs/htmlparse"
	"github.com/aerth/tabs/jsparse"
)

func main() {
	flag.Parse()
	filenames := flag.Args()
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatalln(err)
		}
		b, err := jsparse.Parse(f)
		if err != nil {
			f.Seek(0, 0)
			b, err = htmlparse.Parse(f)
			log.Fatalln(err)
		}
		filename = strings.TrimSuffix(filename, ".html")
		fmt.Fprintf(os.Stdout, "<h1>%s</h1>\n\n", filename)
		os.Stdout.Write(b)
		fmt.Fprintf(os.Stdout, "\n\n\n\n")
	}
}
