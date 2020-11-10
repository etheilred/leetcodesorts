func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] < nums[len(nums)-1] {
			if target <= nums[len(nums)-1] {
				if nums[m] < target {
					l = m + 1
				} else {
					r = m
				}
			} else {
				r = m
			}
		} else {
			if target <= nums[len(nums)-1] {
				l = m + 1
			} else {
				if nums[m] < target {
					l = m + 1
				} else {
					r = m
				}
			}
		}
	}
	if nums[r] == target {
		return r
	}
	return -1
}