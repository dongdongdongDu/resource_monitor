package mem

import (
	"github.com/shirou/gopsutil/mem"
	"log"
	"math"
	"resource_monitor/getInfo/disk"
)

type MemInfo struct {
	Total       string  `json:"total"`
	Available   string  `json:"available"`
	Used        string  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
	Free        string  `json:"free"`
}

func GetMemInfo() *MemInfo {
	var meminfo MemInfo
	memInfos, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("get mem info failed, err: %v\n", err)
	} else {
		meminfo.Total = disk.FormatCap(memInfos.Total)
		meminfo.Available = disk.FormatCap(memInfos.Available)
		meminfo.Used = disk.FormatCap(memInfos.Used)
		meminfo.UsedPercent = math.Round(memInfos.UsedPercent)
		meminfo.Free = disk.FormatCap(memInfos.Free)
	}
	return &meminfo
}

func UpdateMemInfo(memInfos *MemInfo) *MemInfo {
	newMemInfos, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("update mem info failed, err: %v\n", err)
	} else {
		memInfos.Available = disk.FormatCap(newMemInfos.Available)
		memInfos.Used = disk.FormatCap(newMemInfos.Used)
		memInfos.UsedPercent = math.Round(newMemInfos.UsedPercent)
		memInfos.Free = disk.FormatCap(newMemInfos.Free)
	}
	return memInfos
}
func GetSwapMemInfo() *MemInfo {
	var meminfo MemInfo
	memInfos, err := mem.SwapMemory()
	if err != nil {
		log.Printf("get mem info failed, err: %v\n", err)
	} else {
		meminfo.Total = disk.FormatCap(memInfos.Total)
		meminfo.Available = disk.FormatCap(memInfos.Free)
		meminfo.Used = disk.FormatCap(memInfos.Used)
		meminfo.UsedPercent = math.Round(memInfos.UsedPercent)
		meminfo.Free = disk.FormatCap(memInfos.Free)
	}
	return &meminfo
}

func UpdateSwapMemInfo(memInfos *MemInfo) *MemInfo {
	newMemInfos, err := mem.SwapMemory()
	if err != nil {
		log.Printf("update mem info failed, err: %v\n", err)
	} else {
		memInfos.Used = disk.FormatCap(newMemInfos.Used)
		memInfos.Available = disk.FormatCap(newMemInfos.Free)
		memInfos.UsedPercent = math.Round(newMemInfos.UsedPercent)
		memInfos.Free = disk.FormatCap(newMemInfos.Free)
	}
	return memInfos
}
