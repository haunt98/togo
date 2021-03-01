package clock

import "time"

type NowFn func() time.Time

func Now() time.Time {
	return time.Now()
}
