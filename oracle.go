package main

import (
	"fmt"
	"strings"
)

// given a slice of responses, determine if the site is vulnerable
func oracleCLTE(u string, headers []string, bodies []string, payload string) {
	isvuln := false

	var matcher = ""

	args := parseOracle(payload)
	rows := strings.Split(args, "\n")
	for i := 0; i < len(rows); i++ {
		words := strings.SplitN(rows[i], " ", 1)
		if words[0] == "match" {
			matcher = words[1]
		}
	}

	for _, body := range headers {
		if strings.Contains(body, matcher) {
			isvuln = true
		}
	}

	if isvuln {
		fmt.Println(u, "is vulnerable, payload:", payload)
	}
}
