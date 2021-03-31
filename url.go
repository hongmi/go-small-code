package main

import (
	"fmt"
	//"net"
	"net/url"
)

func main() {
	p := fmt.Println

	s := "mongodb://root:password@localhost:9922/test?ssl=true#r"

	u, _ := url.Parse(s)

	p(u)
	p(u.EscapedFragment())
	p(u.EscapedPath())
	p(u.ForceQuery)
	p(u.Fragment)
	p(u.Host)
	p(u.Hostname())
	p(u.IsAbs())
	p(u.Opaque)
	p(u.Query())
	p(u.RawQuery)
	p(u.Redacted())
	p(u.User)

}
