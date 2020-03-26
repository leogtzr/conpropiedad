package main

import "strings"

func validateOptions(options ...string) bool {
	r := true

	for _, opt := range options {
		if len(strings.TrimSpace(opt)) == 0 {
			r = false
			break
		}
	}
	return r
}
