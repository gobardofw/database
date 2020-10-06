package database

// NewPaginator create a new paginator instance
func NewPaginator(queryString string) Paginator {
	p := new(pagination)
	p.init([]uint8{10, 25, 50, 100}, 25, []string{"id"}, "id", queryString)
	return p
}

// NewPaginatorWithDefaults create a new paginator instance
func NewPaginatorWithDefaults(limits []uint8, defaultLimit uint8, sorts []string, defaultSort string, queryString string) Paginator {
	p := new(pagination)
	p.init(limits, defaultLimit, sorts, defaultSort, queryString)
	return p
}
