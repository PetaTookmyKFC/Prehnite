package database

import "path/filepath"

// Save Information about depos into config
func (DD *DBDepot) SaveConfig() {

	configFile := filepath.Join(DD.WorkingDir, "config.yaml")

}
func (DD *DBDepot) LoadConfig() {

}
