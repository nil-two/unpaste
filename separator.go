package main

func toDelimiters(list string) []string {
	var a []string
	for _, ch := range list {
		a = append(a, string(ch))
	}
	if len(a) == 0 {
		return []string{""}
	}
	return a
}
