package valueobject

var roleCategory PermissionCategory = "نقش‌ها"

var (
	PermissionRoleList = Permission{
		category: roleCategory,
		value:    "role.list",
		title:    "مشاهده فهرست نقش‌ها",
	}

	PermissionRoleCreate = Permission{
		category: roleCategory,
		value:    "role.create",
		title:    "ایجاد نقش",
	}
)

var RolePermissions = buildPermissionMap(
	PermissionRoleList,
	PermissionRoleCreate,
)
