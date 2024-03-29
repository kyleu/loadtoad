package loadtoad

import (
	"crypto/tls"
	"net/http"
	"net/http/httptrace"

	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/util"
)

func wireReq(req *http.Request, timing *har.PageTimings) *http.Request {
	start := util.TimerStart()
	var connect, dns, tlsHandshake *util.Timer
	timing.IsMicros = true

	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) { dns = util.TimerStart() },
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			timing.DNS = float64(dns.End())
		},

		TLSHandshakeStart: func() { tlsHandshake = util.TimerStart() },
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			timing.SSL = float64(tlsHandshake.End())
		},

		ConnectStart: func(network, addr string) { connect = util.TimerStart() },
		ConnectDone: func(network, addr string, err error) {
			timing.Connect = float64(connect.End())
		},

		GotFirstResponseByte: func() {
			timing.Receive = float64(start.End())
		},
	}
	return req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
}
