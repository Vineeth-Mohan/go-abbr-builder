package main

import (
	"fmt"
	"regexp"
	"strings"
)

// AbbreviationProcessor :- Class to process text and extract possible abbrevation
type AbbreviationProcessor struct {
	regex *regexp.Regexp
}

// ProcessText :- Process a sentence and extract short form and its full form
func (abbrProc *AbbreviationProcessor) ProcessText(text string) map[string]string {
	abbrProc.regex, _ = regexp.Compile("\\(\\s*[a-zA-Z]+\\s*\\)")
	matches := abbrProc.regex.FindAllStringIndex(text, -1)
	startOffset := 0
	abbrevations := make(map[string]string)
	for _, match := range matches {
		shortForm := text[match[0]+1 : match[1]-1]
		shortForm = strings.TrimSpace(shortForm)
		if !abbrProc.isValidShortForm(shortForm) {
			continue
		}
		fullFormText := text[startOffset:match[0]]
		fullFormText, err := abbrProc.extractFullForm(fullFormText, len(shortForm))
		if err != nil {
			fmt.Println(err)
		}
		if !abbrProc.checkIfValidMatch(shortForm, fullFormText) {
			continue
		}
		abbrevations[shortForm] = fullFormText
		startOffset = match[1]
	}
	return abbrevations
}

// extractFullForm :- Function to extract full from from end
func (abbrProc *AbbreviationProcessor) isValidShortForm(shortForm string) bool {
	if len(shortForm) <= 1 {
		return false
	}
	return true
}

// extractFullForm :- Function to extract full from from end
func (abbrProc *AbbreviationProcessor) extractFullForm(text string, numOfWords int) (string, error) {
	var fullForm string
	regexString := abbrProc.buildRegexForFullForm(numOfWords)
	regex, err := regexp.Compile(regexString)
	if err != nil {
		return "", err
	}
	fullForm = regex.FindString(text)
	return fullForm, nil
}

func (abbrProc *AbbreviationProcessor) buildRegexForFullForm(numOfWords int) string {
	stopWordRegex := "((or|the|of|for|and)\\s+)*"
	regex := strings.Repeat("[a-zA-Z]+\\s+"+stopWordRegex, numOfWords-1)
	regex = regex + "[a-zA-Z]+\\s*$"
	regex = "(?i)" + regex
	return regex
}
func (abbrProc *AbbreviationProcessor) checkIfValidMatch(shortForm string, fullForm string) bool {
	stopWordRegex := "((of|for|and)\\s+)*"
	fmt.Println(shortForm, fullForm)
	numOfWords := len(shortForm)
	regex := "(?i)"
	for i := 0; i < numOfWords; i++ {
		firstCharacter := string(shortForm[i])
		regex = regex + firstCharacter + "[a-zA-Z]+\\s+" + stopWordRegex
	}
	fmt.Println("Regex ->", regex)
	isMatch, _ := regexp.MatchString(regex, fullForm)
	return isMatch
}
