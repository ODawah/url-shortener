package encoders

import "github.com/eknkc/basex"

const base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Base62Encode(input string) (string, error) {
	encoder, err := basex.NewEncoding(base62Alphabet)
	if err != nil {
		return "", err
	}
	encoded := encoder.Encode([]byte(input))
	if err != nil {
		return "", err
	}
	return encoded, nil
}
