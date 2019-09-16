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

// #cgo LDFLAGS: -ldl
/*
#include <stddef.h>
#include <dlfcn.h>
#include <stdlib.h>

#include "nvml.h"

// nvmlHandle is the handle for dynamically loaded libnvidia-ml.so
void *nvmlHandle;

nvmlReturn_t (*nvmlInitFunc)(void);

nvmlReturn_t (*nvmlShutdownFunc)(void);

const char* (*nvmlErrorStringFunc)(nvmlReturn_t result);
const char* nvmlErrorString(nvmlReturn_t result) {
  if (nvmlErrorStringFunc == NULL) {
    return "nvmlErrorString Function Not Found";
  }
  return nvmlErrorStringFunc(result);
}

nvmlReturn_t (*nvmlSystemGetDriverVersionFunc)(char *version, unsigned int length);
nvmlReturn_t nvmlSystemGetDriverVersion(char *version, unsigned int length) {
  if (nvmlSystemGetDriverVersionFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlSystemGetDriverVersionFunc(version, length);
}

nvmlReturn_t (*nvmlDeviceGetCountFunc)(unsigned int *deviceCount);
nvmlReturn_t nvmlDeviceGetCount(unsigned int *deviceCount) {
  if (nvmlDeviceGetCountFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetCountFunc(deviceCount);
}

nvmlReturn_t (*nvmlDeviceGetHandleByIndexFunc)(unsigned int index, nvmlDevice_t *device);
nvmlReturn_t nvmlDeviceGetHandleByIndex(unsigned int index, nvmlDevice_t *device) {
  if (nvmlDeviceGetHandleByIndexFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetHandleByIndexFunc(index, device);
}

nvmlReturn_t (*nvmlDeviceGetMinorNumberFunc)(nvmlDevice_t device, unsigned int *minorNumber);
nvmlReturn_t nvmlDeviceGetMinorNumber(nvmlDevice_t device, unsigned int *minorNumber) {
  if (nvmlDeviceGetMinorNumberFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetMinorNumberFunc(device, minorNumber);
}

nvmlReturn_t (*nvmlDeviceGetUUIDFunc)(nvmlDevice_t device, char *uuid, unsigned int length);
nvmlReturn_t nvmlDeviceGetUUID(nvmlDevice_t device, char *uuid, unsigned int length) {
  if (nvmlDeviceGetUUIDFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetUUIDFunc(device, uuid, length);
}

nvmlReturn_t (*nvmlDeviceGetNameFunc)(nvmlDevice_t device, char *name, unsigned int length);
nvmlReturn_t nvmlDeviceGetName(nvmlDevice_t device, char *name, unsigned int length) {
  if (nvmlDeviceGetNameFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetNameFunc(device, name, length);
}

nvmlReturn_t (*nvmlDeviceGetMemoryInfoFunc)(nvmlDevice_t device, nvmlMemory_t *memory);
nvmlReturn_t nvmlDeviceGetMemoryInfo(nvmlDevice_t device, nvmlMemory_t *memory) {
  if (nvmlDeviceGetMemoryInfoFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetMemoryInfoFunc(device, memory);
}

nvmlReturn_t (*nvmlDeviceGetUtilizationRatesFunc)(nvmlDevice_t device, nvmlUtilization_t *utilization);
nvmlReturn_t nvmlDeviceGetUtilizationRates(nvmlDevice_t device, nvmlUtilization_t *utilization) {
  if (nvmlDeviceGetUtilizationRatesFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetUtilizationRatesFunc(device, utilization);
}

nvmlReturn_t (*nvmlDeviceGetPowerUsageFunc)(nvmlDevice_t device, unsigned int *power);
nvmlReturn_t nvmlDeviceGetPowerUsage(nvmlDevice_t device, unsigned int *power) {
  if (nvmlDeviceGetPowerUsageFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetPowerUsageFunc(device, power);
}

nvmlReturn_t (*nvmlDeviceGetTemperatureFunc)(nvmlDevice_t device, nvmlTemperatureSensors_t sensorType, unsigned int *temp);
nvmlReturn_t nvmlDeviceGetTemperature(nvmlDevice_t device, nvmlTemperatureSensors_t sensorType, unsigned int *temp) {
  if (nvmlDeviceGetTemperatureFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetTemperatureFunc(device, sensorType, temp);
}

nvmlReturn_t (*nvmlDeviceGetFanSpeedFunc)(nvmlDevice_t device, unsigned int *speed);
nvmlReturn_t nvmlDeviceGetFanSpeed(nvmlDevice_t device, unsigned int *speed) {
  if (nvmlDeviceGetFanSpeedFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetFanSpeedFunc(device, speed);
}

nvmlReturn_t (*nvmlDeviceGetEncoderUtilizationFunc)(nvmlDevice_t device, unsigned int* utilization, unsigned int* samplingPeriodUs);
nvmlReturn_t nvmlDeviceGetEncoderUtilization(nvmlDevice_t device, unsigned int* utilization, unsigned int* samplingPeriodUs) {
  if (nvmlDeviceGetEncoderUtilizationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetEncoderUtilizationFunc(device, utilization, samplingPeriodUs);
}

nvmlReturn_t (*nvmlDeviceGetDecoderUtilizationFunc)(nvmlDevice_t device, unsigned int* utilization, unsigned int* samplingPeriodUs);
nvmlReturn_t nvmlDeviceGetDecoderUtilization(nvmlDevice_t device, unsigned int* utilization, unsigned int* samplingPeriodUs) {
  if (nvmlDeviceGetDecoderUtilizationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetDecoderUtilizationFunc(device, utilization, samplingPeriodUs);
}

nvmlReturn_t (*nvmlSystemGetProcessNameFunc)(unsigned int pid, char *name, unsigned int length);
nvmlReturn_t nvmlSystemGetProcessName(unsigned int pid, char *name, unsigned int length) {
  if (nvmlSystemGetProcessNameFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlSystemGetProcessNameFunc(pid, name, length);
}

nvmlReturn_t (*nvmlDeviceGetAccountingModeFunc)(nvmlDevice_t device, nvmlEnableState_t *mode);
nvmlReturn_t nvmlDeviceGetAccountingMode(nvmlDevice_t device, nvmlEnableState_t *mode) {
  if (nvmlDeviceGetAccountingModeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetAccountingModeFunc(device, mode);
}

nvmlReturn_t (*nvmlDeviceGetAccountingStatsFunc)(nvmlDevice_t device, unsigned int pid, nvmlAccountingStats_t *stats);
nvmlReturn_t nvmlDeviceGetAccountingStats(nvmlDevice_t device, unsigned int pid, nvmlAccountingStats_t *stats) {
  if (nvmlDeviceGetAccountingStatsFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetAccountingStatsFunc(device, pid, stats);
}


nvmlReturn_t (*nvmlDeviceGetAccountingPidsFunc)(nvmlDevice_t device, unsigned int *count, unsigned int *pids);
nvmlReturn_t nvmlDeviceGetAccountingPids(nvmlDevice_t device, unsigned int *count, unsigned int *pids) {
  if (nvmlDeviceGetAccountingPidsFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetAccountingPidsFunc(device, count, pids);
}

nvmlReturn_t (*nvmlDeviceGetAccountingBufferSizeFunc)(nvmlDevice_t device, unsigned int* bufferSize);
nvmlReturn_t nvmlDeviceGetAccountingBufferSize(nvmlDevice_t device, unsigned int* bufferSize) {
  if (nvmlDeviceGetAccountingBufferSizeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetAccountingBufferSizeFunc(device, bufferSize);
}

nvmlReturn_t (*nvmlDeviceGetProcessUtilizationFunc)(nvmlDevice_t device, nvmlProcessUtilizationSample_t *utilization,
                                              unsigned int *processSamplesCount, unsigned long long lastSeenTimeStamp);
nvmlReturn_t nvmlDeviceGetProcessUtilization(nvmlDevice_t device, nvmlProcessUtilizationSample_t *utilization,
                                              unsigned int *processSamplesCount, unsigned long long lastSeenTimeStamp){
  if (nvmlDeviceGetProcessUtilizationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetProcessUtilizationFunc(device, utilization, processSamplesCount, lastSeenTimeStamp);
}

nvmlReturn_t (*nvmlDeviceGetSamplesFunc)(nvmlDevice_t device, nvmlSamplingType_t type, unsigned long long lastSeenTimeStamp, nvmlValueType_t *sampleValType, unsigned int *sampleCount, nvmlSample_t *samples);

// Loads the "libnvidia-ml.so.1" shared library.
// Loads all symbols needed and initializes NVML.
// Call this before calling any other methods.
nvmlReturn_t nvmlInit_dl(void) {
  nvmlHandle = dlopen("libnvidia-ml.so.1", RTLD_LAZY);
  if (nvmlHandle == NULL) {
    return NVML_ERROR_LIBRARY_NOT_FOUND;
  }
  nvmlInitFunc = dlsym(nvmlHandle, "nvmlInit_v2");
  if (nvmlInitFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlShutdownFunc = dlsym(nvmlHandle, "nvmlShutdown");
  if (nvmlShutdownFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlErrorStringFunc = dlsym(nvmlHandle, "nvmlErrorString");
  if (nvmlErrorStringFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlSystemGetDriverVersionFunc = dlsym(nvmlHandle, "nvmlSystemGetDriverVersion");
  if (nvmlSystemGetDriverVersionFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetCountFunc = dlsym(nvmlHandle, "nvmlDeviceGetCount_v2");
  if (nvmlDeviceGetCountFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetHandleByIndexFunc = dlsym(nvmlHandle, "nvmlDeviceGetHandleByIndex_v2");
  if (nvmlDeviceGetHandleByIndexFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetMinorNumberFunc = dlsym(nvmlHandle, "nvmlDeviceGetMinorNumber");
  if (nvmlDeviceGetMinorNumberFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetUUIDFunc = dlsym(nvmlHandle, "nvmlDeviceGetUUID");
  if (nvmlDeviceGetUUIDFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetNameFunc = dlsym(nvmlHandle, "nvmlDeviceGetName");
  if (nvmlDeviceGetNameFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetMemoryInfoFunc = dlsym(nvmlHandle, "nvmlDeviceGetMemoryInfo");
  if (nvmlDeviceGetMemoryInfoFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetUtilizationRatesFunc = dlsym(nvmlHandle, "nvmlDeviceGetUtilizationRates");
  if (nvmlDeviceGetUtilizationRatesFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetPowerUsageFunc = dlsym(nvmlHandle, "nvmlDeviceGetPowerUsage");
  if (nvmlDeviceGetPowerUsageFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetTemperatureFunc = dlsym(nvmlHandle, "nvmlDeviceGetTemperature");
  if (nvmlDeviceGetTemperatureFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetFanSpeedFunc = dlsym(nvmlHandle, "nvmlDeviceGetFanSpeed");
  if (nvmlDeviceGetFanSpeedFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetSamplesFunc = dlsym(nvmlHandle, "nvmlDeviceGetSamples");
  if (nvmlDeviceGetSamplesFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetEncoderUtilizationFunc = dlsym(nvmlHandle, "nvmlDeviceGetEncoderUtilization");
  if (nvmlDeviceGetEncoderUtilizationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetDecoderUtilizationFunc = dlsym(nvmlHandle, "nvmlDeviceGetDecoderUtilization");
  if (nvmlDeviceGetDecoderUtilizationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlSystemGetProcessNameFunc = dlsym(nvmlHandle, "nvmlSystemGetProcessName");
  if (nvmlSystemGetProcessNameFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetAccountingModeFunc = dlsym(nvmlHandle, "nvmlDeviceGetAccountingMode");
  if (nvmlDeviceGetAccountingModeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetAccountingStatsFunc = dlsym(nvmlHandle, "nvmlDeviceGetAccountingStats");
  if (nvmlDeviceGetAccountingStatsFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetAccountingPidsFunc = dlsym(nvmlHandle, "nvmlDeviceGetAccountingPids");
  if (nvmlDeviceGetAccountingPidsFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetAccountingBufferSizeFunc = dlsym(nvmlHandle, "nvmlDeviceGetAccountingBufferSize");
  if (nvmlDeviceGetAccountingBufferSizeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetProcessUtilizationFunc = dlsym(nvmlHandle, "nvmlDeviceGetProcessUtilization");
  if (nvmlDeviceGetProcessUtilizationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }

  nvmlReturn_t result = nvmlInitFunc();
  if (result != NVML_SUCCESS) {
    dlclose(nvmlHandle);
    nvmlHandle = NULL;
    return result;
  }
  return NVML_SUCCESS;
}

// Shuts down NVML and decrements the reference count on the dynamically loaded
// "libnvidia-ml.so.1" library.
// Call this once NVML is no longer being used.
nvmlReturn_t nvmlShutdown_dl(void) {
  if (nvmlHandle == NULL) {
    return NVML_SUCCESS;
  }
  if (nvmlShutdownFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlReturn_t r = nvmlShutdownFunc();
  if (r != NVML_SUCCESS) {
    return r;
  }
  return (dlclose(nvmlHandle) ? NVML_ERROR_UNKNOWN : NVML_SUCCESS);
}

// This function is here because the API provided by NVML is not very user
// friendly. This function can be used to get average utilization.gpu or
// power.draw.
//
// `device`: The identifier of the target device.
// `type`: Type of sampling event. Only NVML_TOTAL_POWER_SAMPLES and NVML_GPU_UTILIZATION_SAMPLES are supported.
// `lastSeenTimeStamp`: Return average using samples with timestamp greather than this timestamp. Unix epoch in micro seconds.
// `averageUsage`: Reference in which average is returned.
//
// In my experiments, I found that NVML_GPU_UTILIZATION_SAMPLES buffer stores
// 100 samples that are uniformly spread with ~6 samples per second. So the
// buffer stores last ~16s of data.
// NVML_TOTAL_POWER_SAMPLES buffer stores 120 samples, but in different runs I
// noticed them to be non-uniformly separated. Sometimes 120 samples only
// consisted of 10s of data and sometimes they were spread over 60s.
//
nvmlReturn_t nvmlDeviceGetAverageUsage(nvmlDevice_t device, nvmlSamplingType_t type, unsigned long long lastSeenTimeStamp, unsigned int* averageUsage) {
  if (nvmlHandle == NULL) {
    return NVML_ERROR_LIBRARY_NOT_FOUND;
  }
  if (nvmlDeviceGetSamplesFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }

  // We don't really use this because both the metrics we support
  // averagePowerUsage and averageGPUUtilization are unsigned int.
  nvmlValueType_t sampleValType;

  // This will be set to the number of samples that can be queried. We would
  // need to allocate an array of this size to store the samples.
  unsigned int sampleCount;

  // Invoking this method with `samples` set to NULL sets the sampleCount.
  nvmlReturn_t r = nvmlDeviceGetSamplesFunc(device, type, lastSeenTimeStamp, &sampleValType, &sampleCount, NULL);
  if (r != NVML_SUCCESS) {
    return r;
  }

  // Allocate memory to store sampleCount samples.
  // In my experiments, the sampleCount at this stage was always 120 for
  // NVML_TOTAL_POWER_SAMPLES and 100 for NVML_GPU_UTILIZATION_SAMPLES
  nvmlSample_t* samples = (nvmlSample_t*) malloc(sampleCount * sizeof(nvmlSample_t));

  r = nvmlDeviceGetSamplesFunc(device, type, lastSeenTimeStamp, &sampleValType, &sampleCount, samples);
  if (r != NVML_SUCCESS) {
    free(samples);
    return r;
  }

  int i = 0;
  unsigned int sum = 0;
  for (; i < sampleCount; i++) {
    sum += samples[i].sampleValue.uiVal;
  }
  *averageUsage = sum/sampleCount;

  free(samples);
  return r;
}
*/
import "C"

