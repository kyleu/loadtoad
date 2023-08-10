package har

import "time"

type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Path     string `json:"path,omitempty"`
	Domain   string `json:"domain,omitempty"`
	Expires  string `json:"expires,omitempty"`
	HTTPOnly bool   `json:"httpOnly,omitempty"`
	Secure   bool   `json:"secure,omitempty"`
	Comment  bool   `json:"comment,omitempty"`
}

func (c *Cookie) Exp() time.Time {
	ret, _ := time.Parse("2006-01-02T15:04:05.000Z", c.Expires)
	return ret
}

type Cookies []*Cookie
