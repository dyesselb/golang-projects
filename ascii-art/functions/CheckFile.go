package student

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func CheckFile(path string) bool {
	checkhash := "2f7f19c7bf266f9cb3ce72683550beec"
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))

	return hash == checkhash
}
