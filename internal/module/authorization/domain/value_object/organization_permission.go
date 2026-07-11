package value_object

var (
	PermissionOrganizationList   = Permission{"organization.list"}
	PermissionOrganizationRead   = Permission{"organization.read"}
	PermissionOrganizationCreate = Permission{"organization.create"}
	PermissionOrganizationUpdate = Permission{"organization.update"}
	PermissionOrganizationDelete = Permission{"organization.delete"}
)

var OrganizationPermissions = map[string]Permission{
	"organization.list":   PermissionOrganizationList,
	"organization.read":   PermissionOrganizationRead,
	"organization.create": PermissionOrganizationCreate,
	"organization.update": PermissionOrganizationUpdate,
	"organization.delete": PermissionOrganizationDelete,
}
