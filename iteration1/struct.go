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

	DiscordPlayers []*DiscordPlayer
}

func (game *Game) AddPlayer(player *DiscordPlayer) {
	if len(game.DiscordPlayers) > 10 {
		// Among us Games are limited to 10 players
		// TODO Log and do something
		println("WARNING : trop de Players !!!\n")
		return 
	}
	game.DiscordPlayers = append(game.DiscordPlayers, player)
}

func (game *Game) GetDiscordPlayerFromSocketEvt(amongUsName string) (*DiscordPlayer){

	for _, discordPlayer := range game.DiscordPlayers {
		if discordPlayer.AmongUsLinkedName == amongUsName {
			return discordPlayer
		}
	}
	return nil
}

func (game *Game) MuteAllPlayer(guildId string, mute bool) {
	fmt.Printf("muteallplayer : %v ...\n", mute)
	for _, discordPlayer := range game.DiscordPlayers {
		fmt.Printf("Muting : %+v\n")
		discordPlayer.Mute(guildId, mute)
	}
}

func TryToLinkAUPlayerWithGame(AUplayerName string, captureCode string) {
	/*
		We want to link an among us player with a discord user.
		Warning: if the discord name and the amongUs name are not
		the same, then we won't link them.
		Maybe we should have better heuristic or method -> TODO
	*/

	curGame, _ := GetGameFromCode(captureCode)
	if curGame == nil {
		// TODO: Log
		fmt.Printf("Not any game linked with this code : %s\n", captureCode)
		return
	}
	// TODO: Je n'ai pas reussi a faire la liste des personne
	// presente dans un chan discord	
}


// ajouter un joueur discord a une partie
func AddDiscordPlayerToGame(gm *GuildManagerType, m *discordgo.MessageCreate) error {
	/*
		Add discord player to Game
	*/	
	var discordPlayer *DiscordPlayer
	if _, ok := gm.AUUsers[m.Message.Author.ID]; ok {
		discordPlayer = gm.AUUsers[m.Message.Author.ID]
	} else {
		discordPlayer = NewDiscordPlayer(
				m.Message.Author.ID, m.Message.Author.Username)
	}

	CHANNELID := "775637465981386792"
	curGame := gm.GetGameByChanId(CHANNELID)
	curGame.AddPlayer(discordPlayer)
	
	return nil
}


// Take a chan Id and verify if the game is related to this chan
// Ptetre un peu overkill mais ca me permet de safiser mes struct
// la par exemple, je peux changer IdChan en IDChan, je ne dois
// update le code que dans cette methode
func (game *Game) isGameRelatedToChan(chanId string) bool {
	return game.IdChan == chanId
}

// TODO : not tested
func NewGame(ChannelId string, GameConfig *GameConfig) *Game {

	generatedCode := randSeq(10)
	// Mock
	generatedCode = "aaaaaaaaaa"
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
	// TODO: Refacto -> C'est pas bon
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

func GetGameFromCode(captureCode string) (*Game, *GuildManagerType) {
	for _, guildManager := range G_Gvm.GuildManagers {
		for _, curGame := range guildManager.AUGames {
			if curGame.CaptureCode == captureCode {
				return curGame, guildManager
			}
		}
	}
	return nil, nil
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

func (gm *GuildManagerType) GetGameByChanId(chanId string) (*Game) {
	for _, curGame :=  range gm.AUGames {
		if curGame.IdChan == chanId {
			return curGame
		}
	}
	return nil
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

// ---------------------------------------- PRIVATE Func
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
