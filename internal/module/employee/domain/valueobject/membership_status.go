package valueobject

import (
	"github.com/Ali127Dev/xerr"
)

var (
	ErrInvalidMembershipStatus = xerr.New(
		xerr.CodeBadRequest,
		xerr.WithMeta("membership_status", xerr.ErrorReasonInvalidFormat),
	)
)

type MembershipStatus struct {
	value string
}

var (
	MembershipStatusActive     = MembershipStatus{"active"}
	MembershipStatusOnLeave    = MembershipStatus{"on_leave"}
	MembershipStatusSuspended  = MembershipStatus{"suspended"}
	MembershipStatusResigned   = MembershipStatus{"resigned"}
	MembershipStatusTerminated = MembershipStatus{"terminated"}
)

func NewMembershipStatus(raw string) (MembershipStatus, error) {
	switch raw {
	case MembershipStatusActive.value:
		return MembershipStatusActive, nil
	case MembershipStatusOnLeave.value:
		return MembershipStatusOnLeave, nil
	case MembershipStatusSuspended.value:
		return MembershipStatusSuspended, nil
	case MembershipStatusResigned.value:
		return MembershipStatusResigned, nil
	case MembershipStatusTerminated.value:
		return MembershipStatusTerminated, nil
	default:
		return MembershipStatus{}, ErrInvalidMembershipStatus
	}
}

func (s MembershipStatus) Value() string {
	return s.value
}
