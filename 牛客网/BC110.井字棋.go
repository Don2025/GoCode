package main

import "fmt"

func main() {
	const n = 3
	a := make([][]rune, n)
	for i := 0; i < n; i++ {
		a[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%c", &a[i][j])
		}
	}
	var flag int                                  //1 Kiki wins, -1 Bobo wins.
	if a[0][0] == a[1][1] && a[0][0] == a[2][2] { //左斜线
		if a[0][0] == 'K' {
			flag = 1
		} else if a[0][0] == 'B' {
			flag = -1
		}
	}
	if a[0][2] == a[1][1] && a[0][2] == a[2][0] { //右斜线
		if a[0][2] == 'K' {
			flag = 1
		} else if a[0][2] == 'B' {
			flag = -1
		}
	}
	for i := 0; i < n; i++ {
		if a[i][0] == a[i][1] && a[i][0] == a[i][2] { //行相等
			if a[i][0] == 'K' {
				flag = 1
			} else if a[i][0] == 'B' {
				flag = -1
			}
		}
		if a[0][i] == a[1][i] && a[0][i] == a[2][i] { //列相等
			if a[0][i] == 'K' {
				flag = 1
			} else if a[0][i] == 'B' {
				flag = -1
			}
		}
	}
	if flag == 1 {
		fmt.Print("KiKi wins!")
	} else if flag == -1 {
		fmt.Print("BoBo wins!")
	} else {
		fmt.Print("No winner!")
	}
}
