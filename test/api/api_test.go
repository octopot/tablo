// +build integration

package api_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Example() {
	fmt.Println("API")
	// Output:
	// API
}

func TestSmoke(t *testing.T) {
	require.True(t, true)
}

func TestIntegration(t *testing.T) {
	assert.True(t, true)
}
