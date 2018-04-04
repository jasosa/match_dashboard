package dashboard

import (
	"errors"
)

//ErrMatchAlreadyStarted error returned when a match is already started
var ErrMatchAlreadyStarted = errors.New("Match already started")

//ErrMatchNotStarted error returned when tries to do an action different than
//start in a non-started match
var ErrMatchNotStarted = errors.New("Match not started")
