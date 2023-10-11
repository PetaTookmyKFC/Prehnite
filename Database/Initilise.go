package database

func StartDatabase(dataStore string) (DD *DBDepot, err error) {

	// Initilise the Depot
	DD = &DBDepot{
		Containers: make(map[string]DBContainer),
	}

	// Add the buildin Depot
	err = DD.AddDBContainer(dataStore)
	if err != nil {
		return nil, err
	}

	return DD, nil
}
