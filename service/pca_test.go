package service

import (
    "testing"
    "reflect"
    "github.com/stretchr/testify/assert"
)

// TestComputePCA tests the PCA computation for accuracy.
func TestComputePCA(t *testing.T) {
    // Example test - adjust based on your actual PCA implementation and expected results.
    expected := []float64{1.0} // Simplified expected result
    result, err := ComputePCA("sample text")

    assert.NoError(t, err)
    assert.True(t, reflect.DeepEqual(expected, result))
}

