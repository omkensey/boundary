package iam

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/hashicorp/watchtower/internal/db"
	dbassert "github.com/hashicorp/watchtower/internal/db/assert"
	"github.com/hashicorp/watchtower/internal/oplog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestNewRole(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := db.TestSetup(t, "postgres")
	defer func() {
		err := cleanup()
		assert.NoError(t, err)
		err = conn.Close()
		assert.NoError(t, err)
	}()
	org, proj := TestScopes(t, conn)
	id := testId(t)

	type args struct {
		scopePublicId string
		opt           []Option
	}
	tests := []struct {
		name            string
		args            args
		wantErr         bool
		wantErrMsg      string
		wantIsErr       error
		wantName        string
		wantDescription string
	}{
		{
			name: "valid",
			args: args{
				scopePublicId: org.PublicId,
				opt:           []Option{WithName(id), WithDescription("description-" + id)},
			},
			wantErr:         false,
			wantName:        id,
			wantDescription: "description-" + id,
		},
		{
			name: "valid-proj",
			args: args{
				scopePublicId: proj.PublicId,
				opt:           []Option{WithName(id), WithDescription("description-" + id)},
			},
			wantErr:         false,
			wantName:        id,
			wantDescription: "description-" + id,
		},
		{
			name: "valid-with-no-options" + id,
			args: args{
				scopePublicId: org.PublicId,
			},
			wantErr: false,
		},
		{
			name: "no-scope",
			args: args{
				opt: []Option{WithName(id)},
			},
			wantErr:    true,
			wantErrMsg: "new role: missing scope id invalid parameter",
			wantIsErr:  db.ErrInvalidParameter,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert, require := assert.New(t), require.New(t)
			got, err := NewRole(tt.args.scopePublicId, tt.args.opt...)
			if tt.wantErr {
				require.Error(err)
				assert.Equal(tt.wantErrMsg, err.Error())
				if tt.wantIsErr != nil {
					assert.True(errors.Is(err, tt.wantIsErr))
				}
				return
			}
			require.NoError(err)
			assert.Equal(tt.wantName, got.Name)
			assert.Equal(tt.wantDescription, got.Description)
			assert.Empty(got.PublicId)
		})
	}
}

func Test_RoleCreate(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := db.TestSetup(t, "postgres")
	defer func() {
		err := cleanup()
		assert.NoError(t, err)
		err = conn.Close()
		assert.NoError(t, err)
	}()
	org, proj := TestScopes(t, conn)
	type args struct {
		role *Role
	}
	tests := []struct {
		name        string
		args        args
		wantDup     bool
		wantErr     bool
		wantErrMsg  string
		wantIsError error
	}{
		{
			name: "valid-with-org",
			args: args{
				role: func() *Role {
					id := testId(t)
					role, err := NewRole(org.PublicId, WithName(id), WithDescription("description-"+id))
					require.NoError(t, err)
					grpId, err := newRoleId()
					require.NoError(t, err)
					role.PublicId = grpId
					return role
				}(),
			},
			wantErr: false,
		},
		{
			name: "valid-with-proj",
			args: args{
				role: func() *Role {
					id := testId(t)
					role, err := NewRole(proj.PublicId, WithName(id), WithDescription("description-"+id))
					require.NoError(t, err)
					grpId, err := newRoleId()
					require.NoError(t, err)
					role.PublicId = grpId
					return role
				}(),
			},
			wantErr: false,
		},
		{
			name: "valid-with-dup-null-names-and-descriptions",
			args: args{
				role: func() *Role {
					role, err := NewRole(org.PublicId)
					require.NoError(t, err)
					roleId, err := newRoleId()
					require.NoError(t, err)
					role.PublicId = roleId
					return role
				}(),
			},
			wantDup: true,
			wantErr: false,
		},
		{
			name: "bad-scope-id",
			args: args{
				role: func() *Role {
					id := testId(t)
					role, err := NewRole(id)
					require.NoError(t, err)
					roleId, err := newRoleId()
					require.NoError(t, err)
					role.PublicId = roleId
					return role
				}(),
			},
			wantErr:    true,
			wantErrMsg: "create: vet for write failed scope is not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert, require := assert.New(t), require.New(t)
			w := db.New(conn)
			if tt.wantDup {
				r := tt.args.role.Clone().(*Role)
				roleId, err := newRoleId()
				require.NoError(err)
				r.PublicId = roleId
				err = w.Create(context.Background(), r)
				require.NoError(err)
			}
			r := tt.args.role.Clone().(*Role)
			err := w.Create(context.Background(), r)
			if tt.wantErr {
				require.Error(err)
				assert.Equal(tt.wantErrMsg, err.Error())
				return
			}
			assert.NoError(err)
			assert.NotEmpty(tt.args.role.PublicId)

			foundGrp := allocRole()
			foundGrp.PublicId = tt.args.role.PublicId
			err = w.LookupByPublicId(context.Background(), &foundGrp)
			require.NoError(err)
			assert.Equal(r, &foundGrp)
		})
	}
}

