package tests

import (
	"app/model"
	"testing"

	"github.com/stretchr/testify/require"
)

var tempUserId uint

func TestSysUserService_Create(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		id, err := testEnv.UserService.Create(1, "test")
		tempUserId = id

		require.NotEqual(t, 0, id)
		require.NoError(t, err)
	})

	t.Run("invalid person id", func(t *testing.T) {
		_, err := testEnv.UserService.Create(0, "test")

		require.Error(t, err)
	})

	t.Run("invalid password", func(t *testing.T) {
		_, err := testEnv.UserService.Create(1, "123")

		require.Error(t, err)
	})
}

func TestSysUserService_Get(t *testing.T) {
	t.Run("amount quantity", func(t *testing.T) {
		q, err := testEnv.UserService.GetAmountQuantity()

		require.NotEqual(t, 0, q)
		require.NoError(t, err)
	})

	t.Run("by limit and offset", func(t *testing.T) {
		users, err := testEnv.UserService.GetByLimitOffset(10, 0)

		require.NoError(t, err)
		require.NotEqual(t, nil, users)
	})

	t.Run("by id", func(t *testing.T) {
		user, err := testEnv.UserService.GetById(tempUserId)

		require.NoError(t, err)
		require.Equal(t, tempUserId, user.GetId())
	})
}

func TestSysUserService_Update(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		user, err := testEnv.UserService.GetById(tempUserId)
		require.NoError(t, err)

		err = testEnv.UserService.Update(user, "test1")

		require.NoError(t, err)
	})

	t.Run("invalid person id", func(t *testing.T) {
		user, err := testEnv.UserService.GetById(tempUserId)
		require.NoError(t, err)

		user.PersonId = 0

		err = testEnv.UserService.Update(user, "test1")

		require.Error(t, err)
	})

	t.Run("invalid password", func(t *testing.T) {
		user, err := testEnv.UserService.GetById(tempUserId)
		require.NoError(t, err)

		err = testEnv.UserService.Update(user, "123")

		require.Error(t, err)
	})
}

func TestSysUserService_CheckPassword(t *testing.T) {
	t.Run("valid password", func(t *testing.T) {
		ok, err := testEnv.UserService.CheckPassword(tempUserId, "test1")

		require.NoError(t, err)
		require.Equal(t, true, ok)
	})

	t.Run("invalid password", func(t *testing.T) {
		ok, err := testEnv.UserService.CheckPassword(tempUserId, "1234")

		require.NoError(t, err)
		require.Equal(t, false, ok)
	})
}

func TestSysUserService_Delete(t *testing.T) {
	t.Run("invalid id", func(t *testing.T) {
		err := testEnv.UserService.Delete(0)

		require.Error(t, err)
	})

	t.Run("valid id", func(t *testing.T) {
		err := testEnv.UserService.Delete(tempUserId)

		require.NoError(t, err)
	})
}

func TestUserService_Errors(t *testing.T) {
	t.Run("get amount quantity", func(t *testing.T) {
		_, err := testEnv.InvalidUserService.GetAmountQuantity()

		require.Error(t, err)
	})

	t.Run("get by limit and offset", func(t *testing.T) {
		_, err := testEnv.InvalidUserService.GetByLimitOffset(10, 0)

		require.Error(t, err)
	})

	t.Run("get by id", func(t *testing.T) {
		_, err := testEnv.InvalidUserService.GetById(1)

		require.Error(t, err)
	})

	t.Run("create", func(t *testing.T) {
		_, err := testEnv.InvalidUserService.Create(1, "1234")

		require.Error(t, err)
	})

	t.Run("update", func(t *testing.T) {
		newData := &model.SysUser{
			Id:       10,
			PersonId: 1,
		}

		err := testEnv.InvalidUserService.Update(newData, "1234")

		require.Error(t, err)
	})

	t.Run("delete", func(t *testing.T) {
		err := testEnv.InvalidUserService.Delete(10)

		require.Error(t, err)
	})

	t.Run("check password", func(t *testing.T) {
		_, err := testEnv.InvalidUserService.CheckPassword(1, "1234")

		require.Error(t, err)
	})
}
