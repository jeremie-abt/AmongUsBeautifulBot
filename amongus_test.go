/// Begin early to make a lot of unit test

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bwmarrin/discordgo"
)

var discordSessionTest = initDiscordCon()

// Initialize discord connection for test
func initDiscordCon() *discordgo.Session {

	readjson_map := make(map[string]string)

	json_config_file, _ := os.Open("config.json")
	ret, _ := ioutil.ReadAll(json_config_file)

	json.Unmarshal(ret, &readjson_map)
	bot_token := readjson_map["bot_token"]

	dg, _ := discordgo.New("Bot " + bot_token)
	return dg
}

func TestNewConfig(t *testing.T) {
	// read les config files

	ret := NewGameConfig("testssrc/config1.json")

	if ret.NbImpostor != 1 {
		t.Errorf(
			"[TestNewConfig] NbImpostor badly initialised : %v, "+
				"but should have been %v\n", ret.NbImpostor, 1)
	}
	// TODO : Test the fieldsrandom_Composition
	cusStg := ret.customRoleSettings
	if cusStg.NbRolesWanted != 3 || cusStg.NbRolesOverlapping != 2 ||
		cusStg.RandomComposition != true ||
		*cusStg.Composition[0] != Talkie ||
		*cusStg.Composition[1] != Voyante ||
		*cusStg.Composition[2] != RoleGiver {
		t.Errorf("[TestNewConfig] Composition Fields badly initialised\n")
	}
	if ret.customRoles == nil {
		t.Errorf("[TestNewConfig] customRoles Fields badly initialised\n")
	}
}

func TestNewGuildManager(t *testing.T) {

	// 766750463524732968 : serveur discord de test
	_ = NewGuildManager(discordSessionTest, "766750463524732968")
	// TODO: des que je fais une vraie fonction faire un test

}

func TestNewDiscordPlayer() {
	// 345187059943735309 : identifiant jejems
	// TODO : check si cet id change de temps a autre ou pas du tout
	// -> Si non, on peut faire des stats plutot cool
	_ = NewDiscordPlayer("345187059943735309")
	// TODO: Faire un vraie test quand jaurai une logique plus pousser
}

/*
**		Integration test (via test server discord)
**		Maybe we should mock all of this but that would
**		be a lot of work, I think It's Ok to fail a build if discord api
**		fail, If that happens a lot, at this time maybe we should
**		change things
 */
