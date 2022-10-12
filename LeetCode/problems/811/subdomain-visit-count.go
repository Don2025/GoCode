package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

// ------------------------Leetcode Problem 811------------------------
// https://leetcode.cn/problems/subdomain-visit-count/
func subdomainVisits(cpdomains []string) []string {
	cnt := make(map[string]int)
	for _, cpdomain := range cpdomains {
		i := strings.IndexByte(cpdomain, ' ')
		count, _ := strconv.Atoi(cpdomain[:i])
		domain := cpdomain[i+1:]
		cnt[domain] += count
		for i := strings.IndexByte(domain, '.'); i != -1; i = strings.IndexByte(domain, '.') {
			domain = domain[i+1:]
			cnt[domain] += count
		}
	}
	ans := make([]string, 0, len(cnt))
	for domain, count := range cnt {
		ans = append(ans, strconv.Itoa(count)+" "+domain)
	}
	return ans
}

// ------------------------Leetcode Problem 811------------------------
/*
 * https://leetcode.cn/problems/subdomain-visit-count/
 * 执行用时：4ms 在所有Go提交中击败了95.14%的用户
 * 占用内存：5.2MB 在所有Go提交中击败了82.49%的用户
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cpdomains := strings.Fields(scanner.Text())
		Printf("Output: %s\n", subdomainVisits(cpdomains))
	}
}
