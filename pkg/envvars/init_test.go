package envvars_test

import (
	"errors"
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit_toCreateDeclarationFile(t *testing.T) {
	// given
	expectedDeclaration := &envvars.Declaration{
		Envvars: envvars.EnvvarCollection{
			&envvars.Envvar{
				Name:    "ECHO",
				Example: "Hello World",
			},
		},
	}
	mockWriter := new(mocks.DeclarationWriter)
	mockWriter.On("Write", expectedDeclaration, false).Return(nil)

	// when
	err := envvars.Init(mockWriter)

	// then
	assert.NoError(t, err)
	mockWriter.AssertExpectations(t)
}

func TestInit_toReturnErrorIfDeclarationExists(t *testing.T) {
	// given
	expectedDeclaration := &envvars.Declaration{
		Envvars: envvars.EnvvarCollection{
			&envvars.Envvar{
				Name:    "ECHO",
				Example: "Hello World",
			},
		},
	}
	expectedError := errors.New("error")
	mockWriter := new(mocks.DeclarationWriter)
	mockWriter.On("Write", expectedDeclaration, false).Return(expectedError)

	// when
	err := envvars.Init(mockWriter)

	// then
	assert.EqualError(t, err, err.Error())
	mockWriter.AssertExpectations(t)
}
