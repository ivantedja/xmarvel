package entity

type CharacterCollection struct {
	Code   uint   `json:"code"`
	Status string `json:"status"`
	Data   CharacterData
}

type CharacterData struct {
	Offset  uint `json:"offset"`
	Limit   uint `json:"limit"`
	Total   uint `json:"total"`
	Count   uint `json:"count"`
	Results []Character
}

type Character struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Modified    string `json:"-"`
}
