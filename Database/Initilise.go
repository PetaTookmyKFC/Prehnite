package database

import "path/filepath"

func StartDatabase(workingDirectory string) (DD *DBDepot, err error) {

	dataStore := filepath.Join(workingDirectory, "DatabaseDepot")

	// Initilise the Depot
	DD = &DBDepot{
		WorkingDir: dataStore,
		Containers: make(map[string]DBContainer),
	}

	// Add the buildin Depot
	err = DD.AddDBContainer(filepath.Join(dataStore, "Default"))
	if err != nil {
		return nil, err
	}

	return DD, nil
}

// Read all Depos within working Directory
// Save info about depos into config
