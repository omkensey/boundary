package target

import (
	"context"
	"testing"

	"github.com/hashicorp/boundary/internal/credential/vault"
	"github.com/hashicorp/boundary/internal/db"
	"github.com/hashicorp/boundary/internal/host/static"
	"github.com/hashicorp/boundary/internal/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_TestTcpTarget(t *testing.T) {
	require := require.New(t)
	conn, _ := db.TestSetup(t, "postgres")
	wrapper := db.TestWrapper(t)
	_, proj := iam.TestScopes(t, iam.TestRepo(t, conn, wrapper))
	cats := static.TestCatalogs(t, conn, proj.PublicId, 1)
	hsets := static.TestSets(t, conn, cats[0].GetPublicId(), 2)
	var sets []string
	for _, s := range hsets {
		sets = append(sets, s.PublicId)
	}
	name := testTargetName(t, proj.PublicId)
	target := TestTcpTarget(t, conn, proj.PublicId, name, WithHostSets(sets))
	require.NotNil(t)
	require.NotEmpty(target.PublicId)
	require.Equal(name, target.Name)

	rw := db.New(conn)
	foundSets, err := fetchSets(context.Background(), rw, target.PublicId)
	require.NoError(err)
	foundIds := make([]string, 0, len(foundSets))
	for _, s := range foundSets {
		foundIds = append(foundIds, s.PublicId)
	}
	require.Equal(sets, foundIds)
}

func Test_TestCredentialLibrary(t *testing.T) {
	assert, require := assert.New(t), require.New(t)
	conn, _ := db.TestSetup(t, "postgres")
	wrapper := db.TestWrapper(t)
	_, proj := iam.TestScopes(t, iam.TestRepo(t, conn, wrapper))

	target := TestTcpTarget(t, conn, proj.PublicId, t.Name())
	store := vault.TestCredentialStores(t, conn, wrapper, proj.GetPublicId(), 1)[0]
	vlibs := vault.TestCredentialLibraries(t, conn, wrapper, store.GetPublicId(), 2)
	var libIds []string
	var libs []*CredentialLibrary
	for _, v := range vlibs {
		libIds = append(libIds, v.GetPublicId())
		lib := TestCredentialLibrary(t, conn, target.GetPublicId(), v.GetPublicId())
		require.NotNil(lib)
		libs = append(libs, lib)
	}

	assert.Len(libs, 2)

	rw := db.New(conn)
	foundLibs, err := fetchLibraries(context.Background(), rw, target.PublicId)
	require.NoError(err)
	foundIds := make([]string, 0, len(foundLibs))
	for _, s := range foundLibs {
		foundIds = append(foundIds, s.CredentialLibraryId)
	}
	require.Equal(libIds, foundIds)
}
