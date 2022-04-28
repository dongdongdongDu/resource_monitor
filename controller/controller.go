package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource_monitor/getInfo/cpu"
	"resource_monitor/getInfo/disk"
	"resource_monitor/getInfo/host"
	"resource_monitor/getInfo/mem"
)

type Info struct {
	Name  string
	Infos []cpu.CpuInfo
}

func InitInfo() (cpuInfo []cpu.CpuInfo) {
	cpuInfo = cpu.GetCpuInfo()
	return cpuInfo
}

//func UpdateInfo(cpuInfo []cpu.CpuInfo) (infos []cpu.CpuInfo) {
//	infos = cpu.UpdateCpuInfo(cpuInfo)
//	//b, err := json.Marshal(cpuInfo)
//	//if err != nil {
//	//	log.Printf("json.Marshal err: %v\n", err)
//	//}
//	return
//}
func UpdateInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"cpu": cpu.UpdateCpuInfo(cpu.GetCpuInfo()),
			//"mem":     mem.UpdateMemInfo(mem.GetMemInfo()),
			//"swapmem": mem.UpdateSwapMemInfo(mem.GetSwapMemInfo()),
			//"disk":    disk.UpdateDiskInfo(disk.GetDiskInfo()),
			//"cpu":     cpu.GetCpuInfo(),
			"cpuper":  cpu.GetCpuPercent(),
			"mem":     mem.GetMemInfo(),
			"swapmem": mem.GetSwapMemInfo(),
			"disk":    disk.GetDiskInfo(),
			"host":    host.GetHostInfo(),
		})
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": string(UpdateInfo(InitInfo())),
	//})
}
