package systemutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type commandParameterGroup struct {
	Arguments            []string
	WorkingDirectory     string
	EnvironmentVariables []string
}

func TestPrepareEnvironmentVariables(t *testing.T) {
	type args struct {
		environmentVariables map[string]string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{
				environmentVariables: map[string]string{},
			},
			want: nil,
		},
		{
			name: "nonempty",
			args: args{
				environmentVariables: map[string]string{
					"KEY_ONE": "value #1",
					"KEY_TWO": "value #2",
				},
			},
			want: []string{"KEY_ONE=value #1", "KEY_TWO=value #2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrepareEnvironmentVariables(tt.args.environmentVariables)

			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestCommand(t *testing.T) {
	if os.Getenv("TEST_COMMAND") != "TRUE" {
		return
	}

	require.GreaterOrEqual(t, len(os.Args), 2)

	var arguments []string
	for _, argument := range os.Args[1 : len(os.Args)-1] {
		if strings.HasPrefix(argument, "test-command") {
			arguments = append(arguments, argument)
		}
	}

	workingDirectory, err := os.Getwd()
	require.NoError(t, err)

	var envs []string
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "TEST_COMMAND") {
			envs = append(envs, env)
		}
	}
	sort.Slice(envs, func(i int, j int) bool {
		return envs[i] < envs[j]
	})

	parameters := commandParameterGroup{
		Arguments:            arguments,
		WorkingDirectory:     workingDirectory,
		EnvironmentVariables: envs,
	}

	parametersBytes, err := json.Marshal(parameters)
	require.NoError(t, err)

	tempFileName := os.Args[len(os.Args)-1]
	err = ioutil.WriteFile(tempFileName, parametersBytes, 0600)
	require.NoError(t, err)

	if os.Getenv("TEST_COMMAND_FAILURE") == "TRUE" {
		failureMessage := os.Getenv("TEST_COMMAND_FAILURE_MESSAGE")
		if failureMessage != "" {
			fmt.Fprint(os.Stderr, failureMessage)
		}

		t.Fail()
	}
}

func TestRunCommand(t *testing.T) {
	type args struct {
		name                 string
		arguments            []string
		workingDirectory     string
		environmentVariables map[string]string
	}

	tests := []struct {
		name           string
		args           args
		wantParameters commandParameterGroup
		wantOutputPart string
		wantErr        assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				name: os.Args[0],
				arguments: []string{
					"-test.run=TestCommand",
					"test-command-argument-one",
					"test-command-argument-two",
				},
				workingDirectory: os.TempDir(),
				environmentVariables: map[string]string{
					"TEST_COMMAND":         "TRUE",
					"TEST_COMMAND_KEY_ONE": "value #1",
					"TEST_COMMAND_KEY_TWO": "value #2",
				},
			},
			wantParameters: commandParameterGroup{
				Arguments: []string{
					"test-command-argument-one",
					"test-command-argument-two",
				},
				WorkingDirectory: os.TempDir(),
				EnvironmentVariables: []string{
					"TEST_COMMAND=TRUE",
					"TEST_COMMAND_KEY_ONE=value #1",
					"TEST_COMMAND_KEY_TWO=value #2",
				},
			},
			wantOutputPart: "PASS",
			wantErr:        assert.NoError,
		},
		{
			name: "error without a message in stderr",
			args: args{
				name: os.Args[0],
				arguments: []string{
					"-test.run=TestCommand",
					"test-command-argument-one",
					"test-command-argument-two",
				},
				workingDirectory: os.TempDir(),
				environmentVariables: map[string]string{
					"TEST_COMMAND":         "TRUE",
					"TEST_COMMAND_KEY_ONE": "value #1",
					"TEST_COMMAND_KEY_TWO": "value #2",
					"TEST_COMMAND_FAILURE": "TRUE",
				},
			},
			wantParameters: commandParameterGroup{
				Arguments: []string{
					"test-command-argument-one",
					"test-command-argument-two",
				},
				WorkingDirectory: os.TempDir(),
				EnvironmentVariables: []string{
					"TEST_COMMAND=TRUE",
					"TEST_COMMAND_FAILURE=TRUE",
					"TEST_COMMAND_KEY_ONE=value #1",
					"TEST_COMMAND_KEY_TWO=value #2",
				},
			},
			wantOutputPart: "",
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.IsType(t, (*exec.ExitError)(nil), err)
			},
		},
		{
			name: "error with a message in stderr",
			args: args{
				name: os.Args[0],
				arguments: []string{
					"-test.run=TestCommand",
					"test-command-argument-one",
					"test-command-argument-two",
				},
				workingDirectory: os.TempDir(),
				environmentVariables: map[string]string{
					"TEST_COMMAND":                 "TRUE",
					"TEST_COMMAND_KEY_ONE":         "value #1",
					"TEST_COMMAND_KEY_TWO":         "value #2",
					"TEST_COMMAND_FAILURE":         "TRUE",
					"TEST_COMMAND_FAILURE_MESSAGE": "failure",
				},
			},
			wantParameters: commandParameterGroup{
				Arguments: []string{
					"test-command-argument-one",
					"test-command-argument-two",
				},
				WorkingDirectory: os.TempDir(),
				EnvironmentVariables: []string{
					"TEST_COMMAND=TRUE",
					"TEST_COMMAND_FAILURE=TRUE",
					"TEST_COMMAND_FAILURE_MESSAGE=failure",
					"TEST_COMMAND_KEY_ONE=value #1",
					"TEST_COMMAND_KEY_TWO=value #2",
				},
			},
			wantOutputPart: "",
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				if !assert.NotNil(t, err) {
					return false
				}

				return assert.Contains(t, err.Error(), "failure")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempFile, err := ioutil.TempFile("", "test.*")
			require.NoError(t, err)
			defer os.Remove(tempFile.Name())
			defer tempFile.Close()

			arguments := append(tt.args.arguments, tempFile.Name())
			got, gotErr := RunCommand(
				tt.args.name,
				arguments,
				tt.args.workingDirectory,
				tt.args.environmentVariables,
			)

			gotParametersBytes, err := ioutil.ReadFile(tempFile.Name())
			require.NoError(t, err)

			var gotParameters commandParameterGroup
			err = json.Unmarshal(gotParametersBytes, &gotParameters)
			require.NoError(t, err)

			assert.Equal(t, tt.wantParameters, gotParameters)
			assert.Contains(t, string(got), tt.wantOutputPart)
			tt.wantErr(t, gotErr)
		})
	}
}
