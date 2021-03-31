package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println
	now := time.Now()

	fmt.Println(now)

	fmt.Println(time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC))

	p(now.Add(time.Hour))

	p(now.Add(time.Hour).Sub(now))

	p(now.Local())
	p(now.In(time.UTC))
	p(now.Unix())
	p(now.UnixNano())

	p(now.Format("2006"))
	p(now.Format(time.RFC3339))
	p(now.Format(time.RFC1123))
	p(now.Format(time.RFC822))
	p(now.Format(time.RFC850))
	p(time.RFC1123)
	p(now.Format("2006-01-02T15:04:05.999999-07:00"))

	t, _ := time.Parse("2006-01-02", "2018-03-04")
	p(t)
	t, _ = time.ParseInLocation("2006-01-02", "2018-03-04", time.Local)
	p(t)
}
