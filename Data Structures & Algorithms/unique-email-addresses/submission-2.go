func numUniqueEmails(emails []string) int {
	counted := make(map[string]bool)

	for _, e := range emails {
		if counted[e] { continue }

		split := strings.Split(e, "@")
		local, domain := split[0], split[1]

		local = strings.Split(local, "+")[0]
		local = strings.ReplaceAll(local, ".", "")

		cleaned := local + "@" + domain

		if counted[cleaned] { continue }

		counted[cleaned] = true
	}

	return len(counted)
}