func Test_RoleUpdate(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := db.TestSetup(t, "postgres")
	defer func() {
		err := cleanup()
		assert.NoError(t, err)
		err = conn.Close()
		assert.NoError(t, err)
	}()
	id := testId(t)
	org, proj := TestScopes(t, conn)
	rw := db.New(conn)
	type args struct {
		name           string
		description    string
		fieldMaskPaths []string
		nullPaths      []string
		ScopeId        string
	}
	tests := []struct {
		name           string
		args           args
		wantRowsUpdate int
		wantErr        bool
		wantErrMsg     string
		wantDup        bool
	}{
		{
			name: "valid",
			args: args{
				name:           "valid" + id,
				fieldMaskPaths: []string{"Name"},
				ScopeId:        org.PublicId,
			},
			wantErr:        false,
			wantRowsUpdate: 1,
		},
		{
			name: "proj-scope-id",
			args: args{
				name:           "proj-scope-id" + id,
				fieldMaskPaths: []string{"ScopeId"},
				ScopeId:        proj.PublicId,
			},
			wantErr:    true,
			wantErrMsg: "update: vet for write failed not allowed to change a resource's scope",
		},
		{
			name: "proj-scope-id-not-in-mask",
			args: args{
				name:           "proj-scope-id" + id,
				fieldMaskPaths: []string{"Name"},
				ScopeId:        proj.PublicId,
			},
			wantErr:        false,
			wantRowsUpdate: 1,
		},
		{
			name: "empty-scope-id",
			args: args{
				name:           "empty-scope-id" + id,
				fieldMaskPaths: []string{"Name"},
				ScopeId:        "",
			},
			wantErr:        false,
			wantRowsUpdate: 1,
		},
		{
			name: "dup-name",
			args: args{
				name:           "dup-name" + id,
				fieldMaskPaths: []string{"Name"},
				ScopeId:        org.PublicId,
			},
			wantErr:    true,
			wantDup:    true,
			wantErrMsg: `update: failed pq: duplicate key value violates unique constraint "iam_role_name_scope_id_key"`,
		},
		{
			name: "set description null",
			args: args{
				name:           "set description null" + id,
				fieldMaskPaths: []string{"Name"},
				nullPaths:      []string{"Description"},
				ScopeId:        org.PublicId,
			},
			wantErr:        false,
			wantRowsUpdate: 1,
		},
		{
			name: "set name null",
			args: args{
				description:    "set description null" + id,
				fieldMaskPaths: []string{"Description"},
				nullPaths:      []string{"Name"},
				ScopeId:        org.PublicId,
			},
			wantDup:        true,
			wantErr:        false,
			wantRowsUpdate: 1,
		},
		{
			name: "set description null",
			args: args{
				name:           "set name null" + id,
				fieldMaskPaths: []string{"Name"},
				nullPaths:      []string{"Description"},
				ScopeId:        org.PublicId,
			},
			wantErr:        false,
			wantRowsUpdate: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert, require := assert.New(t), require.New(t)
			if tt.wantDup {
				r := TestRole(t, conn, org.PublicId, WithName(tt.args.name))
				_, err := rw.Update(context.Background(), r, tt.args.fieldMaskPaths, tt.args.nullPaths)
				require.NoError(err)
			}

			id := testId(t)
			role := TestRole(t, conn, org.PublicId, WithDescription(id), WithName(id))

			updateRole := allocRole()
			updateRole.PublicId = role.PublicId
			updateRole.ScopeId = tt.args.ScopeId
			updateRole.Name = tt.args.name
			updateRole.Description = tt.args.description

			updatedRows, err := rw.Update(context.Background(), &updateRole, tt.args.fieldMaskPaths, tt.args.nullPaths)
			if tt.wantErr {
				require.Error(err)
				assert.Equal(0, updatedRows)
				assert.Equal(tt.wantErrMsg, err.Error())
				err = db.TestVerifyOplog(t, rw, role.PublicId, db.WithOperation(oplog.OpType_OP_TYPE_UPDATE), db.WithCreateNotBefore(10*time.Second))
				require.Error(err)
				assert.Equal("record not found", err.Error())
				return
			}
			require.NoError(err)
			assert.Equal(tt.wantRowsUpdate, updatedRows)
			assert.NotEqual(role.UpdateTime, updateRole.UpdateTime)
			foundRole := allocRole()
			foundRole.PublicId = role.GetPublicId()
			err = rw.LookupByPublicId(context.Background(), &foundRole)
			require.NoError(err)
			assert.True(proto.Equal(updateRole, foundRole))
			if len(tt.args.nullPaths) != 0 {
				dbassert := dbassert.New(t, rw)
				for _, f := range tt.args.nullPaths {
					dbassert.IsNull(&foundRole, f)
				}
			}
		})
	}
	t.Run("update dup names in diff scopes", func(t *testing.T) {
		assert, require := assert.New(t), require.New(t)
		id := testId(t)
		_ = TestRole(t, conn, org.PublicId, WithDescription(id), WithName(id))
		projRole := TestRole(t, conn, proj.PublicId, WithName(id))
		updatedRows, err := rw.Update(context.Background(), projRole, []string{"Name"}, nil)
		require.NoError(err)
		assert.Equal(1, updatedRows)

		foundRole := allocRole()
		foundRole.PublicId = projRole.GetPublicId()
		err = rw.LookupByPublicId(context.Background(), &foundRole)
		require.NoError(err)
		assert.Equal(id, projRole.Name)
	})
}

