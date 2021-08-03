package tquotes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func main() {
	// Since the server gives error frequently
	// we try calling api again specific amount of time until we get value
	// res is to check if we have received value from api "res = 1"
	// count is to keep track of number of attempts to call api
	res, count := 0, 0
	// loop will run until we get proper quote from api till 15 attempts
	for res == 0 {
		quotes := Quote() // quote function
		count += 1        //counting attempts
		// Check if received api value is not empty
		// and
		// Check if OS is Linux, Since i have applied text styling for linux terminals only
		if quotes["quoteText"] != "" && runtime.GOOS == "linux" {
			// Output for Linux
			res = 1 // make loop condition false
			print("\n\033[1m\033[36m" + quotes["quoteText"] + "✨\033[0m\n")
			if quotes["quoteAuthor"] != "" { // If author is not empty
				print(" - " + "\033[93m" + quotes["quoteAuthor"] + "\033[0m\n")
			} else { // author is found empty then Anonymous
				print(" - " + "\033[93m" + "Anonymous" + "\033[0m\n")
			}
		} else if quotes["quoteText"] != "" {
			// Output for OS other than linux
			res = 1
			print("\n" + quotes["quoteText"] + "✨\n")
			if quotes["quoteAuthor"] != "" {
				print(" - " + quotes["quoteAuthor"] + "\n")
			} else {
				print(" - " + "Anonymous\n")
			}
		}
		if count > 15 { // Terminate loop if retry attempts exceed 15 times
			res = 1 // make loop condition false
			print("\n Something went wrong. Try again! \n")
		}
	}
}

// Fetch quotes from API and returns JSON like slices
func Quote() map[string]string {
	var client = &http.Client{}

	var data = strings.NewReader(`method=getQuote&format=json&param=ms&lang=en`)

	req, err := http.NewRequest("POST", "http://forismatic.com/api/1.0/", data)

	if err != nil {
		fatal()
	}

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36 Edg/92.0.902.62")
	req.Header.Set("DNT", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "http://forismatic.com")
	req.Header.Set("Referer", "http://forismatic.com/en/")
	req.Header.Set("Accept-Language", "en,hi;q=0.5")
	req.Header.Set("sec-gpc", "1")

	resp, err := client.Do(req)
	if err != nil {
		fatal()
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fatal()
	}
	values := make(map[string]string)

	err = json.Unmarshal(bodyText, &values)
	if err != nil {
		fatal()
	}

	return values
}

func fatal() {
	count := 0
	for count > 3 {
		main()
		count++
		if count > 3 {
			print("\n Something went wrong. Try again! \n")
		}
	}
}
