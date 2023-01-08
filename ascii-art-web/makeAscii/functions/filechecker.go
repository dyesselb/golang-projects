package functions

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func FileCheck(filename string) bool {
	fileHashes := make(map[string]string)
	fileHashes["standard"] = "ac85e83127e49ec42487f272d9b9db8b"
	fileHashes["shadow"] = "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	fileHashes["thinkertoy"] = "86d9947457f6a41a18cb98427e314ff8"
	data, err := os.ReadFile("makeAscii/formats/" + filename + ".txt")
	if err != nil {
		return false
	}
	if !(fileHashes[filename] == GetMD5Hash(string(data))) {
		return false
	}
	return true
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
