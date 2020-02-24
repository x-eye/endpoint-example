package storage

import (
	"context"
	pb "example/pkg/pb/example"
	"fmt"
	"net/http"
)

type partner struct {
	queue  chan *pb.ExampleRequest
	len    int
	cap    int
	client *pb.ExampleService_httpClient
}

func NewPartner(httpClient *http.Client, addr string, cap int) *partner {
	return &partner{
		cap:    cap,
		queue:  make(chan *pb.ExampleRequest, cap),
		client: pb.NewExampleServiceHTTPClient(httpClient, addr),
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
			p.len--
			_, err := p.client.ExampleEndpoint(ctx, populate(r))
			if err != nil {
				// todo: log error
			}
		}
	}
}
