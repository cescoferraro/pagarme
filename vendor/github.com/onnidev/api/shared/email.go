package shared

import "strings"

// NormalizeEmail TODO: NEEDS COMMENT INFO
func NormalizeEmail(email string) string {
	mail := strings.Replace(email, " ", "", -1)
	mail = strings.Replace(mail, "\n", "", -1)
	mail = strings.ToLower(mail)
	return mail
}
