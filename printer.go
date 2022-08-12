package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func printJSON(result *Result, avoidSpoilers bool) {
	cleanResult := ignoreFields(result, "Result", "Error", "Response")
	if avoidSpoilers {
		cleanResult = ignoreFields(cleanResult, "Plot", "Genre", "Runtime", "Actors", "Country", "Language")
	}
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
	metacritic := color.New(color.Bold, color.FgGreen).PrintFunc()
	rt := color.New(color.Bold, color.FgRed).PrintFunc()
	snob := color.New(color.Bold, color.FgCyan).PrintFunc()

	ratings := []int{}

	if result.ImdbRating != "N/A" {
		imdb("IMDb")
		fmt.Println(": " + result.ImdbRating)
		rating, _ := strconv.ParseFloat(result.ImdbRating, 64)
		ratings = append(ratings, int(rating*10))
	}

	if result.Metascore != "N/A" {
		metacritic("Metascore")
		fmt.Println(": " + result.Metascore)
		rating, _ := strconv.Atoi(result.Metascore)
		ratings = append(ratings, rating)
	}

	for _, rating := range result.Ratings {
		if rating.Source == "Rotten Tomatoes" {
			rt("Rotten Tomatoes")
			fmt.Println(": " + rating.Value)
			rating, _ := strconv.Atoi(strings.TrimSuffix(rating.Value, "%"))
			ratings = append(ratings, rating)
		}
	}

	if len(ratings) > 0 {
		avg := average(ratings)
		snob("Snob Consensus")
		printSnobConsensus(avg)
	}
}

func printSnobConsensus(avg int) {
	print(": ")
	label := "Bad - Most likely not worth your time"
	labelColor := color.FgRed

	if avg >= 60 {
		label = "Average - Might be worth watching"
		labelColor = color.FgMagenta
	}

	if avg >= 70 {
		label = "Good - Worth watching"
		labelColor = color.FgBlue
	}

	if avg >= 80 {
		label = "Great - You should watch this"
		labelColor = color.FgGreen
	}

	if avg >= 85 {
		label = "Excellent - Must watch"
		labelColor = color.FgYellow
	}

	color.New(color.Bold, labelColor).Print(label)
	fmt.Printf(" (%d%%)\n", avg)
}

func printItem(label string, text string) {
	if text != "N/A" {
		labelPrint := color.New(color.Bold).PrintFunc()
		labelPrint(label + ": ")
		fmt.Println(text)
	}
}

func printPretty(result *Result, avoidSpoilers bool) {
	title := color.New(color.Bold, color.Underline).PrintlnFunc()

	title(result.Title)
	fmt.Println(result.Rated + " " + result.Type + " - " + result.Released)

	printRatings(result)

	fmt.Println("---")

	if !avoidSpoilers {
		printItem("Plot", result.Plot)
		printItem("Genre", result.Genre)
		printItem("Runtime", result.Runtime)
	}
	printItem("Director", result.Director)
	printItem("Writers", result.Writer)
	if !avoidSpoilers {
		printItem("Actors", result.Actors)
		printItem("Country", result.Country)
		printItem("Languages", result.Language)
	}
	printItem("IMDb URL", "https://www.imdb.com/title/"+result.ImdbID)
}
