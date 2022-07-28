package main

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
)

func printJSON(result *Result) {
	cleanResult := ignoreFields(result, "Result", "Error")
	j, _ := json.MarshalIndent(cleanResult, "", "  ")
	fmt.Println(string(j))
}

func ignoreFields(result interface{}, fields ...string) map[string]interface{} {
	jsonResult, _ := json.Marshal(result)
	mapResult := map[string]interface{}{}
	json.Unmarshal([]byte(string(jsonResult)), &mapResult)

	for _, field := range fields {
		delete(mapResult, field)
	}

	return mapResult
}

func printRatings(result *Result) {
	imdb := color.New(color.Bold, color.FgYellow).PrintFunc()
	metacritic := color.New(color.Bold, color.FgCyan).PrintFunc()
	rt := color.New(color.Bold, color.FgRed).PrintFunc()

	if result.ImdbRating != "N/A" {
		imdb("IMDb")
		fmt.Println(": " + result.ImdbRating)
	}

	if result.Metascore != "N/A" {
		metacritic("Metascore")
		fmt.Println(": " + result.Metascore)
	}

	for _, rating := range result.Ratings {
		if rating.Source == "Rotten Tomatoes" {
			rt("Rotten Tomatoes")
			fmt.Println(": " + rating.Value)
		}
	}
}

func printItem(label string, text string) {
	if text != "N/A" {
		labelPrint := color.New(color.Bold).PrintFunc()
		labelPrint(label + ": ")
		fmt.Println(text)
	}
}

func printPretty(result *Result) {
	title := color.New(color.Bold, color.Underline).PrintlnFunc()

	title(result.Title)
	fmt.Println(result.Rated + " " + result.Type + " - " + result.Released)

	printRatings(result)

	fmt.Println("---")

	printItem("Plot", result.Plot)
	printItem("Genre", result.Genre)
	printItem("Runtime", result.Runtime)
	printItem("Director", result.Director)
	printItem("Writers", result.Writer)
	printItem("Actors", result.Actors)
	printItem("Country", result.Country)
	printItem("Languages", result.Language)
	printItem("IMDb URL", "https://www.imdb.com/title/"+result.ImdbID)
}
