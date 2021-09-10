package client

import (
	"github.com/ci4rail/io4edge-client-go/core"
)

func NewCliClient(device string) (*core.Client, error) {
	c, err := core.NewClientFromSocketAddress(device)
	return c, err
}
