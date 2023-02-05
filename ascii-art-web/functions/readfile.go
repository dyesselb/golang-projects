package functions

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"
)

func MakeMapOfBanner(bannerName string) (map[rune][]string, error) {
	hashSum := []string{"86d9947457f6a41a18cb98427e314ff8", "a49d5fcb0d5c59b2e77674aa3ab8bbb1", "ac85e83127e49ec42487f272d9b9db8b"}

	banner := map[rune][]string{}
	banFileName := "./banners/" + bannerName + ".txt"
	hash, err := CalculateHashsum(banFileName)
	if err != nil {
		return nil, err
	}
	if !CheckHash(hash, hashSum) {
		return nil, &http.ProtocolError{}
	}
	openedfile, err := os.OpenFile(banFileName, os.O_RDONLY, 0)
	if err != nil {
		return banner, err
	}
	fileScanner := bufio.NewScanner(openedfile)
	fileScanner.Split(bufio.ScanLines)
	character := []string{}
	key := rune(31)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			banner[key] = character
			character = []string{}
			key++
		}
		character = append(character, line)
	}
	banner[key] = character
	delete(banner, rune(31))
	openedfile.Close()
	return banner, nil
}

func CalculateHashsum(banFileName string) (string, error) {
	f, err := os.Open(banFileName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	hasher := md5.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func CheckHash(hash string, hashSum []string) bool {
	for _, v := range hashSum {
		if v == hash {
			return true
		}
	}
	return false
}
