package templates

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

// PrimeSieve generates a list of prime numbers of size n.
// 采用goroutine的并发组合，生成一个大小为n的素数列表。
// 程序从素数2开始，依次为每个素数建立一个goroutine，用于作为筛除该素数的倍数。
// ch 指向当前最新输出素数所位于的筛子goroutine的源channel。
func PrimeSieve(n int) (nums []int) {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < n; i++ {
		prime := <-ch
		nums = append(nums, prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	return
}
