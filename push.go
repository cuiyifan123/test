package utils

func Push(s []any, value any) []any {
	s = append(s, value)
	return s
}
