func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] < nums[r] {
			if target <= nums[r] {
				if nums[m] < target {
					l = m + 1
				} else {
					r = m
				}
			} else {
				r = m
			}
		} else if nums[m] > nums[r] {
			if target <= nums[r] {
				l = m + 1
			} else {
				if nums[m] < target {
					l = m + 1
				} else {
					r = m
				}
			}
		} else {
			r--
		}
	}
	if r >= 0 && nums[r] == target {
		return true
	}
	return false
}