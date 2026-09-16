package customfuse

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/yaml"
)

// examplesDir is the demo tree, relative to this package. Nothing else in the
// repository reads it: the scripts are baked into images or handed over in a
// ConfigMap and then run by mount-proxy, so a mistake in one only surfaces as a
// volume that will not come up.
const examplesDir = "../../examples/customfuse"

// danglingContinuation matches a line-continuation backslash with something after it.
// Bash takes the backslash literally there, so the arguments on the following lines are
// dropped from the command rather than passed to it. `bash -n` accepts that, so it needs
// a check of its own.
var danglingContinuation = regexp.MustCompile(`\\(?:[ \t]+|#)`)

// The two things `bash -n` does not look at, and the two most likely ways a demo breaks
// for whoever copies it: a script with no shebang is ENOEXEC at execve, and one that
// starts the client without exec leaves the shell in front of it, so mount-proxy waits
// on a process that is not the FUSE daemon and the client outlives the thing tracking it.
var (
	shebang         = regexp.MustCompile(`^#!`)
	clientAsLastAct = regexp.MustCompile(`^exec\s+\S`)
)

func TestExampleEntrypointsAreValidShell(t *testing.T) {
	if _, err := os.Stat(examplesDir); err != nil {
		t.Skipf("%s is not present: %v", examplesDir, err)
	}
	bash, err := exec.LookPath("bash")
	if err != nil {
		t.Skipf("bash is not installed: %v", err)
	}

	scripts := exampleScripts(t)
	require.NotEmpty(t, scripts, "no script found under %s", examplesDir)
	assert.Contains(t, scriptNames(scripts), filepath.Join(examplesDir, "4-configmap/configmap-production.yaml:entrypoint.sh"),
		"demo 4 keeps its script inside a ConfigMap; if that one is missing the walk is not reading manifests")

	for _, script := range scripts {
		t.Run(script.name, func(t *testing.T) {
			for i, line := range strings.Split(script.body, "\n") {
				assert.NotRegexp(t, danglingContinuation, line,
					"line %d: nothing may follow a continuation backslash", i+1)
			}

			path := filepath.Join(t.TempDir(), "entrypoint.sh")
			require.NoError(t, os.WriteFile(path, []byte(script.body), 0o600))
			out, err := exec.Command(bash, "-n", path).CombinedOutput()
			assert.NoError(t, err, "bash -n: %s", out)
		})
	}
}

func TestExampleEntrypointsHandOverToTheClient(t *testing.T) {
	if _, err := os.Stat(examplesDir); err != nil {
		t.Skipf("%s is not present: %v", examplesDir, err)
	}

	scripts := exampleScripts(t)
	require.NotEmpty(t, scripts, "no script found under %s", examplesDir)

	for _, script := range scripts {
		t.Run(script.name, func(t *testing.T) {
			assert.Regexp(t, shebang, strings.Split(script.body, "\n")[0],
				"the first line has to name an interpreter")
			assert.Regexp(t, clientAsLastAct, lastCommand(script.body),
				"the client has to be the entrypoint's last act, so it is the process mount-proxy waits on")
		})
	}
}

// lastCommand is the final line that is neither blank nor a comment: what the script
// actually does last, which is what mount-proxy ends up waiting on.
func lastCommand(body string) string {
	lines := strings.Split(body, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		if line := strings.TrimSpace(lines[i]); line != "" && !strings.HasPrefix(line, "#") {
			return line
		}
	}
	return ""
}

// A check that has stopped running looks exactly like a clean tree. Each of the three
// above has to be shown failing on a script written to fail it.
func TestExampleChecksRejectABrokenScript(t *testing.T) {
	const (
		noShebang  = "set -e\nexec /usr/local/bin/fuse-client \"$source\" \"$mountpoint\"\n"
		keepsGoing = "#!/bin/bash\nset -e\n/usr/local/bin/fuse-client \"$source\" \"$mountpoint\"\nwait\n"
		// The space between the backslash and the newline is the bug being caught:
		// bash takes that backslash literally instead of joining the lines.
		danglingLine = "format \\ \n    --storage=oss \\\n"
		unparsable   = "#!/bin/bash\nif true; then\n"
	)

	assert.NotRegexp(t, shebang, strings.Split(noShebang, "\n")[0], "no interpreter named")
	assert.NotRegexp(t, clientAsLastAct, lastCommand(keepsGoing), "the shell is still in front of the client")
	assert.Regexp(t, danglingContinuation, strings.Split(danglingLine, "\n")[0], "a continuation with something after it")

	bash, err := exec.LookPath("bash")
	if err != nil {
		t.Skipf("bash is not installed: %v", err)
	}
	path := filepath.Join(t.TempDir(), "unparsable.sh")
	require.NoError(t, os.WriteFile(path, []byte(unparsable), 0o600))
	assert.Error(t, exec.Command(bash, "-n", path).Run(), "the syntax check itself")
}

type exampleScript struct {
	name string
	body string
}

// exampleScripts collects every script a demo can hand to the driver: the entrypoint.sh
// files sitting next to their Dockerfile, and the ones embedded in a ConfigMap's data,
// which is how an administrator supplies a script without building an image at all.
func exampleScripts(t *testing.T) []exampleScript {
	var scripts []exampleScript
	require.NoError(t, filepath.WalkDir(examplesDir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".sh" {
			scripts = append(scripts, exampleScript{path, string(raw)})
			return nil
		}
		if ext := filepath.Ext(path); ext != ".yaml" && ext != ".yml" {
			return nil
		}
		// A file may hold several manifests, and most of them are not ConfigMaps.
		for _, doc := range strings.Split(string(raw), "\n---") {
			var manifest struct {
				Data map[string]string `json:"data"`
			}
			if err := yaml.Unmarshal([]byte(doc), &manifest); err != nil {
				continue
			}
			for key, body := range manifest.Data {
				if filepath.Ext(key) == ".sh" {
					scripts = append(scripts, exampleScript{path + ":" + key, body})
				}
			}
		}
		return nil
	}))
	return scripts
}

func scriptNames(scripts []exampleScript) []string {
	names := make([]string, 0, len(scripts))
	for _, script := range scripts {
		names = append(names, script.name)
	}
	return names
}