import (
	"errors"
	"fmt"
	"time"
)

const (
	szDriver = C.NVML_SYSTEM_DRIVER_VERSION_BUFFER_SIZE
	szName   = C.NVML_DEVICE_NAME_BUFFER_SIZE
	szUUID   = C.NVML_DEVICE_UUID_BUFFER_SIZE
)

var errLibraryNotLoaded = errors.New("could not load NVML library")

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
	if ret == C.NVML_ERROR_LIBRARY_NOT_FOUND || C.nvmlHandle == nil {
		return errLibraryNotLoaded
	}
	err := C.GoString(C.nvmlErrorString(ret))
	return fmt.Errorf("nvml: %v", err)
}

// SystemDriverVersion returns the the driver version on the system.
func SystemDriverVersion() (string, error) {
	if C.nvmlHandle == nil {
		return "", errLibraryNotLoaded
	}
	var driver [szDriver]C.char
	r := C.nvmlSystemGetDriverVersion(&driver[0], szDriver)
	return C.GoString(&driver[0]), errorString(r)
}

// DeviceCount returns the number of nvidia devices on the system.
func DeviceCount() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetCount(&n)
	return uint(n), errorString(r)
}

// Device is the handle for the device.
// This handle is obtained by calling DeviceHandleByIndex().
type Device struct {
	dev C.nvmlDevice_t
}

