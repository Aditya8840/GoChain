package types

type Block struct {
	Index int
	Timestamp string
	Hash []byte
	Data []byte
	Prev_hash []byte
}