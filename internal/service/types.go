package service

type CreateParams struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}

type UpdateParams struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Cost        *float64 `json:"cost"`
}

type DeleteParams struct {
	Id int `json:"id"`
}

type GetParams struct {
	Id int `json:"id"`
}

type Item struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}