// Utilization is Structure to store utilization value and process Id
type Utilization struct {
	Pid       uint   //!< PID of process
	timeStamp uint64 //!< CPU Timestamp in microseconds
	SMUtil    uint   //!< SM (3D/Compute) Util Value
	MemUtil   uint   //!< Frame Buffer Memory Util Value
	EncUtil   uint   //!< Encoder Util Value
	DecUtil   uint   //!< Decoder Util Value
}

// AccountingStats is a Structire to Store accounting Stats for every process
type AccountingStats struct {
	GPUUtilization uint
	//!< Percent of time over the process's lifetime during which one or more kernels was executing on the GPU.
	//! Utilization stats just like returned by \ref nvmlDeviceGetUtilizationRates but for the life time of a
	//! process (not just the last sample period).
	//! Set to NVML_VALUE_NOT_AVAILABLE if nvmlDeviceGetUtilizationRates is not supported

	MemoryUtilization uint
	//!< Percent of time over the process's lifetime during which global (device) memory was being read or written.
	//! Set to NVML_VALUE_NOT_AVAILABLE if nvmlDeviceGetUtilizationRates is not supported

	MaxMemoryUsage uint64
	//!< Maximum total memory in bytes that was ever allocated by the process.
	//! Set to NVML_VALUE_NOT_AVAILABLE if nvmlProcessInfo_t->usedGpuMemory is not supported

	Time uint64
	//!< Amount of time in ms during which the compute context was active. The time is reported as 0 if
	//!< the process is not terminated

	StartTime uint64
	//!< CPU Timestamp in usec representing start time for the process

	IsRunning bool
	//!< Flag to represent if the process is running (1 for running, 0 for terminated)

	Reserved [5]uint
	// Reserved for
}

