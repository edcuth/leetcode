//119. Pascal's Triangle II
//https://leetcode.com/problems/pascals-triangle-ii/
func getRow(rowIndex int) []int {
	result := make([][]int, 0)
	for i := 0; i < rowIndex+1; i++ {
		row := make([]int, i+1)
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}
		result = append(result, row)
	}
	return result[rowIndex]

}