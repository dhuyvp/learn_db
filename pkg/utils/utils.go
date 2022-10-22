package utils

import "math/rand"

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"error_message"`
	Data       interface{} `json:"data"`
}

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	result := make([]byte, n)

	LengthLetter := len(letterBytes)

	for i := range result {
		result[i] = letterBytes[rand.Intn(LengthLetter)%(LengthLetter-10)+10]
	}

	return string(result)
}

func RandInt(n int) string {
	result := make([]byte, n)

	LengthLetter := len(letterBytes)

	for i := range result {
		result[i] = letterBytes[rand.Intn(LengthLetter)%10]
	}

	return string(result)
}
