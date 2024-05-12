package utils

import "time"

type DateProvider interface {
	Today() time.Time
}

type currentDateProvider struct{}

func NewDateProvider() DateProvider {
	return &currentDateProvider{}
}

func (cds *currentDateProvider) Today() time.Time {
	return TruncTime(time.Now())
}
