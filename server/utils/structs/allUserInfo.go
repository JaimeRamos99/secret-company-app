package structs

type AllUserInfo struct {
	GeneralInfo         InfoArray    `json:"generalInfo,omitempty"`
	SameIps             DataSameIps  `json:"sameIps,omitempty"`
	RecommendedProducts []TopProduct `json:"recommendedProducts,omitempty"`
}

func NewAllUserInfo(generalInfo *InfoArray, sameIps *DataSameIps, recommendedProducts []TopProduct) *AllUserInfo {
	return &AllUserInfo{
		GeneralInfo:         *generalInfo,
		SameIps:             *sameIps,
		RecommendedProducts: recommendedProducts,
	}
}
