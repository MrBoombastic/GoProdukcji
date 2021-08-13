package events

import (
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/client/state"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/BOOMfinity-Developers/bfcord/discord/colors"
	"github.com/BOOMfinity-Developers/bfcord/gateway"
	"github.com/BOOMfinity-Developers/bfcord/other"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"goprodukcji/config"
	"log"
	"math"
	"time"
)

var uptime = time.Now()
var GitCommitHash string

func HandleMessageCreate(c state.IState, config config.RunMode) func(message gateway.MessageCreateEvent) {
	return func(message gateway.MessageCreateEvent) {
		channel, _ := message.Channel().Get()
		//Repost all announcements/tweets
		if channel.Type == discord.GuildNewsChannel {
			err := message.Crosspost()
			if err != nil {
				log.Fatal(err)
			}
		}

		//Stats command
		if message.Content == config.Prefix+"stats" || message.Content == config.Prefix+"ping" {
			memory, _ := mem.VirtualMemory()
			pc, _ := host.Info()
			proc, _ := cpu.Info()
			_, sendErr := message.Channel().SendMessage(&discord.MessageCreateOptions{
				Embed: &discord.MessageEmbed{
					Title: "GoProdukcji Stats",
					Description: fmt.Sprintf("Gateway ping: %vms\n"+
						"Wersja: [%v](https://github.com/MrBoombastic/GoProdukcji/commit/%v)\n"+
						"%v\n"+
						"Uptime: %v\n"+
						"RAM (całego serwera): %vMB/%vMB (%v%%)\n\n"+
						"%v %v \n%v %v (wątków: %v)",
						c.Manager().AveragePing(), GitCommitHash, GitCommitHash, other.Version(), time.Since(uptime).String(),
						memory.Used/1024/1024, memory.Total/1024/1024, math.Round(memory.UsedPercent), pc.Platform, pc.KernelVersion, pc.Hostname, proc[0].ModelName, proc[0].Cores),
					Color: colors.Orange,
				}})

			if sendErr != nil {
				log.Fatal(sendErr)
			}
		}

		if message.Content == config.Prefix+"help" {
			_, sendErr := message.Channel().SendMessage(&discord.MessageCreateOptions{
				Embed: &discord.MessageEmbed{
					Title: "GoProdukcji Help",
					Description: fmt.Sprintf("`%vhelp` - wyświetla niniejszą pomoc\n"+
						"`%vstats` - wyświetla statystyki oraz ping bota",
						config.Prefix, config.Prefix),
					Color: colors.Orange,
					Footer: &discord.EmbedFooter{
						Text: "Bot w głównej mierze ma pomagać w automatyce niektórych rzeczy, stąd mało komend",
					},
				}})

			if sendErr != nil {
				log.Fatal(sendErr)
			}
		}
	}
}
