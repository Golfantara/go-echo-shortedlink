package helpers

import "math/rand"

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type RandomURLInterface interface {
	RandomURL(size int) string
}

type randomURL struct{}

func NewGenerateRandomURL() RandomURLInterface{
	return &randomURL{}
}


func(r randomURL) RandomURL(size int) string {
	str := make([]rune, size)
	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}
	return string(str)
}