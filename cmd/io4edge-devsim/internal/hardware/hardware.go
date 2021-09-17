/*
Copyright © 2021 Ci4Rail GmbH
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

package hardware

import (
	"fmt"

	api "github.com/ci4rail/io4edge-client-go/core/v1alpha1"
)

type serialNumber struct {
	Hi uint64
	Lo uint64
}

type hardwareID struct {
	RootArticle  string
	SerialNumber serialNumber
	MajorVersion uint32
}

var (
	hwID = hardwareID{"S101-CPU01UC", serialNumber{0x1234567887654321, 0x4567456745674567}, 1}
)

// IdentifyHardware reports the current hardware inventory data
func IdentifyHardware() *api.CoreResponse {

	res := &api.CoreResponse{
		Id:            api.CommandId_IDENTIFY_HARDWARE,
		Status:        api.Status_OK,
		RestartingNow: false,
		Data: &api.CoreResponse_IdentifyHardware{
			IdentifyHardware: &api.IdentifyHardwareResponse{
				RootArticle: hwID.RootArticle,
				SerialNumber: &api.SerialNumber{
					Hi: hwID.SerialNumber.Hi,
					Lo: hwID.SerialNumber.Lo,
				},
				MajorVersion: hwID.MajorVersion,
			},
		},
	}
	return res
}

// ProgramHardwareIdentification instructs the device to program the new hardware identification in c
func ProgramHardwareIdentification(c *api.ProgramHardwareIdentificationCommand) *api.CoreResponse {
	// TODO: Check signature
	fmt.Printf("ProgramHardwareIdentification %s\n", c.RootArticle)
	hwID.RootArticle = c.RootArticle
	hwID.SerialNumber.Hi = c.SerialNumber.Hi
	hwID.SerialNumber.Lo = c.SerialNumber.Lo
	hwID.MajorVersion = c.MajorVersion

	res := &api.CoreResponse{
		Id:            api.CommandId_IDENTIFY_HARDWARE,
		Status:        api.Status_OK,
		RestartingNow: false,
	}
	return res
}
