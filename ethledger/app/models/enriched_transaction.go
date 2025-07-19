package models

type EnrichedTransaction struct {
	Transaction
	Category string `csv:"Category"`
}
