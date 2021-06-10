package main

type ByteSize float64

//定义数量级
const (
	_           = iota             //用下划线忽略第一个变量,此时iota=0
	KB ByteSize = 1 << (10 * iota) //1 << (10*1),此时iota=1
	MB                             //1 << (10*2),此时iota=2
	GB                             //1 << (10*3),此时iota=3
	TB                             //1 << (10*4),此时iota=4
	PB                             //1 << (10*5),此时iota=5
	EB                             //1 << (10*6),此时iota=6
	ZB                             //1 << (10*7),此时iota=7
	YB                             //1 << (10*8),此时iota=8
)

func main() {
	println("KB =", KB)
	println("MB =", MB)
	println("GB =", GB)
	println("TB =", TB)
	println("PB =", PB)
	println("EB =", EB)
	println("ZB =", ZB)
	println("YB =", YB)
}
