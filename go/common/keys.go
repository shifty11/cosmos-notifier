package common

type ContextKey string

const (
	// ContextKeyTx is the key for the transaction in the context
	ContextKeyTx ContextKey = "tx"

	// ContextKeyUser is the key for the user in the context
	ContextKeyUser ContextKey = "user"
)
