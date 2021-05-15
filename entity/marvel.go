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

type ErrNotFound struct {
	Message string
}

func (e ErrNotFound) Error() string {
	return e.Message
}

type ErrBadRequest struct {
	Message string
}

func (e ErrBadRequest) Error() string {
	return e.Message
}
