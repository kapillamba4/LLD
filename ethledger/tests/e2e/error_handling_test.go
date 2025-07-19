package e2e

import (
	"os/exec"
	"path/filepath"
	"testing"
)

const LONG_ADDRESS = "0x000000000000000000000000000000000000dead123"

func TestEthledgerErrorHandling(t *testing.T) {
	tempDir := t.TempDir()
	outputFile := filepath.Join(tempDir, "test.csv")

	copyCmd := exec.Command("cp", "../../ethledger", filepath.Join(tempDir, "ethledger_test"))
	if err := copyCmd.Run(); err != nil {
		t.Fatalf("Failed to copy binary: %v", err)
	}

	copyEnvCmd := exec.Command("cp", "../../.env", tempDir)
	if err := copyEnvCmd.Run(); err != nil {
		t.Logf("Warning: Failed to copy .env file: %v", err)
	}

	testCases := []struct {
		name        string
		address     string
		expectError bool
	}{
		{
			name:        "Invalid address format",
			address:     "invalid_address",
			expectError: true,
		},
		{
			name:        "Too short address",
			address:     "0x123",
			expectError: true,
		},
		{
			name:        "Too long address",
			address:     LONG_ADDRESS,
			expectError: true,
		},
		{
			name:        "Valid address",
			address:     "0xa39b189482f984388a34460636fea9eb181ad1a6",
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			exportCmd := exec.Command("./ethledger_test", "export", "--wallet", tc.address, "--outfile", outputFile)
			exportCmd.Dir = tempDir

			err := exportCmd.Run()
			if tc.expectError && err == nil {
				t.Errorf("Expected error for address %s, but command succeeded", tc.address)
			}

			if !tc.expectError && err != nil {
				t.Errorf("Expected success for address %s, but got error: %v", tc.address, err)
			}
		})
	}
}
