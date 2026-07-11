package value_object

var (
	PermissionEmployeeList = Permission{
		value: "employee.list",
		title: "مشاهده فهرست کارکنان",
	}

	PermissionEmployeeRead = Permission{
		value: "employee.read",
		title: "مشاهده اطلاعات کارکنان",
	}

	PermissionEmployeeCreate = Permission{
		value: "employee.create",
		title: "ایجاد کارمند",
	}

	PermissionEmployeeUpdate = Permission{
		value: "employee.update",
		title: "ویرایش کارمند",
	}

	PermissionEmployeeDelete = Permission{
		value: "employee.delete",
		title: "حذف کارمند",
	}
)

var EmployeePermissions = map[string]Permission{
	"employee.list":   PermissionEmployeeList,
	"employee.read":   PermissionEmployeeRead,
	"employee.create": PermissionEmployeeCreate,
	"employee.update": PermissionEmployeeUpdate,
	"employee.delete": PermissionEmployeeDelete,
}
