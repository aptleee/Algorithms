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
