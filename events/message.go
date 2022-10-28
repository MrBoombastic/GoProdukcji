package events

import (
	"github.com/BOOMfinity/bfcord/client"
	"github.com/BOOMfinity/bfcord/discord"
	"goprodukcji/utils"
)

func HandleMessage(c client.Client, message discord.Message) {

	me, _ := c.CurrentUser()
	channel, _ := c.Channel(message.ChannelID).Get()
	// Todo: not implemented in bfcord yet
	//guild := message.Guild()

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
			embed := utils.MentionEmbed("") //guild.IconURL(&discord.ImageOptions{})) //Todo: not implemented in bfcord yet
			message := c.Channel(message.ChannelID).SendMessage()
			message.Embed(embed)
			_, err := message.Execute(c)
			if err != nil {
				panic(err)
			}
		}
	}
}
