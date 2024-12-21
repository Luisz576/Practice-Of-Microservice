package config

import "fmt"

const CATEGORY_API_ADDRESS = "http://host.docker.internal:5000"

type ApiEndpoint int

const CategoryApiEndpointGetAll ApiEndpoint = 1
const CategoryApiEndpointCreate ApiEndpoint = 2

func GetCategoryEndpoint(endpoint ApiEndpoint) string {
	switch endpoint {
	// 1 - get all
	// 2 - create
	case 1:
		fallthrough
	case 2:
		return fmt.Sprintf("%s/categories", CATEGORY_API_ADDRESS)
	}
	return CATEGORY_API_ADDRESS
}
