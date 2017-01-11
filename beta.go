package main

import (
    "github.com/bwmarrin/discordgo"
    "github.com/sn0w/Karen/logger"
    "time"
)

var BETA_GUILDS = [...]string{
    "259831440064118784", // FADED's Sandbox        (Me)
    "180818466847064065", // Karen's Sandbox        (Me)
    "172041631258640384", // P0WERPLANT             (Me)
    "161637499939192832", // Coding Lounge          (Devsome)
    "225168913808228352", // Emily's Space          (Kaaz)
    "267186654312136704", // Shinda Sekai Sensen    (黒ゲロロロ)
    "244110097599430667", // S E K A I              (Senpai /「 ステ 」Abuse)
    "266326434505687041", // Bot Test               (quoththeraven)
    "267658193407049728", // Bot Creation           (quoththeraven)
}

// Automatically leaves guilds that are not registered beta testers
func autoLeaver(session *discordgo.Session) {
    for {
        for _, guild := range session.State.Guilds {
            match := false

            for _, betaGuild := range BETA_GUILDS {
                if guild.ID == betaGuild {
                    match = true
                    break
                }
            }

            if !match {
                logger.WARNING.L("beta", "Leaving guild " + guild.ID + " (" + guild.Name + ") because it didn't apply for the beta")
                session.GuildLeave(guild.ID)
            }
        }

        time.Sleep(10 * time.Second)
    }
}