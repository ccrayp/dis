package tests

import (
	"app/model"
	"testing"

	"github.com/stretchr/testify/require"
)

var tempId uint

func TestPersonService_Create(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		id, err := testEnv.PersonService.Create("test", "89216939154", 20)
		tempId = id

		require.NotEqual(t, 0, id)
		require.NoError(t, err)
	})

	t.Run("invalid phone", func(t *testing.T) {
		_, err := testEnv.PersonService.Create("test", "892169391564", 20)

		require.Error(t, err)
	})

	t.Run("invalid age", func(t *testing.T) {
		_, err := testEnv.PersonService.Create("test", "89216939154", -2)

		require.Error(t, err)
	})
}

func TestPersonService_Get(t *testing.T) {
	t.Run("amount quantity", func(t *testing.T) {
		q, err := testEnv.PersonService.GetAmountQuantity()

		require.NotEqual(t, 0, q)
		require.NoError(t, err)
	})

	t.Run("by limit and offset", func(t *testing.T) {
		persons, err := testEnv.PersonService.GetByLimitOffset(10, 0)

		require.NoError(t, err)
		require.NotEqual(t, nil, persons)
	})

	t.Run("by id", func(t *testing.T) {
		person, err := testEnv.PersonService.GetById(tempId)

		require.NoError(t, err)
		require.Equal(t, tempId, person.GetId())
	})
}

func TestPersonService_Update(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		newValidData := &model.Person{
			Id:    tempId,
			Name:  "test1",
			Phone: "89216931954",
			Age:   29,
		}

		err := testEnv.PersonService.Update(newValidData)

		require.NoError(t, err)
	})

	t.Run("invalid phone", func(t *testing.T) {
		newInvalidData := &model.Person{
			Id:    1000000,
			Name:  "test1",
			Phone: "+799216931954",
			Age:   29,
		}

		err := testEnv.PersonService.Update(newInvalidData)

		require.Error(t, err)
	})

	t.Run("invalid age", func(t *testing.T) {
		newInvalidData := &model.Person{
			Id:    1000000,
			Name:  "test1",
			Phone: "+79216931954",
			Age:   0,
		}

		err := testEnv.PersonService.Update(newInvalidData)

		require.Error(t, err)
	})
}

func TestPersonService_Delete(t *testing.T) {
	t.Run("invalid id", func(t *testing.T) {
		err := testEnv.PersonService.Delete(0)

		require.Error(t, err)
	})

	t.Run("valid id", func(t *testing.T) {
		err := testEnv.PersonService.Delete(tempId)

		require.NoError(t, err)
	})
}

func TestPersonService_Errors(t *testing.T) {
	t.Run("get amount quantity", func(t *testing.T) {
		_, err := testEnv.InvalidPersonService.GetAmountQuantity()

		require.Error(t, err)
	})

	t.Run("get by limit and offset", func(t *testing.T) {
		_, err := testEnv.InvalidPersonService.GetByLimitOffset(10, 0)

		require.Error(t, err)
	})

	t.Run("get by id", func(t *testing.T) {
		_, err := testEnv.InvalidPersonService.GetById(1)

		require.Error(t, err)
	})

	t.Run("create", func(t *testing.T) {
		_, err := testEnv.InvalidPersonService.Create("invalid", "89216931954", 20)

		require.Error(t, err)
	})

	t.Run("update", func(t *testing.T) {
		newData := &model.Person{
			Id:    10,
			Name:  "test1",
			Phone: "+79216931954",
			Age:   50,
		}

		err := testEnv.InvalidPersonService.Update(newData)

		require.Error(t, err)
	})

	t.Run("delete", func(t *testing.T) {
		err := testEnv.InvalidPersonService.Delete(10)

		require.Error(t, err)
	})
}
