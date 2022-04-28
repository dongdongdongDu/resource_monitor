package main

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"log"
	"time"
)

type MemInfo struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
	Free        uint64  `json:"free"`
}

func getCpu() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		log.Printf("get cpu info failed, err: %v\n", err)
	} else {
		fmt.Println(cpuInfos)
	}

	percentall, _ := cpu.Percent(time.Second, false)
	percent, _ := cpu.Percent(time.Second, true)

	fmt.Println(percentall)
	fmt.Println(percent)

}
func getMem() *MemInfo {
	var meminfo MemInfo
	memInfos, err := mem.VirtualMemory()
	fmt.Println(memInfos)
	if err != nil {
		log.Printf("get mem info failed, err: %v\n", err)
	} else {
		meminfo.Total = memInfos.Total
		meminfo.Available = memInfos.Available
		meminfo.Used = memInfos.Used
		meminfo.UsedPercent = memInfos.UsedPercent
		meminfo.Free = memInfos.Free
	}
	return &meminfo
}

func getDisk() {
	diskstat, _ := disk.Partitions(true)
	fmt.Println(diskstat)
	for _, stat := range diskstat {
		mapStat, _ := disk.Usage(stat.Device)
		fmt.Println(mapStat)
	}

	//for name, stat := range mapStat {
	//	fmt.Println(name)
	//	data, _ := json.MarshalIndent(stat, "", "  ")
	//	fmt.Println(string(data))
	//}
}

func getHost() {
	timestamp, _ := host.BootTime()
	t := time.Unix(int64(timestamp), 0)
	fmt.Println(t.Local().Format("2006-01-02 15:04:05"))
	fmt.Println(timestamp)
	users, err := host.Users()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
	for _, user := range users {
		data, _ := json.MarshalIndent(user, "", " ")
		fmt.Println(string(data))
	}

	version, _ := host.KernelVersion()
	fmt.Println("version:", version)

	platform, family, version, _ := host.PlatformInformation()
	fmt.Println("platform:", platform) // 操作系统信息
	fmt.Println("family:", family)
	fmt.Println("version:", version)

}

func getProcess() {
	processes, _ := process.Processes()
	fmt.Println(processes)
}

func main() {
	//fmt.Println(getMem())
	//getDisk()
	//getCpu()
	//getHost()
	getProcess()
}
