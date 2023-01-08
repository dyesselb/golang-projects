package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	text_file := os.Args[1]
	data, _ := os.ReadFile(text_file)
	var number_string string
	for i := 0; i < len(data); i++ {
		if data[i] > 47 && data[i] < 58 {
			number_string += string(data[i])
		} else if data[i] == 10 && len(data)-1 != i {
			number_string += " "
		}
	}
	strs := strings.Split(number_string, " ")
	ary := make([]float64, len(strs))
	for i := range ary {
		ary[i], _ = strconv.ParseFloat(strs[i], 64)
	}
	fmt.Println("Average:", int(math.Round(Average(ary))))
	fmt.Println("Median:", int(math.Round(Median(ary))))
	fmt.Println("Variance:", int(math.Round(Variance(ary, Average(ary)))))
	fmt.Println("Standard Deviation:", int(math.Round(Standard(ary, Average(ary)))))
}

func Average(arr []float64) float64 {
	var sum float64 = 0
	for i := range arr {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr))
}

func Median(s []float64) float64 {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	// Median
	if len(s)%2 == 1 {
		index := (len(s) / 2)
		return float64(s[index])
	} else {
		index := len(s) / 2
		sum := math.Round(float64(s[index])+float64(s[index-1])) / 2
		return float64(sum)
	}
}

func Standard(arr []float64, average float64) float64 {
	var result float64
	var super_average float64
	var total float64
	for i := 0; i < len(arr); i++ {
		total += float64(arr[i])
	}
	super_average = math.Round(float64(total) / float64(len(arr)))
	result = float64(Variance(arr, float64(super_average)))
	return math.Sqrt(float64(result))
}

func Variance(arr []float64, average float64) float64 {
	var variance float64
	for i := 0; i < len(arr); i++ {
		variance += float64((arr[i] - average) * (arr[i] - average))
	}
	return math.Round(variance / float64(len(arr)))
}
