package main

import (
	"fmt"
	"strings"
)

// given a slice of responses, determine if the site is vulnerable
func oracleCLTE(u string, headers []string, bodies []string, payload string) {
	isvuln := false

	var matcher = ""

	matcher = parseMatcher(payload)

	for _, header := range headers {
		//fmt.Println(header)
		if strings.Contains(header, matcher) {
			isvuln = true
		}
	}
	for _, body := range bodies {
		if strings.Contains(body, matcher) {
			isvuln = true
		}
	}

	if isvuln {
		fmt.Println(u, "is vulnerable, payload:", payload)
	}
}
