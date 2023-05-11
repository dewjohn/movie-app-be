package dto

import "time"

type StatisticsMovieDto struct {
	Month []time.Month
	Value []int64
}

type StatisticsAllDataDto struct {
	CountMovie   int64
	CountUser    int64
	CountReply   int64
	CountComment int64
}

type StatisticsAllDataDtoResultDto struct {
	Type  []string
	Value []int64
}
