package identity

import "go-identity-manager/internal/platform/acl"

type IdentityService struct {
	ACAPyACL acl.ACAPyACL
}

func NewService(ACAPyACL acl.ACAPyACL) *IdentityService {
	return &IdentityService{
		ACAPyACL: ACAPyACL,
	}
}

func (s IdentityService) Create() {}
