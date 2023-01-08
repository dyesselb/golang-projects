package functions

func DefineAsciiSymbol() map[int]int {
	symbols := make(map[int]int)
	count := 2
	for i := 32; i < 127; i++ {
		symbols[i] = count
		count += 9
	}
	return symbols
}

func MakeArt(value string, data []byte) [][]string {
	symbolStartingPoint := DefineAsciiSymbol()
	asciiWordArr := [][]string{}
	for i := 0; i < len(value); i++ {
		subArr := []string{}
		count := 1
		word := ""

		if i != len(value)-1 && string(value[i]) == "\\" && string(value[i+1]) == "n" {
			subArr = append(subArr, "\n")
			i = i + 1
		} else if value[i] == 13 {
			subArr = append(subArr, "\n")
		} else if value[i] != 10 {
			for _, a_char := range data {
				if a_char == '\n' {
					count++
				}
				if count >= symbolStartingPoint[int(value[i])] && count < symbolStartingPoint[int(value[i])]+9 {
					if a_char == '\n' {
						subArr = append(subArr, word)
						word = ""
					} else {
						word += string(a_char)
					}
				}
			}
		}
		if len(subArr) != 0 {
			asciiWordArr = append(asciiWordArr, subArr)
		}
	}

	return asciiWordArr
}
