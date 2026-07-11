package value_object

var (
	PermissionEmployeeList   = Permission{"employee.list"}
	PermissionEmployeeRead   = Permission{"employee.read"}
	PermissionEmployeeCreate = Permission{"employee.create"}
	PermissionEmployeeUpdate = Permission{"employee.update"}
	PermissionEmployeeDelete = Permission{"employee.delete"}
)

var EmployeePermissions = map[string]Permission{
	"employee.list":   PermissionEmployeeList,
	"employee.read":   PermissionEmployeeRead,
	"employee.create": PermissionEmployeeCreate,
	"employee.update": PermissionEmployeeUpdate,
	"employee.delete": PermissionEmployeeDelete,
}
