package client

import (
	"fmt"
	"strconv"

	e "github.com/ci4rail/io4edge-cli/cmd/io4edge-cli/internal/errors"
	"github.com/ci4rail/io4edge-client-go/core"
)

func getIpAddressPort(instanceName string, serviceName string, timeout int) (string, error) {
	ipAddress, port, err := core.GetAddressFromService(instanceName, serviceName, timeout)
	e.ErrChk(err)
	ipAddrPort := ipAddress + ":" + strconv.Itoa(int(port))

	return ipAddrPort, err
}

// NewCliClient creates the io4edge core client from the device address
func NewCliClient(device string, service string) (*core.Client, error) {
	ipAddrPort, err := getIpAddressPort(device, service, 3)
	e.ErrChk(err)
	fmt.Println("Try to connect with ", ipAddrPort)
	c, err := core.NewClientFromSocketAddress(ipAddrPort)
	return c, err
}
