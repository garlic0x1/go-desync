package main

import (
	"fmt"
	"strings"
)

// given a slice of responses, determine if the site is vulnerable
func oracleCLTE(u string, headers []string, bodies []string) {

	for _, body := range headers {
		if strings.Contains(body, "GPOST") {
			fmt.Println(u, "Is vulnerable to CLTE")
		}
	}
}
