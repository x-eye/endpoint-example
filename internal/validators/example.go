package validators

import (
	"errors"
	pb "example/pkg/pb/example"
	"fmt"
	"regexp"
)

type exampleValidator struct {
	re *regexp.Regexp
}

func NewExampleValidator(regexp *regexp.Regexp) *exampleValidator {
	return &exampleValidator{re: regexp}
}

func (v exampleValidator) matchId(id string) bool {
	return v.re.Match([]byte(id))
}

func (v exampleValidator) Validate(request *pb.ExampleRequest) error {
	if !v.matchId(request.ApId) {
		return errors.New("app_id doesn't match")
	}
	if len(request.ProbeRequests) == 0 {
		return errors.New("empty probe_requests")
	}
	for i, probe := range request.ProbeRequests {
		if !v.matchId(probe.GetMac()) {
			return fmt.Errorf("mac doesn't match on %d probe_request", i)
		}
	}
	return nil
}
