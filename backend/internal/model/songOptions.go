package model

type SongOptions struct {
	SpeedId        uint `json:"speedId"`
	IsVanishOn     bool `json:"isVanishOn"`
	IsInverseOn    bool `json:"isInverseOn"`
	RandomId       uint `json:"randomId"`
	IsSkipOn       bool `json:"isSkipOn"`
	IsVoiceOn      bool `json:"isVoiceOn"`
	SelectedToneId uint `json:"selectedToneId"`
	NotesPosition  int  `json:"notesPosition"`
}