// DeviceHandleByIndex returns the device handle for a particular index.
// The indices range from 0 to DeviceCount()-1. The order in which NVML
// enumerates devices has no guarantees of consistency between reboots.
func DeviceHandleByIndex(idx uint) (Device, error) {
	if C.nvmlHandle == nil {
		return Device{}, errLibraryNotLoaded
	}
	var dev C.nvmlDevice_t
	r := C.nvmlDeviceGetHandleByIndex(C.uint(idx), &dev)
	return Device{dev}, errorString(r)
}

// MinorNumber returns the minor number for the device.
// The minor number for the device is such that the Nvidia device node
// file for each GPU will have the form /dev/nvidia[minor number].
func (d Device) MinorNumber() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetMinorNumber(d.dev, &n)
	return uint(n), errorString(r)
}

// UUID returns the globally unique immutable UUID associated with this device.
func (d Device) UUID() (string, error) {
	if C.nvmlHandle == nil {
		return "", errLibraryNotLoaded
	}
	var uuid [szUUID]C.char
	r := C.nvmlDeviceGetUUID(d.dev, &uuid[0], szUUID)
	return C.GoString(&uuid[0]), errorString(r)
}

// Name returns the product name of the device.
func (d Device) Name() (string, error) {
	if C.nvmlHandle == nil {
		return "", errLibraryNotLoaded
	}
	var name [szName]C.char
	r := C.nvmlDeviceGetName(d.dev, &name[0], szName)
	return C.GoString(&name[0]), errorString(r)
}

