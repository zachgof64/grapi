package utils

import (
	"io/ioutil"
	"os"
)

//GetByte will Return []byte version of the filepath given
func GetByte(filePath string) []byte {
	file, fileErr := os.Open(filePath)
	Check(fileErr)
	b, bErr := ioutil.ReadAll(file)
	Check(bErr)
	file.Close()
	return b
}