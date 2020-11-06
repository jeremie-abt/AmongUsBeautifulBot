package main

import (
	//	"fmt"

	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Game struct {
	IdChan string

	GameConfig *GameConfig
}

// Take an chan Id and verify if the game is related to this chan
// Ptetre un peu overkill mais ca me permet de safiser mes struct
// la par exemple, je peux changer IdChan en IDChan, je ne dois
// update le code que dans cette methode
func (game *Game) isGameRelatedToChan(chanId string) bool {
	return game.IdChan == chanId
}

// TODO : not tested
func NewGame(ChannelId string, GameConfig *GameConfig) *Game {

	return &Game{
		IdChan:     ChannelId,
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
	AUUsers map[string]*DiscordPlayer
	AUGames []*Game

	GuildId string
}

func NewGuildManager(s *discordgo.Session, GuildId string) *GuildManagerType {
	amongUsUser := make(map[string]*DiscordPlayer)
	// TODO : Comment faire ca plus propre ??
	return &GuildManagerType{
		AUUsers: amongUsUser,
		GuildId: GuildId,
	}
}

// TODO: on track tous les joueurs,
//	Attention a ca tout de meme, ca fait beaucoup de calcul pour rien
func (gm *GuildManagerType) HandleVoiceChange(m *discordgo.VoiceState) {
	fmt.Printf("Handlevoice ...\n")
	if _, ok := gm.AUUsers[m.UserID]; !ok {
		fmt.Printf("je rentre cond 1\n")
		newPlayer := NewDiscordPlayer(m.UserID)
		gm.AUUsers[m.UserID] = newPlayer
		//for game := range gm.AUGames {
		//	if game.isGameRelatedToChan(m.ChannelID) {

		//	}
		//}
	} else {

		fmt.Printf("je rentre cond 1\n")
		playerToUpdate := gm.AUUsers[m.UserID]
		playerToUpdate.isMute = m.SelfMute
		playerToUpdate.isDeaf = m.SelfDeaf
		playerToUpdate.channelID = m.ChannelID
		// TODO : verif pour voir si le quota de joueur n'est pas
		// depasse
	}
}

// TODO : A voir si ca a vraiment du sens d'impl cette methode
// pour le guild manager, plutot que de faire une simple methode
// qui prend un guild ID en param ?
func (gm *GuildManagerType) GetChanId(s *discordgo.Session, chanName string) (string, bool) {
	/*
	** Utile pour savoir si un chan exist
	 */

	allChan, err := s.GuildChannels(gm.GuildId)
	if err != nil {
		// TODO : Manage errors
		panic("err verifying if chan exists\n\n")
	}

	for _, curChan := range allChan {
		if curChan.Name == chanName &&
			curChan.Type == discordgo.ChannelTypeGuildVoice {
			return curChan.ID, true
		}
	}
	return "", false
}

func (gm *GuildManagerType) RemoveGame(
	s *discordgo.Session, chanName string) (err error) {

	var indexToRemove int
	var curerror error

	toDeleteChanId, isExisting := gm.GetChanId(s, chanName)
	if isExisting == true {
		// TODO:  Check si nous avons bien cree le cannel, sinon
		// il ne faut pas le delete
		err := DeleteChan(s, toDeleteChanId)
		if err != nil {
			curerror = fmt.Errorf(
				"could not delete the following chan : %s\n"+
					"maybe you should delete It yourself if you want "+
					"To keep your Discord clean", chanName)
		}
	}
	// TODO: Notifier quand le mec done une game qui nexiste pas
	for index, curGame := range gm.AUGames {
		if toDeleteChanId == curGame.IdChan {
			indexToRemove = index
			break
		}
	}

	if len(gm.AUGames) >= 1 {
		gm.AUGames[indexToRemove] = gm.AUGames[len(gm.AUGames)-1]
		gm.AUGames = gm.AUGames[:len(gm.AUGames)-1]
	}
	return curerror
}

func (gm *GuildManagerType) AttachGame(game *Game) {
	// verify if already existing
	isExisting := false
	for _, it := range gm.AUGames {
		if it.IdChan == game.IdChan {
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
func (gvmanager *GlobalVarManagerType) getGuildObj(GuildId string) *GuildManagerType {
	for _, guildManager := range gvmanager.GuildManagers {
		if guildManager.GuildId == GuildId {
			return guildManager
		}
	}
	return nil
}
