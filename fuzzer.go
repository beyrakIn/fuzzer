package main

import (
	"bufio"
	"github.com/fatih/color"
	"net/http"
	"os"
)

var (
	//wg         = sync.WaitGroup{}
	outputFile *os.File
)

// Fuzzer is a function that fuzzes the given domain with the given wordlist
func Fuzzer(domain, wordlist, output string) {
	// open wordlist
	file, err := os.Open(wordlist)
	if err != nil {
		color.Red("[!] Error: %s", err)
		os.Exit(0)
	}
	// close file
	defer closeFile(file)

	// create and open output file if it is not nil
	if output != "" {
		// create file
		createFile(output)
		// clear file
		clearFile(output)
		// open file
		outputFile = openFile(output)
	}

	// read wordlist
	scanner := bufio.NewScanner(file)
	// loop through wordlist
	for scanner.Scan() {
		//wg.Add(1)

		word := scanner.Text()

		func(word string) {
			url := domain + word

			// check if url is valid
			if isValid(url) {
				color.Green("[+] %s", url)

				// write url to output file
				if output != "" {
					_, err := writeToFile(url, outputFile)
					if err != nil {
						color.Red("[!] Error: %s", err)
						os.Exit(0)
					}
				}
			}
			//wg.Done()
		}(word)
	}
	//wg.Wait()

	// close output file
	if output != "" {
		f, err := writeToFile("", outputFile)
		if err != nil {
			color.Red("[!] Error: %s", err)
			os.Exit(0)
		}
		// close file
		closeFile(f)
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
		os.Exit(0)
	}
	// set request header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	// create response
	resp, err := client.Do(req)
	if err != nil {
		// print error
		color.Red("[!] Error: %s", err)
		os.Exit(0)
	}
	// check if response status code is 200
	if resp.StatusCode == 200 {
		return true
	}

	return false
}

// writeToFile is a function that writes the given line to the given file
func writeToFile(line string, file *os.File) (*os.File, error) {
	// append line to file
	_, err := file.WriteString(line + "\n")
	if err != nil {
		return nil, err
	}

	return file, nil
}

// openFile is a function that opens the given file and returns a pointer to the file
func openFile(file string) *os.File {
	// open file
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// print error
		color.Red("[!] Error: %s", err)
		os.Exit(0)
	}

	return f
}

// closeFile is a function that closes the given file and handles errors
func closeFile(file *os.File) {
	// close file
	err := file.Close()
	if err != nil {
		color.Red("[!] Error: %s", err)
	}
}

// clearFile is a function that clears the given file
func clearFile(file string) {
	// open file
	f, err := os.OpenFile(file, os.O_RDWR, 0644)
	if err != nil {
		color.Red("[!] Error: %s", err)
		os.Exit(0)
	}
	// clear file
	f.Truncate(0)
	closeFile(f)
}

func createFile(file string) {
	// create file
	f, err := os.Create(file)
	if err != nil {
		color.Red("[!] Error: %s", err)
		os.Exit(0)
	}
	closeFile(f)
}
