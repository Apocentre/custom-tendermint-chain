package types

import "time"

const (
	// ModuleName defines the module name
	ModuleName = "checkers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkers"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SystemInfoKey = "SystemInfo-value-"
)

const (
	GameCreatedEventType      = "new-game-created" // Indicates what event type to listen to
	GameCreatedEventCreator   = "creator"          // Subsidiary information
	GameCreatedEventGameIndex = "game-index"       // What game is relevant
	GameCreatedEventBlack     = "black"            // Is it relevant to me?
	GameCreatedEventRed       = "red"              // Is it relevant to me?
)

const (
	MovePlayedEventType      = "move-played"
	MovePlayedEventCreator   = "creator"
	MovePlayedEventGameIndex = "game-index"
	MovePlayedEventCapturedX = "captured-x"
	MovePlayedEventCapturedY = "captured-y"
	MovePlayedEventWinner    = "winner"
	MovePlayedEventBoard     = "board"
)

const (
	GameRejectedEventType      = "game-rejected"
	GameRejectedEventCreator   = "creator"
	GameRejectedEventGameIndex = "game-index"
)

// There must be an "ID" that indicates no game. Use "-1", which you save as a constant:
const (
	NoFifoIndex = "-1"
)

// On each update the deadline will always be now plus a fixed duration. In this context, now refers to
// the block's time. Declare this duration as a new constant, plus how the date is to be represented - encoded
// in the saved game as a string
const (
	MaxTurnDuration = time.Duration(24 * 3_600 * 1000_000_000) // 1 day
	DeadlineLayout  = "2006-01-02 15:04:05.999999999 +0000 UTC"
)

const (
	GameForfeitedEventType      = "game-forfeited"
	GameForfeitedEventGameIndex = "game-index"
	GameForfeitedEventWinner    = "winner"
	GameForfeitedEventBoard     = "board"
)

const (
	GameCreatedEventWager = "wager"
)

// To get a rule-of-thumb idea of how much gas is already consumed without your additions,
// look back at your previous transactions
const (
	CreateGameGas       = 15000
	PlayMoveGas         = 1000
	RejectGameRefundGas = 14000
)
