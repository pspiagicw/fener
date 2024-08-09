package vm

type Frame struct {
	Instructions []byte
	IP           int
}

func NewFrame(bytes []byte) *Frame {
	return &Frame{
		Instructions: bytes,
		IP:           0,
	}
}
