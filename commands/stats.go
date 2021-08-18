package commands

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/discord/colors"
	"github.com/BOOMfinity-Developers/bfcord/other"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"goprodukcji/utils"
	"log"
	"math"
	"runtime"
	"time"
)

var uptime = time.Now()
var GitCommitHash string

func StatsHandler(ctx Context) {
	rss := utils.GetMemory()
	// Golang runtime memory stats
	var rmem runtime.MemStats
	runtime.ReadMemStats(&rmem)
	memory, _ := mem.VirtualMemory()
	pc, _ := host.Info()
	proc, _ := cpu.Info()
	_, err := ctx.Message.Reply(&discord.MessageCreateOptions{
		Embed: &discord.MessageEmbed{
			Title: "GoProdukcji Stats",
			Description: fmt.Sprintf(`Gateway ping: %vms
Wersja: [%v](https://github.com/MrBoombastic/GoProdukcji/commit/%v)
%v
Uptime: %v
RAM (całego serwera): %v / %v (%v%%)

heapInuse / heapTotal: %v / %v
GC: %v
STW: %.2fms
RSS: %v

%v
%v %v (wątków: %v)`, ctx.Client.Manager().AveragePing(), GitCommitHash, GitCommitHash,
				other.Version(), time.Since(uptime).String(), utils.FormatBytes(memory.Used), utils.FormatBytes(memory.Total),
				math.Round(memory.UsedPercent), utils.FormatBytes(rmem.HeapInuse), utils.FormatBytes(rmem.HeapSys-rmem.HeapReleased), rmem.NumGC, float64(time.Duration(rmem.PauseTotalNs))/float64(time.Millisecond), utils.FormatBytes(rss),
				pc.Platform, pc.KernelVersion, proc[0].ModelName, proc[0].Cores),
			Color: colors.Orange,
		}})

	if err != nil {
		log.Fatal(err)
	}
	return "XYZ"
}
