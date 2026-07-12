package port

import "github.com/TrueFlowDev/Backend/internal/module/organization/domain/valueobject"

type OrganizationIDGenerator interface {
	Generate() valueobject.OrganizationID
}
