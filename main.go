package main

import (
	"github.com/fatih/color"
	"os"
	"strings"
)

func main() {
	// read parameters
	domain, wordlist, output := readParameters()
	// fuzzer
	Fuzzer(makeUrl(domain), wordlist, output)
}

// function read these parameters from cl -u https://google.com -w wordlist.txt -o output.txt
func readParameters() (string, string, string) {
	var url, wordlist, output string
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-u":
			url = os.Args[i+1]
			i++
		case "-w":
			wordlist = os.Args[i+1]
			i++
		case "-o":
			output = os.Args[i+1]
			i++
		}
	}
	// exit if url or wordlist is empty
	if url == "" || wordlist == "" {
		// print usage
		printUsage()
		os.Exit(0)
	}
	return url, wordlist, output
}

// function print usage
func printUsage() {
	help()
	// print usage
	color.Green("[!] Usage: %s -u https://google.com -w wordlist.txt -o output.txt", os.Args[0])
	color.Green("[!] Usage: %s -u google.com -w wordlist.txt -o output.txt", os.Args[0])
	// -u url is required yellow color
	color.Yellow("[!] -u url is required")
	// -w wordlist is required yellow color
	color.Yellow("[!] -w wordlist is required")
	// -o output is optional
	color.Yellow("[!] -o output is optional")
	// -h help
	color.Yellow("[!] -h help")
	// Thanks for using this tool
	color.Green("[+] Thanks for using this tool")
	// exit
	os.Exit(0)

}

// function that make url with domain start with https:// and end with /
func makeUrl(domain string) (url string) {
	url = domain
	// check if domain has https://
	if !strings.HasPrefix(domain, "https://") {
		url = "https://" + domain
	}
	// check if domain has /
	if !strings.HasSuffix(domain, "/") {
		url = url + "/"
	}
	return
}

func help() {
	h := `
.__           .__          
|  |__   ____ |  | ______  
|  |  \_/ __ \|  | \____ \ 
|   Y  \  ___/|  |_|  |_> >
|___|  /\___  >____/   __/ 
     \/     \/     |__|    
`
	color.Green(h)
}
