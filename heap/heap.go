package heap

type Heap struct {
	Data    []interface{}
	Compare func(a, b interface{}) int
	Size    int
}

func New(compare func(a, b interface{}) int) *Heap {
	return &Heap{
		Compare: compare,
	}
}

func (h *Heap) Push(data interface{}) {
	h.Size++

	h.Data = append(h.Data, data)
	h.bubbleUp()
}

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
		if h.Compare(h.Data[index], h.Data[index/2]) >= 0 {
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
		if baseIndex+1 < h.Size && h.Compare(h.Data[baseIndex], h.Data[baseIndex+1]) > 0 {
			child = baseIndex + 1
		}

		if h.Compare(h.Data[index], h.Data[child]) <= 0 {
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
