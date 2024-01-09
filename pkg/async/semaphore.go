package async

type Semaphore struct {
	c chan struct{}
}

func (s *Semaphore) Acquire() {
	s.c <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.c
}

func (s *Semaphore) Close() {
	close(s.c)
}

func NewSemaphore(size int) *Semaphore {
	return &Semaphore{
		c: make(chan struct{}, size),
	}
}
