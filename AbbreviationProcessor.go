package main

import (
	"fmt"
	"regexp"
	"strings"
)

// AbbreviationProcessor :- Class to process text and extract possible abbrevation
type AbbreviationProcessor struct {
	regex         *regexp.Regexp
	isCaps        *regexp.Regexp
	stopWordRegex string
}

// Init :- Initializes the values
func (abbrProc *AbbreviationProcessor) Init() {
	abbrProc.regex, _ = regexp.Compile("\\(\\s*[a-zA-Z]+\\s*\\)")
	abbrProc.isCaps = regexp.MustCompile(`^[A-Z]+$`)
	abbrProc.stopWordRegex = "((the|of|for|and|in|on|de)\\s+)*"
}

func cleanText(text string) string {
	text = strings.Replace(text, " 's", "'s", -1)
	return text
}

// ProcessText :- Process a sentence and extract short form and its full form
func (abbrProc *AbbreviationProcessor) ProcessText(text string) map[string]string {
	text = cleanText(text)
	matches := abbrProc.regex.FindAllStringIndex(text, -1)
	startOffset := 0
	abbrevations := make(map[string]string)
	for _, match := range matches {
		shortForm := text[match[0]+1 : match[1]-1]
		shortForm = strings.TrimSpace(shortForm)
		if !abbrProc.isValidShortForm(shortForm) {
			continue
		}
		text := text[startOffset:match[0]]
		fullFormText, err := abbrProc.extractFullForm(text, shortForm)
		if err != nil {
			fmt.Println(err)
		}
		if len(fullFormText) == 0 {
			continue
		}
		// if abbrProc.isCaps.MatchString(shortForm) {
		abbrevations[shortForm] = text
		// }
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
func (abbrProc *AbbreviationProcessor) extractFullForm(text string, shortForm string) (string, error) {
	var fullForm string
	regexString := abbrProc.buildRegexForFullForm(shortForm)
	regex, err := regexp.Compile(regexString)
	if err != nil {
		return "", err
	}
	fullForm = strings.TrimSpace(regex.FindString(text))
	return fullForm, nil
}

func (abbrProc *AbbreviationProcessor) buildRegexForFullForm(shortForm string) string {
	numOfWords := len(shortForm)
	regex := "(?i)"
	for i := 0; i < numOfWords; i++ {
		firstCharacter := string(shortForm[i])
		regex = regex + firstCharacter + "[a-zA-Z']+[\\s,-]+" + abbrProc.stopWordRegex
	}
	regex = regex + "$"
	//fmt.Println("Regex -> ", regex)
	return regex
}
