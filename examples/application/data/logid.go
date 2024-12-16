package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	version  = "02"
	alphabet = "1234567890abcdef"
)

func genLogID() string {

	timeStamp := time.Now().UnixNano() / int64(time.Millisecond)
	rand := randomHex(6, alphabet)
	rand2 := randomHex(32, alphabet)

	return fmt.Sprintf("%s%d%s%s", version, timeStamp, rand2, rand)
}

func randomHex(n int, alphabet string) string {
	var sb strings.Builder
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(len(alphabet))])
	}

	return sb.String()
}
