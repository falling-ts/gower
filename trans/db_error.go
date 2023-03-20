package trans

import "gower/services"

var DBError = services.TransCategory{
	"mysql": services.TransMap{
		`^Error 1046 \(3D000\):.*?$`:                        `没有选择数据库`,
		`^Error 1048 \(23000\): Column '(.*?)'.*?$`:         `字段${1}不能是NULL`,
		`^Error 1054 \(42S22\): Unknown column '(.*?)'.*?$`: `未知字段${1}`,
		`^Error 1062 \(23000\): Duplicate entry '(.*?)'.*$`: `${1}已存在`,
		`^Error 1146 \(42S02\): Table '(.*?)'.*?$`:          `数据表${1}不存在`,
		`^Error 1216 \(23000\):.*?$`:                        `无法添加或更新子行:外键约束失败`,
		`^Error 1217 \(23000\):.*?$`:                        `不能删除或更新父行:外键约束失败`,
		`^Error 1364 \(HY000\): Field '%s'.*?$`:             `字段${1}没有默认值`,
		`^Error 1451 \(23000\):.*?$`:                        `无法删除或更新父行:外键约束失败`,
		`^Error 1452 \(23000\):.*?$`:                        `无法添加或更新子行:外键约束失败`,
		`^Error 1064 \(42000\):.*?$`:                        `SQL语句语法错误`,
		`^Error 1100 \(HY000\):.*?$`:                        `在对一个未加锁的表进行操作时，需要先锁定表`,
		`^Error 1104 \(42000\):.*?$`:                        `SELECT操作失败，可能是由于使用了错误的列名或表名`,
		`^Error 1171 \(42000\):.*?$`:                        `BLOB或TEXT类型的列不能有默认值`,
		`^Error 1292 \(22007\):.*?$`:                        `插入或更新操作中，某个列的值与该列的数据类型不匹配`,
		`^Error 1406 \(22001\):.*?$`:                        `插入或更新操作中，某个列的值超过了该列的最大长度限制`,
	},
}
