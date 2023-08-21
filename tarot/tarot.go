package tarot

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"newTarot/tarots"
	"os"
	"strconv"
	"time"
)

var tarotsMap map[string]*TarotCard

type TarotCard struct {
	Id       string `json:"id"`
	Tarot    string `json:"tarot"`
	Detailed string `json:"detailed"`
}

type tarotCardList struct {
	TarotCards []TarotCard `json:"tarots"`
}

type CardLog struct {
	CardTime string `json:"cardTime"`
	CardId   string `json:"cardId"`
}

func JsonToMap() {
	file, err := tarots.TarotJson.Open("json/tarot.json")
	if err != nil {
		panic("open tarot cardjson err:" + err.Error())
	}

	data, err := io.ReadAll(file)
	if err != nil {
		panic("readall err:" + err.Error())
	}

	ts := tarotCardList{TarotCards: make([]TarotCard, 0, 52)}

	err = json.Unmarshal(data, &ts)
	if err != nil {
		panic("Unmarshal err:" + err.Error())
	}

	tarotsMap = make(map[string]*TarotCard, 52)

	for _, v := range ts.TarotCards {
		v := v
		tarotsMap[v.Id] = &TarotCard{
			Id:       v.Id,
			Tarot:    v.Tarot,
			Detailed: v.Detailed,
		}

	}

}

func GetTarotCard() *TarotCard {

	if cl, ok := CheckLog(); !ok {
		randomNum := rand.Intn(51) + 1
		sid := "tarot" + strconv.Itoa(randomNum) + ".jpg"
		AppendLog(time.Now(), tarotsMap[sid].Id)
		return tarotsMap[sid]
	} else {
		return &TarotCard{
			Id:       cl.CardId,
			Tarot:    tarotsMap[cl.CardId].Tarot,
			Detailed: tarotsMap[cl.CardId].Detailed,
		}
	}

}

func AppendLog(Logtime time.Time, CardId string) {

	cl := CardLog{
		CardTime: Logtime.Format(time.DateTime),
		CardId:   CardId,
	}

	bytes, err := json.Marshal(&cl)
	if err != nil {
		fmt.Println("Append log err:", err)
		return
	}

	s := string(bytes) + "\n"

	file, err := os.OpenFile("checkLog.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)

	if err != nil {
		fmt.Println("OpenFile log err:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(s)
	if err != nil {
		fmt.Println("File WriteString log err:", err)
		return
	}

}

func CheckLog() (*CardLog, bool) {

	line := GetLastLine()
	if len(line) <= 0 {
		return nil, false
	}
	cl := new(CardLog)
	json.Unmarshal(line, &cl)

	//刷新时间
	timeformat := time.Now().Format(time.DateOnly)
	timeformat += "04:00:00"

	parseNow, _ := time.Parse(time.DateTime, timeformat)

	parseLog, _ := time.Parse(time.DateTime, cl.CardTime)

	if parseLog.Before(parseNow) {
		if parseNow.Day() == parseLog.Day() {
			return cl, true
		}
		return nil, false
	}

	return cl, true

}

func GetLastLine() []byte {

	file, err := os.OpenFile("checkLog.txt", os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("checkLog open file err", err)
		return nil
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	//一行一行的读取
	bytes := []byte{}
	for reader.Scan() {
		bytes = reader.Bytes()
	}

	return bytes
}
