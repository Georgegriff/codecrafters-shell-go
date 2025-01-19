package command

import (
	"testing"

	"github.com/codecrafters-io/shell-starter-go/cmd/testutils"
)

func TestParseArgs(t *testing.T) {
	command, args := parseArgs(`cat "/tmp/qux/f\n42" "/tmp/qux/f\78" "/tmp/qux/f'\'36"`)
	testutils.ExpectToMatchString(t, command, "cat")
	testutils.ExpectToMatchInt(t, len(args), 3)
	testutils.ExpectToMatchString(t, args[0], "/tmp/qux/f\\n42")
	testutils.ExpectToMatchString(t, args[1], `/tmp/qux/f\78`)
	testutils.ExpectToMatchString(t, args[2], `/tmp/qux/f'\'36`)

}

func TestParseArgs2(t *testing.T) {
	command, args := parseArgs(`echo \'\"script hello\"\'`)
	testutils.ExpectToMatchString(t, command, "echo")
	testutils.ExpectToMatchInt(t, len(args), 2)
	testutils.ExpectToMatchString(t, args[0], `'"script`)
	testutils.ExpectToMatchString(t, args[1], `hello"'`)

}
