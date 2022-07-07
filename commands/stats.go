package commands

import (
	"fmt"
	"github.com/BOOMfinity/bfcord/discord"
	"github.com/BOOMfinity/bfcord/discord/colors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"goprodukcji/utils"
	"math"
	"runtime"
	"runtime/debug"
	"time"
)

var uptime = time.Now()
var GitCommitHash string

var StatsCommand = CommandData{
	Command:     runStats,
	Description: "wyświetla statystyki oraz ping bota",
	Usage:       "",
	Aliases:     []string{"ping"},
}

func runStats(ctx Context) {
	rss := utils.GetMemory()
	// Golang runtime memory stats
	var rmem runtime.MemStats
	runtime.ReadMemStats(&rmem)
	memory, _ := mem.VirtualMemory()
	pc, _ := host.Info()
	proc, _ := cpu.Info()

	buildInfo, _ := debug.ReadBuildInfo()
	deps := buildInfo.Deps
	bfcordVersion := ""
	for _, dep := range deps {
		if dep.Path == "github.com/BOOMfinity/bfcord" {
			bfcordVersion = dep.Version
			break
		}
	}

	message := ctx.Interaction.SendMessageReply()
	message.Embed(discord.MessageEmbed{
		Title: "GoProdukcji Stats",
		Description: fmt.Sprintf(`Gateway ping: %vms
Wersja: [%v](https://github.com/MrBoombastic/GoProdukcji/commit/%v)
bfcord %v
%v
Uptime: %v
RAM (całego serwera): %v / %v (%v%%)

heapInuse / heapTotal: %v / %v
GC: %v
STW: %.2fms
RSS: %v

%v %v
%v (wątków: %v)`, ctx.Client.AvgLatency(), GitCommitHash, GitCommitHash,
			bfcordVersion, runtime.Version(), time.Since(uptime).String(), utils.FormatBytes(memory.Used), utils.FormatBytes(memory.Total),
			math.Round(memory.UsedPercent), utils.FormatBytes(rmem.HeapInuse), utils.FormatBytes(rmem.HeapSys-rmem.HeapReleased), rmem.NumGC, float64(time.Duration(rmem.PauseTotalNs))/float64(time.Millisecond), utils.FormatBytes(rss),
			pc.Platform, pc.KernelVersion, proc[0].ModelName, proc[0].Cores),
		Color: colors.Orange,
	})
	err := message.Execute()
	if err != nil {
		panic(err)
	}
}