// MemoryInfo returns the total and used memory (in bytes) of the device.
func (d Device) MemoryInfo() (uint64, uint64, error) {
	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var memory C.nvmlMemory_t
	r := C.nvmlDeviceGetMemoryInfo(d.dev, &memory)
	return uint64(memory.total), uint64(memory.used), errorString(r)
}

// UtilizationRates returns the percent of time over the past sample period during which:
// utilization.gpu: one or more kernels were executing on the GPU.
// utilization.memory: global (device) memory was being read or written.
func (d Device) UtilizationRates() (uint, uint, error) {
	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var utilization C.nvmlUtilization_t
	r := C.nvmlDeviceGetUtilizationRates(d.dev, &utilization)
	return uint(utilization.gpu), uint(utilization.memory), errorString(r)
}

// PowerUsage returns the power usage for this GPU and its associated circuitry
// in milliwatts. The reading is accurate to within +/- 5% of current power draw.
func (d Device) PowerUsage() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetPowerUsage(d.dev, &n)
	return uint(n), errorString(r)
}

// AveragePowerUsage returns the power usage for this GPU and its associated circuitry
// in milliwatts averaged over the samples collected in the last `since` duration.
func (d Device) AveragePowerUsage(since time.Duration) (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	lastTs := C.ulonglong(time.Now().Add(-1*since).UnixNano() / 1000)
	var n C.uint
	r := C.nvmlDeviceGetAverageUsage(d.dev, C.NVML_TOTAL_POWER_SAMPLES, lastTs, &n)
	return uint(n), errorString(r)
}

