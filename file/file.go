package file

import (
	elog "github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
)

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Write(path string, data []byte) bool {
	if err := ioutil.WriteFile(path, data, 0600); err != nil {
		elog.Errorf("%v", err.Error())
		return false
	}
	return true
}

func Read(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return content
}

func Delete(path string) bool {
	if err := os.Remove(path); err != nil {
		elog.Errorf("%v", err.Error())
		return false
	}
	return true
}
