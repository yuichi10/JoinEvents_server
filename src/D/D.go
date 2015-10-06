package D

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

const (
	TableUsers int = iota
)

const (
	//the charactors to add password before md5
	HashKey string = "96A5SFDBgK3B5Jz32nWE"
)

func GetMd5Hash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text + HashKey))
	return hex.EncodeToString(hash.Sum(nil))
}

//check that is match policy
//The id allowed to use alphabet, capital letter and number
//the number of chars is 8 <= id <= 16
func CheckAlphanumeric(text string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z0-9]{8,16}$", text); !m {
		return false
	} else {
		return true
	}
}
