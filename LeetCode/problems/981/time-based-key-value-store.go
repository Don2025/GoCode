package _81

// https://leetcode.cn/problems/time-based-key-value-store/
//------------------------Leetcode Problem 981------------------------

type pair struct {
	Timestamp int
	Value     string
}

type TimeMap struct {
	timeMap map[string][]pair
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{make(map[string][]pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timeMap[key] = append(this.timeMap[key], pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	var ans string
	values := this.timeMap[key]
	if len(values) == 0 || values[0].Timestamp > timestamp {
		return ans
	}
	l, r := 0, len(values)-1
	for l <= r {
		mid := l + (r-l)>>1
		if values[mid].Timestamp > timestamp {
			r = mid - 1
		} else if values[mid].Timestamp <= timestamp {
			ans = values[mid].Value
			l = mid + 1
		}
	}
	return ans
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

//------------------------Leetcode Problem 981------------------------
/*
 * https://leetcode.cn/problems/time-based-key-value-store/
 * 执行用时：448ms 在所有Go提交中击败了67.35%的用户
 * 占用内存：65.4MB 在所有Go提交中击败了81.63%的用户
**/
