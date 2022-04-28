package cpu

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"log"
	"math"
	"time"
)

type CpuInfo struct {
	CPU       int32   `json:"cpu"`
	Cores     int32   `json:"cores"`
	ModelName string  `json:"modelName"`
	Mhz       float64 `json:"mhz"`
	Percent   float64 `json:"percent"`
}

func GetCpuInfo() (CpuInfos []CpuInfo) {
	var cpuinfo CpuInfo
	cpuInfos, err := cpu.Info()
	if err != nil {
		log.Printf("get cpu info failed, err: %v\n", err)
	} else {
		for _, ci := range cpuInfos {
			cpuinfo.CPU = ci.CPU
			cpuinfo.Cores = ci.Cores
			cpuinfo.ModelName = ci.ModelName
			cpuinfo.Mhz = ci.Mhz
		}
	}
	CpuInfos = append(CpuInfos, cpuinfo)
	return
}

func UpdateCpuInfo(cpuInfos []CpuInfo) []CpuInfo {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Printf("update cpu percent failed, err: %v\n", err)
	} else {
		if len(cpuInfos) != len(percent) {
			log.Println("update cpu percent failed, err: len(cpuInfos) != len(percent)")
		} else {
			for i := 0; i < len(cpuInfos); i++ {
				cpuInfos[i].Percent = math.Round(percent[i])
			}
		}
	}
	return cpuInfos
}
func GetCpuPercent() (cpuPercent []map[string]float64) {
	percent, err := cpu.Percent(time.Second, true)
	if err != nil {
		log.Printf("update cpu percent failed, err: %v\n", err)
	} else {
		for i, v := range percent {
			cpuPer := make(map[string]float64)
			cpuPer[fmt.Sprintf("cpu%v", i)] = math.Round(v)
			cpuPercent = append(cpuPercent, cpuPer)
		}
	}
	return
}
