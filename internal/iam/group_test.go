package iam

import (
	"context"
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/watchtower/internal/db"
	"gotest.tools/assert"
)

func Test_NewGroup(t *testing.T) {
	db.StartTest()
	t.Parallel()
	cleanup, url := db.SetupTest(t, "../db/migrations/postgres")
	defer cleanup()
	defer db.CompleteTest() // must come after the "defer cleanup()"
	conn, err := db.TestConnection(url)
	assert.NilError(t, err)
	defer conn.Close()

	t.Run("valid", func(t *testing.T) {
		w := db.GormReadWriter{Tx: conn}
		s, err := NewScope(OrganizationScope)
		assert.NilError(t, err)
		assert.Check(t, s.Scope != nil)
		err = w.Create(context.Background(), s)
		assert.NilError(t, err)
		assert.Check(t, s.Id != 0)

		rootUser, err := NewUser(s, AsRootUser(true))
		assert.NilError(t, err)
		err = w.Create(context.Background(), rootUser)
		assert.NilError(t, err)

		grp, err := NewGroup(s, rootUser, WithDescription("this is a test group"))
		assert.NilError(t, err)
		assert.Check(t, grp != nil)
		assert.Equal(t, rootUser.Id, grp.OwnerId)
		assert.Equal(t, grp.Description, "this is a test group")
		assert.Equal(t, s.Id, grp.PrimaryScopeId)
		err = w.Create(context.Background(), grp)
		assert.NilError(t, err)
		assert.Check(t, grp.Id != 0)
	})
}

func Test_GroupMembers(t *testing.T) {
	db.StartTest()
	t.Parallel()
	cleanup, url := db.SetupTest(t, "../db/migrations/postgres")
	defer cleanup()
	defer db.CompleteTest() // must come after the "defer cleanup()"
	conn, err := db.TestConnection(url)
	assert.NilError(t, err)
	defer conn.Close()

	t.Run("valid", func(t *testing.T) {
		w := db.GormReadWriter{Tx: conn}
		s, err := NewScope(OrganizationScope)
		assert.NilError(t, err)
		assert.Check(t, s.Scope != nil)
		err = w.Create(context.Background(), s)
		assert.NilError(t, err)
		assert.Check(t, s.Id != 0)

		rootUser, err := NewUser(s, AsRootUser(true))
		assert.NilError(t, err)
		err = w.Create(context.Background(), rootUser)
		assert.NilError(t, err)

		grp, err := NewGroup(s, rootUser, WithDescription("this is a test group"))
		assert.NilError(t, err)
		assert.Check(t, grp != nil)
		assert.Equal(t, rootUser.Id, grp.OwnerId)
		assert.Equal(t, grp.Description, "this is a test group")
		assert.Equal(t, s.Id, grp.PrimaryScopeId)
		err = w.Create(context.Background(), grp)
		assert.NilError(t, err)
		assert.Check(t, grp.Id != 0)

		gm, err := NewGroupMember(s, grp, rootUser)
		assert.NilError(t, err)
		assert.Check(t, gm != nil)
		err = w.Create(context.Background(), gm)
		assert.NilError(t, err)

		meth, err := NewAuthMethod(s, rootUser, AuthUserPass)
		assert.NilError(t, err)
		assert.Check(t, meth != nil)
		err = w.Create(context.Background(), meth)
		assert.NilError(t, err)

		id, err := uuid.GenerateUUID()
		assert.NilError(t, err)
		alias, err := NewUserAlias(s, rootUser, meth, id)
		assert.NilError(t, err)
		assert.Check(t, alias != nil)
		err = w.Create(context.Background(), alias)
		assert.NilError(t, err)

		gm2, err := NewGroupMember(s, grp, alias)
		assert.NilError(t, err)
		assert.Check(t, gm2 != nil)
		err = w.Create(context.Background(), gm2)
		assert.NilError(t, err)

		members, err := grp.Members(context.Background(), &w)
		assert.NilError(t, err)
		assert.Check(t, members != nil)
		assert.Check(t, len(members) == 2)
		for _, m := range members {
			if m.GetMemberId() != alias.Id && m.GetMemberId() != rootUser.Id {
				t.Error("members not one of the known ids")
			}
		}
	})
}
