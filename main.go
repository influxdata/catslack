package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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
	resp, err := http.Post(slackURL, "application/json", bytes.NewBuffer(p))

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
	// If no flag was given, use the $CATSLACK_URL env var.
	if slackURL == "" {
		slackURL = os.Getenv("CATSLACK_URL")
	}
}

func main() {

	// Exit if not passed a url
	if slackURL == "" {
		fmt.Println("Must pass a working Slack integration URL using the `-url` flag")
		fmt.Println("or by setting the $CATSLACK_URL env variable.")
		os.Exit(1)
	}

	// Start reading `stdin`
	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println("Error reading stdin")
	} else {
		postStdinToSlack(string(bytes))
	}
}
