package oalib

/**
* @name: auth
* @author: veypi <i@veypi.com>
* @date: 2021-11-17 16:49
* @description：auth
**/

// AuthLevel 权限等级
// 对于操作类权限
// 0 禁止执行
// 1 允许执行
// 对于资源类权限
// 0 相当于没有
// 1 有限读权限
// 2 读权限
// 3 创建权限
// 4 修改权限
// 5 删除权限
// 6 赋予其余人权限
type AuthLevel uint

const (
	AuthNone AuthLevel = 0
	AuthDo   AuthLevel = 1
	// AuthPart TODO: 临时权限
	AuthPart   AuthLevel = 1
	AuthRead   AuthLevel = 2
	AuthCreate AuthLevel = 3
	AuthUpdate AuthLevel = 4
	AuthDelete AuthLevel = 5
	AuthAll    AuthLevel = 6
)

func (a AuthLevel) Upper(b AuthLevel) bool {
	return a > b
}

func (a AuthLevel) CanDo() bool {
	return a > AuthNone
}

func (a AuthLevel) CanRead() bool {
	return a >= AuthRead
}

func (a AuthLevel) CanCreate() bool {
	return a >= AuthCreate
}

func (a AuthLevel) CanUpdate() bool {
	return a >= AuthUpdate
}

func (a AuthLevel) CanDelete() bool {
	return a >= AuthDelete
}

func (a AuthLevel) CanDoAny() bool {
	return a >= AuthAll
}
