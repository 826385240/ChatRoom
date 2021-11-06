package benchmark

import (
	"flag"
	"ChatRoom/src/lib/logger"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
)

const (
	GO_PPROF_CPU_PROFILE = 0x1
	GO_PPROF_MEM_PROFILE = 0x2
)

type Gopprof struct {
	flags uint64
}

func NewGopprof() *Gopprof {
	return &Gopprof{}
}

func (this *Gopprof) WebProfile(ip string, port string) {
	logger.FATALLN(http.ListenAndServe(ip+":"+port, nil))
}

func (this *Gopprof) LocalProfile() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

	flag.Parse()
	//初始化cpu分析文件
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			logger.FATAL("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			logger.FATAL("could not start CPU profile: ", err)
		} else {
			logger.DEBUG("Sucess,start CPU profile!")
		}

		this.flags = this.flags | GO_PPROF_CPU_PROFILE
	}

	//初始化内存分析文件
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			logger.FATAL("could not create memory profile: ", err)
		}
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			logger.FATAL("could not write memory profile: ", err)
		} else {
			logger.DEBUG("Sucess,start memory profile!")
		}

		this.flags = this.flags | GO_PPROF_MEM_PROFILE
		f.Close()
	}
}

func (this *Gopprof) DeferLocalProfile() {
	if this.flags&GO_PPROF_CPU_PROFILE > 0 {
		pprof.StopCPUProfile()
	}
}

var g_gopprof *Gopprof = NewGopprof()

func WebProfile(ip string, port string) {
	g_gopprof.WebProfile(ip, port)
}

func LocalProfile() {
	g_gopprof.LocalProfile()
}

func DeferLocalProfile() {
	g_gopprof.DeferLocalProfile()
}
