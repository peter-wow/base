package main

import (
	"example/gopl.io/ch8/work8.5/mandelbrot"
	"image/png"
	"os"
	"runtime"
)

func main() {
	// -1 不设置
	workers := runtime.GOMAXPROCS(-1)
	img := mandelbrot.ConcurrentRender(workers)
	png.Encode(os.Stdout, img)
}
