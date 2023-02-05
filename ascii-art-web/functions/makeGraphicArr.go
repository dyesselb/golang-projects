package functions

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func MakeGraphicWord(charsMap map[rune][]string, word string) (string, error) {
	// word = strings.ReplaceAll(word, "\\n", "\n")
	word = strings.ReplaceAll(word, "\r\n", "\n")
	// word = strings.ReplaceAll(word, "\\t", "     ")
	// re_up := regexp.MustCompile(`(\\)([a-zA-Z0-9!?']+)`)
	// word = re_up.ReplaceAllString(word, "$2")

	arrOfword := strings.Split(word, "\n")
	graphics := ""
	count := 0
	for _, v := range arrOfword {
		if v == "" {
			count++
		}
	}
	if count == len(arrOfword) {
		for j := 0; j < count-1; j++ {
			graphics += "\n"
		}
	} else {
		for i := 0; i < len(arrOfword); i++ {
			if len(arrOfword[i]) != 0 {
				for index := 1; index <= 8; index++ {
					for _, char := range arrOfword[i] {
						if val, ok := charsMap[char]; ok {
							graphics += val[index]
						} else {
							return "", fmt.Errorf("\ncharacter in argument is not allowed\n")
						}
					}
					graphics += "\n"
				}
			} else {
				graphics += "\n"
			}
		}
	}

	return graphics, nil
}

func PrintGraphics(charsMap map[rune][]string, word string) {
	arrOfword := strings.Split(word, "\n")
	count := 0
	for _, v := range arrOfword {
		if v == "" {
			count++
		}
	}
	if count == len(arrOfword) {
		for j := 0; j < count-1; j++ {
			fmt.Println()
		}
	} else {
		for i := 0; i < len(arrOfword); i++ {
			if len(arrOfword[i]) != 0 {
				for index := 1; index <= 8; index++ {
					for _, char := range arrOfword[i] {
						if val, ok := charsMap[char]; ok {
							fmt.Print(val[index])
						} else {
							log.Println("\ncharacter in argument is not allowed")
							os.Exit(0)
						}
					}
					fmt.Println()
				}
			} else {
				fmt.Println()
			}
		}
	}
}
