package exphist

type ExpHistVector struct {
	mergeSize int
	buckets   VectorBuckets
}

func NewForVector(m int) *ExpHistVector {
	return &ExpHistVector{
		mergeSize: m + 1,
		buckets:   [][]VectorBucket{[]VectorBucket{}},
	}
}

func (e *ExpHistVector) Add(x []float64) {
	bucket := VectorBucket{
		contents: x,
		capacity: 1,
	}
	e.buckets[0] = append(e.buckets[0], bucket)

	e.merge()
}

func (e *ExpHistVector) Drop() {
	e.tail()
}

func (e *ExpHistVector) Size() int {
	return e.buckets.Size()
}

func (e *ExpHistVector) Sum() []float64 {
	return e.buckets.Sum()
}

func (e *ExpHistVector) Tail() VectorBuckets {
	if len(e.buckets[len(e.buckets)-1]) == 1 {
		return [][]VectorBucket{e.buckets[len(e.buckets)-1]}
	} else {
		return [][]VectorBucket{[]VectorBucket{e.buckets[len(e.buckets)-1][0]}}
	}
}

func (e *ExpHistVector) Scale(gamma float64) {
	for i := 0; i < len(e.buckets); i++ {
		for j := 0; j < len(e.buckets[i]); j++ {
			for k, _ := range e.buckets[i][j].contents {
				e.buckets[i][j].contents[k] *= gamma
			}
		}
	}
}

func (e *ExpHistVector) merge() {
	for i, _ := range e.buckets {
		if len(e.buckets[i]) < e.mergeSize {
			continue
		}
		if i == len(e.buckets)-1 {
			e.buckets = append(e.buckets, []VectorBucket{})
		}
		contents := make([]float64, len(e.buckets[i][0].contents))
		for j, _ := range e.buckets[i][0].contents {
			contents[j] = e.buckets[i][0].contents[j] + e.buckets[i][1].contents[j]
		}
		e.buckets[i+1] = append(e.buckets[i+1], VectorBucket{
			contents: contents,
			capacity: e.buckets[i][0].capacity + e.buckets[i][1].capacity,
		})
		e.buckets[i] = e.buckets[i][2:]
	}
}

func (e *ExpHistVector) tail() {
	if len(e.buckets[len(e.buckets)-1]) == 1 {
		e.buckets = e.buckets[0 : len(e.buckets)-1]
	} else {
		e.buckets[len(e.buckets)-1] = e.buckets[len(e.buckets)-1][1:]
	}
}

type VectorBuckets [][]VectorBucket

func (b VectorBuckets) Size() int {
	sum := 0
	for i, _ := range b {
		for j, _ := range b[i] {
			sum += b[i][j].capacity
		}
	}
	return sum
}

func (b VectorBuckets) Sum() []float64 {
	var sum []float64
	for i, _ := range b {
		for j, _ := range b[i] {
			if sum == nil {
				sum = make([]float64, len(b[i][j].contents))
			}
			for k, _ := range b[i][j].contents {
				sum[k] += b[i][j].contents[k]
			}
		}
	}
	return sum
}

type VectorBucket struct {
	contents []float64
	capacity int
}
