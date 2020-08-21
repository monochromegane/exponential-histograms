package exphist

type ExpHistRealNumber struct {
	mergeSize int
	buckets   Buckets
}

func NewForRealNumber(m int) *ExpHistRealNumber {
	return &ExpHistRealNumber{
		mergeSize: m + 1,
		buckets:   [][]Bucket{[]Bucket{}},
	}
}

func (e *ExpHistRealNumber) Add(x float64) {
	bucket := Bucket{
		content:  x,
		capacity: 1,
	}
	e.buckets[0] = append(e.buckets[0], bucket)

	e.merge()
}

func (e *ExpHistRealNumber) Drop() {
	e.tail()
}

func (e *ExpHistRealNumber) Size() int {
	return e.buckets.Size()
}

func (e *ExpHistRealNumber) Sum() float64 {
	return e.buckets.Sum()
}

func (e *ExpHistRealNumber) Tail() Buckets {
	if len(e.buckets[len(e.buckets)-1]) == 1 {
		return [][]Bucket{e.buckets[len(e.buckets)-1]}
	} else {
		return [][]Bucket{[]Bucket{e.buckets[len(e.buckets)-1][0]}}
	}
}

func (e *ExpHistRealNumber) Scale(gamma float64) {
	for i := 0; i < len(e.buckets); i++ {
		for j := 0; j < len(e.buckets[i]); j++ {
			e.buckets[i][j].content *= gamma
		}
	}
}

func (e *ExpHistRealNumber) merge() {
	for i, _ := range e.buckets {
		if len(e.buckets[i]) < e.mergeSize {
			continue
		}
		if i == len(e.buckets)-1 {
			e.buckets = append(e.buckets, []Bucket{})
		}
		e.buckets[i+1] = append(e.buckets[i+1], Bucket{
			content:  e.buckets[i][0].content + e.buckets[i][1].content,
			capacity: e.buckets[i][0].capacity + e.buckets[i][1].capacity,
		})
		e.buckets[i] = e.buckets[i][2:]
	}
}

func (e *ExpHistRealNumber) tail() {
	if len(e.buckets[len(e.buckets)-1]) == 1 {
		e.buckets = e.buckets[0 : len(e.buckets)-1]
	} else {
		e.buckets[len(e.buckets)-1] = e.buckets[len(e.buckets)-1][1:]
	}
}

type Buckets [][]Bucket

func (b Buckets) Size() int {
	sum := 0
	for i, _ := range b {
		for j, _ := range b[i] {
			sum += b[i][j].capacity
		}
	}
	return sum
}

func (b Buckets) Sum() float64 {
	sum := 0.0
	for i, _ := range b {
		for j, _ := range b[i] {
			sum += b[i][j].content
		}
	}
	return sum
}

type Bucket struct {
	content  float64
	capacity int
}
