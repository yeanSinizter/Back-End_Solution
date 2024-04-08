package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// Struct to hold meat counts
type MeatSummary struct {
	Beef map[string]int `json:"beef"`
}

func main() {
	http.HandleFunc("/beef/summary", beefSummaryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func beefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch text from URL
	text, err := fetchTextFromURL("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse and count meat types
	meats := countMeatTypes(text)

	// Create MeatSummary struct
	summary := MeatSummary{Beef: meats}

	// Marshal summary to JSON
	responseJSON, err := json.Marshal(summary)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set JSON response header and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func fetchTextFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func countMeatTypes(text string) map[string]int {
	meats := make(map[string]int)
	// Regular expression to match meat names
	re := regexp.MustCompile(`\b(?:fatback|t-bone|pastrami|pork|meatloaf|jowl|enim|bresaola)\b`)
	// Find all matches in text
	matches := re.FindAllString(text, -1)
	// Count occurrences of each meat
	for _, match := range matches {
		meats[match]++
	}
	return meats
}
