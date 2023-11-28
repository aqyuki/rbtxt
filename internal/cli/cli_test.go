package cli_test

import (
	"testing"

	"github.com/aqyuki/rbtxt/internal/cli"
	"github.com/stretchr/testify/assert"
)

func TestCreateDefaultCLI(t *testing.T) {
	actual := cli.CreateDefaultCLI()
	assert.NotNil(t, actual, "should return CLI instance but return nil")
}
