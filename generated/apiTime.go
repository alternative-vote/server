package generated

import (
	"fmt"
	"strings"
	"time"
)

const (
	time_format = "2006-01-02T15:04:05.999Z0700"
)

type APITime struct {
	time.Time
}

func NewAPITime(t time.Time) APITime {
	return APITime{
		Time: t,
	}
}

func ParseAPITime(s string) (APITime, error) {
	t, err := time.Parse(time_format, s)
	return APITime{
		Time: t,
	}, err
}

func (t *APITime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(time_format))
	return []byte(stamp), nil
}

func (t *APITime) UnmarshalJSON(data []byte) error {
	t = new(APITime)

	s := strings.Replace(string(data), "\"", "", -1)

	var err error

	t.Time, err = time.Parse(time_format, s)
	if err == nil {
		return nil
	}

	t.Time, err = time.Parse("2006-01-02", s)
	if err == nil {
		return nil
	}

	return err
}
