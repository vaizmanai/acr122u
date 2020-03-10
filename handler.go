package acr122u

// Handler is the interface that handles each card when present in the field.
type Handler interface {
	OnPresent(Card) error // executed when card is presented to the reader
	OnRelease(Card) error // executed when card is released from the reader
}
