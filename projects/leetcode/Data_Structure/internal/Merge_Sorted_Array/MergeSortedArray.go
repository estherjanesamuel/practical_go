package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// i := m - 1
	// j := n - 1
	// k := m + n - 1

	// for j >= 0 {
	// 	if i < 0 || nums1[i] < nums2[j] {
	// 		k--
	// 		j--
	// 		nums1[k] = nums2[j]
	// 	} else {
	// 		k--
	// 		i--
	// 		nums1[k] = nums1[i]
	// 	}
	// }
	// return []int{}

	i, m, n := len(nums1) - 1, m -1, n -1
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
		i--
	}
	for m >= 0 {
		nums1[i] = nums1[m]
		m--
		i--
	}

	for n >= 0 {
		nums1[i] = nums2[n]
		n--
		i--
	}
	i= i+1
	nums1 = nums1[i:]
	return nums1
}