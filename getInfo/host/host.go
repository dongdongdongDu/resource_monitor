package host

import (
	"github.com/shirou/gopsutil/host"
	"log"
	"time"
)

type HostInfo struct {
	BootTime string `json:"bootTime"`
	Users    int    `json:"users"`
	Kernel   string `json:"kernel"`
	Platform string `json:"platform"`
	Family   string `json:"family"`
	Version  string `json:"version"`
}

func GetHostInfo() (hostinfo HostInfo) {
	timestamp, err := host.BootTime()
	if err != nil {
		log.Printf("get host bootTime info err:%v", err)
	} else {
		t := time.Unix(int64(timestamp), 0)
		hostinfo.BootTime = t.Local().Format("2006-01-02 15:04:05")
	}

	users, err := host.Users()
	if err != nil {
		log.Printf("get host users info err:%v", err)
	} else {
		hostinfo.Users = len(users)
	}

	kernel, err := host.KernelVersion()
	if err != nil {
		log.Printf("get host kernel info err:%v", err)
	} else {
		hostinfo.Kernel = kernel
	}

	platform, family, version, err := host.PlatformInformation()
	if err != nil {
		log.Printf("get host platform info err:%v", err)
	} else {
		hostinfo.Platform = platform
		hostinfo.Family = family
		hostinfo.Version = version
	}
	return
}