// AverageGPUUtilization returns the utilization.gpu metric (percent of time
// one of more kernels were executing on the GPU) averaged over the samples
// collected in the last `since` duration.
func (d Device) AverageGPUUtilization(since time.Duration) (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	lastTs := C.ulonglong(time.Now().Add(-1*since).UnixNano() / 1000)
	var n C.uint
	r := C.nvmlDeviceGetAverageUsage(d.dev, C.NVML_GPU_UTILIZATION_SAMPLES, lastTs, &n)
	return uint(n), errorString(r)
}

// Temperature returns the temperature for this GPU in Celsius.
func (d Device) Temperature() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetTemperature(d.dev, C.NVML_TEMPERATURE_GPU, &n)
	return uint(n), errorString(r)
}

// FanSpeed returns the temperature for this GPU in the percentage of its full
// speed, with 100 being the maximum.
func (d Device) FanSpeed() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetFanSpeed(d.dev, &n)
	return uint(n), errorString(r)
}

// EncoderUtilization returns the percent of time over the last sample period during which the GPU video encoder was being used.
// The sampling period is variable and is returned in the second return argument in microseconds.
func (d Device) EncoderUtilization() (uint, uint, error) {
	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var n C.uint
	var sp C.uint
	r := C.nvmlDeviceGetEncoderUtilization(d.dev, &n, &sp)
	return uint(n), uint(sp), errorString(r)
}

// DecoderUtilization returns the percent of time over the last sample period during which the GPU video decoder was being used.
// The sampling period is variable and is returned in the second return argument in microseconds.
func (d Device) DecoderUtilization() (uint, uint, error) {
	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var n C.uint
	var sp C.uint
	r := C.nvmlDeviceGetDecoderUtilization(d.dev, &n, &sp)
	return uint(n), uint(sp), errorString(r)
}

// DeviceGetAccountingMode Queries process's accounting stats
// @return mode                                 Reference in which to return the current accounting mode
func (d Device) AccountingMode() (C.nvmlEnableState_t, error) {
	var stats C.nvmlEnableState_t
	if C.nvmlHandle == nil {
		return stats, errLibraryNotLoaded
	}
	r := C.nvmlDeviceGetAccountingMode(d.dev, &stats)
	return stats, errorString(r)
}

