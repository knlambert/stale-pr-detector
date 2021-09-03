package pkg

import "time"

//Time describes a wrapper around the time package.
type Time interface {
	Now() time.Time
}
