package utils

import (
	"log"
	"net/url"
	"strings"
)

func ContainsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func AppendIfNotExists(slice []string, str string) []string {
	if !ContainsString(slice, str) {
		log.Println("appending", str)
		slice = append(slice, str)
	}
	return slice
}

func IsWorkUrl(work string) bool {
	u, err := url.Parse(work)
	if err != nil {
		log.Println(err)
	}
	if u.Host == "archiveofourown.org" {
		if strings.Contains(u.Path, "works") {
			return true
		}
	}
	return false
}
