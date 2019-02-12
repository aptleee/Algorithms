package lock

var turn int = 0
var flag = [2]int{0, 0}

func lock(i int) {
	flag[i] = 1
	turn = 1 - i
	for flag[1-i] == 1 && turn == 1-i {
	}
}

func unlock(i int) {
	flag[i] = 0
}

// simple spin lock implementation
// flag = 0
// func test&set(flag)  {
// 	rv = flag
//	flag = 1
//	return rv
// }

// func lock() {
//	for test&set(flag) == 1 {}
// }
//
// func unlock() {
// 	flag = 0
// }

// func compare&swap(v ,expected, new_v int) {
//	tmp := v
//	if v == expected {
//		v = new_v
//	}
// 	return tmp
// }
//
// lock = 0
// func lock(){
// 	for compare&swap(v, 0, 1) != 0 {}
// }
// func unlock() {
//	lock = 0
// }

