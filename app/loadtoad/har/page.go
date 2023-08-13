package har

import (
	"fmt"
	"strings"

	"github.com/kyleu/loadtoad/app/util"
)

type Page struct {
	StartedDateTime string      `json:"startedDateTime"`
	ID              string      `json:"id"`
	Title           string      `json:"title"`
	PageTiming      *PageTiming `json:"pageTiming"`
	Comment         string      `json:"comment,omitempty"`
}

type Pages []*Page

type PageTiming struct {
	OnContentLoad int    `json:"onContentLoad"`
	OnLoad        int    `json:"onLoad"`
	Comment       string `json:"comment"`
}

type PageTimings struct {
	IsMicros bool   `json:"isMicros,omitempty"`
	Total    int    `json:"total,omitempty"`
	Blocked  int    `json:"blocked,omitempty"`
	DNS      int    `json:"dns,omitempty"`
	Connect  int    `json:"connect,omitempty"`
	Send     int    `json:"send,omitempty"`
	Wait     int    `json:"wait,omitempty"`
	Receive  int    `json:"receive,omitempty"`
	SSL      int    `json:"ssl,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

func (p *PageTimings) String() string {
	mult := 1000
	if p.IsMicros {
		mult = 1
	}
	var ret []string
	if p.Total > 0 {
		ret = append(ret, fmt.Sprintf("Total: %s", util.MicrosToMillis(p.Total*mult)))
	}
	if p.Blocked > 0 {
		ret = append(ret, fmt.Sprintf("Blocked: %s", util.MicrosToMillis(p.Blocked*mult)))
	}
	if p.DNS > 0 {
		ret = append(ret, fmt.Sprintf("DNS: %s", util.MicrosToMillis(p.DNS*mult)))
	}
	if p.Connect > 0 {
		ret = append(ret, fmt.Sprintf("Connect: %s", util.MicrosToMillis(p.Connect*mult)))
	}
	if p.Send > 0 {
		ret = append(ret, fmt.Sprintf("Send: %s", util.MicrosToMillis(p.Send*mult)))
	}
	if p.Wait > 0 {
		ret = append(ret, fmt.Sprintf("Wait: %s", util.MicrosToMillis(p.Wait*mult)))
	}
	if p.Receive > 0 {
		ret = append(ret, fmt.Sprintf("Receive: %s", util.MicrosToMillis(p.Receive*mult)))
	}
	if p.SSL > 0 {
		ret = append(ret, fmt.Sprintf("SSL: %s", util.MicrosToMillis(p.SSL*mult)))
	}
	return strings.Join(ret, ", ")
}
