package util

import "math/rand"

var Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int, seed int64) string {
	r := rand.New(rand.NewSource(seed))
	b := make([]rune, n)
	for i := range b {
		b[i] = Letters[r.Intn(len(Letters))]
	}
	return string(b)
}
