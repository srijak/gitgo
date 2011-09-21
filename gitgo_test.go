package gitgo

import (
	"testing"
	"io/ioutil"
)

func assertEqual(t *testing.T, o, e interface{}) {
	if o != e {
		t.Errorf("expected %d, got %d", e, o)
	}
}

func assertNil(t *testing.T, o interface{}) {
	if o != nil {
		t.Error("expected nil, got", o)
	}
}

func assertNotNil(t *testing.T, o interface{}) {
	if o == nil {
		t.Error("expected NOT nil, got", o)
	}
}

func getTestGit() *Git {
	dir, err := ioutil.TempDir("", "gitgo_test")
	if err != nil {
		panic(err)
	}
	g := NewGit(dir)
	g.Init()

	return g
}
func Test_NewGitDir_NoFiles(t *testing.T) {
	g := getTestGit()
	files, err := g.ListFiles()
	assertEqual(t, len(files), 0)
	assertNil(t, err)
}

func Test_NewGitDir_AddFile(t *testing.T) {
	g := getTestGit()
	err := g.WriteFile("a", "contents of a")
	if err != nil {
		panic(err)
	}
	g.CommitAll("added a")

	files, err := g.ListFiles()
	assertEqual(t, len(files), 1)
	assertNil(t, err)
}

func Test_NewGitDir_ListFilesChangedSince(t *testing.T) {
	g := getTestGit()
	err := g.WriteFile("a", "contents of a")
	if err != nil {
		panic(err)
	}
	g.CommitAll("added a")
	hash, _ := g.GetCurrentCommitHash()

	g.WriteFile("b", "contents of b")
	g.WriteFile("c", "contents of c")
	g.CommitAll("added b, c")

	files, err := g.ListFiles()
	assertEqual(t, len(files), 3)
	assertNil(t, err)

	files, err = g.ListFilesChangedSince(hash)
	assertEqual(t, len(files), 2)
	assertNil(t, err)
}
