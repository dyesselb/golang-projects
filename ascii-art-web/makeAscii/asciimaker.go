package askiied

import (
	"ascii-art-web/makeAscii/functions"
	"bufio"
	"os"
)

func Askiied(str string, version string) (res string, err int) {
	versions := []string{"shadow", "thinkertoy", "standard"}

	for _, value := range str {
		if !(value > 0 && value <= 255) {
			return "", 400
		}
	}

	for _, value := range versions {
		if value == version {
			validFile := functions.FileCheck(value)
			if !validFile {
				return "", 500
			}
			file, _ := os.Open("makeAscii/formats/" + value + ".txt")
			defer file.Close()
			reader := bufio.NewScanner(file)
			art := ""
			for reader.Scan() {
				art += reader.Text() + "\n"
			}
			arr := functions.MakeArt(str, []byte(art))
			str := functions.Write(arr)
			return str, 0
		}
	}
	return "", 400
}
