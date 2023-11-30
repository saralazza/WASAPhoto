package api

import (
	"regexp"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

func CheckAuthentication(auth string, uid uint64) error {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	stringToken := re.FindAllString(auth, -1)
	if len(stringToken) == 0 {
		return database.ErrNotAuthorized
	}
	authuid, _ := strconv.Atoi(stringToken[0])

	if uint64(authuid) != uid {
		return database.ErrNotAuthorized
	}
	return nil
}
