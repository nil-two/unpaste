package main

func toDelimiters(list string) []string {
	var a []string
	for _, ch := range list {
		a = append(a, string(ch))
	}
	return a
}
