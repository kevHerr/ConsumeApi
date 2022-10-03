package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var championName = "Garen"

type Response struct {
	Tipo    string `json:"type"`
	Format  string `json:"format"`
	Version string `json:"version"`
	Data    Data   `json:"data"`
}

type Data struct {
	Champion map[string]Champion
}
type Champion struct {
	Version string   `json:"version"`
	Id      string   `json:"id"`
	Key     int      `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Spells  []Spells `json:"spells"`
}
type Info struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty"`
}

type Stats struct {
	Hp                   int `json:"hp"`
	HpPerlevel           int `json:"hpperlevel"`
	Mp                   int `json:"mp"`
	MpPerlevel           int `json:"mpperlevel"`
	Movespeed            int `json:"movespeed"`
	Armor                int `json:"armor"`
	ArmorPerlevel        int `json:"armorperlevel"`
	SpellBlock           int `json:"spellblock"`
	SpellBlockPerlevel   int `json:"spellblockperlevel"`
	AttackRange          int `json:"attacktrange"`
	HpRegen              int `json:"hpregen"`
	HpRegenPerlevel      int `json:"hpregenperlevel"`
	MpRegen              int `json:"mpregen"`
	MpRegenPerlevel      int `json:"mpregenperlevel"`
	Crit                 int `json:"crit"`
	CritPerlevel         int `json:"critperlevel"`
	AttackDamage         int `json:"attackdamage"`
	AttackDamagePerlevel int `json:"attackdamageperlevel"`
	AttackSpeedPerlevel  int `json:"attackspeedperlevel"`
	AttackSpeed          int `json:"attackspeed"`
}

type Spells struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tooltip     string `json:"tooltip"`
}

func find(s string, d map[string]interface{}) string {
	r := ""
	for k, v := range d {
		if k == s {
			r += fmt.Sprintf("%s", v) + ";"
		}
		if c, ok := v.(map[string]interface{}); ok {
			r += find(s, c)
		}
	}
	return r
}

func main() {
	responseOne, err := http.Get("http://ddragon.leagueoflegends.com/cdn/12.17.1/data/en_US/champion/Aatrox.json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseAll, err := http.Get("http://ddragon.leagueoflegends.com/cdn/12.17.1/data/en_US/champion.json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(responseOne.Body)
	if err != nil {

		log.Fatal(err)

	}

	responsaDataAll, err := ioutil.ReadAll(responseAll.Body)
	if err != nil {
		log.Fatal(err)
	}

	//var ObjectResponseMany Response
	d := map[string]interface{}{}
	json.Unmarshal([]byte(responsaDataAll), &d)

	fmt.Println(find("data", d))

	var ObjectResponse Response
	json.Unmarshal([]byte(responseData), &ObjectResponse)

	//fmt.Println(ObjectResponse.Data.Champion.Name)
	//fmt.Println(ObjectResponse.Data.Champion.Tags)
	//fmt.Println(ObjectResponse.Data.Champion.Partype)
	//fmt.Println(ObjectResponse.Data.Champion.Spells)
}
