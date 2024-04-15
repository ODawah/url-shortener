package encoders

import (
	"crypto/md5"
	"fmt"
)

func MD5Encode(input string) (string, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(input))
	if err != nil {
		return "", err
	}
	encodedString := hash.Sum(nil)

	return fmt.Sprintf("%x", encodedString), nil
}
