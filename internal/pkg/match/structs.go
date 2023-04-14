package match

type UserRes struct {
	UserId uint   `json:"UserId" structs:"userId"`
	Name   string `json:"name" structs:"name"`
	Photo  uint   `json:"avatar" structs:"avatar"`
	Shown  bool   `json:"shown" structs:"shown"`
}

type ReactionInp struct {
	EvaluatedUserId uint   `json:"evaluatedUserId"`
	Reaction        string `json:"reaction"`
}

type ChatAnswer struct {
	Uint  string `json:"email" structs:"email"`
	Name  string `json:"name" structs:"name"`
	Photo uint   `json:"avatar" structs:"avatar"`
}
