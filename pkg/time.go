package pkg

import "time"

//timeWrapper describes a wrapper around the time package.
type timeWrapper interface {
	Now() time.Time
}
