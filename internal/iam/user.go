package iam

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/watchtower/internal/db"
	"github.com/hashicorp/watchtower/internal/iam/store"
	"google.golang.org/protobuf/proto"
)

// User defines watchtower users which are scoped to an Organization
type User struct {
	*store.User
	tableName string `gorm:"-"`
}

// ensure that User implements the interfaces of: Resource, ClonableResource and db.VetForWriter
var _ Resource = (*User)(nil)
var _ Clonable = (*User)(nil)
var _ db.VetForWriter = (*User)(nil)

// NewUser creates a new in memory user and allows options:
// WithName - to specify the user's friendly name
func NewUser(organizationPublicId string, opt ...Option) (*User, error) {
	opts := getOpts(opt...)
	withName := opts.withName
	publicId, err := db.NewPublicId("u")
	if err != nil {
		return nil, fmt.Errorf("error generating public ID %w for new user", err)
	}
	u := &User{
		User: &store.User{
			PublicId: publicId,
			Name:     withName,
			ScopeId:  organizationPublicId,
		},
	}
	return u, nil
}

func allocUser() User {
	return User{
		User: &store.User{},
	}
}

// Clone creates a clone of the User
func (u *User) Clone() interface{} {
	cp := proto.Clone(u.User)
	return &User{
		User: cp.(*store.User),
	}
}

// Roles gets the roles for the user (we should/can support options to include roles associated with the user's groups)
func (u *User) Roles(ctx context.Context, r db.Reader, opt ...Option) (map[string]*Role, error) {
	if r == nil {
		return nil, errors.New("reader is nil for getting the user's roles")
	}
	if u.PublicId == "" {
		return nil, errors.New("error user id is unset for finding roles")
	}
	where := "public_id in (select role_id from iam_assigned_role ipr where principal_id  = ? and type = ?)"
	roles := []*Role{}
	if err := r.SearchWhere(ctx, &roles, where, u.PublicId, UserRoleType.String()); err != nil {
		return nil, err
	}
	results := map[string]*Role{}
	for _, r := range roles {
		results[r.PublicId] = r
	}
	return results, nil
}

// Grants finds the grants for the user and supports options:
// WithGroupGrants which will get the grants assigned to the user's groups as well
func (u *User) Grants(ctx context.Context, r db.Reader, opt ...Option) ([]*RoleGrant, error) {
	const (
		whereBase = `
role_id in (select role_id from iam_assigned_role ipr where principal_id  = ? and type = ?)
`

		whereWithGrpGrants = `
select 
	rg.*
from
	iam_role_grant rg,
	iam_assigned_role ipr, 
	iam_group grp, 
	iam_group_member gm 
where 
	rg.role_id = ipr.role_id and 
	ipr.principal_id = grp.public_id and 
	grp.public_id = gm.group_id and 
	gm.member_id = $1 and gm.type = 'user' and
	ipr."type" = 'group'
union
select 
	rg.*
from 
	iam_role_grant rg,
	iam_assigned_role ipr 
where 
	ipr.role_id  = rg.role_id and 
	ipr.principal_id  = $2 and ipr.type = 'user'
`
	)
	if r == nil {
		return nil, errors.New("error reader is nil for getting the user's grants")
	}
	opts := getOpts(opt...)
	withGrpGrants := opts.withGroupGrants
	if u.PublicId == "" {
		return nil, errors.New("error user id is unset for finding roles")
	}
	if withGrpGrants {
		grants := []*RoleGrant{}
		tx, err := r.DB()
		if err != nil {
			return nil, err
		}
		rows, err := tx.Query(whereWithGrpGrants, u.PublicId, u.PublicId)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			g := allocRoleGrant()
			if err := r.ScanRows(rows, &g); err != nil {
				return nil, err
			}
			grants = append(grants, &g)
		}
		return grants, nil
	}

	grants := []*RoleGrant{}
	if err := r.SearchWhere(ctx, &grants, whereBase, u.PublicId, UserRoleType.String()); err != nil {
		return nil, err
	}
	return grants, nil

}

// Groups will get the user's groups
func (u *User) Groups(ctx context.Context, r db.Reader) ([]*Group, error) {
	if r == nil {
		return nil, errors.New("error reader is nil for getting the user's groups")
	}
	if u.PublicId == "" {
		return nil, errors.New("error user id is unset for finding user groups")
	}
	where := "public_id in (select distinct group_id from iam_group_member where member_id = ? and type = ?)"
	groups := []*Group{}
	if err := r.SearchWhere(ctx, &groups, where, u.PublicId, UserMemberType.String()); err != nil {
		return nil, err
	}
	return groups, nil
}

// VetForWrite implements db.VetForWrite() interface
func (u *User) VetForWrite(ctx context.Context, r db.Reader, opType db.OpType, opt ...db.Option) error {
	if u.PublicId == "" {
		return errors.New("error public id is empty string for user write")
	}
	if u.ScopeId == "" {
		return errors.New("error scope id not set for user write")
	}
	// make sure the scope is valid for users
	if err := u.scopeIsValid(ctx, r); err != nil {
		return err
	}
	return nil
}

func (u *User) scopeIsValid(ctx context.Context, r db.Reader) error {
	ps, err := LookupScope(ctx, r, u)
	if err != nil {
		return err
	}
	if ps.Type != OrganizationScope.String() {
		return errors.New("error scope is not an organization")
	}
	return nil
}

// GetScope returns the scope for the User
func (u *User) GetScope(ctx context.Context, r db.Reader) (*Scope, error) {
	return LookupScope(ctx, r, u)
}

// ResourceType returns the type of the User
func (*User) ResourceType() ResourceType { return ResourceTypeUser }

// Actions returns the  available actions for Users
func (*User) Actions() map[string]Action {
	return CrudActions()
}

// TableName returns the tablename to override the default gorm table name
func (u *User) TableName() string {
	if u.tableName != "" {
		return u.tableName
	}
	return "iam_user"
}

// SetTableName sets the tablename and satisfies the ReplayableMessage interface
func (u *User) SetTableName(n string) {
	if n != "" {
		u.tableName = n
	}
}