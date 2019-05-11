package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestbuildRegexForFullForm(t *testing.T) {
	stopWordRegex := "(of|for|and)*"
	match, _ := regexp.MatchString(stopWordRegex, "")
	fmt.Println("Vm")
	fmt.Println(match)

}
