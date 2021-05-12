package ch07

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
    *c  += ByteCounter(len(p))
    return len(p), nil
}
