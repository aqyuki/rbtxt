package info_test

import (
	"testing"

	"github.com/aqyuki/rbtxt/internal/info"
	"github.com/stretchr/testify/assert"
)

func TestGetApplicationInformation(t *testing.T) {
	info, err := info.GetApplicationInformation()
	assert.NoError(t, err, "should be no error but return error")
	assert.NotNil(t, info, "should be not nil but return nil")
}
