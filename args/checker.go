package args

// CheckLength checks if the required number of arguments is given
func CheckLength(args []string, length int) bool {
	return len(args) == length
}
