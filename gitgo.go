package gitgo

import (
	"exec"
	"strings"
	"os"
	"path"
	"io/ioutil"
)

type Git struct {
	bin      string
	dir      string
	root_dir string
}

func NewGit(rootPath string) *Git {
	g := new(Git)
	g.root_dir = rootPath
	g.dir = path.Join(g.root_dir, ".git")
	return g
}

func (g Git) getBin() string {
	if g.bin != "" {
		return g.bin
	}
	return "git"
}

func (g Git) exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func (g Git) Init() os.Error {
	if g.exists(g.dir) {
		return nil
	}
	if err := os.MkdirAll(g.root_dir, 0777); err != nil {
		panic(err)
	}
	_, err := g.run("init")
	return err
}
func (g Git) WriteFile(file, contents string) os.Error {
	absPath := path.Join(g.root_dir, file)
	err := ioutil.WriteFile(absPath, []byte(contents), 0777)
	if err != nil {
		return err
	}
	return g.Add(file)
}

func (g Git) Add(file string) os.Error {
	_, err := g.run("add", file)
	return err
}

func (g Git) CommitAll(msg string) os.Error {
	//escape msg at some point.
	_, err := g.run("commit", "-aqm", "\""+msg+"\"")
	return err
}
func (g Git) GetCurrentCommitHash() (string, os.Error) {
	out, err := g.run("rev-parse", "HEAD")
	if err != nil {
		return "", err
	}
	return out[0], nil
}

func (g Git) ListFilesChangedSince(commitHash string) ([]string, os.Error) {
	return g.run("diff", "--name-only", commitHash, "HEAD")
}

func (g Git) ListFiles() ([]string, os.Error) {
	return g.run("ls-files")
}

func isNewline(rune int) bool {
	return rune == '\n'
}

func (g Git) run(args ...string) ([]string, os.Error) {
	cmd := exec.Command(g.getBin(), args...)
	cmd.Dir = g.root_dir
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	lines := strings.TrimSpace(string(out))
	// must be a better way to split on \n 
	// *and* return empty array if empty string?
	return strings.FieldsFunc(lines, isNewline), err
}
