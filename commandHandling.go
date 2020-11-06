package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

/*
**	All bot command will have a function like this
**	to treat the command
 */
// TODO : Reflechir a un pattern de test
func HandleCreateGame(
	s *discordgo.Session,
	msg *discordgo.Message,
	gm *GuildManagerType) {

	if len(msg.Content) <= 11 {
		panic("nom trop court\n")
		return // TODO : Gerer la gestion d'erreur
		// Probablement renvoyer un usage
	} else {
		gameChanName := msg.Content[12:]

		newChanId, isExisting := gm.GetChanId(s, gameChanName)
		if isExisting == false {
			// Create It if not existings
			createdChan := CreateNewChan(
				s, gm.GuildId, gameChanName)
			newChanId = createdChan.ID
		}

		newGame := NewGame(newChanId, NewGameConfig("config_role.json"))
		gm.AttachGame(newGame)
	}
}

func HandleStopGame(
	s *discordgo.Session,
	msg *discordgo.Message,
	gm *GuildManagerType) {
	/*
	**	Stop game, maybe ask if we want to destroy data ?
	**	or print data in General before destroying ?
	**	Not sure about what to do with this data,
	**	But at least stop listening this chan, delete the chan
	**	if our bot has created It and stop track this chan
	 */

	if len(msg.Content) <= 9 {
		NotifyError(s, msg.Author.ID, fmt.Errorf(USAGE_STOPGAME))
		return
	} else {
		gameChanName := msg.Content[10:]

		err := gm.RemoveGame(s, gameChanName)
		if err != nil {
			NotifyError(s, msg.Author.ID, err)
		}
		fmt.Printf("sucessfully stoped game\n")
	}
}
