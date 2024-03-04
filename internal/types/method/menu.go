package method

type Menu interface {
	// with tree as (
	//     select id,
	//            parent_id,
	//            convert(nvarchar(150),CONCAT('.',id,'.')) as path
	//     from @@table where parent_id = 0
	//     union all
	//     select data.id,
	//            data.parent_id,
	//            Convert(nvarchar(150),CONCAT(tree.Path , '-','.',data.id,'.')) as Path
	//     from @@table data
	//     join tree tree on data.parent_id = tree.id and tree.path not like '%' + CONCAT('.',data.Id,'.') + '%'
	// )
	// select distinct convert(bigint, replace(s.value, '.', '')) as id from tree t
	// cross apply STRING_SPLIT(t.path, '-') s
	// where t.id in @menuIDs
	QueryAllRelated(menuIDs []uint) ([]uint, error)
}
