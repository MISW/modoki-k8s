package handler

import (
	"context"

	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/store/apps"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"github.com/modoki-paas/modoki-k8s/pkg/rbac/permissions"
	"github.com/modoki-paas/modoki-k8s/pkg/types"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppServer struct {
	Context *ServerContext
}

var _ api.AppServer = &AppServer{}

func (s *AppServer) Create(ctx context.Context, req *api.AppCreateRequest) (res *api.AppCreateResponse, err error) {
	if err := auth.IsAuthorized(ctx, permissions.AppCreate); err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	tx, err := s.Context.DB.BeginTxx(ctx, nil)

	if err != nil {
		return nil, xerrors.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			if err := tx.Commit(); err != nil {
				err = xerrors.Errorf("failed to commit transaction: %w", err)
				res = nil
			}
		}
	}()

	store := apps.NewAppStore(tx)

	app := &types.App{
		Owner: req.Spec.Owner,
		Name:  req.Spec.Name,
		Spec:  (*types.AppSpec)(req.Spec),
	}

	seq, err := store.AddApp(app)

	if err != nil {
		return nil, xerrors.Errorf("failed to store app config in db: %w", err)
	}

	app, err = store.GetApp(seq)

	if err != nil {
		return nil, xerrors.Errorf("failed to get app config in db: %w", err)
	}

	return &api.AppCreateResponse{
		Id:   app.ID,
		Spec: req.GetSpec(),
	}, nil
}
