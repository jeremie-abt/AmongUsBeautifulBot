# AmongUsBeautifulBot

Feature:
Commands can be executed in whichever chan you want
	- .creategame gamename(string) : create a new discord chan within which there will be an among us game
	- .stopgame gamename(string) : stop listen for a game (delete chan if We created It)
	- .addConfig (json file) : add config file (see config file)
	- .stat : get all the stat for the current player within a channel


Vocabulary :
	- config file : json file which will permit to manage the game, how many imposter, add one or several custome role
	- Custom role : some extra role which are not present within the original game, like clairvoyant
	- Game : listening for a discord channel, recording a getting a state for all the player, asking the admin to record his game
	and mute / unmute people at the begining of game or if they're dead
	- admin : The one which has created the game.

Thoughts:
I don't know If I'll do all of this part, cause at the basics these was only a litle project to get into go devellopment,
But here is some idea that could be great to implement.
	- Storing a lot of info and make a nice website to log all the data, because there isn't any on the game -> I would love to make a graphql backend
	to get into this subject but not sure this is the better for this project.
	- Making a guild contest, like you invite your friend from another guild, you flag them with a role saying they're foreigner
	and you can play guild v guild, and have stat on your guild vs your friend guild (not sure this is possible via discord api)
	- making a speaking time within the config file to organizate the debate (like 5 secs each turn)
	- Maybe get some points each time you win a game, giving you some bonus like an extra speaking time (to define)
	- I'll probably add more Idea in the future

I'm not sure yet of all the design code, That's my first real go project so we'll see ;) !
