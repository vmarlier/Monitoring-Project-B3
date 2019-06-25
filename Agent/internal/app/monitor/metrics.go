package monitor

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/net"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type metric struct {
	Date       time.Time `json:"Date"`
	Hostname   string    `json:"hostname"`
	Ipv4       string    `json:"ipv4"`
	OS         string    `json:"os"`
	Uptime     string    `json:"uptime"`
	Process    string    `json:"process"`
	Memory     memory    `json:"memory"`
	Disks      []dsk     `json:"disk"`
	CPUs       []cp      `json:"cpus"`
	Interfaces []intf    `json:"interfaces"`
}

type memory struct {
	Total    string `json:"total"`
	Free     string `json:"free"`
	UsedPerc string `json:"usedPerc"`
}

type dsk struct {
	Path     string `json:"path"`
	Total    string `json:"total"`
	Free     string `json:"free"`
	Used     string `json:"used"`
	UsedPerc string `json:"usedPerc"`
}

type cp struct {
	Index     int       `json:"index"`
	VendorID  string    `json:"vendorId"`
	Family    string    `json:"family"`
	NbCores   int       `json:"nbCores"`
	Model     string    `json:"model"`
	Speed     float64   `json:"speed"`
	ThreadUse []float64 `json:"threaduse"`
}

type threadUse struct {
	Index    int    `json:"index"`
	UsedPerc string `json:"usedPerc"`
}

type intf struct {
	Name  string   `json:"name"`
	Flags []string `json:"flags"`
}

func dealwithErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func GetMetrics(diskPath []string) []byte {
	h := getHostInfo()
	runtimeOS := runtime.GOOS

	d := []dsk{}
	for _, path := range diskPath {
		d = append(d, getDiskInfo(path))
	}

	c := getCPUInfo()
	n := getNetInfo()

	jsn, _ := json.Marshal(&metric{
		Hostname:   h[0],
		Ipv4:       "0",
		OS:         runtimeOS,
		Uptime:     h[1],
		Process:    h[2],
		Memory:     getMemInfo(),
		Disks:      d,
		CPUs:       c,
		Interfaces: n,
	})

	return jsn
}

func getMemInfo() memory {
	// memory
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)

	total := strconv.FormatUint(vmStat.Total, 10)
	free := strconv.FormatUint(vmStat.Available, 10)
	usedPerc := strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64)

	return memory{Total: total, Free: free, UsedPerc: usedPerc}
}

func getDiskInfo(path string) dsk {
	// disk - start from "/" mount point for Linux
	diskStat, err := disk.Usage("/")
	dealwithErr(err)

	total := strconv.FormatUint(diskStat.Total, 10)
	free := strconv.FormatUint(diskStat.Free, 10)
	used := strconv.FormatUint(diskStat.Used, 10)
	usedPerc := strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64)

	return dsk{Path: path, Total: total, Free: free, Used: used, UsedPerc: usedPerc}
}

func getCPUInfo() []cp {
	// cpu - get CPU number of cores and speed
	cpuStat, err := cpu.Info()
	dealwithErr(err)

	percentage, err := cpu.Percent(0, true)
	dealwithErr(err)

	result := []cp{}

	for i := 0; i < len(cpuStat); i++ {
		result = append(result, cp{
			Index:     int(cpuStat[i].CPU),
			VendorID:  cpuStat[i].VendorID,
			Family:    cpuStat[i].Family,
			NbCores:   int(cpuStat[i].Cores),
			Model:     cpuStat[i].ModelName,
			Speed:     cpuStat[0].Mhz,
			ThreadUse: percentage,
		})
	}

	return result
}

func getNetInfo() (interf []intf) {
	// get interfaces MAC/hardware address
	interfStat, err := net.Interfaces()
	dealwithErr(err)

	for _, itf := range interfStat {
		interf = append(interf, intf{
			Name:  itf.Name,
			Flags: itf.Flags,
		})
	}

	return interf
}

func getHostInfo() [3]string {
	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	dealwithErr(err)

	return [3]string{os.Getenv("HOSTNAME"), strconv.FormatUint(hostStat.Uptime, 10), strconv.FormatUint(hostStat.Procs, 10)}
}
