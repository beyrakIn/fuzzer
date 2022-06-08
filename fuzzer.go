package main

import (
	"bufio"
	"github.com/fatih/color"
	"net/http"
	"os"
)

// Fuzzer is a function that fuzzes the given domain with the given wordlist
func Fuzzer(domain, wordlist, output string) {
	// open wordlist
	file, err := os.Open(wordlist)
	if err != nil {
		// print error
		color.Red("[!] Error: %s", err)
		// exit
		os.Exit(0)
	}
	// close file
	defer file.Close()

	if output != "" {
		clearFile(output)
	}

	// read wordlist
	scanner := bufio.NewScanner(file)
	// loop through wordlist
	for scanner.Scan() {
		// get word
		word := scanner.Text()
		// create url
		url := domain + word
		// check if url is valid
		if isValid(url) {
			// print valid
			color.Green("[+] %s", url)
			// write url to output file

			if output != "" {
				writeToFile(url, output)
			}
		}
	}
}

// function check if url is valid
func isValid(url string) bool {
	// create client
	client := &http.Client{}
	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// print error
		color.Red("[!] Error: %s", err)
		// exit
		os.Exit(0)
	}
	// set request header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	// create response
	resp, err := client.Do(req)
	if err != nil {
		// print error
		color.Red("[!] Error: %s", err)
		// exit
		os.Exit(0)
	}
	// check if response status code is 200
	if resp.StatusCode == 200 {
		// return true
		return true
	}
	// return false
	return false
}

// writeToFile is a function that writes the given url to the given file
func writeToFile(url, output string) {
	// open file
	file, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// print error
		color.Red("[!] Error: %s", err)
		// exit
		os.Exit(0)
	}
	// write url to file
	_, err = file.WriteString(url + "\n")
	if err != nil {
		// print error
		color.Red("[!] Error: %s", err)
		// exit
		os.Exit(0)
	}
	// close file
	defer file.Close()
}

// function clear file
func clearFile(file string) {
	// open file
	f, err := os.OpenFile(file, os.O_RDWR, 0644)
	if err != nil {
		// print error
		color.Red("[!] Error: %s", err)
		// exit
		os.Exit(0)
	}
	// clear file
	f.Truncate(0)
	// close file
	defer f.Close()
}