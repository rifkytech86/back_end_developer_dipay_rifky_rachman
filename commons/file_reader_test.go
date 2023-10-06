package commons

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockFileReader struct{}

func (m *MockFileReader) ReadFile(filename string) ([]byte, error) {
	// Return mock data based on the filename
	if filename == "test.txt" {
		return []byte("mocked content"), nil
	}
	return nil, errors.New("File not found")
}
func TestFileReader_ReadFile(t *testing.T) {
	// Create an instance of the mock
	mockFileReader := &MockFileReader{}

	// Define the filename and the expected result
	filename := "test.txt"
	expectedResult := []byte("mocked content")

	// Call the function being tested
	result, err := mockFileReader.ReadFile(filename)

	// Check if the result and error match the expectations
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !bytes.Equal(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}
}

func TestNewFileReader(t *testing.T) {
	// Call the function being tested
	fileReader := NewFileReader()

	// Check if the returned type is correct
	assert.Implements(t, (*IFileReader)(nil), fileReader)
}
