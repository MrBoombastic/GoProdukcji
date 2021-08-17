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
	"runtime"
	"strings"
	"time"
)

var uptime = time.Now()
var GitCommitHash string

func HandleMessageCreate(c state.IState, config config.RunMode) func(message gateway.MessageCreateEvent) {
	return func(message gateway.MessageCreateEvent) {
		if !strings.HasPrefix(message.Content, config.Prefix) {
			return
		}
		channel, _ := message.Channel().Get()
		args := strings.Fields(strings.TrimPrefix(message.Content, config.Prefix))
		command := args[0]
		args = args[1:]

		//Repost all announcements/tweets
		if channel.Type == discord.GuildNewsChannel {
			err := message.Crosspost()
			if err != nil {
				log.Fatal(err)
			}
		}

		//Stats command
		if command == "stats" || command == "ping" {
			rss := getMemory()
			// Golang runtime memory stats
			var rmem runtime.MemStats
			runtime.ReadMemStats(&rmem)
			memory, _ := mem.VirtualMemory()
			pc, _ := host.Info()
			proc, _ := cpu.Info()
			_, sendErr := message.Channel().SendMessage(&discord.MessageCreateOptions{
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

			if sendErr != nil {
				log.Fatal(sendErr)
			}
		}

		if command == "help" {
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
		if command == "search" {
			if len(args) == 0 {
				_, err := message.Channel().SendMessage(&discord.MessageCreateOptions{Content: "Musisz podać tytuł artykułu do wyszukania!"})
				if err != nil {
					return
				}
				return
			}

			foundArticle, foundErr := SearchArticle(strings.Join(args, " "))
			if foundErr != nil {
				_, err := message.Channel().SendMessage(&discord.MessageCreateOptions{Content: "Błąd: " + foundErr.Error() + "!"})
				if err != nil {
					return
				}
				return
			}
			_, sendErr := message.Channel().SendMessage(&discord.MessageCreateOptions{
				Embed: &discord.MessageEmbed{
					Title: foundArticle.Title,
					URL:   foundArticle.URL,
					Image: &discord.EmbedImage{
						Url: foundArticle.FeatureImage,
					},
					Description: strings.ReplaceAll(foundArticle.Excerpt, "\n", " ") + " (...)",
					Footer: &discord.EmbedFooter{
						Text: foundArticle.PublishedAt.Format(time.RFC822),
					},
				}})

			if sendErr != nil {
				log.Fatal(sendErr)
			}
		}
	}
}
