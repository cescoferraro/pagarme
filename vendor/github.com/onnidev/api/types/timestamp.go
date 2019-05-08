package types

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Timestamp TODO: NEEDS COMMENT INFO
type Timestamp time.Time

// MarshalJSON TODO: NEEDS COMMENT INFO
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts * 1000)
	return []byte(stamp), nil
}

// UnmarshalJSON TODO: NEEDS COMMENT INFO
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	gg := time.Unix(int64(ts/1000), 0)
	*t = Timestamp(gg)
	return nil
}

// Time TODO: NEEDS COMMENT INFO
func (t *Timestamp) Time() time.Time {
	return time.Time(*t)
}

// GetBSON TODO: NEEDS COMMENT INFO
func (t *Timestamp) GetBSON() (interface{}, error) {
	if t.Time().IsZero() {
		return nil, nil
	}
	return t.Time(), nil
}

// SetBSON TODO: NEEDS COMMENT INFO
func (t *Timestamp) SetBSON(raw bson.Raw) error {
	var tm time.Time

	if err := raw.Unmarshal(&tm); err != nil {
		return err
	}

	*t = Timestamp(tm)

	return nil

}

// String TODO: NEEDS COMMENT INFO
func (t *Timestamp) String() string {
	return strconv.Itoa(int(time.Time(*t).Unix() / 100000000))
}
