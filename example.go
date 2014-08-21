package fondot

import (
	"github.com/nanoeru/tcol256"
)

//	 単色
func Mono(px, py, x, y int, flag bool) string {
	if flag {
		return tcol256.BgString(" ", tcol256.GREEN)
	} else {
		return " "
	}
}

//	単純カラー
func SimpleCol(px, py, x, y int, flag bool) string {
	if flag {
		return tcol256.BgString(" ", px%(tcol256.COL_MAX-1)+1) //tcol256.Bg256String(" ", (px+x)*2+64, (px+y)*3+160, (x+y)*3+96)
	} else {
		return " "
	}
}

//	 グラデーション
func Grad(px, py, x, y int, flag bool) string {
	if flag {
		return tcol256.Bg256String(" ", x*6+32, (x/3+y)*8+24, (x+y)*7+128)
	} else {
		return " "
	}
}

//	 グラデーション
func GradMono(px, py, x, y int, flag bool) string {
	if flag {
		return tcol256.Bg256String(" ", (32+x/2+y)%128+16, (96+x/2+y*2)%64+192, (64+x/4+y*3)%128+32)
	} else {
		return " "
	}
}

//func Grad2(px, py, x, y int, flag bool) string {
//	if flag {
//		return tcol256.Bg256String(" ", (px+x)*16+128, (py+y)*16+128, (px+py+x+y)*16+128)
//	} else {
//		return " "
//	}
//}
