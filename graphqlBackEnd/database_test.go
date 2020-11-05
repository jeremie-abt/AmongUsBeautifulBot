package main

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

var db *sql.DB
var mock sqlmock.Sqlmock

// TODO: Test de mes fonctions d'updates / voir si c'est des bons tests
func TestUpdateWinHisto(t *testing.T) {

	//	mock.ExpectBegin()
	//	mock.ExpectCommit() // rien a foutre ? C'est pour les transac ca ??

	mock.ExpectPrepare("INSERT INTO GamesWinHisto").ExpectExec().WithArgs(-1, -1, -1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectPrepare("INSERT INTO GamesWinHisto").ExpectExec().WithArgs(-1, -1, 57984).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectPrepare("INSERT INTO GamesWinHisto").ExpectExec().WithArgs(-1, 68569, 57984).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectPrepare("INSERT INTO GamesWinHisto").ExpectExec().WithArgs(156, 68569, 57984).WillReturnResult(sqlmock.NewResult(1, 1))

	// TODO : Faire un constructeur, car la je dois mettre moi meme
	// a -1 -> Pas ouf
	newWinObj := &GameWinHisto{
		IdPlayer: -1,
		IdGuild:  -1,
		IdChan:   -1,
	}

	// Est-ce des bon tests unitaire ???
	InsertNewWin(db, newWinObj)
	newWinObj.IdChan = 57984
	InsertNewWin(db, newWinObj)
	newWinObj.IdGuild = 68569
	InsertNewWin(db, newWinObj)
	newWinObj.IdPlayer = 156
	InsertNewWin(db, newWinObj)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("update winHisto failed : %s\n", err)
	}
}

func TestUpdatePlayer(t *testing.T) {
	mock.ExpectPrepare("INSERT INTO Player").ExpectExec().WithArgs(-1, "").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectPrepare("INSERT INTO Player").ExpectExec().WithArgs(654611, "").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectPrepare("INSERT INTO Player").ExpectExec().WithArgs(654611, "[15,18]").WillReturnResult(sqlmock.NewResult(1, 1))

	newPlayerObj := &Player{
		DiscordId:   -1,
		GuildPlayed: "",
	}

	InsertNewPlayer(db, newPlayerObj)
	newPlayerObj.DiscordId = 654611
	InsertNewPlayer(db, newPlayerObj)
	newPlayerObj.GuildPlayed = "[15,18]"
	InsertNewPlayer(db, newPlayerObj)
}

func TestUpdateGuild(t *testing.T) {
	mock.ExpectPrepare("INSERT INTO Guild").ExpectExec().WithArgs(-1, -1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectPrepare("INSERT INTO Guild").ExpectExec().WithArgs(654833, -1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectPrepare("INSERT INTO Guild").ExpectExec().WithArgs(654833, 150).WillReturnResult(sqlmock.NewResult(1, 1))

	newGuildObj := &Guild{
		DiscordId:     -1,
		NbGamesFailed: -1,
	}

	InsertNewGuild(db, newGuildObj)
	newGuildObj.DiscordId = 654833
	InsertNewGuild(db, newGuildObj)
	newGuildObj.NbGamesFailed = 150
	InsertNewGuild(db, newGuildObj)
}

func init() {
	db, mock, _ = sqlmock.New()
}
