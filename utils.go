package main

func average(ratings []int) int {
	var sum int
	for _, rating := range ratings {
		sum += rating
	}
	return sum / len(ratings)
}
