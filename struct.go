package main

import (
	//	"fmt"

	"fmt"
	"github.com/bwmarrin/discordgo"

	"math/rand"
)


type Game struct {
	IdChan string

	GameConfig *GameConfig
	CaptureCode string

	// Does someone in the lobby has plug the amonguscapture ??
	IsCapturedConn bool
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
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

	generatedCode := randSeq(10)
	fmt.Printf("Voici le code : %s\n", generatedCode)

	return &Game{
		IdChan:     ChannelId,
		GameConfig: GameConfig,
		CaptureCode: generatedCode,
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

// TODO : implementer des systemes de lock
// pour a terme toujours passer par des fonctions pour toucher
// a cette struct

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

// TODO: Test et definir cette partie car je ne suis meme pas sur
// de ce que ca doit faire
func (gm *GuildManagerType) HandleVoiceChange(m *discordgo.VoiceState) {
	/*
		Est-ce quon doit faire quelque chose ??
	*/	


	fmt.Printf("Handlevoice ...\n")
	if _, ok := gm.AUUsers[m.UserID]; !ok {
		newPlayer := NewDiscordPlayer(m.UserID)
		gm.AUUsers[m.UserID] = newPlayer
	} else {
		//TODO : ??
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
**		GLOBAL var manager :
**			Global struct containing all the state the bot has
**			in-memory
*/

type GlobalVarManagerType struct {
	GuildManagers 	[]*GuildManagerType
	Sess			*discordgo.Session
}

func NewGlobalVarManager(s *discordgo.Session) GlobalVarManagerType {
	return GlobalVarManagerType{
		Sess: s,
	}
}

var G_Gvm GlobalVarManagerType

// TODO: Trouver comment revenir a la ligne
func (gvmanager *GlobalVarManagerType) getGuildObj(GuildId string) *GuildManagerType {
	for _, guildManager := range gvmanager.GuildManagers {
		if guildManager.GuildId == GuildId {
			return guildManager
		}
	}
	return nil
}

func (gvm *GlobalVarManagerType) AddGuild(guildID string) *GuildManagerType {
	// TODO: lock
	newguild := NewGuildManager(gvm.Sess, guildID)

	gvm.GuildManagers = append(gvm.GuildManagers, newguild)
	return newguild
}

// TODO(optionnel): implementer un generateur par type iterable
// par exmple plutot que de loop sur les guilds, on appel une fonction
// qui return un chan ou une list et hop
func (gvm *GlobalVarManagerType) GetGameCode(code string) *Game {
	for _, guildManager := range gvm.GuildManagers {
		// iteration sur les games
		for _, game := range guildManager.AUGames {
			fmt.Printf("Game : %+v\n", game)
			if code == game.CaptureCode {
				return game
			}
		}
	}
	return nil
}
