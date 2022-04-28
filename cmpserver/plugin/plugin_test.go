package plugin

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/argoproj/argo-cd/v2/cmpserver/apiclient"
	"github.com/argoproj/argo-cd/v2/test"
)

func newService(configFilePath string) (*Service, error) {
	config, err := ReadPluginConfig(configFilePath)
	if err != nil {
		return nil, err
	}

	initConstants := CMPServerInitConstants{
		PluginConfig: *config,
	}

	service := &Service{
		initConstants: initConstants,
	}
	return service, nil
}

type pluginOpt func(*CMPServerInitConstants)

func withDiscover(d Discover) pluginOpt {
	return func(cic *CMPServerInitConstants) {
		cic.PluginConfig.Spec.Discover = d
	}
}

func buildPluginConfig(opts ...pluginOpt) *CMPServerInitConstants {
	cic := &CMPServerInitConstants{
		PluginConfig: PluginConfig{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ConfigManagementPlugin",
				APIVersion: "argoproj.io/v1alpha1",
			},
			Metadata: metav1.ObjectMeta{
				Name: "some-plugin",
			},
			Spec: PluginConfigSpec{
				Version: "v1.0",
			},
		},
	}
	for _, opt := range opts {
		opt(cic)
	}
	return cic
}

func TestMatchRepository(t *testing.T) {
	type fixture struct {
		service *Service
		path    string
	}
	setup := func(t *testing.T, opts ...pluginOpt) *fixture {
		t.Helper()
		cic := buildPluginConfig(opts...)
		path := filepath.Join(test.GetTestDir(t), "testdata", "kustomize")
		s := NewService(*cic)
		return &fixture{
			service: s,
			path:    path,
		}
	}
	t.Run("will match plugin by filename", func(t *testing.T) {
		// given
		d := Discover{
			FileName: "kustomization.yaml",
		}
		f := setup(t, withDiscover(d))

		// when
		match, err := f.service.matchRepository(context.Background(), f.path)

		// then
		assert.NoError(t, err)
		assert.True(t, match)
	})
	t.Run("will not match plugin by filename if file not found", func(t *testing.T) {
		// given
		d := Discover{
			FileName: "not_found.yaml",
		}
		f := setup(t, withDiscover(d))

		// when
		match, err := f.service.matchRepository(context.Background(), f.path)

		// then
		assert.NoError(t, err)
		assert.False(t, match)
	})
	t.Run("will match plugin by glob", func(t *testing.T) {
		// given
		d := Discover{
			Find: Find{
				Glob: "**/*/plugin.yaml",
			},
		}
		f := setup(t, withDiscover(d))

		// when
		match, err := f.service.matchRepository(context.Background(), f.path)

		// then
		assert.NoError(t, err)
		assert.True(t, match)
	})
	t.Run("will not match plugin by glob if not found", func(t *testing.T) {
		// given
		d := Discover{
			Find: Find{
				Glob: "**/*/not_found.yaml",
			},
		}
		f := setup(t, withDiscover(d))

		// when
		match, err := f.service.matchRepository(context.Background(), f.path)

		// then
		assert.NoError(t, err)
		assert.False(t, match)
	})
	t.Run("will match plugin by command when returns any output", func(t *testing.T) {
		// given
		d := Discover{
			Find: Find{
				Command: Command{
					Command: []string{"echo", "test"},
				},
			},
		}
		f := setup(t, withDiscover(d))

		// when
		match, err := f.service.matchRepository(context.Background(), f.path)

		// then
		assert.NoError(t, err)
		assert.True(t, match)
	})
	t.Run("will not match plugin by command when returns no output", func(t *testing.T) {
		// given
		d := Discover{
			Find: Find{
				Command: Command{
					Command: []string{"echo"},
				},
			},
		}
		f := setup(t, withDiscover(d))

		// when
		match, err := f.service.matchRepository(context.Background(), f.path)

		// then
		assert.NoError(t, err)
		assert.False(t, match)
	})
	t.Run("will not match plugin by command when command fails", func(t *testing.T) {
		// given
		d := Discover{
			Find: Find{
				Command: Command{
					Command: []string{"cat", "nil"},
				},
			},
		}
		f := setup(t, withDiscover(d))

		// when
		match, err := f.service.matchRepository(context.Background(), f.path)

		// then
		assert.Error(t, err)
		assert.False(t, match)
	})
}

