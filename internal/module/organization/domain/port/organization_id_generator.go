package port

import "github.com/TrueFlowDev/Backend/internal/module/organization/domain/value_object"

type OrganizationIDGenerator interface {
	Generate() value_object.OrganizationID
}