func Test_RoleDelete(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := db.TestSetup(t, "postgres")
	defer func() {
		err := cleanup()
		assert.NoError(t, err)
		err = conn.Close()
		assert.NoError(t, err)
	}()

	rw := db.New(conn)
	id := testId(t)
	org, _ := TestScopes(t, conn)

	tests := []struct {
		name            string
		role            *Role
		wantRowsDeleted int
		wantErr         bool
		wantErrMsg      string
	}{
		{
			name:            "valid",
			role:            TestRole(t, conn, org.PublicId),
			wantErr:         false,
			wantRowsDeleted: 1,
		},
		{
			name:            "bad-id",
			role:            func() *Role { r := allocRole(); r.PublicId = id; return &r }(),
			wantErr:         false,
			wantRowsDeleted: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert, require := assert.New(t), require.New(t)
			deleteRole := allocRole()
			deleteRole.PublicId = tt.role.GetPublicId()
			deletedRows, err := rw.Delete(context.Background(), &deleteRole)
			if tt.wantErr {
				require.Error(err)
				return
			}
			require.NoError(err)
			if tt.wantRowsDeleted == 0 {
				assert.Equal(tt.wantRowsDeleted, deletedRows)
				return
			}
			assert.Equal(tt.wantRowsDeleted, deletedRows)
			foundRole := allocRole()
			foundRole.PublicId = tt.role.GetPublicId()
			err = rw.LookupByPublicId(context.Background(), &foundRole)
			require.Error(err)
			assert.True(errors.Is(db.ErrRecordNotFound, err))
		})
	}
}

func TestRole_Actions(t *testing.T) {
	assert := assert.New(t)
	r := &Role{}
	a := r.Actions()
	assert.Equal(a[ActionCreate.String()], ActionCreate)
	assert.Equal(a[ActionUpdate.String()], ActionUpdate)
	assert.Equal(a[ActionRead.String()], ActionRead)
	assert.Equal(a[ActionDelete.String()], ActionDelete)
}

func TestRole_ResourceType(t *testing.T) {
	assert := assert.New(t)
	r := &Role{}
	ty := r.ResourceType()
	assert.Equal(ty, ResourceTypeRole)
}

func TestRole_GetScope(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := db.TestSetup(t, "postgres")
	defer func() {
		err := cleanup()
		assert.NoError(t, err)
		err = conn.Close()
		assert.NoError(t, err)
	}()
	org, proj := TestScopes(t, conn)

	t.Run("valid-org", func(t *testing.T) {
		assert, require := assert.New(t), require.New(t)
		w := db.New(conn)
		role := TestRole(t, conn, org.PublicId)
		scope, err := role.GetScope(context.Background(), w)
		require.NoError(err)
		assert.True(proto.Equal(org, scope))
	})
	t.Run("valid-proj", func(t *testing.T) {
		assert, require := assert.New(t), require.New(t)
		w := db.New(conn)
		role := TestRole(t, conn, proj.PublicId)
		scope, err := role.GetScope(context.Background(), w)
		require.NoError(err)
		assert.True(proto.Equal(proj, scope))
	})
}

func TestRole_Clone(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := db.TestSetup(t, "postgres")
	defer func() {
		err := cleanup()
		assert.NoError(t, err)
		err = conn.Close()
		assert.NoError(t, err)
	}()
	org, _ := TestScopes(t, conn)
	t.Run("valid", func(t *testing.T) {
		assert := assert.New(t)
		role := TestRole(t, conn, org.PublicId, WithDescription("this is a test role"))
		cp := role.Clone()
		assert.True(proto.Equal(cp.(*Role).Role, role.Role))
	})
	t.Run("not-equal", func(t *testing.T) {
		assert := assert.New(t)
		role := TestRole(t, conn, org.PublicId)
		role2 := TestRole(t, conn, org.PublicId)
		cp := role.Clone()
		assert.True(!proto.Equal(cp.(*Role).Role, role2.Role))
	})
}