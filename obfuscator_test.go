package obfuscator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewObfuscator_Success(t *testing.T) {
	// Arrange & Act
	obf, err := NewObfuscator()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, obf)
}

func TestNewObfuscator_CachedData(t *testing.T) {
	// Arrange & Act
	obf, err := NewObfuscator()

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, obf.CachedData)
}

func TestNewObfuscator_MultipleObfuscations(t *testing.T) {
	// Arrange
	obf, err := NewObfuscator()
	require.NoError(t, err)

	jsCode := "function test() { return 'Hello World'; }"

	for i := range 10 {
		result, err := obf.Obfuscate(jsCode)
		assert.NoError(t, err, "Obfuscation #%d ", i+1)
		assert.NotEmpty(t, result, "obfuscation #%d", i+1)
	}
}
