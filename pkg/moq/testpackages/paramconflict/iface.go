package paramconflict

import "time"

type Interface interface {
	Method(string, bool, string, bool, int, int32, int64, float32, float64, time.Time, time.Time)
}
