// +build use_external_db

package store

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/modoki-paas/modoki-k8s/apiserver/testutil"
)

func TestAddUser(t *testing.T) {
	t.Run("success_normal", func(t *testing.T) {
		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewDB(db)

		u, err := store.User().AddUser("test-id", "test-name", UserNormal, UserRoleAdmin)

		if err != nil {
			t.Fatalf("failed to add user: %v", err)
		}

		if u.SeqID <= 0 {
			t.Errorf("id should be >0, but got %v", u.SeqID)
		}
		if u.ID != "test-id" {
			t.Errorf("id should be test-id, but got %v", u.ID)
		}
		if u.UserType != UserNormal {
			t.Errorf("type should be UserNormal, but got %v", u.UserType)
		}
		if u.SystemRole == UserRoleAdmin {
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

		store := NewDB(db)

		u, err := store.User().AddUser("test-id", "test-name", UserNormal, UserRoleAdmin)

		if err != nil {
			t.Fatalf("failed to add user: %v", err)
		}

		if u.SeqID <= 0 {
			t.Errorf("id should be >0, but got %v", u.SeqID)
		}
		if u.ID != "test-id" {
			t.Errorf("id should be test-id, but got %v", u.ID)
		}
		if u.UserType != UserOrganization {
			t.Errorf("type should be UserOrganization, but got %v", u.UserType)
		}
		if u.SystemRole == UserRoleAdmin {
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

func TestGetUserFromToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db := testutil.NewSQLConn(t)
		defer db.Close()

		store := NewDB(db)

		u, err := store.User().AddUser("test-id", "test-name", UserNormal, UserRoleAdmin)

		if err != nil {
			t.Fatalf("failed to add user: %v", err)
		}
		token, err := store.Token().AddToken(&Token{
			Token:           "my-token",
			Owner:           u.SeqID,
			Author:          u.SeqID,
		})
		if err != nil {
			t.Fatalf("failed to add token: %v", err)
		}

		ru, rt, err := store.User().GetUserFromToken(token.Token)

		if err != nil {
			t.Fatalf("failed to get user from token: %v", err)
		}

		if !cmp.Equal(ru, u) {
			t.Errorf("user does not match: want %v, got %v", *u, *ru)
		}
		if !cmp.Equal(rt, token) {
			t.Errorf("token does not match: want %v, got %v", *token, *rt)
		}
	})
}
