package main

func removeDuplicated(s []string) []string {
	if len(s) < 2 {
		return s
	}

	if s[0] == s[1] {
		if len(s) > 2 {
			return removeDuplicated(append(s[:1], removeDuplicated(s[2:])...))
		}
		return removeDuplicated(s[:1])
	}
	return append(s[:1], removeDuplicated(s[1:])...)
}
