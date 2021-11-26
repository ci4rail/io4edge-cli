/*
Copyright © 2021 Ci4Rail GmbH <engineering@ci4rail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"time"

	"github.com/ci4rail/io4edge-cli/cmd/io4edge-cli/internal/client"
	e "github.com/ci4rail/io4edge-cli/cmd/io4edge-cli/internal/errors"
	api "github.com/ci4rail/io4edge-client-go/core/v1alpha1"
	"github.com/spf13/cobra"
)

var programDeviceIdentificationCmd = &cobra.Command{
	Use:     "program-devid INSTANCE_NAME",
	Aliases: []string{"devid"},
	Short:   "Program default mdns instance name",
	Long: `Program default mdns instance name in the flash of the device.
After a restart of the device, it will show up with this name in the mdns browser.
Example:
io4edge-cli program-devid S101-IOU04-USB-EXT-1`,
	Run:  programDeviceIdentification,
	Args: cobra.ExactArgs(1),
}

func programDeviceIdentification(cmd *cobra.Command, args []string) {
	instanceName := args[0]

	id := &api.ProgramDeviceIdentificationCommand{
		InstanceName: instanceName,
	}

	c, err := client.NewCliClientFromIp(ipAddrPort)
	e.ErrChk(err)

	err = c.ProgramDeviceIdentification(id, time.Duration(timeoutSecs)*time.Second)
	e.ErrChk(err)
}

func init() {
	rootCmd.AddCommand(programDeviceIdentificationCmd)
}
