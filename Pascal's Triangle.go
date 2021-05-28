//118. Pascal's Triangle
//https://leetcode.com/problems/pascals-triangle/
func generate(numRows int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}
		result = append(result, row)
	}
	return result
}