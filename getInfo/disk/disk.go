package disk

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"math"
	"strconv"
)

type DiskInfo struct {
	Device      string  `json:"device"`
	Total       string  `json:"total"`
	Free        string  `json:"free"`
	Used        string  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
	Mountpoint  string  `json:"mountpoint"`
}

func FormatCap(x uint64) (cap string) {

	m := float64(x / 1024 / 1024)
	g := m / 1024
	g, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", g), 64)
	if g >= 1 {
		cap = fmt.Sprintf("%gG", g)
	} else {
		cap = fmt.Sprintf("%.0fM", m)
	}
	return
}

func GetDiskInfo() (DiskInfos []DiskInfo) {
	var diskinfo DiskInfo
	diskstat, _ := disk.Partitions(true)
	for _, stat := range diskstat {
		device := stat.Device
		mapStat, _ := disk.Usage(stat.Mountpoint)
		if mapStat != nil && mapStat.Total != 0 {
			diskinfo.Device = device
			diskinfo.Total = FormatCap(mapStat.Total)
			diskinfo.Used = FormatCap(mapStat.Used)
			diskinfo.Free = FormatCap(mapStat.Free)
			diskinfo.UsedPercent = math.Round(mapStat.UsedPercent)
			diskinfo.Mountpoint = mapStat.Path
			DiskInfos = append(DiskInfos, diskinfo)
		}
	}
	return
}

func UpdateDiskInfo(diskInfos []DiskInfo) []DiskInfo {
	diskstat, _ := disk.Partitions(false)
	for _, stat := range diskstat {
		mapStat, _ := disk.Usage(stat.Mountpoint)
		for i := 0; i < len(diskInfos); i++ {
			diskInfos[i].Used = FormatCap(mapStat.Used)
			diskInfos[i].Free = FormatCap(mapStat.Free)
			diskInfos[i].UsedPercent = math.Round(mapStat.UsedPercent)
		}
	}
	return diskInfos
}
