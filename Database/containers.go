package database

import (
	"fmt"
	"os"
)

type DBContainer struct {
	Database map[string]Database
}
type Database struct {
	Name     string
	Location string
}

func (DD *DBContainer) AddDatabase(Name string, Location string) (id string, err error) {
	RI := Database{
		Name:     Name,
		Location: Location,
	}

	// Check if name already in Container - Repeat ( Max Repeat 100)
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
		fmt.Println("Too Many Databases with the same name")
		return "", os.ErrExist
	}

	id = Name
	// Add database to Map with CustomID
	DD.Database[id] = RI
	return id, nil
}
