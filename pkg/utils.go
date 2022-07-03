package pkg

import "fmt"

func shortenSentence(rawSent string, maxLength int) string {
	if len(rawSent) < maxLength {
		return rawSent
	}
	return fmt.Sprintf("%s..", rawSent[:maxLength])
}
