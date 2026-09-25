// emails[i] consist of lowercase English letters, '+', '.' and '@'.
// xactly one '@' character

// Domain names end with the ".com" suffix.

// . in local name is ignored
// + in local name, ignore after +

// first:
// separate email by local and domain name at @

// domain name stay as is, only local needs clean up

// go through each local char
// - ignore and remove . when see
// - ignore right side if see + and stop

// reconstruct email

// if already exist, dont count, else count 

// return count

func numUniqueEmails(emails []string) int {
	counted := make(map[string]bool)

	for _, e := range emails {
		if counted[e] { continue }

		split := strings.Split(e, "@")
		local, domain := split[0], split[1]

		if plusIndex := strings.Index(local, "+"); plusIndex != -1 { local = local[:plusIndex] }

		local = strings.ReplaceAll(local, ".", "")

		cleaned := local + "@" + domain

		if counted[cleaned] { continue }

		counted[cleaned] = true
	}

	return len(counted)
}