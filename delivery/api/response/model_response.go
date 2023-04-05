package response

import "github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type SingleResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type PagedResponse struct {
	Status Status        `json:"status"`
	Data   []interface{} `json:"data"`
	Paging dto.Paging    `json:"paging,omitempty"`
}
