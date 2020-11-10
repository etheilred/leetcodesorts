import "fmt"

func sortArray(nums []int) []int {
	sort(nums, 0, len(nums)-1, 16)
	return nums
}

var med [3]int

func sort(nums []int, l, r int, depth int) {
	if r-l == 1 {
		if nums[l] > nums[r] {
			t := nums[l]
			nums[l] = nums[r]
			nums[r] = t
		}
	} else if l < r-1 {
		if r-l <= 64 {
			insertionSort(nums[l : r+1])
			return
		}
		if depth < 0 {
			heapSort(nums[l : r+1])
			return
		}
		p := partition(nums, l, r)
		sort(nums, l, p, depth-1)
		sort(nums, p+1, r, depth-1)
	}
}

func partition(nums []int, l, r int) int {
	m := nums[(l+r)/2]
	i := l
	j := r
	for i <= j {
		for nums[i] < m {
			i++
		}
		for nums[j] > m {
			j--
		}
		if i >= j {
			break
		}
		t := nums[i]
		nums[i] = nums[j]
		nums[j] = t
		i++
		j--
	}
	return j
}

func insertionSort(nums []int) {
	var t int
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			t = nums[j]
			nums[j] = nums[j-1]
			nums[j-1] = t
		}
	}
}

func heapSort(nums []int) {
	heapSize := len(nums)
	// for i := 0; i < len(nums); i++ {
	// 	heapSize++
	// 	siftUp(nums[:heapSize], i)
	// }
	for i := len(nums) / 2; i >= 0; i-- {
		siftDown(nums, i)
	}
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		nums[0], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[0]
		heapSize--
		siftDown(nums[:heapSize], 0)
	}
}

func siftDown(heap []int, i int) {
	for 2*i+1 < len(heap) {
		left := 2*i + 1
		right := 2*i + 2
		j := left
		if right < len(heap) && heap[right] > heap[left] {
			j = right
		}
		if heap[i] >= heap[j] {
			break
		}
		heap[i], heap[j] = heap[j], heap[i]
		i = j
	}
}

func siftUp(heap []int, i int) {
	for heap[i] > heap[(i-1)/2] {
		heap[i], heap[(i-1)/2] = heap[(i-1)/2], heap[i]
		i = (i - 1) / 2
	}
}