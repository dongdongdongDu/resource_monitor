package disk

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"math"
	"strconv"
)

type DiskInfo struct {
	Path        string  `json:"path"`
	Total       string  `json:"total"`
	Free        string  `json:"free"`
	Used        string  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
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
		mapStat, _ := disk.Usage(stat.Device)
		diskinfo.Path = mapStat.Path
		diskinfo.Total = FormatCap(mapStat.Total)
		diskinfo.Used = FormatCap(mapStat.Used)
		diskinfo.Free = FormatCap(mapStat.Free)
		diskinfo.UsedPercent = math.Round(mapStat.UsedPercent)
		DiskInfos = append(DiskInfos, diskinfo)
	}
	return
}

func UpdateDiskInfo(diskInfos []DiskInfo) []DiskInfo {
	diskstat, _ := disk.Partitions(true)
	for _, stat := range diskstat {
		mapStat, _ := disk.Usage(stat.Device)
		for i := 0; i < len(diskInfos); i++ {
			diskInfos[i].Used = FormatCap(mapStat.Used)
			diskInfos[i].Free = FormatCap(mapStat.Free)
			diskInfos[i].UsedPercent = math.Round(mapStat.UsedPercent)
		}
	}
	return diskInfos
}
