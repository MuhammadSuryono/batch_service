package utilscustome

import "os"

type IUtils interface {
	fileExists(filename string) bool
}

type UtilsHandle struct {
}

func UtilsCustome() IUtils {
	return &UtilsHandle{}
}

func (*UtilsHandle) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
