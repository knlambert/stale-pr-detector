package std

import "time"

func CreateTime() *timeImpl {
	return &timeImpl{}
}

type timeImpl struct{}

func (t *timeImpl) Now() time.Time {
	return time.Now()
}
