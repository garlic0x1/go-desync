package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	LIMIT        = 5
	USECUSTOM    = false
	CUSTOMHEADER = ""
)

func main() {
	var (
		uflag    string
		urls     []string
		urlsfile string
		tempdir  string
		proxy    string
		nthreads int
		timeout  int
	)

	flag.StringVar(&uflag, "u", "", "Target URL")
	flag.StringVar(&urlsfile, "urls", "", "List of URLs")
	flag.StringVar(&tempdir, "templates", "templates/", "Directory of YAML templates to test")
	flag.StringVar(&CUSTOMHEADER, "header", "", "Custom header to add to requests, example: '-header \"User-Agent: garlic0x1\"'")
	flag.IntVar(&LIMIT, "tries", 5, "Number of requests to send to test each template")
	flag.IntVar(&nthreads, "threads", 5, "Number of concurrent targets to test")
	flag.IntVar(&timeout, "timeout", 10, "Timeout")
	flag.StringVar(&proxy, "proxy", "", "Set the environment proxy, for example: http://example.com:8080")
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

	if CUSTOMHEADER != "" {
		USECUSTOM = true
	}

	// get templates from folder
	files, err := ioutil.ReadDir(tempdir)
	if err != nil {
		log.Fatal(err)
	}

	// limit concurrency
	sem := make(chan struct{}, nthreads)

	// creat type for channel
	type response struct {
		URL      string
		Headers  []string
		Bodies   []string
		Template string
	}

	for _, file := range files {
		filename := tempdir + file.Name()

		c := make(chan response)

		// targets looped inside templates to distribute load
		for _, u := range urls {
			select {
			case sem <- struct{}{}:
				go func(u string, filename string, timeout int) {
					headers, bodies := testTemplate(u, filename, timeout)

					// the listener loop will already be running, so responding to c will be immediate
					c <- response{
						URL:      u,
						Headers:  headers,
						Bodies:   bodies,
						Template: filename,
					}
					// add to the count of routines
					<-sem
				}(u, filename, timeout)
			default:
				testTemplate(u, filename, timeout)
			}
		}

		// listen for responses
		for _, _ = range urls {
			respStruct := <-c
			oracle(respStruct.URL, respStruct.Headers, respStruct.Bodies, respStruct.Template)
		}
	}
}
