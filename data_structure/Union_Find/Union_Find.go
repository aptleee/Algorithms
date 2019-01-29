package unionFind

type UF struct {
	count int
	friends []int
}

func New(n int) *UF {
	u := &UF {
		count:n,
		friends:make([]int, n),
	}
	for i := range u.friends{
		u.friends[i] = i
	}
	return u
}
func (u *UF) find(i int) int{
	for i != u.friends[i]{
		i = u.friends[i]
	}
	return i
}
func (u *UF) union(i, j int) {
	ir := u.find(i)
	jr := u.find(j)
	if ir != jr {
		u.count--
		u.friends[ir] = jr
	}
}
func findCircleNum(M [][]int) int {
	n := len(M)
	u := New(n)
	for i:=0; i<n; i++ {
		for j:=i+1; j<n; j++ {
			if M[i][j] == 1{
				u.union(i, j)
			}
		}
	}
	return u.count
}

func main() {

}