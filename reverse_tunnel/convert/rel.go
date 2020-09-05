package convert

import "daily/reverse_tunnel/convert/file"

const (
	HACONFIG  = "ha"
	NATCONFIG = "nat"
	LINK      = "link"
)

var (
	RUNJSON = file.NewRunJSON()
)

var Rel = map[string][]file.FileHandler{
	HACONFIG:  {RUNJSON},
	NATCONFIG: {RUNJSON},
}
