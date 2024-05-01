package data_model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataModel(t *testing.T) {
	t.Run("valid data model", func(t *testing.T) {
		dataModel := &UserDataModel{
			Id:   "1",
			Name: "Test",
		}
		err := dataModel.validateDataModel()
		assert.NoError(t, err)
	})

	t.Run("invalid data model - name too short", func(t *testing.T) {
		dataModel := &UserDataModel{
			Id:   "1",
			Name: "Te",
		}
		err := dataModel.validateDataModel()
		assert.Error(t, err)
	})

	t.Run("invalid data model - name missing", func(t *testing.T) {
		dataModel := &UserDataModel{
			Id:   "1",
			Name: "",
		}
		err := dataModel.validateDataModel()
		assert.Error(t, err)
	})
}