func Test_Negative_ConfigFile_DoesnotExist(t *testing.T) {
	configFilePath := "./testdata/kustomize-neg/config"
	service, err := newService(configFilePath)
	require.Error(t, err)
	require.Nil(t, service)
}

func TestGenerateManifest(t *testing.T) {
	configFilePath := "./testdata/kustomize/config"
	service, err := newService(configFilePath)
	require.NoError(t, err)

	res1, err := service.generateManifest(context.Background(), "", nil)
	require.NoError(t, err)
	require.NotNil(t, res1)

	expectedOutput := "{\"apiVersion\":\"v1\",\"data\":{\"foo\":\"bar\"},\"kind\":\"ConfigMap\",\"metadata\":{\"name\":\"my-map\"}}"
	if res1 != nil {
		require.Equal(t, expectedOutput, res1.Manifests[0])
	}
}

// TestRunCommandContextTimeout makes sure the command dies at timeout rather than sleeping past the timeout.
func TestRunCommandContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 990*time.Millisecond)
	defer cancel()
	// Use a subshell so there's a child command.
	command := Command{
		Command: []string{"sh", "-c"},
		Args:    []string{"sleep 5"},
	}
	before := time.Now()
	_, err := runCommand(ctx, command, "", []string{})
	after := time.Now()
	assert.Error(t, err) // The command should time out, causing an error.
	assert.Less(t, after.Sub(before), 1*time.Second)
}

func Test_getParametersAnnouncement_empty_command(t *testing.T) {
	staticYAML := `
- name: static-a
- name: static-b
`
	static := &[]Static{}
	err := yaml.Unmarshal([]byte(staticYAML), static)
	require.NoError(t, err)
	command := Command{
		Command: []string{"echo"},
		Args:    []string{`[]`},
	}
	res, err := getParametersAnnouncement(context.Background(), "", *static, command)
	require.NoError(t, err)
	assert.Equal(t, []*apiclient.ParameterAnnouncement{{Name: "static-a"}, {Name: "static-b"}}, res.ParameterAnnouncements)
}

func Test_getParametersAnnouncement_no_command(t *testing.T) {
	staticYAML := `
- name: static-a
- name: static-b
`
	static := &[]Static{}
	err := yaml.Unmarshal([]byte(staticYAML), static)
	require.NoError(t, err)
	command := Command{}
	res, err := getParametersAnnouncement(context.Background(), "", *static, command)
	require.NoError(t, err)
	assert.Equal(t, []*apiclient.ParameterAnnouncement{{Name: "static-a"}, {Name: "static-b"}}, res.ParameterAnnouncements)
}

func Test_getParametersAnnouncement_static_and_dynamic(t *testing.T) {
	staticYAML := `
- name: static-a
- name: static-b
`
	static := &[]Static{}
	err := yaml.Unmarshal([]byte(staticYAML), static)
	require.NoError(t, err)
	command := Command{
		Command: []string{"echo"},
		Args:    []string{`[{"name": "dynamic-a"}, {"name": "dynamic-b"}]`},
	}
	res, err := getParametersAnnouncement(context.Background(), "", *static, command)
	require.NoError(t, err)
	expected := []*apiclient.ParameterAnnouncement{
		{Name: "dynamic-a"},
		{Name: "dynamic-b"},
		{Name: "static-a"},
		{Name: "static-b"},
	}
	assert.Equal(t, expected, res.ParameterAnnouncements)
}

func Test_getParametersAnnouncement_invalid_json(t *testing.T) {
	command := Command{
		Command: []string{"echo"},
		Args:    []string{`[`},
	}
	_, err := getParametersAnnouncement(context.Background(), "", []Static{}, command)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

func Test_getParametersAnnouncement_bad_command(t *testing.T) {
	command := Command{
		Command: []string{"exit"},
		Args:    []string{"1"},
	}
	_, err := getParametersAnnouncement(context.Background(), "", []Static{}, command)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error executing dynamic parameter output command")
}

func Test_getTempDirMustCleanup(t *testing.T) {
	tempDir := t.TempDir()
	workDir, cleanup, err := getTempDirMustCleanup(tempDir)
	require.NoError(t, err)
	require.DirExists(t, workDir)

	// Induce a cleanup error to verify panic behavior.
	err = os.Chmod(tempDir, 0000)
	require.NoError(t, err)
	assert.Panics(t, func() {
		cleanup()
	}, "cleanup must panic to protect from directory traversal vulnerabilities")

	err = os.Chmod(tempDir, 0700)
	require.NoError(t, err)
	cleanup()
	assert.NoDirExists(t, workDir)
}
