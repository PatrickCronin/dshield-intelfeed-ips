package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"
	"os"
)

const intelFeedURL string = "https://isc.sans.edu/api/intelfeed?json"

func main() {
	emailAddr := parseArgs()

	ips, err := getIntelFeedIPs(emailAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get intelfeed: %+v", err)
		os.Exit(-1)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func parseArgs() string {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintf(
			fs.Output(),
			"%s\n"+
				"    Retrieve and output the IPs on DShield API's Intelfeed endpoint\n"+
				"    (https://isc.sans.edu/api/#intelfeed).\n"+
				"\n"+
				"Usage: (-h|-email xxx@yyy.com)",
			os.Args[0],
		)
	}

	emailStr := fs.String("email", "", "DShield's API wants to know who's asking")
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(fs.Output(), "parse command line args: %+v", err)
		os.Exit(-1)
	}

	email, err := mail.ParseAddress(*emailStr)
	if err != nil {
		fmt.Fprintf(fs.Output(), "-email: must be a valid email address\n")
		os.Exit(-1)
	}

	return email.Address
}

type intelFeedEntry struct {
	IP string `json:"ip"`
}

func getIntelFeedIPs(emailAddr string) ([]string, error) {
	ctx := context.Background()
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", intelFeedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("User-Agent", emailAddr)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	entries := make([]intelFeedEntry, 0)
	if err := json.Unmarshal(body, &entries); err != nil {
		return nil, fmt.Errorf("parse response body as JSON: %w", err)
	}

	ips := make([]string, len(entries))
	for i, entry := range entries {
		ips[i] = entry.IP
	}

	return ips, nil
}
