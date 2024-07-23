package model

// Student represents the structure of the document in the index
type Student struct {
	FavoriteSubject string `json:"Favorite Subject"`
	Gender          string `json:"Gender"`
	House           string `json:"House"`
	Name            string `json:"Name"`
	Specialty       string `json:"Specialty"`
	Spell           string `json:"Spell"`
	Story           string `json:"Story"`
	WandType        string `json:"Wand Type"`
	Year            int    `json:"Year"`
}

type QueryByIdRequest struct {
	DocId string `json:"id"`
}

type QueryByParametersRequest struct {
	House     string `json:"House"`
	Specialty string `json:"Specialty"`
	Spell     string `json:"Spell"`
	WandType  string `json:"Wand Type"`
}
