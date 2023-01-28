package fullname_parser

import (
	"log"
	"regexp"
	"strings"
)

type ParsedName struct {
	First string
	Last  string
}

func ParseFullname(fullname string) (parsedName ParsedName) {
	nameParts := []string{}
	//split name to parts and store commas
	nameParts = splitName(fullname, nameParts)

	// First name: remove and store first part as first name
	if len(nameParts) > 0 {
		parsedName.First, nameParts = findFirstname(nameParts)
	}
	// Last name: remove and store last name
	if len(nameParts) > 0 {
		log.Print("find last name")
		parsedName.Last = findLastname(nameParts)
	}

	return
}

func splitName(fullname string, nameParts []string) []string {
	re := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	fullname = re.ReplaceAllLiteralString(fullname, " ")
	nameParts = strings.Split(strings.TrimSpace(fullname), " ")
	log.Print(nameParts)
	return nameParts
}

func findLastname(nameParts []string) (lastname string) {
	lastname = strings.Join(nameParts, " ")
	return
}

func findFirstname(nameParts []string) (string, []string) {
	firstname := nameParts[0]
	nameParts = nameParts[1:]
	return firstname, nameParts
}
