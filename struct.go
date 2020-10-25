package main

import (
//	"fmt"

	"github.com/bwmarrin/discordgo"
//	"fmt"
)

type Game struct {
	IdChan string

	GameConfig *GameConfig
	
}

func NewGame(ChannelId string, GameConfig *GameConfig) (*Game) {

	return &Game{
		IdChan: ChannelId,
		GameConfig: GameConfig,
	}
}
/// pas d'appartenance a une game, on va avoir ca direct via
/// le channel Id, ce qi

/// Est-ce qu'une game a besoins d'avoir une liste de plyaer ?
/// Peut-etre plus tard si j'ai souvent besoins de check
/// mais pour l'instant je ne pense pas
/// tu peux retrouver les players a partir d'une game
/// et a partir d'un joueru une game, par contre si on
/// a souvent besoins de recup les players d'une game
/// la il faudrat en effet faire quelque chose


type GuildManagerType struct {
	// AU -> Amongus ;)
	// TODO: DiscordPlayer, pointers ??
	AUUsers map[string]*DiscordPlayer
	AUGames []*Game

	GuildId string
}

func NewGuildManager(s *discordgo.Session, GuildId string) (*GuildManagerType) {
	amongUsUser := make(map[string]*DiscordPlayer)
	// TODO : Comment faire ca plus propre ??
	return &GuildManagerType{
		AUUsers: amongUsUser,
		GuildId: GuildId,
	}
}

func (gm *GuildManagerType) HandleVoiceChange(m *discordgo.VoiceState) {
	if _, ok := gm.AUUsers[m.UserID]; !ok {
		newPlayer := NewDiscordPlayer(m.UserID)
		gm.AUUsers[m.UserID] = newPlayer
	} else {
		playerToUpdate := gm.AUUsers[m.UserID]
		playerToUpdate.isMute = m.SelfMute
		playerToUpdate.isDeaf = m.SelfDeaf
		// TODO: Ici il faut egalement faire ce quil se doit
		// pour les changer correctement de game etc ...
		playerToUpdate.channelID = m.ChannelID
	}

	// verif s'il existe
	// si oui :
	// 	faire les updates de state
	//	si changement de chan faire le nexessate
	// si non:
	// 	Lajouter -> DONe
	//'
}

func (gm *GuildManagerType) AttachGame (game *Game) {
	// verify if already existing
	isExisting := false
	for	_, it := range(gm.AUGames) {
		if (it.IdChan == game.IdChan) {
			isExisting = true
		}
	}
	if isExisting {
		// TODO : manage errors
		panic("tu dois gerer les errors")
	}
	gm.AUGames = append(gm.AUGames, game)
}

/*
**	GlobalVarManager is a struct which will be global
**	Its purpose is to hold all the data needed within
**	event handler ?
*/

type GlobalVarManagerType struct {
	GuildManagers []*GuildManagerType
}

// TODO: Trouver comment revenir a la ligne
func (gvmanager *GlobalVarManagerType) getGuildObj(GuildId string) (*GuildManagerType) {
	for _, guildManager := range gvmanager.GuildManagers {
		if guildManager.GuildId == GuildId {
			return guildManager
		}
	}
	return nil
}
