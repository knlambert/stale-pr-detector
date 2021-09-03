package pkg

import "time"

//TimeWrapper describes a wrapper around the time package.
type TimeWrapper interface {
	Now() time.Time
}
