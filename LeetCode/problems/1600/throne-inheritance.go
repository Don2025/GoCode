package _600

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */

type ThroneInheritance struct {
	king  string
	edges map[string][]string
	dead  map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{kingName, map[string][]string{}, map[string]bool{}}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.edges[parentName] = append(this.edges[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.dead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	var ans []string
	var preorder func(string)
	preorder = func(name string) {
		if !this.dead[name] {
			ans = append(ans, name)
		}
		for _, x := range this.edges[name] {
			preorder(x)
		}
	}
	preorder(this.king)
	return ans
}

//------------------------Leetcode Problem 1600------------------------
/*
 * https://leetcode.cn/problems/throne-inheritance/
 * 执行用时：576ms 在所有Go提交中击败了38.89%的用户
 * 占用内存：52.8MB 在所有Go提交中击败了88.89%的用户
**/
