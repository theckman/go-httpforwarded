package httpforwarded_test

import (
	httpforwarded "github.com/theckman/go-httpforwarded"
	. "gopkg.in/check.v1"
)

func (*TestSuite) TestFormat(c *C) {
	var out string
	var err error

	params := map[string][]string{
		"for":   []string{"192.0.2.1", "192.0.2.4"},
		"by":    []string{"192.0.2.200", "192.0.2.202"},
		"proto": []string{"http"},
	}

	out, err = httpforwarded.Format(params)
	c.Assert(err, IsNil)
	c.Check(out, Equals, "by=192.0.2.200, by=192.0.2.202; for=192.0.2.1, for=192.0.2.4; proto=http")
}