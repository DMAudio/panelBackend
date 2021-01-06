package typeUtils

func SliceIndexOfString(input []string, el string) int {
	if input == nil {
		return -1
	}
	if len(input) == 0 {
		return -1
	}
	for i, item := range input {
		if item == el {
			return i
		}
	}
	return -1

}
