package bot

import (
	"github.com/bwmarrin/discordgo"
)

type DcUsers map[string]*DcUser

func NewDcUsers() DcUsers {
	return make(map[string]*DcUser)
}

type DcUser struct {
	Name      string
	DiscordId string
	IsLinked  bool
}

func NewDcUser(name string, discordId string) *DcUser {

	// TODO: Verifier si un among us user existe
	return &DcUser{
		Name:      name,
		DiscordId: discordId,
		IsLinked:  false,
	}
}

// --------- Implementation entityType

func (dcusrs DcUsers) Mute(mute bool) {
	// TODO: Mute on discor
	return
}

func (dcusrs DcUsers) GetEntity(id string) *DcUser {
	if entity, found := dcusrs[id]; found {
		return entity
	}
	return nil
}

func (dcusrs DcUsers) GetEntityId(name string) string {

	for _, dcusr := range dcusrs {
		if dcusr.Name == name {
			dcusr.IsLinked = true
			return dcusr.DiscordId
		}
	}
	return ""
}

func (dcusrs DcUsers) AddEntity(entity interface{}) {

	if len(dcusrs) >= 10 {
		return
	}
	dcusr := entity.(*DcUser)
	if _, found := dcusrs[dcusr.DiscordId]; found {
		// TODO log
		return
	}
	for _, element := range dcusrs {
		if element.Name == dcusr.Name {
			return
		}
	}
	dcusrs[dcusr.DiscordId] = dcusr
}

func (dcusrs DcUsers) RenameEntity(id string, newName string) {
	dcusr := dcusrs.GetEntity(id)
	if dcusr == nil {
		return
	}

	for _, element := range dcusrs {
		if element.Name == newName {
			// Name already existing
			return
		}
	}
	dcusr.Name = newName
}

func (dcusrs DcUsers) RemoveEntity(id string) {
	delete(dcusrs, id)
}
