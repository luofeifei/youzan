package baseMongo

type (
	IndexKey struct {
		Name string `json:"name"` // 字段名
		Key  string `json:"key"`  // 1(ASC) -1 (DESC) 2D 2DSphere GeoHayStack Hashed Text
	}
	IndexData struct {
		Keys   []IndexKey `json:"keys"`
		Unique bool       `json:"unique"` // 是否唯一
		Exp    int32      `json:"exp"`    // 过期时间
	}
)