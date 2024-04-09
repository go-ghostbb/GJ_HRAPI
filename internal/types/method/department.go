package method

import "gorm.io/gen"

type Department interface {
	// WITH tree AS (
	// SELECT * from @@table WHERE id = @id
	// UNION ALL
	// SELECT b.* FROM tree INNER JOIN @@table b ON b.id = tree.parent_id)
	// SELECT * FROM tree
	QueryUp(id uint) ([]*gen.T, error)
}
