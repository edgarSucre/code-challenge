package collection

type (
	collectible interface {
		comparable
	}

	Collection[T collectible] struct {
		data      []T
		instances map[T]int
	}

	iterator[T collectible] func(val T) bool
)

func New[T collectible](a []T) Collection[T] {
	c := Collection[T]{}
	c.data = a
	c.instances = make(map[T]int)

	for _, v := range a {
		c.instances[v]++
	}

	return c
}

func (c Collection[T]) Contains(val T) bool {
	_, ok := c.instances[val]
	return ok
}

func (c Collection[T]) Size() int {
	return len(c.data)
}

func (c *Collection[T]) Push(val T) {
	c.data = append(c.data, val)
	c.instances[val]++
}

func (c *Collection[T]) Pop() (last T) {
	last = c.data[c.Size()-1]
	c.data = c.data[:c.Size()-1]
	if c.instances[last] > 1 {
		c.instances[last]--
	} else {
		delete(c.instances, last)
	}

	return last
}

func (c Collection[T]) Filter(fn iterator[T]) Collection[T] {
	arr := make([]T, 0, c.Size())

	for _, v := range c.data {
		if fn(v) {
			arr = append(arr, v)
		}
	}

	return New(arr)
}
