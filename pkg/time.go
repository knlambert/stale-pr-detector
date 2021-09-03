package pkg

import "time"

//Time describe a wrapper around the time package.
type Time interface {
	Now() time.Time
}
