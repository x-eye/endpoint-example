package storage

import (
	"bytes"
	"context"
	"encoding/json"
	pb "example/pkg/pb/example"
	"fmt"
	"net/http"
)

type partner struct {
	queue      chan *pb.ExampleRequest
	len        int
	cap        int
	target     string
	httpClient *http.Client
}

func NewPartner(httpClient *http.Client, target string, cap int) *partner {
	return &partner{
		cap:        cap,
		queue:      make(chan *pb.ExampleRequest, cap),
		target:     target,
		httpClient: httpClient,
	}
}

const (
	defaultBssid = "FF-FF-FF-FF-FF-FF"
	defaultSsid  = "Unknown"
)

func (p *partner) Send(request *pb.ExampleRequest) error {
	if p.len == p.cap {
		return fmt.Errorf("partner queue overflow")
	}
	p.len++
	p.queue <- request
	return nil
}

func populate(r *pb.ExampleRequest) *pb.ExampleRequest {
	res := &pb.ExampleRequest{
		ApId:          r.ApId,
		ProbeRequests: make([]*pb.Probe, 0, len(r.ProbeRequests)),
	}
	for _, p := range r.ProbeRequests {
		probe := *p
		if probe.Bssid == "" {
			probe.Bssid = defaultBssid
		}
		if probe.Ssid == "" {
			probe.Ssid = defaultSsid
		}
		res.ProbeRequests = append(res.ProbeRequests, &probe)
	}
	return res
}

func (p *partner) Consume(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case r := <-p.queue:
			data, err := json.Marshal(populate(r))
			if err != nil {
				p.queue <- r
				// todo: log error
				break
			}
			req, err := http.NewRequest(http.MethodPost, p.target, bytes.NewReader(data))
			resp, err := p.httpClient.Do(req)
			if err != nil {
				p.queue <- r
				// todo: log error
				break
			}
			if resp.StatusCode != http.StatusOK {
				// todo: use fallback strategy
				break
			}
			p.len--
		}
	}
}
