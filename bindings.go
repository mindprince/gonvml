/*
Copyright 2017 Google Inc.

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

package gonvml

// #cgo LDFLAGS: -ldl -Wl,--unresolved-symbols=ignore-in-object-files
// #include "nvml/nvml_dl.c"
import "C"

import (
	"errors"
	"fmt"
)

const (
	szDriver = C.NVML_SYSTEM_DRIVER_VERSION_BUFFER_SIZE
	szName   = C.NVML_DEVICE_NAME_BUFFER_SIZE
)

// Initialize initializes NVML.
// Call this before calling any other methods.
func Initialize() error {
	return errorString(C.nvmlInit_dl())
}

// Shutdown shuts down NVML.
// Call this once NVML is no longer being used.
func Shutdown() error {
	return errorString(C.nvmlShutdown_dl())
}

// errorString takes a nvmlReturn_t and converts it into a golang error.
// It uses a nvml method to convert to a user friendly error message.
func errorString(ret C.nvmlReturn_t) error {
	if ret == C.NVML_SUCCESS {
		return nil
	}
	// We need to special case this because if nvml library is not found
	// nvmlErrorString() method will not work.
	if ret == C.NVML_ERROR_LIBRARY_NOT_FOUND {
		return errors.New("could not load NVML library")
	}
	err := C.GoString(C.nvmlErrorString(ret))
	return fmt.Errorf("nvml: %v", err)
}

// SystemDriverVersion returns the the driver version on the system.
func SystemDriverVersion() (string, error) {
	var driver [szDriver]C.char

	r := C.nvmlSystemGetDriverVersion(&driver[0], szDriver)
	return C.GoString(&driver[0]), errorString(r)
}

// DeviceCount returns the number of nvidia devices on the system.
func DeviceCount() (uint, error) {
	var n C.uint

	r := C.nvmlDeviceGetCount(&n)
	return uint(n), errorString(r)
}

// Device is the handle for the device.
// This handle is obtained by calling DeviceHandleByIndex().
type Device struct {
	dev C.nvmlDevice_t
}

// DeviceHandleByIndex returns the device handle for a particular index.
// The indices range from 0 to DeviceCount()-1. The order in which NVML
// enumerates devices has no guarantees of consistency between reboots.
func DeviceHandleByIndex(idx uint) (Device, error) {
	var dev C.nvmlDevice_t

	r := C.nvmlDeviceGetHandleByIndex(C.uint(idx), &dev)
	return Device{dev}, errorString(r)
}

// Name returns the product name of the device.
func (d Device) Name() (string, error) {
	var name [szName]C.char

	r := C.nvmlDeviceGetName(d.dev, &name[0], szName)
	return C.GoString(&name[0]), errorString(r)
}

// MemoryInfo returns the total and used memory (in bytes) of the device.
func (d Device) MemoryInfo() (uint64, uint64, error) {
	var memory C.nvmlMemory_t

	r := C.nvmlDeviceGetMemoryInfo(d.dev, &memory)
	return uint64(memory.total), uint64(memory.used), errorString(r)
}

// UtilizationRates returns the percent of time over the past sample period during which:
// utilization.gpu: one or more kernels were executing on the GPU.
// utilizatoin.memory: global (device) memory was being read or written.
func (d Device) UtilizationRates() (uint, uint, error) {
	var utilization C.nvmlUtilization_t

	r := C.nvmlDeviceGetUtilizationRates(d.dev, &utilization)
	return uint(utilization.gpu), uint(utilization.memory), errorString(r)
}
