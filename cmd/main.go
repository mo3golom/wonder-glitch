package main

import (
	"bytes"
	"fmt"
	"github.com/mo3golom/wonder-glitch/pkg/loader"
	"github.com/mo3golom/wonder-glitch/wonderGlitchDTO"
	"github.com/mo3golom/wonder-glitch/wonderGlitchService"
	"image/png"
	"io/ioutil"
)

func main() {
	img, _ := loader.LoadImage("test.png")
	glitch := wonderGlitchService.NewGlitchService(
		wonderGlitchService.NewEffectHandlerBus(
			wonderGlitchDTO.GetTypesList(),
		),
	)

	inputEffects := []wonderGlitchDTO.InputEffect{
		{Id: "colorRose", Sort: 0},
		{Id: "colorGoldenFizz", Sort: 2},
		{Id: "shift", Sort: 10},
	}

	for i := 10; i <= 20; i++ {
		newImg := glitch.SetDest(img).SetFactor(8.0).Glitchify(inputEffects)

		buf := new(bytes.Buffer)
		_ = png.Encode(buf, newImg)
		_ = ioutil.WriteFile(fmt.Sprintf("result%v.png", i), buf.Bytes(), 0644)
	}
}

//