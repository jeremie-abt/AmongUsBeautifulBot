package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type GameWinHisto struct {
	Uid       int64  `json:"uid"`
	IdPlayer  int    `json:"idplayer"`
	IdGuild   int    `json:"idguild"`
	IdChan    int    `json:"idchan"`
	CreatedAt string `json:"idchan"`
}

type Player struct {
	Uid         int64  `json:"uid"`
	DiscordId   int    `json:"discordId"`
	GuildPlayed string `json:"guildPlayed"`
}

type Guild struct {
	Uid           int64 `json:"uid"`
	DiscordId     int   `json:"discordId"`
	NbGamesFailed int   `json:"nbGamesFailed"`
}

func InsertNewWin(db *sql.DB, newWin *GameWinHisto) int64 {

	statement, err := db.Prepare(`
		INSERT INTO GamesWinHisto (IdPlayer, idGuild, idChan)
		Values (?, ?, ?)
	`)
	if err != nil {
		// TODO: Gerer les erreurs
		fmt.Printf("err : %+v\n", err)
		panic("gere les erreurs")
	}

	result, err := statement.Exec(
		newWin.IdPlayer, newWin.IdGuild, newWin.IdChan)

	if err != nil {
		fmt.Printf("Err : %+v\n")
		panic("gere les erreurs")
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Err : %+v\n")
		panic("gere les erreurs")
	}
	return lastInsertedId
}

func InsertNewPlayer(db *sql.DB, newPlayer *Player) int64 {

	statement, err := db.Prepare(`
		INSERT INTO Player (guildPlayed, discordId) VALUES (?, ?)
	`)

	if err != nil {
		fmt.Printf("err : %+v\n", err)
		panic("gere les erreurs")
	}
	result, err := statement.Exec(newPlayer.GuildPlayed, newPlayer.DiscordId)
	if err != nil {
		fmt.Printf("error : %+v\n\n", err)
		panic("gere les errors")
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("error : %+v\n\n", err)
		panic("gere les errors")
	}

	return lastInsertedId
}

func InsertNewGuild(db *sql.DB, newGuild *Guild) int64 {

	statement, err := db.Prepare(`
		INSERT INTO Guild (nbGamesFailed, discordId) VALUES (?, ?)
	`)

	if err != nil {
		fmt.Printf("err : %+v\n", err)
		panic("gere les erreurs")
	}
	ret, err := statement.Exec(newGuild.NbGamesFailed, newGuild.DiscordId)
	if err != nil {
		fmt.Printf("error : %+v\n\n", err)
		panic("gere les errors")
	}

	lastInsertedId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("error : %+v\n\n", err)
		panic("gere les errors")
	}
	return lastInsertedId
}
