package console

//MsgNotRecognizedCommandInMatch is a message sent while the match is started and there is an error
//parsing the command
var MsgNotRecognizedCommandInMatch = "input error - please type 'print' for game details"

//MsgNotRecognizedCommandNotInMatch is an error sent while the match is NOT started and there is an error
//parsing the command
var MsgNotRecognizedCommandNotInMatch = "input error - please start a game through typing 'Start: '' vs. ''"

//MsNoGameInProgress error sent if a game is not in progress
var MsNoGameInProgress = "No game currently in progress"

//MsgGameInProgress error sent if a game is not in progress
var MsgGameInProgress = "Game currently in progress - to start a new game finish the current one  through typing 'End'"
