package LeetBook

func equationsPossible(equations []string) bool {
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}
	for _, str := range equations {
		if str[1] == '=' {
			idx1 := int(str[0] - 'a')
			idx2 := int(str[3] - 'a')
			union(parent, idx1, idx2)
		}
	}
	for _, str := range equations {
		if str[1] == '!' {
			idx1 := int(str[0] - 'a')
			idx2 := int(str[3] - 'a')
			if find(parent, idx1) == find(parent, idx2) {
				return false
			}
		}
	}
	return true
}

func union(parent []int, idx1, idx2 int) {
	parent[find(parent, idx1)] = find(parent, idx2)
}

func find(parent []int, idx int) int {
	for parent[idx] != idx {
		parent[idx] = parent[parent[idx]]
		idx = parent[idx]
	}
	return idx
}

/*
 * 执行用时：4ms 在所有Go提交中击败了75.93%的用户
 * 占用内存：2.8MB 在所有Go提交中击败了100.00%的用户
**/
