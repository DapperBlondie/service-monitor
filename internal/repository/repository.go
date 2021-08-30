package repository

import "github.com/DapperBlondie/service-monitor/internal/models"

// DatabaseRepo is the database repository
type DatabaseRepo interface {
	// AllPreferences preferences
	AllPreferences() ([]models.Preference, error)
	SetSystemPref(name, value string) error
	InsertOrUpdateSitePreferences(pm map[string]string) error

	// GetUserById users and authentication
	GetUserById(id int) (models.User, error)
	InsertUser(u models.User) (int, error)
	UpdateUser(u models.User) error
	DeleteUser(id int) error
	UpdatePassword(id int, newPassword string) error
	Authenticate(email, testPassword string) (int, string, error)
	AllUsers() ([]*models.User, error)
	InsertRememberMeToken(id int, token string) error
	DeleteToken(token string) error
	CheckForToken(id int, token string) bool
	InsertHost(h models.Host) (int, error)
	GetHostByID(id int) (*models.Host, error)
	UpdateHost(h *models.Host) error
}
