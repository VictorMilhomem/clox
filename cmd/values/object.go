package values

type ObjString struct {
	Length int
	Chars  string
}

func CopyString(chars string) *ObjString {
	return &ObjString{
		Length: len(chars),
		Chars:  chars,
	}
}
