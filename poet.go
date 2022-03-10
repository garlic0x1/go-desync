package main

import (
	"bytes"
	"log"
	"net/url"
	"regexp"
	"strings"
	"text/template"
)

var LIMIT = 5

// sends LIMIT requests to u with specified template
func testTemplate(u string, templatefile string, timeout int) {
	parsed, err := url.Parse(u)
	if err != nil {
		log.Println("failed to parse url", u, err)
	}
	host := parsed.Host
	path := parsed.Path

	// template stuff
	type tpl struct {
		Host string
		Path string
	}
	//temp := template.Must(template.Parse(parseRequest(templatefile)))
	//temp := template.Must(template.ParseFiles(templatefile))

	// get the request out of the yaml file
	temp, err := template.New("request").Parse(parseRequest(templatefile))

	var res bytes.Buffer
	err = temp.Execute(&res, tpl{
		Host: host,
		Path: path,
	})

	// replace newlines with \r\n
	reg := regexp.MustCompile(`\n`)
	result := reg.ReplaceAllString(res.String(), "\r\n")
	//fmt.Println(result)

	// insert custom header if needed
	if USECUSTOM {
		result = insertHeader(result, CUSTOMHEADER)
	}

	// make a type for the channel to use
	type response struct {
		Header string
		Body   string
	}

	// do all the requests concurrently for best chance of success
	// listen on a channel instead of using a waitgroup to prevent race conditions
	c := make(chan response)

	var headers []string
	var bodies []string
	for i := 0; i < LIMIT; i++ {
		go func(host string, result string, timeout int, c chan response) {
			header, body := socketreq(host, result, timeout)
			c <- response{
				Header: header,
				Body:   body,
			}
		}(host, result, timeout, c)
	}

	// listen for responses
	for i := 0; i < LIMIT; i++ {
		respStruct := <-c
		headers = append(headers, respStruct.Header)
		bodies = append(bodies, respStruct.Body)
	}
	// send to oracle
	oracleCLTE(u, headers, bodies, templatefile)
}

func insertHeader(message string, header string) string {
	lines := strings.Split(message, "\n")
	result := ""
	for i := 0; i < len(lines); i++ {
		result = result + lines[i] + "\n"
		if i == 1 {
			result = result + header + "\r\n"
		}
	}
	log.Println(result)
	return result
}
