package Depot

import "path/filepath"

func StartDatabase(workingDirectory string) (DD *DBDepot, err error) {

	dataStore := filepath.Join(workingDirectory, "DatabaseDepot")

	// Try and load config file
	worked, depot := CreateFromConfig(dataStore)
	// Check if the config loaded
	if worked {
		// Save the returned depot the to global store
		DD = depot
	} else {
		// Initilise a blank depot
		DD = &DBDepot{
			WorkingDir: dataStore,
			Containers: make(map[string]DBContainer),
		}
		// Add the buildin Depot
		err = DD.AddDBContainer(filepath.Join(dataStore, "Default"))
		if err != nil {
			return nil, err
		}
		DD.SaveConfig()
	}

	return DD, nil
}
