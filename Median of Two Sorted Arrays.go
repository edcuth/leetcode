//4. Median of Two Sorted Arrays
//https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := make([]int, len(nums1)+len(nums2))
	j := 0
	l := 0
	x := 0
	for i := 0; i < len(arr); i++ {
		if j == len(nums1) || l == len(nums2) {
			x = i
			break
		}
		if nums1[j] > nums2[l] {
			arr[i] = nums2[l]
			l++
		} else {
			arr[i] = nums1[j]
			j++
		}
	}
	if j < len(nums1) {
		for i := j; i < len(nums1); i++ {
			arr[x] = nums1[i]
			x++
		}
	}
	if l < len(nums2) {
		for i := l; i < len(nums2); i++ {
			arr[x] = nums2[i]
			x++
		}
	}
	fmt.Println(arr)
	if len(arr)%2 == 0 {
		return (float64(arr[len(arr)/2]+arr[len(arr)/2-1]) / 2)
	} else {
		return (float64(arr[len(arr)/2]))
	}
}