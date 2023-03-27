package domain

type Role struct {
	Name       string           // 唯一
	Permission []PermissionItem //  允许访问的资源列表
}

type PermissionItem = string
