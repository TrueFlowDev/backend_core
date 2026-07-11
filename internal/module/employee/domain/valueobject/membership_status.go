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

var membershipStatuses = buildMembershipStatusMap(
	MembershipStatusActive,
	MembershipStatusOnLeave,
	MembershipStatusSuspended,
	MembershipStatusResigned,
	MembershipStatusTerminated,
)

func buildMembershipStatusMap(statuses ...MembershipStatus) map[string]MembershipStatus {
	m := make(map[string]MembershipStatus, len(statuses))
	for _, s := range statuses {
		m[s.value] = s
	}
	return m
}

func ParseMembershipStatus(raw string) (MembershipStatus, error) {
	s, ok := membershipStatuses[raw]
	if !ok {
		return MembershipStatus{}, ErrInvalidMembershipStatus
	}
	return s, nil
}

func (s MembershipStatus) Value() string {
	return s.value
}
