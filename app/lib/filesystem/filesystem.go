// Content managed by Project Forge, see [projectforge.md] for details.
package filesystem

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mandelsoft/vfs/pkg/memoryfs"
	"github.com/mandelsoft/vfs/pkg/osfs"
	"github.com/mandelsoft/vfs/pkg/readonlyfs"
	"github.com/mandelsoft/vfs/pkg/vfs"
	"github.com/pkg/errors"
)

type FileSystem struct {
	Mode     string `json:"mode,omitempty"`
	ReadOnly bool   `json:"readOnly,omitempty"`
	root     string
	f        vfs.FileSystem
}

var _ FileLoader = (*FileSystem)(nil)

func NewFileSystem(root string, readonly bool, mode string) (*FileSystem, error) {
	if mode == "" {
		mode = "file"
	}
	var f vfs.FileSystem
	switch mode {
	case "memory":
		f = memoryfs.New()
	case "file":
		f = osfs.New()
	case "":
		switch runtime.GOOS {
		case "js":
			f = memoryfs.New()
		default:
			f = osfs.New()
		}
	default:
		return nil, errors.Errorf("invalid filesystem mode [%s]", mode)
	}
	if readonly {
		f = readonlyfs.New(f)
	}
	return &FileSystem{root: root, f: f}, nil
}

func (f *FileSystem) getPath(ss ...string) string {
	s := filepath.Join(ss...)
	if strings.HasPrefix(s, f.root) {
		return s
	}
	return filepath.Join(f.root, s)
}

func (f *FileSystem) Root() string {
	return f.root
}

func (f *FileSystem) Clone() FileLoader {
	ret, _ := NewFileSystem(f.root, f.ReadOnly, f.Mode)
	return ret
}

func (f *FileSystem) Stat(path string) (*FileInfo, error) {
	p := f.getPath(path)
	s, err := f.f.Stat(p)
	if err != nil {
		return nil, err
	}
	return FileInfoFromFS(s), nil
}

func (f *FileSystem) SetMode(path string, mode FileMode) error {
	p := f.getPath(path)
	return f.f.Chmod(p, mode.ToFS())
}

func (f *FileSystem) Exists(path string) bool {
	x, _ := f.Stat(path)
	return x != nil
}

func (f *FileSystem) IsDir(path string) bool {
	s, err := f.Stat(path)
	if s == nil || err != nil {
		return false
	}
	return s.IsDir
}

func (f *FileSystem) CreateDirectory(path string) error {
	p := f.getPath(path)
	if err := f.f.MkdirAll(p, DirectoryMode.ToFS()); err != nil {
		return errors.Wrapf(err, "unable to create data directory [%s]", p)
	}
	return nil
}

func (f *FileSystem) String() string {
	return fmt.Sprintf("fs://%s", f.root)
}
