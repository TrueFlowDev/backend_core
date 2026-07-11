package valueobject

var employeeCategory = PermissionCategory("کارکنان")

var (
	PermissionEmployeeList = Permission{
		category: employeeCategory,
		value:    "employee.list",
		title:    "مشاهده فهرست کارکنان",
	}

	PermissionEmployeeRead = Permission{
		category: employeeCategory,
		value:    "employee.read",
		title:    "مشاهده اطلاعات کارکنان",
	}

	PermissionEmployeeCreate = Permission{
		category: employeeCategory,
		value:    "employee.create",
		title:    "ایجاد کارمند",
	}

	PermissionEmployeeUpdate = Permission{
		category: employeeCategory,
		value:    "employee.update",
		title:    "ویرایش کارمند",
	}

	PermissionEmployeeDelete = Permission{
		category: employeeCategory,
		value:    "employee.delete",
		title:    "حذف کارمند",
	}
)

var EmployeePermissions = buildPermissionMap(
	PermissionEmployeeList,
	PermissionEmployeeRead,
	PermissionEmployeeCreate,
	PermissionEmployeeUpdate,
	PermissionEmployeeDelete,
)
