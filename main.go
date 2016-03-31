package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var slackURL string

type SlackPostJSON struct {
	Text string `json:"text"`
}

// Makes userfacing output pretty!
func formatString(text string) string {
	return fmt.Sprintf("```%v```", text)
}

// Posts a string to the slackURL
func postStdinToSlack(text string) {

	// Make the datastructure with the text already there
	slackJSON := SlackPostJSON{Text: formatString(text)}

	// Marshall the JSON
	p, err := json.Marshal(&slackJSON)

	// Check the error
	if err != nil {
		fmt.Println("Failed to marshal JSON while posting to slack")
	}

	// Make the post request
	resp, err := http.Post(slackURL, "applicaion/json", bytes.NewBuffer(p))

	// Make sure the URL is right!
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request returned a non-200 status code. Check your Slack URL...")
	}

	// Check the error
	if err != nil {
		fmt.Printf("Error making post request to Slack: %v\n", err)
	}
}

func init() {
	// Set up the -url flag
	flag.StringVar(&slackURL, "url", "", "Pass the full url from the slack integration page")

	// Parse all flags
	flag.Parse()
	if slackURL == "" {
		slackURL = os.Getenv("URL")
	}
}

func main() {

	// Exit if not passed a url
	if slackURL == "" {
		fmt.Println("Must pass a working Slack integration URL using the `-url` flag.")
		return
	}

	// Start reading `stdin`
	scanner := bufio.NewScanner(os.Stdin)

	// Initialize a string to hold the output
	var t string

	// Start scanning through input
	for scanner.Scan() {
		if scanner.Text() != "EOF" {
			// Add new lines coming in to the output string
			t += fmt.Sprintf("%v\n", scanner.Text())
		} else {
			// If a line comes in that is just `EOF` send the text and exit
			postStdinToSlack(t)
			break
		}
	}
}
