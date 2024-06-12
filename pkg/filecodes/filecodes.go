package filecodes

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var FileMap = map[string]string{}

func createCode() string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func AddFileCode(filename string) string {
	code := createCode()
	FileMap[code] = filename
	return code
}

func GetFile(code string) string {
	return FileMap[code]
}
