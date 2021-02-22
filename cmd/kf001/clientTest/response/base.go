package response

type UserRegion struct {
	ID       int32        `json:"id"`
	ParentId int32        `json:"parent_id"`
	Name     string       `json:"name"`
	Loc      []float64    `json:"loc"`
	Children []UserRegion `json:"children"`
}