// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

package datetime

import "time"

type theTime struct {
	unix int64
}

func NewUnixNow() *theTime {
	return &theTime{unix: time.Now().Unix()}
}

func NewUnix(unix int64) *theTime {
	return &theTime{unix: unix}
}

func NewFormat(t string) (*theTime, error) {
	timeLayout := "2006-01-02 15:04:05"
	loc := time.FixedZone("CST", 8*3600)
	tt, err := time.ParseInLocation(timeLayout, t, loc)
	if err != nil {
		return nil, err
	}
	return &theTime{unix: tt.Unix()}, nil
}

func NewISO8601(iso8601 string) (*theTime, error) {
	t, err := time.ParseInLocation(time.RFC3339, iso8601, time.UTC)
	if err != nil {
		return nil, err
	}
	return &theTime{unix: t.Unix()}, nil
}

func (t *theTime) ToUnix() int64 {
	return t.unix
}

func (t *theTime) ToFormat() string {
	return time.Unix(t.unix, 0).Format("2006-01-02 15:04:05")
}

func (t *theTime) ToFormatForTpl(tpl string) string {
	return time.Unix(t.unix, 0).Format(tpl)
}

func (t *theTime) ToIso8601() string {
	return time.Unix(t.unix, 0).Format(time.RFC3339)
}
