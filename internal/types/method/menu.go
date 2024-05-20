package method

type Menu interface {
	// with tree as (
	//	   select id,
	//	          parent_id,
	//	          convert(nvarchar(150), concat('.', id, '.')) as path
	//	   from @@table where parent_id = 0
	//	   union all
	//	   select data.id,
	//	          data.parent_id,
	//	          convert(nvarchar(150), concat(tree.path, '-', '.', data.id, '.')) as path
	//	   from @@table data
	//	   join tree tree on data.parent_id = tree.id and tree.path not like '%' + concat('.', data.id, '.') + '%'
	// )
	// select distinct convert(bigint, replace(s.value, '.', '')) as id from tree t
	// cross apply string_split(t.path, '-') s
	// where t.id in @menuIDs
	QueryAllRelated(menuIDs []uint) ([]uint, error)
}
