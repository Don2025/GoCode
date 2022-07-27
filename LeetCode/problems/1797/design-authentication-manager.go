package _797

// https://leetcode.cn/problems/design-authentication-manager/
//------------------------Leetcode Problem 1797------------------------

type AuthenticationManager struct {
	Manager    map[string]int
	TimeToLive int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{map[string]int{}, timeToLive}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.Manager[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if x, ok := this.Manager[tokenId]; ok {
		if x+this.TimeToLive > currentTime {
			this.Manager[tokenId] = currentTime
		}
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var cnt int
	for _, x := range this.Manager {
		if x+this.TimeToLive > currentTime {
			cnt++
		}
	}
	return cnt
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
//------------------------Leetcode Problem 1797------------------------
/*
 * https://leetcode.cn/problems/design-authentication-manager/
 * 执行用时：100ms 在所有Go提交中击败了14.63%的用户
 * 占用内存：7.2MB 在所有Go提交中击败了82.93%的用户
**/
