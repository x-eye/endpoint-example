package validators

import (
	"encoding/json"
	pb "example/pkg/pb/example"
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	idRe    = "^([0-9A-Fa-f]{2}[-]){5}([0-9A-Fa-f]{2})$"
	good    = "test_data/good.json"
	invalid = "test_data/invalid.json"
)

func TestExampleValidator_Validate(t *testing.T) {
	re, err := regexp.Compile(idRe)
	assert.NoError(t, err)
	v := NewExampleValidator(re)
	t.Run("good", func(t *testing.T) {
		data, err := ioutil.ReadFile(good)
		assert.NoError(t, err)
		request := &pb.ExampleRequest{}
		assert.NoError(t, json.Unmarshal(data, request))
		assert.NoError(t, v.Validate(request))
	})
	t.Run("invalid", func(t *testing.T) {
		data, err := ioutil.ReadFile(invalid)
		assert.NoError(t, err)
		request := &pb.ExampleRequest{}
		assert.NoError(t, json.Unmarshal(data, request))
		assert.Error(t, v.Validate(request))
	})
}
