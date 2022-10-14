package main

func reverse(s string) string {
	rs := make([]rune, len(s))

	i := len(s) - 1
	for _, c := range s {
		rs[i] = c
		i--
	}

	return string(rs)
}
