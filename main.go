package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {

	hostname := flag.String("u", "", "A hostname to replace the URLs with")
	outputFile := flag.String("o", "output.txt", "Path to the output file (optional)")

	flag.Parse()

	if *hostname == "" {
		log.Fatal("Error: Please provide a hostname using the -u flag.")
	}

	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parsedURL, err := url.Parse(line)
		if err != nil {
			log.Printf("Error parsing URL: %v. Skipping line.", err)
			continue
		}

		newHostURL, err := url.Parse(*hostname)
		if err != nil {
			log.Fatalf("Error parsing provided hostname: %v", err)
		}

		parsedURL.Scheme = newHostURL.Scheme
		parsedURL.Host = newHostURL.Host

		lines = append(lines, parsedURL.String())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}

	updatedContent := strings.Join(lines, "\n")

	if *outputFile != "" {
		err := os.WriteFile(*outputFile, []byte(updatedContent), 0644)
		if err != nil {
			log.Fatalf("Error writing to output file: %v", err)
		}
	} else {
		fmt.Println(updatedContent)
	}

	fmt.Println("File updated successfully")
}
