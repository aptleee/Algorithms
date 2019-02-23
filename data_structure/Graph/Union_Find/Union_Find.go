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

func (u *UF) connected(i, j int) bool {
	return u.friends[i] == u.friends[j]
}

func (u *UF) find(i int) int{
	for i != u.friends[i]{
		i = u.friends[i]
	}
	return i
}

func (u *UF) quickUnion(i, j int) { //quick union
	ir := u.find(i)
	jr := u.find(j)
	if ir != jr {
		u.count--
		u.friends[ir] = jr
	}
}

func (u *UF) quickFind(i int) int {
	return u.friends[i]
}

func (u *UF) union(i, j int) {
	ir := u.quickFind(i)
	jr := u.quickFind(j)
	if ir != jr {
		for k := 0; k < len(u.friends); k++ {
			if u.friends[k] == ir {
				u.friends[k] = jr
			}
		}
	}
	u.count--
}
