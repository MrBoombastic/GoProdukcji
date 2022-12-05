package events

import (
	"github.com/BOOMfinity/bfcord/api/cdn"
	"github.com/BOOMfinity/bfcord/client"
	"github.com/BOOMfinity/bfcord/discord"
	"goprodukcji/utils"
)

func HandleMessage(c client.Client, message discord.Message) {

	me, _ := c.CurrentUser()
	channel, _ := c.Channel(message.ChannelID).Get()

	//Repost all announcements/tweets
	if channel.Type == discord.ChannelTypeNews {
		err := c.API().Channel(channel.ID).Message(message.ID).CrossPost()
		if err != nil {
			panic(err)
		}
	}

	//React to pinging
	mentionedUsers := message.Mentions
	if len(mentionedUsers) > 0 {
		if mentionedUsers[0].ID == me.ID {
			guild, err := message.Guild(c).Get()
			if err != nil {
				panic(err)
			}
			embed := utils.MentionEmbed(guild.IconURL(cdn.ImageFormatPNG, cdn.ImageSize512))
			message := c.Channel(message.ChannelID).SendMessage()
			message.Embed(embed)
			_, err = message.Execute(c)
			if err != nil {
				panic(err)
			}
		}
	}
}
