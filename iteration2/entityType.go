package bot

/*
	This is the only interface that AuPlayer should know about
	Here this represent a discord user, but in reality this could be
	anything, like Slack user or msn user why not ;)
*/
type PlayerEntities interface {
	GetEntityId(string) string
	Mute(bool)
	/*
		je remove cette fonction car je veux remove toutes les fonctions
		qui ne sont pas call par le monde exterieur
	*/
	//	GetEntity(string) interface{}
	AddEntity(interface{})
	RemoveEntity(string)
}
