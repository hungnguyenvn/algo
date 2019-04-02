package union

type Union struct {
	arr []int
}

func NewUnion(size int) *Union {
	return &Union{
		arr: make([]int, size),
	}
}

func (u *Union) initArr() {
	for i := 0; i < len(u.arr); i++ {
		u.arr[i] = i
	}
}

func (u *Union) union(p int, q int) {
	t := u.arr[p]

	if t == u.arr[q] {
		return
	}

	for j := 0; j < len(u.arr); j++ {
		if u.arr[j] == t {
			u.arr[j] = u.arr[q]
		}
	}
}

func (u *Union) root(k int) int {
	if u.arr[k] == k {
		return k
	}

	return u.root(u.arr[u.arr[k]])
}

func (u *Union) quickUnion(p int, q int) {
	rootP := u.root(p)
	rootQ := u.root(q)

	u.arr[rootQ] = rootP
}

func (u *Union) quickConnect(p int, q int) bool {
	return u.root(p) == u.root(q)
}

func (u *Union) connected(p int, q int) bool {
	return u.arr[p] == u.arr[q]
}

func sampleTest() {
	// var arr []int

	// fmt.Printf("Enter size >_ ")
	// var size int
	// _, err := fmt.Scan(&size)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// arr = make([]int, size)
	// for {
	// 	var p, q int
	// 	_, err := fmt.Scan(&p, &q)
	// 	if err != nil {
	// 		break
	// 	}

	// 	union(p, q)
	// }

	// for {
	// 	var p, q int
	// 	_, err := fmt.Scan(&p, &q)
	// 	if err != nil {
	// 		break
	// 	}

	// 	fmt.Println(connected(p, q))
	// }
}
