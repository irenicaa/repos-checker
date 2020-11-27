package config

import (
	"encoding/json"
	"io"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/irenicaa/repos-checker/loader"
	"github.com/irenicaa/repos-checker/loader/sources/bitbucket"
	"github.com/irenicaa/repos-checker/loader/sources/external"
	filesystem "github.com/irenicaa/repos-checker/loader/sources/file-system"
	"github.com/irenicaa/repos-checker/loader/sources/github"
	"github.com/irenicaa/repos-checker/loader/sources/gitlab"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		reader io.Reader
		logger loader.Logger
	}

	tests := []struct {
		name              string
		args              args
		wantSources       []loader.Source
		wantReferenceName string
		wantErr           assert.ErrorAssertionFunc
	}{
		{
			name: "success (empty)",
			args: args{
				reader: strings.NewReader("[]"),
				logger: &MockLogger{},
			},
			wantSources:       nil,
			wantReferenceName: "",
			wantErr:           assert.NoError,
		},
		{
			name: "success (nonempty, without a specified reference)",
			args: args{
				reader: strings.NewReader(`[
					{
						"name": "github",
						"options": {
							"owner": "one",
							"pageSize": 23
						}
					},
					{
						"name": "bitbucket",
						"options": {
							"workspace": "two",
							"pageSize": 42
						}
					}
				]`),
				logger: &MockLogger{},
			},
			wantSources: []loader.Source{
				&github.Source{
					Owner:    "one",
					PageSize: 23,
					Logger:   &MockLogger{},
				},
				&bitbucket.Source{
					Workspace: "two",
					PageSize:  42,
					Logger:    &MockLogger{},
				},
			},
			wantReferenceName: "",
			wantErr:           assert.NoError,
		},
		{
			name: "success (nonempty, with a specified reference)",
			args: args{
				reader: strings.NewReader(`[
					{
						"name": "github",
						"options": {
							"owner": "one",
							"pageSize": 23
						}
					},
					{
						"name": "bitbucket",
						"isReference": true,
						"options": {
							"workspace": "two",
							"pageSize": 42
						}
					}
				]`),
				logger: &MockLogger{},
			},
			wantSources: []loader.Source{
				&github.Source{
					Owner:    "one",
					PageSize: 23,
					Logger:   &MockLogger{},
				},
				&bitbucket.Source{
					Workspace: "two",
					PageSize:  42,
					Logger:    &MockLogger{},
				},
			},
			wantReferenceName: "bitbucket",
			wantErr:           assert.NoError,
		},

		// errors
		{
			name: "error (with config reading)",
			args: args{
				reader: iotest.TimeoutReader(strings.NewReader(`[
					{
						"name": "github",
						"options": {
							"owner": "one",
							"pageSize": 23
						}
					},
					{
						"name": "bitbucket",
						"isReference": true,
						"options": {
							"workspace": "two",
							"pageSize": 42
						}
					}
				]`)),
				logger: &MockLogger{},
			},
			wantSources:       nil,
			wantReferenceName: "",
			wantErr:           assert.Error,
		},
		{
			name: "error (with config unmarshalling)",
			args: args{
				reader: strings.NewReader("["),
				logger: &MockLogger{},
			},
			wantSources:       nil,
			wantReferenceName: "",
			wantErr:           assert.Error,
		},
		{
			name: "error (with multiple references)",
			args: args{
				reader: strings.NewReader(`[
					{
						"name": "github",
						"isReference": true,
						"options": {
							"owner": "one",
							"pageSize": 23
						}
					},
					{
						"name": "bitbucket",
						"isReference": true,
						"options": {
							"workspace": "two",
							"pageSize": 42
						}
					}
				]`),
				logger: &MockLogger{},
			},
			wantSources:       nil,
			wantReferenceName: "",
			wantErr:           assert.Error,
		},
		{
			name: "error (with source loading)",
			args: args{
				reader: strings.NewReader(`[
					{
						"name": "github",
						"options": {
							"owner": "one",
							"pageSize": 23
						}
					},
					{
						"name": "unknown"
					}
				]`),
				logger: &MockLogger{},
			},
			wantSources:       nil,
			wantReferenceName: "",
			wantErr:           assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSources, gotReferenceName, err :=
				LoadConfig(tt.args.reader, tt.args.logger)

			tt.args.logger.(*MockLogger).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.wantSources, gotSources)
			assert.Equal(t, tt.wantReferenceName, gotReferenceName)
			tt.wantErr(t, err)
		})
	}
}

