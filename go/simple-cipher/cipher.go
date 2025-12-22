package cipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type vigenere struct {
	plain  string
	cipher string
	key    string
}

type shift int
