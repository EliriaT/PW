package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	e "github.com/EliriaT/News-Tg-Bot/lib/error"
	"github.com/EliriaT/News-Tg-Bot/storage"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// perm to write and read to all users
const defaultPerm = 0774

type Storage struct {
	basePath string
}

func NewFileStorage(basePath string) Storage {
	return Storage{basePath: basePath}
}

// basepath/usernamefolder/filenamehash
func (s Storage) Save(link *storage.Link) (err error) {
	// wrap the error before returning it
	defer func() { err = e.WrapIfErr("can't save to disk", err) }()

	filePath := filepath.Join(s.basePath, link.ID)

	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}

	newFileName, err := fileName(link)
	if err != nil {
		return err
	}

	filePath = filepath.Join(filePath, newFileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	if err = gob.NewEncoder(file).Encode(link); err != nil {
		return err
	}
	return nil
}

func (s Storage) PickRandom(userName string) (page *storage.Link, err error) {
	defer func() { err = e.WrapIfErr("can't pick random page ", err) }()
	filesPath := filepath.Join(s.basePath, userName)

	//check that the username folder indeed exists
	//

	files, err := os.ReadDir(filesPath)
	if err != nil {
		return nil, err
	}
	// the bot will show that there are no links
	if len(files) == 0 {
		return nil, storage.ErrNoSavedLinks
	}
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(files))

	randomFile := files[i]
	return s.decodePage(filepath.Join(filesPath, randomFile.Name()))
}

func (s Storage) Remove(page *storage.Link) (err error) {
	defer func() { err = e.WrapIfErr("can't delete page ", err) }()
	fName, err := fileName(page)
	if err != nil {
		return err
	}

	removablePath := filepath.Join(s.basePath, page.ID, fName)

	if err := os.Remove(removablePath); err != nil {
		return e.Wrap(fmt.Sprintf("can't delete page %s", removablePath), err)
	}
	return nil
}

func (s Storage) IsPresent(link *storage.Link) (present bool, err error) {
	defer func() { err = e.WrapIfErr("can't check presence of page ", err) }()
	fName, err := fileName(link)
	if err != nil {
		return false, err
	}

	path := filepath.Join(s.basePath, link.ID, fName)

	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		return false, e.Wrap(fmt.Sprintf("%s", path), err)
	}
	return true, nil
}

func fileName(l *storage.Link) (string, error) {
	return l.Hash()
}

func (s Storage) decodePage(filepath string) (link *storage.Link, err error) {
	defer func() { err = e.WrapIfErr("can't decode page ", err) }()
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	if err := gob.NewDecoder(file).Decode(link); err != nil {
		return nil, err
	}
	return
}
