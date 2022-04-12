package common

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestAesDecrypt(t *testing.T) {
	key := "44a7e2d6a8fdbaae"

	file, err := os.Open("../testData/1.ts")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	decryptData, _ := AesDecrypt(content, key)

	ioutil.WriteFile("../testData/test.mp4", decryptData, 0644)
}
