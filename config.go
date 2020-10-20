package main

// Config for a entire discord server
// TODO: Maybe support per discord channel / among game
// custom settings, but now only one config for one
// discord server
type GameConfig struct {
	NbImpostor int
	LightRange int
	// TODO, faire un tour dans les configs de amongus

	// Custom config
	// 	Some custom config ...
	// Role (see role.txt)
	WantTalkie bool
}

// Cette methode est amener a beaucoup changer,
// c'est pourquoi il est important de passer par la
// pour instantier un GameConfig
func NewGameConfig ()(*GameConfig) {
	// TODO : Some stuff to correctly get some configs

	// (Either stored in a file or via command discord)
	// ( To define ! )

	// TODO: unmock
	return &GameConfig{
		NbImpostor: 1,
		WantTalkie: true,
	}
}
