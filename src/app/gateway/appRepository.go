package gateway

import "app/infra/database"

//GetApp get app info
func (dbRepo *DatabaseRepository) GetApp(appName string) (database.App, error) {
	var fitnessApp database.App
	err := dbRepo.GetDBHandle().Where("name=?", appName).First(&fitnessApp).Error
	if err != nil {
		return database.App{}, err
	}
	return fitnessApp, nil
}

//CreateApp create App
func (dbRepo *DatabaseRepository) CreateApp(app *database.App) error {
	err := dbRepo.GetDBHandle().Model(database.App{}).Save(&app).Error
	if err == nil {
		return err
	}
	return nil
}
