package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		uflag    string
		urls     []string
		urlsfile string
		proxy    string
	)

	flag.StringVar(&uflag, "u", "", "Target URL")
	flag.StringVar(&urlsfile, "urls", "", "List of URLs")
	flag.StringVar(&proxy, "proxy", "", "Set the Golang proxy, for example: http://example.com:8080")
	flag.Parse()

	// get urls from flags
	if uflag == "" && urlsfile == "" {
		fmt.Println("Must include a target with -u or a list of targets with -urls")
		os.Exit(0)
	} else if uflag == "" && urlsfile != "" {
		file, err := os.Open(urlsfile)
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		file.Close()
	} else {
		urls = append(urls, uflag)
	}

	if proxy != "" {
		fmt.Println("Setting up proxy at", proxy)
		os.Setenv("HTTP_PROXY", proxy)
	}

	for _, u := range urls {
		log.Println(u)
		testTemplate(u, "template.txt")
		testTemplate(u, "template2.txt")
		testTemplate(u, "template3.txt")
	}
}
