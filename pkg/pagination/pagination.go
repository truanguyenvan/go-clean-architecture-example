package pagination

import (
	"fmt"
	"go-clean-architecture-example/pkg/utils/structure"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ExpressionFilter is struct for add filtering in slice or array
type ExpressionFilter struct {
	PropertyName string      `json:"property" validate:"required"`
	Value        interface{} `json:"value" validate:"required"`
	Comparison   Comparison  `json:"comparison" validate:"required"`
}

type ExpressionFilters struct {
	ExpressionFilters []ExpressionFilter `json:"expression_filters,omitempty"`
}

// Pagination query
type PaginationQuery struct {
	Size              int    `json:"size,omitempty"`
	Page              int    `json:"page,omitempty"`
	OrderBy           string `json:"order_by,omitempty"`
	IsDescending      bool   `json:"is_descending"`
	AndLogic          bool   `json:"and_logic,omitempty"`
	ExpressionFilters *ExpressionFilters
}

// Set page size
func (q *PaginationQuery) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		q.Size = defaultSize
		return nil
	}
	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}
	q.Size = n

	return nil
}

// Set page number
func (q *PaginationQuery) SetPage(pageQuery string) error {
	if pageQuery == "" {
		q.Size = 0
		return nil
	}
	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}
	q.Page = n

	return nil
}

// Set order by
func (q *PaginationQuery) SetOrderBy(orderByQuery string) {
	q.OrderBy = orderByQuery
}

// Set order by
func (q *PaginationQuery) SetIsDescending(isDescending bool) {
	q.IsDescending = isDescending
}

// Get offset
func (q *PaginationQuery) GetOffset() int {
	if q.Page == 0 {
		return 0
	}
	return (q.Page - 1) * q.Size
}

// Get limit
func (q *PaginationQuery) GetLimit() int {
	return q.Size
}

// Get OrderBy
func (q *PaginationQuery) GetOrderBy() string {
	return q.OrderBy
}

// Get OrderBy
func (q *PaginationQuery) GetPage() int {
	return q.Page
}

// Get OrderBy
func (q *PaginationQuery) GetSize() int {
	return q.Size
}

func (q *PaginationQuery) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&orderBy=%s", q.GetPage(), q.GetSize(), q.GetOrderBy())
}

// Get pagination query struct from
func GetPaginationFromCtx(c *fiber.Ctx, validator structure.Validator) (*PaginationQuery, error) {
	q := &PaginationQuery{}
	//binding query param
	if err := q.SetPage(c.Query("page")); err != nil {
		return nil, err
	}
	if err := q.SetSize(c.Query("size")); err != nil {
		return nil, err
	}
	q.SetOrderBy(c.Query("orderBy"))
	isDescending, err := strconv.ParseBool(c.Query("is-descending"))
	if err != nil {
		return nil, err
	}
	q.SetIsDescending(isDescending)

	filter := new(ExpressionFilters)
	if err := c.BodyParser(&filter); err != nil {
		return nil, err
	}
	if err := validator.Validate(filter); err != nil {
		return nil, err
	}
	q.ExpressionFilters = filter
	return q, nil
}

// Get total pages int
func GetTotalPages(totalCount int, pageSize int) int {
	d := float64(totalCount) / float64(pageSize)
	return int(math.Ceil(d))
}

// Get has more
func GetHasMore(currentPage int, totalCount int, pageSize int) bool {
	return currentPage < totalCount/pageSize
}
