package main

import (
	"net/url"
	"testing"
)

func TestLinkMatches(t *testing.T) {
	//
	domain := Replacement{Path: "/route", HostSource: "source.domain", HostTarget: "target.domain"}
	incomingLink, _ := url.Parse("https://source.domain/route/0x1234")
	expectedResult := true

	actualResult := LinkMatches(incomingLink, domain)
	if actualResult != expectedResult {
		t.Errorf("Incorrect matching result\nExpected: %t\nGot:    %t", expectedResult, actualResult)
	}

	//
	domain = Replacement{Path: "/route", HostSource: "source.domain", HostTarget: "target.domain"}
	incomingLink, _ = url.Parse("https://anoher.domain/route/0x1234")
	expectedResult = false

	actualResult = LinkMatches(incomingLink, domain)
	if actualResult != expectedResult {
		t.Errorf("Incorrect matching result\nExpected: %t\nGot:    %t", expectedResult, actualResult)
	}

	//
	domain = Replacement{Path: "/route", HostSource: "source.domain", HostTarget: "target.domain"}
	incomingLink, _ = url.Parse("https://source.domain/something-else/0x1234")
	expectedResult = false

	actualResult = LinkMatches(incomingLink, domain)
	if actualResult != expectedResult {
		t.Errorf("Incorrect matching result\nExpected: %t\nGot:    %t", expectedResult, actualResult)
	}
}

func TestRewriteLink(t *testing.T) {
	// 1
	domain := Replacement{Path: "/", HostSource: "source.domain", HostTarget: "target.domain"}
	incomingLink, _ := url.Parse("https://source.domain/")
	expectedLink := "https://target.domain/"

	actualLink := RewriteLink(incomingLink, domain)
	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot:    %s", expectedLink, actualLink)
	}

	// 2
	domain = Replacement{Path: "/", HostSource: "source.domain", HostTarget: "target.domain"}
	incomingLink, _ = url.Parse("https://source.domain/0x1234")
	expectedLink = "https://target.domain/#/0x1234"

	actualLink = RewriteLink(incomingLink, domain)
	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot:    %s", expectedLink, actualLink)
	}

	// 3
	domain = Replacement{Path: "/route", HostSource: "source.domain", HostTarget: "target.domain"}
	incomingLink, _ = url.Parse("https://source.domain/route/0x1234")
	expectedLink = "https://target.domain/route#/0x1234"

	actualLink = RewriteLink(incomingLink, domain)
	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot:    %s", expectedLink, actualLink)
	}

	// 4
	domain = Replacement{Path: "/route", HostSource: "source.domain", HostTarget: "target.domain"}
	incomingLink, _ = url.Parse("https://source.domain/route/0x1234")
	expectedLink = "https://target.domain/route#/0x1234"

	actualLink = RewriteLink(incomingLink, domain)
	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot:    %s", expectedLink, actualLink)
	}

	// 5
	domain = Replacement{Path: "/route/here/", HostSource: "source.domain", HostTarget: "target.domain"}
	incomingLink, _ = url.Parse("https://source.domain/route/here/0x1234/0x5678")
	expectedLink = "https://target.domain/route/here/#/0x1234/0x5678"

	actualLink = RewriteLink(incomingLink, domain)
	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot:    %s", expectedLink, actualLink)
	}

	// 6
	domain = Replacement{Path: "/route/here", HostSource: "source.domain", HostTarget: "target.domain"}
	incomingLink, _ = url.Parse("https://source.domain/route/here/0x1234/0x5678")
	expectedLink = "https://target.domain/route/here#/0x1234/0x5678"

	actualLink = RewriteLink(incomingLink, domain)
	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot:    %s", expectedLink, actualLink)
	}

	// 7
	domain = Replacement{Path: "/route", HostSource: "source.domain:1234", HostTarget: "target.domain"}
	incomingLink, _ = url.Parse("https://source.domain:1234/route/0x1234/0x5678")
	expectedLink = "https://target.domain/route#/0x1234/0x5678"

	actualLink = RewriteLink(incomingLink, domain)
	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot:    %s", expectedLink, actualLink)
	}
}
