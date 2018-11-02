package integration_test

import (
	"fmt"
	"math/rand"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func pseudoRandString() string {

	rand.Seed(time.Now().UnixNano())

	prefix := make([]byte, 5)
	for i := range prefix {
		prefix[i] = chars[rand.Intn(len(chars))]
	}

	return fmt.Sprintf("%s-%d", string(chars), time.Now().Unix())
}
