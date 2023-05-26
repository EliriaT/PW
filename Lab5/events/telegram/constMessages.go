package events

const (
	AlreadySavedMessage = "This link was already saved 😉"
	HelpMessage         = `
	I am a bot and I can assist you with reading news!

	Here are the commands I was taught to execute:

▪️  /start                                       - if you want to greet me again
▪️	/latest_news [topic]                - to view 5 latest news
▪️  /save_news 	https://url         - to save an url to your saved news
▪️  /saved_news                           - to view the previously saved news
▪️  /help                                        - to view the commands
`

	StartMessage = "Hi 👋 ! \n\n" + HelpMessage

	UnKnownCommandMessage = "Invalid command 🐸 ! I am sorry I was not taught what to do in this situations."
	NoSavedNewsMessage    = "You have no saved news. 🤷🏽‍♀️ "
	SavedNewsMessage      = "You successfully saved a message!🤸🏻‍♀️ To view all saved messages: /saved_news ."
)
