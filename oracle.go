package main

import (
	"fmt"
	"strings"
)

// given a slice of responses, determine if the site is vulnerable
func oracleCLTE(u string, headers []string, bodies []string, payload string) {
	isvuln := false

	for _, body := range headers {
		if strings.Contains(body, "GPOST") {
			isvuln = true
		}
	}

	if isvuln {
		fmt.Println(u, "is vulnerable, payload:", payload)
	}
}
