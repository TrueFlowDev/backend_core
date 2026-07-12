package valueobject

var organizationCategory = PermissionCategory("سازمان")

var (
	PermissionOrganizationList = Permission{
		category: organizationCategory,
		value:    "organization.list",
		title:    "مشاهده فهرست سازمان‌ها",
	}

	PermissionOrganizationRead = Permission{
		category: organizationCategory,
		value:    "organization.read",
		title:    "مشاهده اطلاعات سازمان",
	}

	PermissionOrganizationCreate = Permission{
		category: organizationCategory,
		value:    "organization.create",
		title:    "ایجاد سازمان",
	}

	PermissionOrganizationUpdate = Permission{
		category: organizationCategory,
		value:    "organization.update",
		title:    "ویرایش سازمان",
	}

	PermissionOrganizationDelete = Permission{
		category: organizationCategory,
		value:    "organization.delete",
		title:    "حذف سازمان",
	}
)

var OrganizationPermissions = buildPermissionMap(
	PermissionOrganizationList,
	PermissionOrganizationRead,
	PermissionOrganizationCreate,
	PermissionOrganizationUpdate,
	PermissionOrganizationDelete,
)
