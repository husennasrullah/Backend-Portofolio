package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func StructToJSON(input interface{}) (output string) {
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		output = ""
		return
	}
	output = string(b)
	return
}

func ConvertStrToInt(number string, defaultResult int) int {
	result, err := strconv.Atoi(number)
	if err != nil {
		return defaultResult
	}
	return result
}

func GetSubID(subscriptionID string) string {
	parts := strings.Split(subscriptionID, "-")
	if len(parts) > 1 {
		parts = parts[1:]
	}
	return strings.Join(parts, "-")
}

func TimeToUnix(t time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}

func EncryptSHA1(text string) string {
	sha := sha1.New() //nolint:gosec
	sha.Write([]byte(text))
	encrypted := sha.Sum(nil)
	encryptedString := hex.EncodeToString(encrypted)
	return encryptedString
}
