package wonderGlitchDTO

import (
	effects "github.com/mo3golom/wonder-glitch/wonderGlitchService/effects"
)

type EffectType struct {
	Id     string
	Effect effects.EffectInterface
}

// GetTypesList Возвращает список доступных эффектов
func GetTypesList() []EffectType {
	mainEffect := &effects.EffectStruct{}

	return []EffectType{
		{Id: "shift", Effect: &effects.ShiftEffect{EffectStruct: mainEffect}},
		{Id: "colorRose", Effect: &effects.ColorEffect{HexColor: "#fe1172"}},
		{Id: "colorDarkBlue", Effect: &effects.ColorEffect{HexColor: "#140ace"}},
		{Id: "colorGoldenFizz", Effect: &effects.ColorEffect{HexColor: "#feec2d"}},
		{Id: "colorGoldenMalachite", Effect: &effects.ColorEffect{HexColor: "#12af4f"}},
		{Id: "colorMoodyBlue", Effect: &effects.ColorEffect{HexColor: "#777fd1"}},
	}
}
