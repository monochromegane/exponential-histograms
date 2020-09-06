package exphist

type ExpHistRealNumber struct {
	exphistV *ExpHistVector
}

func NewForRealNumber(m int) *ExpHistRealNumber {
	return &ExpHistRealNumber{
		exphistV: NewForVector(m),
	}
}

func (e *ExpHistRealNumber) Add(x float64) {
	e.exphistV.Add([]float64{x})
}

func (e *ExpHistRealNumber) Drop() {
	e.exphistV.Drop()
}

func (e *ExpHistRealNumber) Size() int {
	return e.exphistV.Size()
}

func (e *ExpHistRealNumber) Sum() float64 {
	return e.exphistV.Sum()[0]
}

func (e *ExpHistRealNumber) Tail() Buckets {
	t := e.exphistV.Tail()

	buckets := make([][]Bucket, len(t))
	for i, _ := range t {
		bs := make([]Bucket, len(t[i]))
		for j, _ := range t[i] {
			bs[j] = Bucket{
				capacity: t[i][j].capacity,
				content:  t[i][j].contents[0],
			}
		}
		buckets[i] = bs
	}
	return buckets
}

func (e *ExpHistRealNumber) Scale(gamma float64) {
	e.exphistV.Scale(gamma)
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
