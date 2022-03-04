package file

import (
	"io/ioutil"
	"os"
)

//Exists check if file exists
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//Read filepath
func Read(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return content
}