// DeviceGetAccountingStats Queries process's accounting stats.
// @param pid                                  Process Id of the target process to query stats for
// @return stats                               Reference in which to return the process's accounting stats
func (d Device) AccountingStats(pid uint) (*AccountingStats, error) {
	if C.nvmlHandle == nil {
		return nil, errLibraryNotLoaded
	}
	var stats C.nvmlAccountingStats_t
	r := C.nvmlDeviceGetAccountingStats(d.dev, C.uint(pid), &stats)

	accountingStats := &AccountingStats{
		GPUUtilization:    uint(stats.gpuUtilization),
		MemoryUtilization: uint(stats.memoryUtilization),
		MaxMemoryUsage:    uint64(stats.maxMemoryUsage),
		Time:              uint64(stats.time),
		StartTime:         uint64(stats.startTime),
		IsRunning:         uint(stats.isRunning) == 1,
	}

	return accountingStats, errorString(r)
}

// DeviceGetAccountingPids Queries list of processes that can be queried for accounting stats. The list of processes returned
// @param count                                Maxnum pids
// @return pids                                Pids result
// @return count                               Queried pids num
func (d Device) AccountingPids(count uint) ([]C.uint, uint, error) {
	// init pids
	cCount := C.uint(count)
	if C.nvmlHandle == nil {
		return nil, 0, errLibraryNotLoaded
	}
	if count == 0 {
		r := C.nvmlDeviceGetAccountingPids(d.dev, &cCount, nil)
		return nil, uint(cCount), errorString(r)
	}

	pids := make([]C.uint, count)
	for index := range pids {
		pids[index] = 0
	}

	r := C.nvmlDeviceGetAccountingPids(d.dev, &cCount, &pids[0])
	return pids, uint(cCount), errorString(r)
}

// DeviceGetAccountingBufferSize Returns the number of processes that the circular buffer with accounting pids can hold.
// @return buffersize                        buffersize
func (d Device) AccountingBufferSize() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var bufferSize C.uint
	r := C.nvmlDeviceGetAccountingBufferSize(d.dev, &bufferSize)
	return uint(bufferSize), errorString(r)
}

// DeviceGetProcessUtilization Retrieves the current utilization and process ID
// @param processCount                      Maxnum process buffersize
// @param since                             The last query time for process
// @return utilizations                     The utilizations for all process
// @return processCount                     The queried utilizations
func (d Device) ProcessUtilization(processCount uint, since time.Duration) ([]*Utilization, error) {
	if C.nvmlHandle == nil {
		return nil, errLibraryNotLoaded
	}
	if processCount <= 0 {
		return nil, errors.New("Process Count Less than zero")
	}

	cUtilizations := make([]C.nvmlProcessUtilizationSample_t, processCount)
	var runningProcess C.uint = C.uint(processCount * C.sizeof_nvmlProcessUtilizationSample_t)

	lastTS := C.ulonglong(time.Now().Add(-1*since).UnixNano() / 1000)
	r := C.nvmlDeviceGetProcessUtilization(d.dev, &cUtilizations[0], &runningProcess, lastTS)
	if errorString(r) != nil {
		return nil, errorString(r)
	}

	statisticsProcess := uint(runningProcess)
	if processCount < uint(runningProcess) {
		statisticsProcess = processCount
	}

	utilizations := make([]*Utilization, statisticsProcess)
	utilCount := 0
	for _, utilization := range cUtilizations[:statisticsProcess] {
		if utilization.pid <= 0 {
			continue
		}
		u := &Utilization{
			Pid:       uint(utilization.pid),
			timeStamp: uint64(utilization.timeStamp),
			SMUtil:    uint(utilization.smUtil),
			MemUtil:   uint(utilization.memUtil),
			EncUtil:   uint(utilization.encUtil),
			DecUtil:   uint(utilization.decUtil),
		}
		utilizations[utilCount] = u
		utilCount++
	}

	return utilizations[:utilCount], errorString(r)
}

// SystemGetProcessName GetProcessName by pid
// @param pid                      Process's id
// @param buffersize               The process name's buffersize
// @return name                    Process name
func SystemGetProcessName(pid, buffersize uint) (string, error) {
	if C.nvmlHandle == nil {
		return "", errLibraryNotLoaded
	}
	c := make([]C.char, buffersize)
	r := C.nvmlSystemGetProcessName(C.uint(pid), &c[0], C.uint(buffersize))
	return C.GoString(&c[0]), errorString(r)
}
