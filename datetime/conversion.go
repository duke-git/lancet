// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

package datetime

import "time"

type theTime struct {
	unix int64
}

// NewUnixNow return unix timestamp of current time
func NewUnixNow() *theTime {
	return &theTime{unix: time.Now().Unix()}
}

// NewUnix return unix timestamp of specified time
func NewUnix(unix int64) *theTime {
	return &theTime{unix: unix}
}

// NewFormat return unix timestamp of specified time string, t should be "yyyy-mm-dd hh:mm:ss"
func NewFormat(t string) (*theTime, error) {
	timeLayout := "2006-01-02 15:04:05"
	loc := time.FixedZone("CST", 8*3600)
	tt, err := time.ParseInLocation(timeLayout, t, loc)
	if err != nil {
		return nil, err
	}
	return &theTime{unix: tt.Unix()}, nil
}

// NewISO8601 return unix timestamp of specified iso8601 time string
func NewISO8601(iso8601 string) (*theTime, error) {
	t, err := time.ParseInLocation(time.RFC3339, iso8601, time.UTC)
	if err != nil {
		return nil, err
	}
	return &theTime{unix: t.Unix()}, nil
}

// ToUnix return unix timestamp
func (t *theTime) ToUnix() int64 {
	return t.unix
}

// ToFormat return the time string 'yyyy-mm-dd hh:mm:ss' of unix time
func (t *theTime) ToFormat() string {
	return time.Unix(t.unix, 0).Local().Format("2006-01-02 15:04:05")
}

// ToFormatForTpl return the time string which format is specified tpl
func (t *theTime) ToFormatForTpl(tpl string) string {
	return time.Unix(t.unix, 0).Local().Format(tpl)
}

// ToFormatForTpl return iso8601 time string
func (t *theTime) ToIso8601() string {
	return time.Unix(t.unix, 0).Local().Format(time.RFC3339)
}
