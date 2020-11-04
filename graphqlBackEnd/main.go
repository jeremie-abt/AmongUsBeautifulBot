package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, _ := sql.Open("sqlite3", "db")

	PlayerObj := graphql.NewObject(graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"Uid": &graphql.Field{
				Type: graphql.Int,
			},
			"DiscordId": &graphql.Field{
				Type: graphql.Int,
			},
			"GuildPlayed": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	GuildObj := graphql.NewObject(graphql.ObjectConfig{
		Name: "guild",
		Fields: graphql.Fields{
			"Uid": &graphql.Field{
				Type: graphql.Int,
			},
			"NbGamesFailed": &graphql.Field{
				Type: graphql.Int,
			},
			"DiscordId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})

	GamesWinHisto := graphql.NewObject(graphql.ObjectConfig{
		Name: "gamesWinHisto",
		Fields: graphql.Fields{
			"Uid": &graphql.Field{
				Type: graphql.Int,
			},
			"IdPlayer": &graphql.Field{
				Type: graphql.Int,
			},
			"IdGuild": &graphql.Field{
				Type: graphql.Int,
			},
			"IdChan": &graphql.Field{
				Type: graphql.Int,
			},
			"CreateAt": &graphql.Field{
				Type: graphql.Int, // Int ??
			},
		},
	})

	// Mutation
	MutationConfig := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"GamesHisto": &graphql.Field{
				Type: GamesWinHisto,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					gameHistoObj := &GameWinHisto{
						IdPlayer:  p.Args["idPlayer"].(int),
						IdGuild:   p.Args["idGuild"].(int),
						IdChan:    p.Args["idChan"].(int),
						CreatedAt: time.Now().String(), // TODO : recup
					}
					fmt.Printf("recording to database ...\n")
					insertedUid := InsertNewWin(db, gameHistoObj)
					gameHistoObj.Uid = insertedUid
					return gameHistoObj, nil
				},
				Args: graphql.FieldConfigArgument{
					"idPlayer": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: -1,
					},
					"idGuild": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: -1,
					},
					"idChan": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: -1,
					},
				},
			},
			"PlayerObj": &graphql.Field{
				Type: PlayerObj,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fmt.Printf("Recording new player to databaase ...\n")
					newPlayerObj := &Player{
						DiscordId:   p.Args["discordId"].(int),
						GuildPlayed: p.Args["guildPlayed"].(string),
					}
					insertedId := InsertNewPlayer(db, newPlayerObj)
					newPlayerObj.Uid = insertedId
					return newPlayerObj, nil
				},
				Args: graphql.FieldConfigArgument{
					"discordId": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: -1,
					},
					"guildPlayed": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
			},
			"GuildObj": &graphql.Field{
				Type: GuildObj,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fmt.Printf("Recording guild player to databaase ...\n")
					newGuildObj := &Guild{
						DiscordId:     p.Args["discordId"].(int),
						NbGamesFailed: p.Args["nbGamesFailed"].(int),
					}
					insertedId := InsertNewGuild(db, newGuildObj)
					newGuildObj.Uid = insertedId
					return newGuildObj, nil
				},
				Args: graphql.FieldConfigArgument{
					"discordId": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: -1,
					},
					"nbGamesFailed": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: -1,
					},
				},
			},
		},
	})

	QueryConfig := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"fuck": &graphql.Field{
				Name: "basicfieldNamedebug",
				Type: graphql.String,
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryConfig,
		Mutation: MutationConfig,
	})

	queryStatement := `
		mutation testmuta {
			GuildObj (discordId: 35546545464, nbGamesFailed:150) {
				Uid
			}
		}
	`
	result := graphql.Do(graphql.Params{
		RequestString: queryStatement,
		Schema:        schema,
	})
	fmt.Printf("result : %+v\n\n", result)

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
