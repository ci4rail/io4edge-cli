package client

import (
	"time"

	"github.com/ci4rail/io4edge-client-go/core"
)

// NewCliClient creates the io4edge core client from the device address
func NewCliClient(serviceAddr string) (*core.Client, error) {
	c, err := core.NewClientFromService(serviceAddr, time.Duration(1)*time.Second)
	return c, err
}
