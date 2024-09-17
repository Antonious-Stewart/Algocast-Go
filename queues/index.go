package queues

type Queue struct {
	Data   []int
	Length int
}

func (q *Queue) Enqueue(value int) *Queue {
	q.Data = append(q.Data, value)
	q.Length++
	return q
}

func (q *Queue) Dequeue() *Queue {
	if q.Length == 0 {
		return nil
	}

	q.Data = append(q.Data[:0], q.Data[1:]...)

	q.Length--
	return q
}

func (q *Queue) Peek() int {
	if q.Length == 0 {
		return 0
	}

	return q.Data[0]
}

func NewQueue() *Queue {
	return &Queue{}
}
