package utils

import (
	"encoding/json"
	"strings"
	"time"
)

type LocalDate struct {
	time.Time
}

func TruncTime(tm time.Time) time.Time {
	parsed, _ := time.Parse(time.DateOnly, tm.Format(time.DateOnly))
	return parsed
}

func FormatTime(tm time.Time) string {
	return tm.Format(time.DateOnly)
}

func (ld *LocalDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return err
	}
	*ld = LocalDate{t}
	return nil
}

func (ld LocalDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(FormatTime(ld.Time))
}
