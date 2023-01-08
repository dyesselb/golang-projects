package functions

func Write(arr [][]string) string{
	str := ""
	startP := 0
	for index, value := range arr {
		if len(value) == 1 && startP == index {
			str += "\n"
			startP += 1
		} else if len(value) == 1 {
			printer(arr, startP, index, &str, false)
			startP = index + 1
		} else if index+1 == len(arr) {
			printer(arr, startP, index+1, &str, true)
		}
	}
	return str[:len(str)-1]
}

func printer(array [][]string, startP int, endP int, str *string, last  bool ) {
	for i := 1; i <= 8; i++ {
		for j := startP; j < endP; j++ {
			*str += array[j][i]
		}
		*str += "\n"
	}
}
