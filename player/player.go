package player

import (
	"fmt"
)

type Type int

const (
	None   Type = 0
	Cross  Type = 1
	Naught Type = 2
)

func Switch(t Type) Type {
	switch t {
	case Cross:
		return Naught
	case Naught:
		return Cross
	}

	panic(fmt.Sprintf("invalid player type: %d", t))
}
