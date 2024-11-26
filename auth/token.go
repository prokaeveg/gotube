package auth

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func CreateTokenForUser(userId int) (string, string) {
	salt := getSalt(20)
	token := getMD5Hash(salt + getMD5Hash(salt+strconv.Itoa(userId)))

	salt = getSalt(20)
	refreshToken := getMD5Hash(salt + getMD5Hash(salt+strconv.Itoa(userId)))

	return token, refreshToken
}

func getSalt(len int) string {
	b := make([]byte, len)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len)]
	}

	return string(b)
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
