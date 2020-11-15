package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type CustomRoleName string

const (
	Talkie  CustomRoleName = ("talkie")
	Voyante                = "voyante"
	// He can choose to assign all the roles to whichever he's willing to
	RoleGiver = "rolegiver"
)

// Compo aleatoire -> Avec un nb predefine de role custom
// Compo specifique (potentiellement plusieurs fois le meme role)
// Certains role ne s'appliquant uniquement aux imposter ou aux gentils
// La superposition de roles
//	-> Ca sera probablement aux gens d'equilibrer la meta
// Des roles payants (Why not)

// customRole -> Settings for one specific custom Role
// Mais la jsuis pas sur que ce soit hyper opti de le faire
// des maintenan -> A voir plus tard
type customRole struct {
	roleName CustomRoleName

	// Compatibility -> 1 == this role can be on a not imposter guy
	// Compatibility -> 2 == this role can be on an imposter guy
	// Compatibility -> 2 == this role can be on both
	compatibility int
}

// Get the all config
type customRoleSettings struct {
	NbRolesWanted      int               `json:"nb_roles_wanted"`
	NbRolesOverlapping int               `json:"nb_roles_overlapping"`
	RandomComposition  bool              `json:"random_composition"`
	Composition        []*CustomRoleName `json:"composition"`
}

// Config for a game, or maybe for the entire guild,
// I'm not sure yet
type GameConfig struct {
	NbImpostor int
	// TODO : a voir si en effet on peut creer depuis le
	// bot les settings dans among

	customRoleSettings *customRoleSettings
	customRoles        []*customRole
}

func (gcf *GameConfig) GenerateCustomRoles() {
	/*
	**	Generate new customRoles in according to
	**	customRoleSettings field
	 */

	// TODO: Unmock, tu peux partir du principe que
	// customRoles est alloc a nbRolesWanted
	gcf.customRoles[0] = &customRole{
		roleName:      "talkie",
		compatibility: 3,
	}
	gcf.customRoles[1] = &customRole{
		roleName:      "talkie",
		compatibility: 3,
	}
}

func NewGameConfig(jsonPath string) *GameConfig {
	// Passer par un json qui va recup direct le json

	// TODO : Get default chan por notifier les erreurs
	jsonConfigRoleFile, _ := os.Open(jsonPath)

	configJson, _ := ioutil.ReadAll(jsonConfigRoleFile)

	var customRoleSgs customRoleSettings
	json.Unmarshal(configJson, &customRoleSgs)

	gameConfig := &GameConfig{
		NbImpostor: 1,

		customRoleSettings: &customRoleSgs,
		customRoles:        make([]*customRole, customRoleSgs.NbRolesWanted),
	}
	gameConfig.GenerateCustomRoles()

	return gameConfig
}
