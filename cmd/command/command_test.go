package command

import (
	"testing"

	"github.com/codecrafters-io/shell-starter-go/cmd/testutils"
)

func TestParseArgs(t *testing.T) {
	args := parseArgs(`cat 'file1.go' file2.go 'file3.go' file`)
	testutils.ExpectToMatchInt(t, len(args), 5)
	testutils.ExpectToMatchString(t, args[0], "cat")
	testutils.ExpectToMatchString(t, args[1], "file1.go")
	testutils.ExpectToMatchString(t, args[2], "file2.go")
	testutils.ExpectToMatchString(t, args[3], "file3.go")
	testutils.ExpectToMatchString(t, args[4], "file")

}
