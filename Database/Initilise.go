package database

func StartDatabase(dataStore string) (DD *DatabaseDepot, err error) {
	// Create Empty Depot
	DD = &DatabaseDepot{
		Database: make(map[string]DatabaseContainer),
	}
	// Load Default Depot
	if DD.AddDepot(dataStore) != nil {
		return nil, err
	}

	return DD, nil
}