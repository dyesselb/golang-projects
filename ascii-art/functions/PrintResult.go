package student

import (
	"fmt"
	"os"
	"strings"
)

func Split(s string) []string {
	s = strings.ReplaceAll(s, "\n", "\\n")
	// s = strings.ReplaceAll(s, "\\!", "!")
	str := strings.Split(s, "\\n")
	return str
}

func CheckOnlyNewLines(str string) (bool, int) {
	str = strings.ReplaceAll(str, "\\n", "\n")
	count := 0
	for _, s := range str {
		if s != '\n' {
			return false, 0
		} else {
			count++
		}
	}
	return true, count
}

func CreateArt(text string) {
	data, _ := os.ReadFile("standard.txt")
	lines := Split(text)
	banner := Split(string(data))
	bool1, count := CheckOnlyNewLines(text)

	if !Check(text) {
		fmt.Println("ERROR: You must enter only ASCII characters!")
		return
	}

	if bool1 {
		for i := 0; i < count; i++ {
			fmt.Println()
		}
		return
	}

	result := make([][]string, 8*len(lines))
	for k := range lines {
		for _, x := range lines[k] {
			for i := 8 * k; i < 8*(k+1); i++ {
				ind := (x-32)*8 + 1 + (x - 32 + rune(i-8*k) + 1)
				result[i] = append(result[i], banner[ind])

			}
		}
	}

	for i, k := range result {
		if i < len(result)-7 && k == nil && result[i+7] == nil {
			result = append(result[:i], result[i+7:]...)
			i = i + 8
		}
	}

	for i := range result {
		fmt.Println(strings.Join(result[i], ""))
	}
}

func Check(s string) bool {
	for _, l := range s {
		if rune(l) < 0 || rune(l) > 257 {
			return false
		}
	}
	return true
}
