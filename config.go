package main

// GameConfig
// This struct should retain information on the overall
// Config of the game :
// Either basic configs like light range or custom config
// Like which role Are needed etc ...
type GameConfig struct {
	NbImpostor int
	 // ...
	 // TODO : Really support these kind of default config
	 // and think of a design not too hard to get into

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
