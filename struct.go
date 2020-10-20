package main

type Games struct {
	IdChan string

	GameConfig *GameConfig
	
}

// detection dun user qui change de chan / partie
// Tracking correcte dun user
type GuildManagerType struct {
	// AU -> Amongus ;)
	// TODO: DiscordPlayer, pointers ??
	AUUsers map[string]DiscordPlayer
	AUGames []Games

	GuildId string
}


/*
**	GlobalVarManager is a struct which will be global
**	Its purpose is to hold all the data needed within
**	event handler ?
*/

type GlobalVarManagerType struct {
	GuildManager GuildManagerType
}
