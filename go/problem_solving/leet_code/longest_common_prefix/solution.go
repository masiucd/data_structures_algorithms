package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	res := ""
	for i := 0; i < len(strs[0]); i++ {
		for _, s := range strs {
			char := strs[0][i]
			if i == len(s) || s[i] != char {
				return res
			}
		}
		res += string(strs[0][i])
	}
	return res
}
