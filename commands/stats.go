package commands

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/client"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/discord/colors"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
	"github.com/BOOMfinity-Developers/bfcord/other"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"log"
	"math"
	"runtime"
	"time"
)

type CommandHandler func(c client.Client, message gateway.MessageCreateEvent)

var uptime = time.Now()
var GitCommitHash string

func StatsHandler(c client.Client, message gateway.MessageCreateEvent) CommandHandler {
	rss := getMemory()
	// Golang runtime memory stats
	var rmem runtime.MemStats
	runtime.ReadMemStats(&rmem)
	memory, _ := mem.VirtualMemory()
	pc, _ := host.Info()
	proc, _ := cpu.Info()
	_, err := message.Reply(&discord.MessageCreateOptions{
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
%v %v (wątków: %v)`, c.Manager().AveragePing(), GitCommitHash, GitCommitHash,
				other.Version(), time.Since(uptime).String(), formatBytes(memory.Used), formatBytes(memory.Total),
				math.Round(memory.UsedPercent), formatBytes(rmem.HeapInuse), formatBytes(rmem.HeapSys-rmem.HeapReleased), rmem.NumGC, float64(time.Duration(rmem.PauseTotalNs))/float64(time.Millisecond), formatBytes(rss),
				pc.Platform, pc.KernelVersion, proc[0].ModelName, proc[0].Cores),
			Color: colors.Orange,
		}})

	if err != nil {
		log.Fatal(err)
	}
}
