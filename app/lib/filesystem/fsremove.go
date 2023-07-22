// Content managed by Project Forge, see [projectforge.md] for details.
package filesystem

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/kyleu/solitaire/app/util"
)

func (f *FileSystem) Remove(path string, logger util.Logger) error {
	p := f.getPath(path)
	logger.Warnf("removing file at path [%s]", p)
	if err := os.Remove(p); err != nil {
		return errors.Wrapf(err, "error removing file [%s]", path)
	}
	return nil
}

func (f *FileSystem) RemoveRecursive(path string, logger util.Logger) error {
	if !f.Exists(path) {
		return nil
	}
	p := f.getPath(path)
	s, err := os.Stat(p)
	if err != nil {
		return errors.Wrapf(err, "unable to stat file [%s]", path)
	}
	if s.IsDir() {
		var files []fs.DirEntry
		files, err = os.ReadDir(p)
		if err != nil {
			logger.Warnf("cannot read path [%s] for removal: %+v", path, err)
		}
		for _, file := range files {
			err = f.RemoveRecursive(filepath.Join(path, file.Name()), logger)
			if err != nil {
				return err
			}
		}
	}
	err = os.Remove(p)
	if err != nil {
		return errors.Wrapf(err, "unable to remove file [%s]", path)
	}
	return nil
}
