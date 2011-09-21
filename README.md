Git bindings for go. 
Just wraps commandline git using exec.
Doesnt handle all errors well.

``` go
import "github.com/srijak/gitgo"
.
.
g := NewGit(dir)
// initialize if you want to create a new git repo
g.Init()

// add a new file to the repo.
g.WriteFile("filename.0", "contents of the file")
g.CommitAll("added filename.0")

files, _ := g.ListFiles()  => ["filename.0"]

// get the current commitHash
hash, _ := g.GetCurrentCommitHash()  => "mADEupHash"

// add a new file to the repo.
g.WriteFile("filename.1", "contents of the file")
g.WriteFile("filename.2", "contents of the file")
g.CommitAll("added filename.1,2")

// get a list of all files in repo
files, _ := g.ListFiles()  => ["filename.0","filename.1", "filename.2"]

// get a list of all files changes since <hash>
files, _ = get.ListFilesChangedSince(hash) => ['filename.1', "filename.2"]

```
