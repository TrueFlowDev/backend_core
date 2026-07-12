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

	PermissionRoleUpdate = Permission{
		category: roleCategory,
		value:    "role.update",
		title:    "ویرایش نقش",
	}

	PermissionRoleDelete = Permission{
		category: roleCategory,
		value:    "role.delete",
		title:    "حذف نقش",
	}
)

var RolePermissions = buildPermissionMap(
	PermissionRoleList,
	PermissionRoleCreate,
	PermissionRoleUpdate,
	PermissionRoleDelete,
)
