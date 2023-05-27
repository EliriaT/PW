package events

const (
	AlreadySavedMessage = "This link was already saved 😉"
	HelpMessage         = `
	I am a bot and I can assist you with reading news!

	Here are the commands I was taught to execute:

🔸  /start
   If you want to greet me again

🔸	/latest_news [topic]
   To view 5 latest news

🔸  /save_news 	url
   To save an url to your saved news

🔸  /saved_news
   To view the previously saved news

🔸  /rnd_news                 
   To view and remove a random saved url 

🔸  /remove_news              
   To view and remove a news link from the saved news

🔸  /help                          
   To view the commands
`

	StartMessage = "Hi 👋 ! \n\n" + HelpMessage

	UnKnownCommandMessage = "Invalid command 🐙 ! I am sorry I was not taught what to do in this situations."
	NoSavedNewsMessage    = "You have no saved news. 🤷‍♀️️ "
	SavedNewsMessage      = "You successfully saved a message!🤸🏻‍♀️ To view all saved messages: /saved_news ."
	SelectNews            = "📌Please select an URL to read and delete it from the saved news."
	NoSuchLink            = "No such link found in saved news! 🙅‍♀️"
	ProvideURL            = "✍🏻 Please provide below a valid an URL to save: "
	NoNews                = "No news available for this topic!🧐"
)
