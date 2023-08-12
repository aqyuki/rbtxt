package args

func CheckLength(args []string, length int) bool {
	return len(args) == length
}
