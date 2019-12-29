package users

import "golang.org/x/xerrors"

var (
	// ErrUserIDDuplicates means user ids duplicate
	ErrUserIDDuplicates = xerrors.New("user ids duplicate")

	// ErrUnknownUser means user is unknown
	ErrUnknownUser = xerrors.New("unknown user")
)
