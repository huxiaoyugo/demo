package main

import (
	"os"
	"fmt"
	"strings"
	"net/http"
	"bytes"
	"image/jpeg"
	"image"
	"image/draw"
	"path"
	"bufio"
	"io"
	"github.com/skip2/go-qrcode"
	"image/png"
	"io/ioutil"
	"image/color"
)

// 图片处理

const (
	BasePath = "/Users/huxiaoyu/gocode/src/demo/imageHandler/"
)


func ImgConfig(r io.Reader) (image.Config, error) {
	cfg,err:= jpeg.DecodeConfig(r)
	if err != nil {
		return cfg, err
	}
	cfg.ColorModel = color.CMYKModel
	return cfg, nil
}


func main() {
	//url := "http://img.sccnn.com/bimg/337/24851.jpg"
	//getImg(url)
	//err := mergeQrCodeImg("https://hemaxiche.com",255, BasePath+"a4.jpg")
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 获取背景图片对象
	bgb, _ := os.Open(BasePath+"a4.jpg")
	defer bgb.Close()
	image.RegisterFormat("jpeg", "\xff\xd8", jpeg.Decode, ImgConfig)
	bg, _ := jpeg.Decode(bgb)

	watermarkb, _ := os.Open(BasePath+"tt.jpg")
	defer watermarkb.Close()
	watermark, _ := jpeg.Decode(watermarkb)

	mergeImg(bg, watermark)
}

func mergeImg(baseImg image.Image, overImg image.Image) error {
	col := color.RGBA{}
	//offset := image.Pt(200, 200)
	b := baseImg.Bounds()
	m := image.NewCMYK(b)  //NewRGBA(b)
	fmt.Printf("%T\n", m.ColorModel().Convert(col))
	//x := (baseImg.Bounds().Max.X - overImg.Bounds().Max.X) / 2
	//y := (baseImg.Bounds().Max.Y - overImg.Bounds().Max.Y) / 2

	//offset := image.Pt(x, y)
	fmt.Println(baseImg.Bounds())
	fmt.Println(overImg.Bounds())

	draw.Draw(m, b, baseImg, image.ZP, draw.Src)
	//draw.Draw(m, b.Add(offset), overImg, image.ZP, draw.Over)

	imgw, err := os.Create(BasePath+"watermarked.jpg")
	if err != nil {
		return err
	}
	jpeg.Encode(imgw, m, &jpeg.Options{100})
	defer imgw.Close()
	return nil
}

func mergeQrCodeImg(qrContent string, qrSize int, bgSrc string) error {

	// 二维码image对象
	img, err := GeneQrcodeImage(qrContent, qrSize)
	if err != nil {
		return err
	}

	// 获取背景图片对象
	bgb, err := os.Open(bgSrc)
	defer bgb.Close()
	if err != nil {
		return err
	}
	watermark, err := jpeg.Decode(bgb)
	if err != nil {
		return err
	}

	err = mergeImg(watermark, img)
	if err != nil {
		return err
	}
	return nil
}

// 生成二维码
func GeneQrcodeImage(content string, size int) (image.Image, error) {

	// 有两种方式：
	// 1、qrcode.WriteFile("http://www.flysnow.org/",qrcode.Medium,256,"./blog_qrcode.png")
	// 2、 qrcode.Encode(content, qrcode.High,size)
	byteArr, err := qrcode.Encode(content, qrcode.Highest, size)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(byteArr)
	return png.Decode(reader)
}

// 图片下载函数
func downImg(url string) {
	// 处理前缀为//的url
	if string([]rune(url)[:2]) == "//" {
		url = "http:" + url
		fmt.Println(url)
	}

	// 解决文件无图片格式后缀的问题
	fileName := path.Base(url)
	if !strings.Contains(fileName, ".png") {
		fileName += ".png"
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer resp.Body.Close()
	reader := bufio.NewReaderSize(resp.Body, 32*1024)

	file, _ := os.Create(BasePath + fileName)
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d\n", written)
}



func getImg(url string) ( reader *bytes.Reader, err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	}
	fmt.Println(name)
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	//n, err = io.Copy(out, bytes.NewReader(pix))

	reader = bytes.NewReader(pix)
	return
}
