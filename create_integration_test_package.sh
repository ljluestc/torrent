#!/bin/bash
set -e

echo "===================================================================="
echo "             Creating isolated test package"
echo "===================================================================="

# Create a directory for isolated tests
mkdir -p isolated_tests

# Create a simple test file that doesn't depend on other packages
cat > isolated_tests/simple_test.go << 'EOF'
package isolated_tests

import (
	"testing"
)

// TestSimple is just a simple test to verify testing works
func TestSimple(t *testing.T) {
	// This test should always pass
	if 1+1 != 2 {
		t.Fatal("Basic math failed")
	}
}
EOF

echo "Created isolated test package."
echo "You can run this test with: go test ./isolated_tests"
