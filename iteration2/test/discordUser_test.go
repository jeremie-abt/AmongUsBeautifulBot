package test

import (
	"testing"

	//	"github.com/bwmarrin/discordgo"
	myBot "github.com/jeremie-abt/AmongUsBeautifulBot/iteration2"
)

var dcusrs myBot.DcUsers
var dcusrssample = []*myBot.DcUser{
	myBot.NewDcUser("jeremie", "546546478984130"),
	myBot.NewDcUser("adrien", "546546478984131"),
	myBot.NewDcUser("aurelie", "546546478984132"),
	myBot.NewDcUser("stephane", "546546478984133"),
	myBot.NewDcUser("alain", "546546478984134"),
	myBot.NewDcUser("xavie", "546546478984135"),
	myBot.NewDcUser("terreurdu39", "546546478984136"),
	myBot.NewDcUser("thomas", "546546478984137"),
	myBot.NewDcUser("delphine", "546546478984138"),
	myBot.NewDcUser("julie", "546546478984139"),
}

func TestGetDiscordEntity(t *testing.T) {
	dcusrs = myBot.NewDcUsers()

	dcusrs.GetEntity(dcusrssample[9].DiscordId)

	dcusrs.AddEntity(dcusrssample[0])
	if dcusr := dcusrs.GetEntity(dcusrssample[0].DiscordId); dcusr == nil {
		t.Errorf("dcusrsample[0] should have been selected")
	}
	if dcusr := dcusrs.GetEntity(dcusrssample[1].DiscordId); dcusr != nil {
		t.Errorf("dcusrsample[1] should not exist")
	}
	return
}

func TestGetDiscordEntityId(t *testing.T) {
	dcusrs = myBot.NewDcUsers()
	fillEntitiesWithSample()

	entityId := dcusrs.GetEntityId("juliasdasde")
	if entityId != "" {
		t.Errorf("entityId should be nil : %s\n", entityId)
	}
	entityId = dcusrs.GetEntityId("julie")
	if entityId != "546546478984139" {
		t.Errorf("Bad entityId %s - should be : 546546478984139\n", entityId)
	}

	return
}

func TestAddDiscordEntity(t *testing.T) {
	dcusrs = myBot.NewDcUsers()

	discordEntity := myBot.NewDcUser(dcusrssample[0].Name, dcusrssample[0].DiscordId)

	dcusrs.AddEntity(discordEntity)
	if _, found := dcusrs[dcusrssample[0].DiscordId]; !found {
		t.Errorf("following discord user not added : %+v\n", dcusrssample[0])
		return
	}
	dcusr, _ := dcusrs[dcusrssample[0].DiscordId]
	if dcusr.Name != dcusrssample[0].Name {
		t.Errorf("following discord user have wrong name %s : %+v\n", dcusr.Name, dcusrssample[0])
	}

	fillEntitiesWithSample()

	if len(dcusrs) != 10 {
		t.Errorf("len(%d) should be 10\n", len(dcusrs))
	}

	for _, value := range dcusrssample {
		if value, found := dcusrs[value.DiscordId]; !found {
			t.Errorf("dcusr not found : %+v\n", value)
		}
	}

	dcusrs.AddEntity(myBot.NewDcUser("jeanjacque-le-shlag", "6546844868464"))
	if len(dcusrs) > 10 {
		t.Errorf("Len should not exceed 10\n")
	}

	dcusrs = myBot.NewDcUsers()
	dcusrs.AddEntity(myBot.NewDcUser("jeanjacque-le-shlag", "6546844868464"))
	dcusrs.AddEntity(myBot.NewDcUser("test", "6546844868464"))
	dcusrs.AddEntity(myBot.NewDcUser("jeanjacque-le-shlag", "65465445"))
	if len(dcusrs) != 1 {
		t.Errorf("Should not add doublon")
	}
}

func TestRenameEntity(t *testing.T) {
	dcusrs = myBot.NewDcUsers()

	fillEntitiesWithSample()
	dcusrs.RenameEntity(dcusrssample[0].DiscordId, "newNameTest")
	dcusr := dcusrs.GetEntity(dcusrssample[0].DiscordId)
	if dcusr.Name != "newNameTest" {
		t.Errorf("discord user %s should have name newNameTest\n", dcusr.Name)
	}

	dcusrs.RenameEntity(dcusrssample[1].DiscordId, "newNameTest")
	dcusr = dcusrs.GetEntity(dcusrssample[1].DiscordId)
	if dcusr.Name == "newNameTest" {
		t.Errorf("should not have name newNameTest - already existing\n")
	}

	dcusrs.RenameEntity(myBot.NewDcUser("aoijdsoadij", "asdljakad").DiscordId, "newname")
}

func TestRemoveDiscordEntity(t *testing.T) {
	dcusrs = myBot.NewDcUsers()
	dcusrs.RemoveEntity(dcusrssample[5].DiscordId)

	dcusrs.AddEntity(dcusrssample[3])
	dcusrs.AddEntity(dcusrssample[4])
	dcusrs.RemoveEntity(dcusrssample[4].DiscordId)
	if dcusr := dcusrs.GetEntity(dcusrssample[3].DiscordId); dcusr == nil {
		t.Errorf("dcusrsample[3] should not have been suppressed")
	}
	if dcusr := dcusrs.GetEntity(dcusrssample[4].DiscordId); dcusr != nil {
		t.Errorf("dcusrsample[4] should have been suppressed")
	}
}

func fillEntitiesWithSample() {
	dcusrs = myBot.NewDcUsers()

	for _, value := range dcusrssample {
		dcusrs.AddEntity(value)
	}

}
