//49. Group Anagrams
//https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		w := SortString(s)
		m[w] = append(m[w], s)
	}
	result := make([][]string, 0)
	for v, _ := range m {
		result = append(result, m[v])
	}
	return result
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}