package dto

import "time"

type StatisticsMovieDto struct {
	Month []time.Month
	Value []int64
}
