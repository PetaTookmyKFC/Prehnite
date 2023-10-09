package database

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type DatabaseDepot struct {
	Database map[string]DatabaseContainer
}

func (DD *DatabaseDepot) AddDatabase(Name string, Location string) (id string, err error) {
	RI := DatabaseContainer{
		Name:     Name,
		Location: Location,
	}

	// Check if name already in Depot - Repeat ( Max Repeat 100)
	Inc := 0
	tooMany := false
	orig := Name
	found := true
	for found && !tooMany {
		if Inc > 100 {
			tooMany = true
		} else {
			if _, ok := DD.Database[Name]; ok {
				// if name exists place _n* at the end
				found = true
				Inc++
				fmt.Printf("NameExists : %s\n", Name)
				Name = fmt.Sprintf("%s-n%d", orig, Inc)
			} else {
				found = false
			}
		}
	}
	if tooMany {
		fmt.Println("TooMany")
		return "", os.ErrExist
	}

	id = Name
	// Add database to Map with CustomID
	DD.Database[id] = RI
	return id, nil
}
func (DD *DatabaseDepot) AddDepot(Location string) error {
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
				_, err = DD.AddDatabase(
					strings.TrimSuffix(tmpName, filepath.Ext(tmpName)),
					filepath.Join(Location, tmpName),
				)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (DD *DatabaseDepot) DatabaseList() {
	for _, value := range DD.Database {
		fmt.Printf("%s : %s", value.Name, value.Location)
	}
}
