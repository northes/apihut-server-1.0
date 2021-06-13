package server

import (
	"apihut-server/config"
	"apihut-server/constant"
	"apihut-server/model"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"

	"github.com/nullrocks/identicon"
)

func GetIdenticonAvatar(i *model.IdenticonAvatar) (string, error) {
	ig, err := identicon.New(
		i.Namespace, // Namespace
		i.Size,      // Number of blocks (Size)
		i.Density,   // Density
		identicon.SetRandom(i.Random),
		identicon.SetBackgroundColorFunction(func(cb []byte, fc color.Color) color.Color {
			if len(i.BackgroundColor) != 0 {
				// 透明
				if i.BackgroundColor == "transparent" {
					return color.Transparent
				}
				// 自定义颜色
				bc, err := hex2rgba(i.BackgroundColor)
				if err != nil {
					return defaultBackgroundColor()
				}
				return bc
			}
			// 默认颜色
			return defaultBackgroundColor()
		}),
		identicon.SetFillColorFunction(func(hashBytes []byte) color.Color {
			// 自定义颜色
			if len(i.FillColor) != 0 {
				fc, err := hex2rgba(i.FillColor)
				if err != nil {
					return defaultFillColor(hashBytes)
				}
				return fc
			}
			// 默认颜色
			return defaultFillColor(hashBytes)
		}),
	)
	if err != nil {
		panic(err)
	}

	ii, err := ig.Draw(i.Hash)
	if err != nil {
		panic(err) // Text is empty
	}

	path := fmt.Sprintf(
		"%s/img/identicon/%s-%d-%d-%d-%s.png",
		config.Conf.Path.Temp,
		i.Namespace,
		i.Size,
		i.Density,
		i.Pixels,
		i.Hash,
	)

	if i.Output == constant.Base64Output {
		out := new(bytes.Buffer)
		ii.Png(300, out) // 300px * 300px
		b64 := base64.StdEncoding.EncodeToString(out.Bytes())
		fmt.Println(b64)
		return "data:image/jpg;base64," + b64, nil
	}

	img, _ := os.Create(path)
	defer img.Close()
	ii.Png(i.Pixels, img) // 300px * 300px
	fmt.Println(img.Name())
	return path, nil
}

func hex2rgba(hexStr string) (color.Color, error) {
	hexc := strings.Split(hexStr, "#")
	hex := hexc[len(hexc)-1:][0]

	if len(hex) != 6 {
		fmt.Println(hexStr)
		return nil, errors.New("非法颜色值")
	}

	r, _ := strconv.ParseInt(hex[:2], 16, 10)
	g, _ := strconv.ParseInt(hex[2:4], 16, 18)
	b, _ := strconv.ParseInt(hex[4:], 16, 10)

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}, nil
}

func defaultFillColor(hashBytes []byte) color.Color {
	cb1, cb2 := uint32(hashBytes[0]), uint32(hashBytes[1])
	h := (cb1 + cb2) % 360
	s := (cb1 % 30) + 60
	l := (cb2 % 20) + 40

	// Some colors in the HSL color model are too bright and don't play well
	// with the default background color. This is a naïve normalization method.
	if (h >= 50 && h <= 85) || (h >= 170 && h <= 190) {
		s = 80
		l -= 20
	} else if h > 85 && h < 170 {
		l -= 10
	}

	return identicon.HSL{h, s, l}
}

func defaultBackgroundColor() color.Color {
	return color.NRGBA{R: 240, G: 240, B: 240, A: 255}
}
