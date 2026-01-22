package hasher

type Hahser interface {
	Encode(id int) (string, error)
	Decode(code string) (int, error)
}
