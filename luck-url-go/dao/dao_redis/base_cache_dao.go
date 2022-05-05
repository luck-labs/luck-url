package dao_redis

type BaseDao interface {
}

type baseDao struct {
}

var baseDaoInstance = &baseDao{}

func NewBaseDao() BaseDao {
	return baseDaoInstance
}
