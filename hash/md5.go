package hash

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func FileMD5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.New(
			fmt.Sprintf("File open err:%v", err))
	}
	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", errors.New(fmt.Sprintf("copy error%v", err))
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func StringMD5(filename string) (string, error) {
	md6 := md5.New()
	md6.Write([]byte(filename))
	return hex.EncodeToString(md6.Sum(nil)), nil

}
