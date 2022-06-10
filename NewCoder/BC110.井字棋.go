package main

import "fmt"

func main() {
	const n = 3
	var t rune
	a := make([][]rune, n)
	for i := 0; i < n; i++ {
		a[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%c", &a[i][j])
			fmt.Scanf("%c", &t) //吃空格
		}
	}
	if a[0][0] == a[1][1] && a[0][0] == a[2][2] { //左斜线
		if a[0][0] == 'K' {
			private_Print(1)
			return
		} else if a[0][0] == 'B' {
			private_Print(-1)
			return
		}
	}
	if a[0][2] == a[1][1] && a[0][2] == a[2][0] { //右斜线
		if a[0][2] == 'K' {
			private_Print(1)
			return
		} else if a[0][2] == 'B' {
			private_Print(-1)
			return
		}
	}
	for i := 0; i < n; i++ {
		if a[i][0] == a[i][1] && a[i][0] == a[i][2] { //行相等
			if a[i][0] == 'K' {
				private_Print(1)
				return
			} else if a[i][0] == 'B' {
				private_Print(-1)
				return
			}
		}
		if a[0][i] == a[1][i] && a[0][i] == a[2][i] { //列相等
			if a[0][i] == 'K' {
				private_Print(1)
				return
			} else if a[0][i] == 'B' {
				private_Print(-1)
				return
			}
		}
	}
	private_Print(0)
	return
}

func private_Print(flag int) {
	if flag == 1 {
		fmt.Print("KiKi wins!")
	} else if flag == -1 {
		fmt.Print("BoBo wins!")
	} else {
		fmt.Print("No winner!")
	}
}

/*
 * 运行时间：7ms 超过0.00%用Go提交的代码
 * 占用内存：888KB 超过100.00%用Go提交的代码
**/
