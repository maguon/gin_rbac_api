package common

import (
	"fmt"
	"strings"
	"time"
)

type CommonDate struct {
	time.Time
}

func (t CommonDate) MarshalJSON() ([]byte, error) {
	date := t.Time.Format("2006-01-02")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *CommonDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	fmt.Println(s)
	fmt.Println(date)
	t.Time = date
	return
}
