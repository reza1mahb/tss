package common

import (
	"encoding/gob"

	"github.com/binance-chain/tss-lib/tss"
)

func init() {
	gob.Register(DummyMsg{})
}

type DummyMsg struct {
	Content string
}

func (m DummyMsg) String() string {
	return m.Content
}

// always broadcast
func (m DummyMsg) GetTo() *tss.PartyID {
	return nil
}

func (m DummyMsg) GetFrom() *tss.PartyID {
	return nil
}

func (m DummyMsg) GetType() string {
	return ""
}

func (m DummyMsg) ValidateBasic() bool {
	return true
}
