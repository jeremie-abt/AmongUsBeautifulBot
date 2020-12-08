package bot

type AuPlayers struct {
	AuPlayers      map[string]*AuPlayer
	playerEntities PlayerEntities
}

type IAuPlayer interface {
	IsAlive()
	SetIsAlive(isAlive bool)
	UnMute()
	Mute(bool)
}

type ColorPlayer int

const (
	YELLOW ColorPlayer = iota
	GREEN
)

type AuPlayer struct {
	Name     string
	LinkedTo string
	isAlive  bool

	color ColorPlayer
}

func NewAuPlayers(playerEntities PlayerEntities) *AuPlayers {

	auplayers := make(map[string]*AuPlayer)
	return &AuPlayers{
		AuPlayers:      auplayers,
		playerEntities: playerEntities,
	}
}

func (auplayers AuPlayers) AddAuPlayer(name string) *AuPlayer {

	if len(auplayers.AuPlayers) == 10 {
		return nil
	}

	auplayer := auplayers.GetAuPlayer(name)
	if auplayer != nil {
		return auplayer
	}
	entityId := auplayers.playerEntities.GetEntityId(name)
	auplayers.AuPlayers[name] = &AuPlayer{
		Name:     name,
		LinkedTo: entityId,
		isAlive:  true,
	}

	return auplayers.AuPlayers[name]
}

func (auplayers AuPlayers) GetAuPlayer(name string) *AuPlayer {
	if auplayer, found := auplayers.AuPlayers[name]; found {
		return auplayer
	}
	return nil
}

func (auplayers AuPlayers) UpdateAuPlayer(name string, newName string) {
	if auPlayer := auplayers.GetAuPlayer(name); auPlayer != nil {
		auplayers.DeleteAuPlayer(name)
		auplayers.AddAuPlayer(newName)
	}
}

func (auplayers AuPlayers) DeleteAuPlayer(name string) {
	if auPlayer := auplayers.GetAuPlayer(name); auPlayer != nil {
		delete(auplayers.AuPlayers, name)
	}
}

func (auplayers AuPlayers) Range() <-chan *AuPlayer {

	auplayerChan := make(chan *AuPlayer, 1)
	go func() {
		defer close(auplayerChan)
		for _, auplayer := range auplayers.AuPlayers {
			auplayerChan <- auplayer
		}
	}()

	return auplayerChan
}

func (player *AuPlayer) Mute(wantMute bool) {
	if player.IsAlive() == true {
		player.Mute(true)
		return
	}
	return
}

func (player *AuPlayer) UnMute() {
	if player.IsAlive() {
		player.Mute(false)
	}
	return
}

func (player *AuPlayer) ChangeColor(color ColorPlayer) error {
	//TODO impl
	return nil
}

func (auPlayer *AuPlayer) SetIsAlive(isAlive bool) {
	auPlayer.isAlive = isAlive
	return
}

func (auPlayer *AuPlayer) IsAlive() bool {
	return auPlayer.isAlive
}
