package main

import (
	"dp_pattern/knapsack"
	"dp_pattern/max_sum"
	"dp_pattern/stdin"
)

func main() {
	switch stdin.StrStdin() {
	case "knapsack":
		knapsack.Knapsack()
	case "maxSum":
		max_sum.MaxSum()
	}
}
