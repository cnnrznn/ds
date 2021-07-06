// Package heap provides a min-heap data structure
package heap

// Heap is a min-heap with the property that a parent node has a
// value less than or equal to all of it's children. The 'Less' function
// is provided to New when constructing a Heap.
type Heap struct {
	Data []interface{}
	Less func(a, b interface{}) bool
	Size int
}

// New returns a fresh heap with the provided compare function.
func New(less func(a, b interface{}) bool) *Heap {
	return &Heap{
		Less: less,
	}
}

// Push an item onto the heap
func (h *Heap) Push(data interface{}) {
	h.Size++

	h.Data = append(h.Data, data)
	h.bubbleUp()
}

// Pop the minimum item from the heap and return it
func (h *Heap) Pop() interface{} {
	h.Size--

	tmp := h.Data[0]

	h.Data[0] = h.Data[h.Size]
	h.bubbleDown()

	return tmp
}

func (h *Heap) bubbleUp() {
	index := h.Size - 1

	for {
		if index == 0 {
			break
		}
		if !h.Less(h.Data[index], h.Data[index/2]) {
			break
		}

		h.swap(index, index/2)
		index = index / 2
	}
}

func (h *Heap) bubbleDown() {
	index := 0

	for {
		baseIndex := index*2 + 1

		if baseIndex >= h.Size {
			// at leaf
			break
		}

		child := baseIndex
		if baseIndex+1 < h.Size && h.Less(h.Data[baseIndex+1], h.Data[baseIndex]) {
			child = baseIndex + 1
		}

		if h.Less(h.Data[index], h.Data[child]) {
			break
		}

		h.swap(index, child)
		index = child
	}
}

func (h *Heap) swap(i, j int) {
	tmp := h.Data[i]
	h.Data[i] = h.Data[j]
	h.Data[j] = tmp
}
