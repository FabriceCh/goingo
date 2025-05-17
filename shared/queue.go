package shared

type Queue[T any] struct {
	Data []T
}

func NewQueue[T any]() *Queue[T] {
	d := make([]T, 0)
	return &Queue[T]{Data: d}
}

func (q *Queue[T]) IsEmpty() bool { return len(q.Data) == 0 }
func (q *Queue[T]) Insert(newItem T) {
	q.Data = append(q.Data, newItem)
}

func (q *Queue[T]) InsertMany(newItems []T) {
	q.Data = append(q.Data, newItems...)
}

func (q *Queue[T]) Pop() T {
	x := q.Data[0]
	q.Data = q.Data[1:]
	return x
}
