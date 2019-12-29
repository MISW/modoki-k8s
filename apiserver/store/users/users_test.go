// +build use_external_db

package users

import (
	"github.com/modoki-paas/modoki-k8s/pkg/rbac/roles"
	"testing"
	"time"

	"github.com/modoki-paas/modoki-k8s/internal/testutil"
	"github.com/modoki-paas/modoki-k8s/pkg/types"
)

func TestAddUser(t *testing.T) {
	t.Run("success_normal", func(t *testing.T) {
		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewUserStore(db)

		seq, err := store.AddUser("test-id", "test-name", types.UserNormal, roles.SystemAdmin.Name)

		if err != nil {
			t.Fatalf("failed to add user: %v", err)
		}

		u, err := store.GetUser(seq)

		if err != nil {
			t.Fatalf("failed to get user: %v", err)
		}

		if u.SeqID <= 0 {
			t.Errorf("id should be >0, but got %v", u.SeqID)
		}
		if u.ID != "test-id" {
			t.Errorf("id should be test-id, but got %v", u.ID)
		}
		if u.UserType != types.UserNormal {
			t.Errorf("type should be UserNormal, but got %v", u.UserType)
		}
		if u.SystemRole != roles.SystemAdmin.Name {
			t.Errorf("role should be UserRoleAdmin(admin), but got %v", u.SystemRole)
		}
		if u.Name != "test-name" {
			t.Errorf("name should be %v, but got %v", "test-name", u.Name)
		}
		if u.CreatedAt == (time.Time{}) {
			t.Error("created_at is not set")
		}
		if u.UpdatedAt == (time.Time{}) {
			t.Error("updated_at is not set")
		}
	})

	t.Run("success_organization", func(t *testing.T) {
		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewUserStore(db)

		seq, err := store.AddUser("test-id", "test-name", types.UserOrganization, roles.SystemAdmin.Name)

		if err != nil {
			t.Fatalf("failed to add user: %v", err)
		}

		u, err := store.GetUser(seq)

		if err != nil {
			t.Fatalf("failed to get user: %v", err)
		}

		if u.SeqID <= 0 {
			t.Errorf("id should be >0, but got %v", u.SeqID)
		}
		if u.ID != "test-id" {
			t.Errorf("id should be test-id, but got %v", u.ID)
		}
		if u.UserType != types.UserOrganization {
			t.Errorf("type should be UserOrganization, but got %v", u.UserType)
		}
		if u.SystemRole != roles.SystemAdmin.Name {
			t.Errorf("role should be UserRoleAdmin(admin), but got %v", u.SystemRole)
		}
		if u.Name != "test-name" {
			t.Errorf("name should be %v, but got %v", "test-name", u.Name)
		}
		if u.CreatedAt == (time.Time{}) {
			t.Error("created_at is not set")
		}
		if u.UpdatedAt == (time.Time{}) {
			t.Error("updated_at is not set")
		}
	})
}
