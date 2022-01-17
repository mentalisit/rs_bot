package botTelegram

func removeDuplicateElementInt(languages []int) []int {
	result := make([]int, 0, len(languages))
	temp := map[int]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
