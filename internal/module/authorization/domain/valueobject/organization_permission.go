package valueobject

var (
	PermissionOrganizationList = Permission{
		value: "organization.list",
		title: "مشاهده فهرست سازمان‌ها",
	}

	PermissionOrganizationRead = Permission{
		value: "organization.read",
		title: "مشاهده اطلاعات سازمان",
	}

	PermissionOrganizationCreate = Permission{
		value: "organization.create",
		title: "ایجاد سازمان",
	}

	PermissionOrganizationUpdate = Permission{
		value: "organization.update",
		title: "ویرایش سازمان",
	}

	PermissionOrganizationDelete = Permission{
		value: "organization.delete",
		title: "حذف سازمان",
	}
)

var OrganizationPermissions = map[string]Permission{
	"organization.list":   PermissionOrganizationList,
	"organization.read":   PermissionOrganizationRead,
	"organization.create": PermissionOrganizationCreate,
	"organization.update": PermissionOrganizationUpdate,
	"organization.delete": PermissionOrganizationDelete,
}
