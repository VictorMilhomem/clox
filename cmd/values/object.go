package values

type ObjType int

const (
	OBJ_STRING ObjType = iota
)

type Obj struct {
	kind ObjType
}

type ObjString struct {
	obj    Obj
	length int
	chars  []rune
}
