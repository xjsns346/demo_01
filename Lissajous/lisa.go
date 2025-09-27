//获取的gif文件无法打开，不会搞，练习1.5，1.6没有做。

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
)

// palette  (n) 调色板
var palette = []color.Color{color.White, color.Black}

const (
	//定义两个常量。
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	//这里的rand.seed()的作用是根据传入的数字，生成一个伪随机数。
	//time.Now()是获取当前时间，UTC是转成世界标准时间，UnixNano是转成Unix时间戳，单位是纳秒。
	//以此来保证每次运行获得的伪随机数不同。

	lissajous(os.Stdout) //输出二进制的gif文件到标准输出。

}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 含义：X 方向振荡器的完整周期数。决定绘制多少个正弦波循环。数值越大，图案越复杂，线条越密。
		res     = 0.001 //含义：角度分辨率（angular resolution）。控制步长 t += res，决定采样点的密集程度。值小 → 点多 → 线条更平滑。值大 → 点少 → 图案更粗糙，但计算更快。
		size    = 100   // 含义：画布的半径.图像的坐标范围是 [-size .. +size]，所以整体画布大小是 (2*size+1) x (2*size+1)。控制 GIF 的分辨率。
		nframes = 64    // 含义：动画的帧数。决定一共生成多少张图。帧数多 → 动画更流畅，但文件更大。
		delay   = 8     // 含义：帧间延迟。单位是 10ms（也就是 8 * 10 = 80ms）。决定每帧显示的时间。
	)
	//freq → y 振荡器的频率，决定曲线的花样复杂度。
	//rand.Float64()是生成一个在[0,1)的随机数，后面的*3.0是将范围变成[0,3)。
	freq := rand.Float64() * 3.0

	//anim → GIF 动画对象，用来存储帧并设置循环次数。
	anim := gif.GIF{LoopCount: nframes}

	//phase → 相位差，随着动画推进而变化，让图案动起来。
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ { //64次循环，每次生成一个帧。
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)   //设置一个画布，左上角坐标为(0,0)，右下角坐标为(2*size+1, 2*size+1)。
		img := image.NewPaletted(rect, palette)        //创建一张 基于调色板的图像。
		for t := 0.0; t < cycles*2*math.Pi; t += res { //循环到时间完成2Π*周期数。
			x := math.Sin(t)                                                          //设定x坐标
			y := math.Sin(t*freq + phase)                                             //设定y坐标，括号内是频率乘时间加上相位差。
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex) //使用黑点渲染图像。
		}
		phase += 0.1                           //递增相位差，使每一帧的图像都不同，来达到动画的效果。
		anim.Delay = append(anim.Delay, delay) //存储延迟时间。
		anim.Image = append(anim.Image, img)   //存储图像。
	}
	err := gif.EncodeAll(out, &anim) //把内存中的 GIF 动画结构体编码成二进制 .gif 格式，并写到输出目标。输出一个err。
	if err != nil {
		log.Fatal(err)
	}

}
