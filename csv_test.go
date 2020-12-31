package gnet

import (
	"fmt"
	"testing"
)

func TestReadConfigFromCSV(t *testing.T) {

	err, data := ReadConfigFromCSV("e:\\npc.csv", 2, 4, &GenConfigObj{
		GenObjFun:   GenNpcObj,
	})

	if err != nil {
		fmt.Println(err)
	}
	for i, v:=range data{
		fmt.Printf("%d %v - %v \n",i+1, v.(*Npc).Id, v.(*Npc).Name)
	}
}

type HeroExp struct {
	Lv uint32 `json:"-"`
	Exp uint32 `json:"-"`
}

func GenHeroExpObj() interface{}{
	return &HeroExp{}
}

type Npc struct {
	Id uint32 `json:"id"`
	Name string `json:"-"`
	Icon string `json:"-"`
	Model string `json:"-"`
	Bloodheight uint32 `json:"-"`
	Numheight uint32 `json:"-"`
	Modelnum uint32 `json:"-"`
	Type uint32 `json:"-"`
	Level uint32 `json:"-"`
	Hp uint32 `json:"-"`
	Attack uint32 `json:"-"`
	Defense uint32 `json:"-"`
	Magic uint32 `json:"-"`
	Mdef uint32 `json:"-"`
	Speed uint32 `json:"-"`
	Crit uint32 `json:"-"`
	Crithurt uint32 `json:"-"`
	Chase uint32 `json:"-"`
	Control uint32 `json:"-"`
	Potential string `json:"-"`
	Camp uint32 `json:"-"`
	Sex uint32 `json:"-"`
	Skill uint32 `json:"-"`
}

func GenNpcObj() interface{}{
	return &Npc{}
}