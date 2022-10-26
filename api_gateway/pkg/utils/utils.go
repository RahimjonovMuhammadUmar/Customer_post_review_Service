package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type QueryParams struct {
	Filters  map[string]string
	Page     int64
	Limit    int64
	Ordering []string
	Search   string
}

func ParseQueryParams(queryParams map[string][]string) (*QueryParams, []string) {
	params := QueryParams{
		Filters:  make(map[string]string),
		Page:     1,
		Limit:    10,
		Ordering: []string{},
		Search:   "",
	}
	var errStr []string
	var err error

	for key, value := range queryParams {
		if key == "page" {
			params.Page, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `page` param")
			}
			continue
		}

		if key == "limit" {
			params.Limit, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `limit` param")
			}
			continue
		}

		if key == "search" {
			params.Search = value[0]
			continue
		}

		if key == "ordering" {
			params.Ordering = strings.Split(value[0], ",")
			continue
		}
		params.Filters[key] = value[0]
	}

	return &params, errStr
}

func RandomNum(num ...int) int {
	var leng int
	if len(num) == 0 {
		leng = 7
	} else {
		leng = num[0]
	}
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	code := 0
	for i := 0; i <= leng; i++ {
		r := rnd.Intn(10)

		code *= 10
		code += r

	}
	return code
}
