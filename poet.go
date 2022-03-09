package main

import (
	"bytes"
	"log"
	"net/url"
	"regexp"
	"sync"
	"text/template"
)

var LIMIT = 5

func testTemplate(u string, templatefile string, timeout int) {
	parsed, err := url.Parse(u)
	if err != nil {
		log.Println("failed to parse url", u, err)
	}
	host := parsed.Host
	path := parsed.Path

	type tpl struct {
		Host string
		Path string
	}
	temp := template.Must(template.ParseFiles(templatefile))

	var res bytes.Buffer
	err = temp.Execute(&res, tpl{
		Host: host,
		Path: path,
	})
	reg := regexp.MustCompile(`\n`)

	result := reg.ReplaceAllString(res.String(), "\r\n")

	var wg sync.WaitGroup

	var headers []string
	var bodies []string
	for i := 0; i < LIMIT; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			header, body := socketreq(host, result, timeout)
			headers = append(headers, header)
			bodies = append(bodies, body)
		}()
	}
	wg.Wait()

	oracleCLTE(u, headers, bodies, templatefile)
}
