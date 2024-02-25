package costrecord

import "cloud.google.com/go/spanner"

type CostRecord struct {
	Date   spanner.NullDate
	Cost   float64
	Amount int64
}
