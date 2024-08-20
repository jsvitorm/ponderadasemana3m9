// db/database_test.go
package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseTestSuite struct {
	suite.Suite
	db *gorm.DB
}

type User struct {
	ID   uint
	Name string
}

func (suite *DatabaseTestSuite) SetupSuite() {
	//dsn := "user=testuser password=testpassword dbname=testdb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	suite.Require().NoError(err, "Error connecting to the test database")
	suite.db = db.Debug()

	err = suite.db.AutoMigrate(&User{})
	suite.Require().NoError(err, "Error auto-migrating database tables")
}

func (suite *DatabaseTestSuite) TestUserInsertion() {
	user := User{Name: "John Doe"}
	err := suite.db.Create(&user).Error
	suite.Require().NoError(err, "Error creating user record")

	var retrievedUser User
	err = suite.db.First(&retrievedUser, "name = ?", "John Doe").Error
	suite.Require().NoError(err, "Error retrieving user record")

	suite.Equal(user.Name, retrievedUser.Name, "Names should match")
}

func (suite *DatabaseTestSuite) TearDownSuite() {
	err := suite.db.Exec("DROP TABLE users;").Error
	suite.Require().NoError(err, "Error dropping test table")

	sqlDB, err := suite.db.DB()
	suite.Require().NoError(err, "Error getting SQL DB")
	err = sqlDB.Close()
	suite.Require().NoError(err, "Error closing the test database")
}

func TestSuite(t *testing.T) {
	if os.Getenv("POSTGRES_DSN") == "" {
		t.Skip("Skipping PostgreSQL tests; provide POSTGRES_DSN environment variable.")
	}
	suite.Run(t, new(DatabaseTestSuite))
}
