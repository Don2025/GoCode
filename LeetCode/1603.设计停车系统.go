package main

type ParkingSystem struct {
	left [4]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{[4]int{0, big, medium, small}}
}

func (s *ParkingSystem) AddCar(carType int) bool {
	if s.left[carType] > 0 {
		s.left[carType]--
		return true
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

/*
 * 执行用时：44ms 在所有Go提交中击败了11.32%的用户
 * 占用内存：7MB 在所有Go提交中击败了71.70%的用户
**/
