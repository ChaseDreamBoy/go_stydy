// 生成萨如图形(Lissajous figures), 一个 gif 图片
// 生成方法 : go build Lissajous.go
// 然后需要将标准输出重定向到一个GIF图像文件, 使用 : Lissajous.exe > out.gif
package main

import (
	// "fmt"
	"image"
	// 当 import 了一个包路径包含有多个单词的 package 时，
	// 比如 image/color（image和color两个单词），通常只需要用最后那个单词表示这个包就可以。
	// 所以当写 color.White 时，这个变量指向的是 image/color 包里的变量
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	// "time"
)

var palette = []color.Color{color.White, color.Black}

const (
	// 白色 调色板
	whiteIndex = 0
	// 黑色
	blackIndex = 1
)

func main() {
	// 利用当前时间的伪随机数发生器
	// rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {

	const (
		// 全x振子转数
		cycles = 5
		// 角分辨率
		res = 0.001
		// 图像画布封面  [-size..+size]
		size = 100
		// 动画帧数
		nFrames = 64
		// 以10毫秒为单位的帧间延迟
		delay = 8
	)

	// y振荡器的相对频率
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nFrames}
	// 相位差
	phase := 0.0
	for i := 0; i < nFrames; i++ {
		// 外层循环 每一次都会生成一个单独的动画帧
		// image.Rect ==> 一个矩形的最小坐标和最大坐标 ==> 201*201大小的图片
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// 它生成了一个包含两种颜色的 201*201 大小的图片，白色和黑色。
		// 所有像素点都会被默认设置为其零值（也就是调色板palette里的第0个值），这里设置的是白色。
		// 每次外层循环都会生成一张新图片，并将一些像素设置为黑色。
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	// NOTE: ignoring encoding errors
	gif.EncodeAll(out, &anim)
}
