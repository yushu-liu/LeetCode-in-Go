package problem0943

import (
	"strings"
)

func shortestSuperstring(A []string) string {
	size := len(A)
	indexs := make([]int, 0, size)
	isUsed := make([]bool, size)
	suffixes := getSuffixes(A)
	res := strings.Repeat("?", 12*20+1)
	for i := 0; i < size; i++ {
		isUsed[i] = true
		greedy(append(indexs, i), len(A[i]), A, isUsed, suffixes, &res)
		isUsed[i] = false
	}
	return res
}

// indexs 按顺序记录了 super string 中单词的 index
// length 记录了 super string 的长度
// 传入 suffixes 是为了避免重复多次计算两个单词之间的重叠关系
func greedy(indexs []int, length int, A []string, isUsed []bool, suffixes [][]int, minRes *string) {
	if len(indexs) == len(A) {
		if len(*minRes) > length {
			// NOTICE: 只有在确定找到了更短的 super string 时，才把它拼接出来。
			*minRes = connect(A, indexs, suffixes)
		}
		return
	}

	tail := indexs[len(indexs)-1]
	// get max suffix length of UNUSED string
	maxLen := -1
	lens := suffixes[tail]
	for i, sl := range lens {
		if maxLen >= sl || isUsed[i] {
			continue
		}
		maxLen = sl
	}

	// only connect string with max suffix length
	for i, sl := range lens {
		if sl < maxLen || isUsed[i] {
			continue
		}
		isUsed[i] = true
		greedy(append(indexs, i), length+len(A[i])-maxLen, A, isUsed, suffixes, minRes)
		isUsed[i] = false
	}
}

// res[i][j] == 3 means A[j][:3] is A[i]'s suffix
func getSuffixes(A []string) [][]int {
	size := len(A)
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}
			res[i][j] = suffix(A[i], A[j])
		}
	}
	return res
}

func suffix(a, b string) int {
	// none is substring of another, so can -1
	i := min(len(a), len(b)) - 1
	for !strings.HasSuffix(a, b[:i]) {
		i--
	}
	return i
}

func connect(A []string, indexs []int, suffixes [][]int) string {
	size := len(A)
	var sb strings.Builder
	sb.Grow(240)
	i := indexs[0]
	sb.WriteString(A[i])
	for k := 1; k < size; k++ {
		j := indexs[k]
		s := suffixes[i][j]
		sb.WriteString(A[j][s:])
		i = j
	}
	return sb.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}