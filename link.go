package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

// LinkMatches returns true when given
func LinkMatches(mainURL *url.URL, replacement Replacement) bool {
	if mainURL.Host != replacement.HostSource {
		return false
	} else if !strings.HasPrefix(mainURL.RequestURI(), replacement.Path) {
		return false
	}
	return true
}

// RewriteLink Replaces the current replacement by the target one, detects the prefix path and converts any further URL parameters into a hash path
func RewriteLink(mainURL *url.URL, replacement Replacement) string {
	if !LinkMatches(mainURL, replacement) {
		log.Fatalln("RewriteLink state is undefined if the link doesn't match the given replacement")
	}

	newHash := strings.Replace(mainURL.RequestURI(), replacement.Path, "", 1)

	if newHash == "" {
		return fmt.Sprintf("https://%s%s", replacement.HostTarget, replacement.Path)
	}

	if newHash[0] != '/' {
		newHash = "/" + newHash
	}
	return fmt.Sprintf("https://%s%s#%s", replacement.HostTarget, replacement.Path, newHash)
}
