package design_principles

// KMP
// - a: 主串
// - b：模式串
// - n：主串长度
// - m：模式串长度
func KMP(a, b []byte, n, m int) int {
	next := getNext(b, m)
	j := 0

	for i := 0; i < n; i++ {
		for j > 0 && a[i] != b[j] { // 一直找到a[i]和b[j]
			j = next[j-1] + 1
		}

		if a[i] == b[j] { // 找到匹配模式串
			j++
		}

		if j == m {
			return i - m + 1
		}
	}

	return -1
}

// getNext
// - b：表示模式串
// - m：表示模式串的长度
func getNext(b []byte, m int) []int {
	next := make([]int, m)
	next[0] = -1
	k := -1

	for i := 1; i < m; i++ {
		for k != -1 && b[k+1] != b[i] {
			k = next[k]
		}

		if b[k+1] == b[i] {
			k++
		}

		next[i] = k
	}

	return next
}