func TestLoadSource(t *testing.T) {
	type args struct {
		sourceConfig SourceConfig
		logger       loader.Logger
	}

	tests := []struct {
		name    string
		args    args
		want    loader.Source
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "github",
			args: args{
				sourceConfig: SourceConfig{
					Name: "github",
					Options: json.RawMessage([]byte(`{
						"owner": "test",
						"pageSize": 23,
						"logger": null
					}`)),
				},
				logger: &MockLogger{},
			},
			want: &github.Source{
				Owner:    "test",
				PageSize: 23,
				Logger:   &MockLogger{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "bitbucket",
			args: args{
				sourceConfig: SourceConfig{
					Name: "bitbucket",
					Options: json.RawMessage([]byte(`{
						"workspace": "test",
						"pageSize": 23,
						"logger": null
					}`)),
				},
				logger: &MockLogger{},
			},
			want: &bitbucket.Source{
				Workspace: "test",
				PageSize:  23,
				Logger:    &MockLogger{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "gitlab",
			args: args{
				sourceConfig: SourceConfig{
					Name: "gitlab",
					Options: json.RawMessage([]byte(`{
						"owner": "test",
						"pageSize": 23,
						"logger": null
					}`)),
				},
				logger: &MockLogger{},
			},
			want: &gitlab.Source{
				Owner:    "test",
				PageSize: 23,
				Logger:   &MockLogger{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "file-system",
			args: args{
				sourceConfig: SourceConfig{
					Name: "file-system",
					Options: json.RawMessage([]byte(`{
						"basePath": "test",
						"logger": null
					}`)),
				},
				logger: &MockLogger{},
			},
			want: &filesystem.Source{
				BasePath: "test",
				Logger:   &MockLogger{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "external",
			args: args{
				sourceConfig: SourceConfig{
					Name: "external",
					Options: json.RawMessage([]byte(`{
						"additionalName": "one",
						"command": "two",
						"arguments": ["three", "four"],
						"workingDirectory": "five",
						"environmentVariables": {"six": "seven", "eight": "nine"}
					}`)),
				},
				logger: &MockLogger{},
			},
			want: &external.Source{
				AdditionalName:       "one",
				Command:              "two",
				Arguments:            []string{"three", "four"},
				WorkingDirectory:     "five",
				EnvironmentVariables: map[string]string{"six": "seven", "eight": "nine"},
			},
			wantErr: assert.NoError,
		},
		{
			name: "multi-source (empty)",
			args: args{
				sourceConfig: SourceConfig{
					Name:       "multi-source",
					SubSources: []SourceConfig{},
				},
				logger: &MockLogger{},
			},
			want:    loader.MultiSource(nil),
			wantErr: assert.NoError,
		},
		{
			name: "multi-source (nonempty)",
			args: args{
				sourceConfig: SourceConfig{
					Name: "multi-source",
					SubSources: []SourceConfig{
						{
							Name: "github",
							Options: json.RawMessage([]byte(`{
								"owner": "test",
								"pageSize": 23,
								"logger": null
							}`)),
						},
						{
							Name: "bitbucket",
							Options: json.RawMessage([]byte(`{
								"workspace": "test",
								"pageSize": 23,
								"logger": null
							}`)),
						},
					},
				},
				logger: &MockLogger{},
			},
			want: loader.MultiSource{
				&github.Source{
					Owner:    "test",
					PageSize: 23,
					Logger:   &MockLogger{},
				},
				&bitbucket.Source{
					Workspace: "test",
					PageSize:  23,
					Logger:    &MockLogger{},
				},
			},
			wantErr: assert.NoError,
		},

		// errors
		{
			name: "unknown source",
			args: args{
				sourceConfig: SourceConfig{Name: "unknown"},
				logger:       &MockLogger{},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "incorrect options",
			args: args{
				sourceConfig: SourceConfig{
					Name:    "github",
					Options: json.RawMessage([]byte("{")),
				},
				logger: &MockLogger{},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "error in a multi-source",
			args: args{
				sourceConfig: SourceConfig{
					Name: "multi-source",
					SubSources: []SourceConfig{
						{
							Name:    "github",
							Options: json.RawMessage([]byte("{")),
						},
					},
				},
				logger: &MockLogger{},
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadSource(tt.args.sourceConfig, tt.args.logger)

			tt.args.logger.(*MockLogger).InnerMock.AssertExpectations(t)
			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}
