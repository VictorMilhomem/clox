package values

type ObjType int

const (
	OBJ_STRING ObjType = iota
)

type Obj struct {
	Kind ObjType
}

type ObjString struct {
	Obj    Obj
	Length int
	Chars  string
}

func CopyString(chars string) *ObjString {
	return &ObjString{
		Obj:    Obj{OBJ_STRING},
		Length: len(chars),
		Chars:  chars,
	}
}
