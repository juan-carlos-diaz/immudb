package server

import (
	"context"
	"github.com/codenotary/immudb/pkg/api/schema"
	"github.com/codenotary/immudb/pkg/auth"
	"github.com/codenotary/immudb/pkg/errors"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *ImmuServer) OpenSession(ctx context.Context, r *schema.OpenSessionRequest) (*schema.OpenSessionResponse, error) {
	if !s.Options.auth {
		return nil, errors.New(ErrAuthDisabled).WithCode(errors.CodProtocolViolation)
	}

	u, err := s.getValidatedUser(r.User, r.Password)
	if err != nil {
		return nil, errors.Wrap(err, ErrInvalidUsernameOrPassword)
	}

	if !u.Active {
		return nil, errors.New(ErrUserNotActive)
	}

	var token string

	s.dbList.GetId(r.DatabaseName)

	if s.multidbmode {
		token, err = auth.GenerateToken(*u, -1, s.Options.TokenExpiryTimeMin)
	} else {
		token, err = auth.GenerateToken(*u, s.dbList.GetId(r.DatabaseName), s.Options.TokenExpiryTimeMin)
	}
	if err != nil {
		return nil, err
	}

	if u.Username == auth.SysAdminUsername {
		u.IsSysAdmin = true
	}

	s.addUserToLoginList(u)

	return &schema.OpenSessionResponse{
		SessionID:  token,
		ServerUUID: s.UUID.String(),
	}, nil
}

func (s *ImmuServer) CloseSession(ctx context.Context, request *schema.CloseSessionRequest) (*empty.Empty, error) {
	if !s.Options.auth {
		return nil, errors.New(ErrAuthDisabled).WithCode(errors.CodProtocolViolation)
	}

	_, user, err := s.getLoggedInUserdataFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	s.removeUserFromLoginList(user.Username)

	_, err = auth.DropTokenKeysForCtx(ctx)

	return new(empty.Empty), err
}