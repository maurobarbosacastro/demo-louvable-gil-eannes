package forms

type FreeTranform struct {
	Id        string  `form:"id" json:"id" binding:"required"`
	Width     int     `form:"width" json:"width" binding:"required"`
	Height    int     `form:"height" json:"height" binding:"required"`
	Blur      bool    `form:"blur" json:"blur"`
	BlurValue float64 `form:"blurValue" json:"blurValue"`
}
