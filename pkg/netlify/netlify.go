package netlify

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const acceptedMembersEnvVar = "ACCEPTED_MEMBERS"

var members Members

func init() {
	members = Members{}

	acceptedMembers := os.Getenv(acceptedMembersEnvVar)

	memberExtractor := regexp.MustCompile(`([a-z]+):(\d+)`)
	if !memberExtractor.MatchString(acceptedMembers) {
		return
	}

	memberMatches := memberExtractor.FindAllStringSubmatch(acceptedMembers, -1)
	for _, match := range memberMatches {
		chatID, err := strconv.Atoi(match[2])
		if err != nil {
			return
		}

		members[chatID] = match[1]
	}
}

type Members map[int]string

func (m Members) NameFromChat(chatID int) (string, error) {
	if name, found := m[chatID]; found {
		return name, nil
	}

	return "", fmt.Errorf("unrecognized chat ID %d", chatID)
}

func FetchMembers() Members {
	return members
}
