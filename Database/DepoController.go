package database

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type DBDepot struct {
	Containers map[string]DBContainer
}

func (DD *DBDepot) AddDBContainer(Location string) error {

	// Create TmpName
	TmpName := filepath.Base(Location)
	// Check Container exists
	Inc := 0
	tooMany := false
	orig := TmpName
	found := true
	for found && !tooMany {
		if Inc > 100 {
			tooMany = true
		} else {
			if _, ok := DD.Containers[TmpName]; ok {
				// if name exists place _n* at the end
				found = true
				Inc++
				fmt.Printf("NameExists : %s\n", TmpName)
				TmpName = fmt.Sprintf("%s-n%d", orig, Inc)
			} else {
				found = false
			}
		}
	}

	// Check Directory Exists
	Inst_Container := DBContainer{
		Database: make(map[string]Database),
	}

	if _, ok := os.Stat(Location); os.IsNotExist(ok) {
		os.MkdirAll(Location, fs.ModeDir)
	} else if ok != nil {
		return ok
	}

	// Scan Directory
	dirItem, err := os.ReadDir(Location)
	if err != nil {
		return err
	}

	// Loop over directory contents
	for _, value := range dirItem {
		// Is item not a folder
		if !value.Type().IsDir() {
			// Is the file a .db file
			tmpName := value.Name()

			if filepath.Ext(tmpName) == ".db" {
				// Add Database to Depot
				_, err = Inst_Container.AddDatabase(
					strings.TrimSuffix(tmpName, filepath.Ext(tmpName)),
					filepath.Join(Location, tmpName),
				)
				if err != nil {
					return err
				}
			}
		}
	}
	DD.Containers[TmpName] = Inst_Container
	return nil
}

func (DD *DBDepot) ListContainers() []string {
	var rString []string
	for id := range DD.Containers {
		rString = append(rString, id)
	}
	return rString
}

// func (DD *DBContainer) DatabaseList() {
// 	for _, value := range DD.Database {
// 		fmt.Printf("%s : %s", value.Name, value.Location)
// 	}
// }
