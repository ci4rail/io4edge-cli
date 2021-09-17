package client

import (
	"github.com/ci4rail/io4edge-client-go/core"
)

// NewCliClient creates the io4edge core client from the device address
func NewCliClient(device string) (*core.Client, error) {
	c, err := core.NewClientFromSocketAddress(device)
	return c, err
}
