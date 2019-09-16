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

package main

import (
	"fmt"
	"time"

	"github.com/mindprince/gonvml"
)

func main() {
	start := time.Now()
	err := gonvml.Initialize()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer gonvml.Shutdown()
	fmt.Printf("Initialize() took %v\n", time.Since(start))

	driverVersion, err := gonvml.SystemDriverVersion()
	if err != nil {
		fmt.Printf("SystemDriverVersion() error: %v\n", err)
		return
	}
	fmt.Printf("SystemDriverVersion(): %v\n", driverVersion)

	numDevices, err := gonvml.DeviceCount()
	if err != nil {
		fmt.Printf("DeviceCount() error: %v\n", err)
		return
	}
	fmt.Printf("DeviceCount(): %v\n", numDevices)

	for i := 0; i < int(numDevices); i++ {
		dev, err := gonvml.DeviceHandleByIndex(uint(i))
		if err != nil {
			fmt.Printf("\tDeviceHandleByIndex() error: %v\n", err)
			return
		}

		minorNumber, err := dev.MinorNumber()
		if err != nil {
			fmt.Printf("\tdev.MinorNumber() error: %v\n", err)
			return
		}
		fmt.Printf("\tminorNumber: %v\n", minorNumber)

		uuid, err := dev.UUID()
		if err != nil {
			fmt.Printf("\tdev.UUID() error: %v\n", err)
			return
		}
		fmt.Printf("\tuuid: %v\n", uuid)

		name, err := dev.Name()
		if err != nil {
			fmt.Printf("\tdev.Name() error: %v\n", err)
			return
		}
		fmt.Printf("\tname: %v\n", name)

		totalMemory, usedMemory, err := dev.MemoryInfo()
		if err != nil {
			fmt.Printf("\tdev.MemoryInfo() error: %v\n", err)
			return
		}
		fmt.Printf("\tmemory.total: %v, memory.used: %v\n", totalMemory, usedMemory)

		gpuUtilization, memoryUtilization, err := dev.UtilizationRates()
		if err != nil {
			fmt.Printf("\tdev.UtilizationRates() error: %v\n", err)
			return
		}
		fmt.Printf("\tutilization.gpu: %v, utilization.memory: %v\n", gpuUtilization, memoryUtilization)

		powerDraw, err := dev.PowerUsage()
		if err != nil {
			fmt.Printf("\tdev.PowerUsage() error: %v\n", err)
			return
		}
		fmt.Printf("\tpower.draw: %v\n", powerDraw)

		averagePowerDraw, err := dev.AveragePowerUsage(10 * time.Second)
		if err != nil {
			fmt.Printf("\tdev.AveragePowerUsage() error: %v\n", err)
			return
		}
		fmt.Printf("\taverage power.draw for last 10s: %v\n", averagePowerDraw)

		averageGPUUtilization, err := dev.AverageGPUUtilization(10 * time.Second)
		if err != nil {
			fmt.Printf("\tdev.AverageGPUUtilization() error: %v\n", err)
			return
		}
		fmt.Printf("\taverage utilization.gpu for last 10s: %v\n", averageGPUUtilization)

		temperature, err := dev.Temperature()
		if err != nil {
			fmt.Printf("\tdev.Temperature() error: %v\n", err)
			return
		}
		fmt.Printf("\ttemperature.gpu: %v C\n", temperature)

		fanSpeed, err := dev.FanSpeed()
		if err != nil {
			fmt.Printf("\tdev.FanSpeed() error: %v\n", err)
			return
		}
		fmt.Printf("\tfan.speed: %v%%\n", fanSpeed)

		encoderUtilization, _, err := dev.EncoderUtilization()
		if err != nil {
			fmt.Printf("\tdev.EncoderUtilization() error: %v\n", err)
			return
		}
		fmt.Printf("\tutilization.encoder: %d\n", encoderUtilization)

		decoderUtilization, _, err := dev.DecoderUtilization()
		if err != nil {
			fmt.Printf("\tdev.DecoderUtilization() error: %v\n", err)
			return
		}
		fmt.Printf("\tutilization.decoder: %d\n", decoderUtilization)

		modeStats, err := dev.AccountingMode()
		if err != nil {
			fmt.Printf("\tdev.DeviceGetAccountingMode() error: %v\n", err)
			return
		}
		fmt.Printf("\taccounting.mode enable: %v\n", modeStats)

		bufferSize, err := dev.AccountingBufferSize()
		if err != nil {
			fmt.Printf("\tdev.DeviceGetAccountingBufferSize() error: %v\n", err)
			return
		}
		fmt.Printf("\taccounting.buffersize: %d\n", bufferSize)

		pids, count, err := dev.AccountingPids(bufferSize)
		if err != nil {
			fmt.Printf("\tdev.DeviceGetAccountingPids() error: %v\n", err)
		} else {
			fmt.Printf("\taccounting.pids.count: %v\n", count)
			for _, pid := range pids[:count] {
				fmt.Printf("\t\tPid: %v", pid)
				stats, err := dev.AccountingStats(uint(pid))
				if err != nil {
					fmt.Printf("\tdev.DeviceGetAccountingStats() error: %v\n", err)
				} else {
					fmt.Printf(", GPUUtilization: %v", stats.GPUUtilization)
					fmt.Printf(", MemoryUtilization: %v", stats.MemoryUtilization)
					fmt.Printf(", MaxMemoryUsage: %v", stats.MaxMemoryUsage)
					fmt.Printf(", Time: %v", stats.Time)
					fmt.Printf(", StartTime: %v", stats.StartTime)
					fmt.Printf(", IsRunning: %v", stats.IsRunning)
					fmt.Println()
				}
			}
		}

		utilizations, err := dev.ProcessUtilization(10, 10*time.Second)
		if err != nil {
			fmt.Printf("\tdev.DeviceGetProcessUtilization() error: %v\n", err)
		} else {
			fmt.Printf("\tProcess count: %v\n", len(utilizations))

			utilizations = utilizations
			for _, sample := range utilizations {
				fmt.Printf("\t\tProcess: %v", sample.Pid)
				fmt.Printf(", SM  util: %v", sample.SMUtil)
				fmt.Printf(", Mem util: %v", sample.MemUtil)
				fmt.Printf(", Enc util: %v", sample.EncUtil)
				fmt.Printf(", Dec util: %v", sample.DecUtil)

				name, err := gonvml.SystemGetProcessName(sample.Pid, 64)
				if err != nil {
					fmt.Printf("\n\tdev.SystemGetProcessName() error: %v\n", err)
				} else {
					fmt.Printf(", Name: %s\n", name)
				}
			}
		}

		fmt.Println()
	}
}
