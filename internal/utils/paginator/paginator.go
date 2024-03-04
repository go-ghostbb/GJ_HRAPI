package paginator

import (
	"gorm.io/gen"
)

type Pagination struct {
	Total  int64
	Limit  int
	Offset int
}

func New(limit int, offset int) *Pagination {
	res := &Pagination{
		Limit:  limit,
		Offset: offset,
	}
	return res
}

func (p *Pagination) Where(conds ...gen.Condition) func(gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		if len(conds) > 0 {
			dao.Where(conds...)
		}
		p.Total, _ = dao.Count()

		if p.Limit == 0 {
			return dao
		}
		return dao.Offset(p.Offset).Limit(p.Limit)
	}
}
