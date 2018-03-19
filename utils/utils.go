package utils

type V int

func Compare(a, b V) int {
	return int(a - b)
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
