package data_api

import (
	"blogv2/common/res"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/sirupsen/logrus"
	"time"
)

type ComputerDataResponse struct {
	CpuPercent  float64 `json:"cpuPercent"`
	MemPercent  float64 `json:"memPercent"`
	DiskPercent float64 `json:"diskPercent"`
}

func (DataApi) ComputerDataView(c *gin.Context) {
	var data = ComputerDataResponse{
		CpuPercent:  GetCpuPercent(),
		MemPercent:  GetMemPercent(),
		DiskPercent: GetDiskPercent(),
	}
	res.SuccessWithData(c, data)
}

func GetCpuPercent() float64 {
	cpuPercent, _ := cpu.Percent(time.Second, false)
	return cpuPercent[0]
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	// 获取所有挂载点的磁盘使用率信息
	partitions, err := disk.Partitions(false)
	if err != nil {
		logrus.Errorf("获取磁盘信息错误 %s", err)
		return 0
	}

	var total uint64
	var used uint64
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			logrus.Error("Error getting usage for %s: %v", partition.Mountpoint, err)
			continue
		}

		total += usage.Total
		used += usage.Used
	}
	// 计算总使用率
	usagePercent := float64(used) / float64(total) * 100

	return usagePercent
}
