package gocache

type ByteView struct {
	b []byte
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) BytesSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}
