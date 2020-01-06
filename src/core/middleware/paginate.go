package core

import (
	"context"
	"net/http"
	"strconv"
)

// PaginateParams Paginate required params
type PaginateParams struct {
	Index int64
	Limit int64
}

const (
	// ParamCtx param context index
	ParamCtx key = iota
)

// PaginateHandler Sets paginate required params to context
func PaginateHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get page query value or set default
		page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
		if page <= 0 || err != nil {
			page = 1
		}

		// Get limit query value or set default
		limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
		if limit <= 0 || err != nil || limit > 100 {
			limit = 5
		}

		// Pagination algorithm
		// f(x)=(x(n))-x
		// Where x = limit and n = page
		index := (limit * page) - limit

		r.URL.Query().Set("limit", strconv.FormatInt(limit, 10))
		ctx := context.WithValue(r.Context(), ParamCtx, &PaginateParams{index, limit})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
