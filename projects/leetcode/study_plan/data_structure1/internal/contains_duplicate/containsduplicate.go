package containsduplicate



func containsDuplicate(nums []int) bool {
	hashTable := make(map[int]bool)

	for _, val := range nums {
		if hashTable[val] {
			return true
		} else {
			hashTable[val] = true
		}
	}
	return false
}

func containsDuplicateII(nums []int) bool {
    set := make(map[int]struct{})
    for _, v := range nums {
        if _, ok := set[v]; ok {
            return true
        } else {
            set[v] = struct{}{}
        }
    }
    return false
}