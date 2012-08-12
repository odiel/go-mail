package email

import (
	"regexp"
	"strings"
)

var emailRegexp *regexp.Regexp = regexp.MustCompile(`[a-zA-Z0-9\-+~_%]+[a-zA-Z0-9\-+~_%.]*@([a-z]+[a-z0-9\\-]*\.)+[a-z][a-z]+`)

func NormalizeAddressLowercase(address string) string {
	return strings.ToLower(strings.TrimSpace(address))
}

func CompareAddressesCaseinsensitive(addressA, addressB string) bool {
	return NormalizeAddressLowercase(addressA) == NormalizeAddressLowercase(addressB)
}

// ValidateAddress uses a simplified regular expression for checking
// email addresses. Please report real world addresses that are not working as bug.
func ValidateAddress(address string) (normalizedAddress string, err error) {
	normalizedAddress = strings.TrimSpace(address)
	valid := emailRegexp.Match([]byte(normalizedAddress))
	if !valid {
		return "", ErrInvalidEmailAddress(address)
	}
	return normalizedAddress, nil
}

type ErrInvalidEmailAddress string

func (self ErrInvalidEmailAddress) Error() string {
	return "Invalid email address: " + string(self)
}
