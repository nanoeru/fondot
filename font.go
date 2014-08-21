// Copyright 2010 The Freetype-Go Authors. All rights reserved.
// Use of this source code is governed by your choice of either the
// FreeType License or the GNU General Public License version 2 (or
// any later version), both of which can be found in the LICENSE file.

package fondot

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	//"io/ioutil"
	"log"
	"os"
	"sync"
	"unicode/utf8"

	"github.com/scpayson/freetype-go/freetype"
)

var (
	fontfile = "misaki_mincho.ttf"
	//fontfile = "PixelMplus12-Regular.ttf"
	//fontfile = "HuiFont109.ttf"
	//	font     *truetype.Font
	font, _ = freetype.ParseFont(nil)
	once    = new(sync.Once)
)

type Drawer func(px, py, x, y int, flag bool) string

func Draw(text string, dotFunc Drawer) {
	//	初期設定
	dpi := 72
	//	size := int(8)
	size := int(16)
	spacing := 1.0
	texts := []string{text}

	once.Do(func() {
		////	フォント読み込み
		//fontBytes, err := ioutil.ReadFile(*fontfile)
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
		//font, err := freetype.ParseFont(fontBytes)
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
		//	フォント読み込み
		fontBytes, err := Asset("data/" + fontfile)
		if err != nil {
			log.Println(err)
			return
		}

		font, err = freetype.ParseFont(fontBytes)
		if err != nil {
			log.Println(err)
			return
		}
	})

	//	描画設定
	fg, bg := image.Black, image.White
	rgba := image.NewRGBA(image.Rect(0, 0, size*utf8.RuneCountInString(text), size))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(font)
	c.SetFontSize(float64(size))
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	// Draw the text.
	pt := freetype.Pt(0, 0+c.FUnitToPixelRU(font.UnitsPerEm()))
	for _, s := range texts {
		_, err := c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFix32(float64(size) * spacing)
	}

	//	画像から文字表示
	for jj, text := range texts {
		for j := jj * size; j < (jj+1)*size; j++ {
			for ii := 0; ii < utf8.RuneCountInString(text); ii++ {
				for i := ii * size; i < (ii+1)*size; i++ {
					r, _, _, _ := rgba.At(i, j).RGBA()
					fmt.Print(dotFunc(ii, jj, i, j, r != 65535))
				}
			}
			fmt.Println()
		}
	}

	//	画像保存
	f, err := os.Create("out.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//fmt.Println("Wrote out.png OK.")
}
