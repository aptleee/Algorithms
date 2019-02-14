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
// 	tmp = flag
//	flag = 1
//	return tmp
// }

// func lock() {
//	for test&set(flag) == 1 {}
// }
//
// func unlock() {
// 	flag = 0
// }

// func compare&swap(flag ,expected, new_v int) {
//	tmp := flag
//	if flag == expected {
//		flag = new_v
//	}
// 	return tmp
// }
//
// flag = 0
// func lock(){
// 	for compare&swap(flag, 0, 1) == 1 {}
// }
// func unlock() {
//	lock = 0
// }

