package model

type CostumeOptions struct {
	CurrentBody     uint `json:"currentBody"`
	CurrentFace     uint `json:"currentFace"`
	CurrentHead     uint `json:"currentHead"`
	CurrentKigurumi uint `json:"currentKigurumi"`
	CurrentPuchi    uint `json:"currentPuchi"`
	ColorBody       uint `json:"colorBody"`
	ColorFace       uint `json:"colorFace"`
	ColorLimb       uint `json:"colorLimb"`
}
