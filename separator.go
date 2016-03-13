package main

func toDelimiters(list string) []string {
	var isEscaping bool
	var a []string
	for _, ch := range list {
		if isEscaping {
			switch ch {
			case '\\':
				a = append(a, "\\")
			case 't':
				a = append(a, "\t")
			case 'n':
				a = append(a, "\n")
			case '0':
				a = append(a, "")
			default:
				a = append(a, string(ch))
			}
			isEscaping = false
			continue
		}
		if ch == '\\' {
			isEscaping = true
			continue
		}
		a = append(a, string(ch))
	}
	if len(a) == 0 {
		return []string{""}
	}
	return a
}
