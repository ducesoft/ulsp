package jsonrpc2

import (
	"context"
	"encoding/json"
	"github.com/ducesoft/ulsp/log"
	"sync"
)

// ConnOpt is the type of function that can be passed to NewConn to
// customize the Conn before it is created.
type ConnOpt func(*Conn)

// OnRecv causes all requests received on conn to invoke f(req, nil)
// and all responses to invoke f(req, resp),
func OnRecv(f func(ctx context.Context, rr *Request, rs *Response)) ConnOpt {
	return func(c *Conn) { c.onRecv = append(c.onRecv, f) }
}

// OnSend causes all requests sent on conn to invoke f(req, nil) and
// all responses to invoke f(nil, resp),
func OnSend(f func(ctx context.Context, rr *Request, rs *Response)) ConnOpt {
	return func(c *Conn) { c.onSend = append(c.onSend, f) }
}

// LogMessages causes all messages sent and received on conn to be
// logged using the provided logger.
func LogMessages() ConnOpt {
	return func(c *Conn) {
		// Remember reqs we have received so we can helpfully show the
		// request method in OnSend for responses.
		var (
			mu         sync.Mutex
			reqMethods = map[ID]string{}
		)
		OnRecv(func(ctx context.Context, rr *Request, rs *Response) {
			switch {
			case rr != nil:
				mu.Lock()
				reqMethods[rr.ID] = rr.Method
				mu.Unlock()

				log.Info(ctx, "")
				params, _ := json.Marshal(rr.Params)

				if rr.Notif {
					log.Info(ctx, "jsonrpc2: --> notif: %s: %s\n", rr.Method, params)
				} else {
					log.Info(ctx, "jsonrpc2: --> request #%s: %s: %s\n", rr.ID, rr.Method, params)
				}

			case rs != nil:
				var method string
				if rr != nil {
					method = rr.Method
				} else {
					method = "(no matching request)"
				}
				switch {
				case rs.Result != nil:
					result, _ := json.Marshal(rs.Result)
					log.Info(ctx, "jsonrpc2: --> result #%s: %s: %s\n", rs.ID, method, result)
				case rs.Error != nil:
					err, _ := json.Marshal(rs.Error)
					log.Info(ctx, "jsonrpc2: --> error #%s: %s: %s\n", rs.ID, method, err)
				}
			}
		})(c)
		OnSend(func(ctx context.Context, rr *Request, rs *Response) {
			switch {
			case rr != nil:
				params, _ := json.Marshal(rr.Params)
				if rr.Notif {
					log.Info(ctx, "jsonrpc2: <-- notif: %s: %s\n", rr.Method, params)
				} else {
					log.Info(ctx, "jsonrpc2: <-- request #%s: %s: %s\n", rr.ID, rr.Method, params)
				}

			case rs != nil:
				mu.Lock()
				method := reqMethods[rs.ID]
				delete(reqMethods, rs.ID)
				mu.Unlock()
				if method == "" {
					method = "(no previous request)"
				}

				if rs.Result != nil {
					result, _ := json.Marshal(rs.Result)
					log.Info(ctx, "jsonrpc2: <-- result #%s: %s: %s\n", rs.ID, method, result)
				} else {
					err, _ := json.Marshal(rs.Error)
					log.Info(ctx, "jsonrpc2: <-- error #%s: %s: %s\n", rs.ID, method, err)
				}
			}
		})(c)
	}
}
