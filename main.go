package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := flag.String("text", "./wiki_gist.txt", "File to parse for Abbrevations")
	debug := flag.Bool("Debug", false, "Enable Debug mode")
	flag.Parse()

	processor := AbbreviationProcessor{}

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lineNumber := 0
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		lineNumber = lineNumber + 1
		if err == io.EOF {
			break
		}
		abbrevations := processor.ProcessText(string(line))
		for shortForm, fullFormText := range abbrevations {
			if *debug {
				fmt.Println(shortForm+"(", lineNumber, ") -> "+fullFormText)
			}
			if !*debug {
				fmt.Println(shortForm + "," + fullFormText)
			}
		}
	}

}
