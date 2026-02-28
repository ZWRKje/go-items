package service

type CreateParams struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Cost        float64 `json:"value"`
}

type UpdateParams struct {
	Id          int      `json:"id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Cost        *float64 `json:"value"`
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
	Cost        float64 `json:"value"`
}
