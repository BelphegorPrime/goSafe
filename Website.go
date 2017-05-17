package main

import "time"

type Website struct {
	ID        int
	Url       string
	UserName  string
	Password  string
	Timestamp time.Time
}

func (w Website) ToString() string {
	return "\n" +
		w.Url + " " +
		w.UserName + " " +
		w.Password + " " +
		w.Timestamp.Format("02 Jan 2006 15:04:05")
}
