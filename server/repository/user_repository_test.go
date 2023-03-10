package repository

import (
	"errors"
	"fmt"
	"gostarter/server/model"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var testTime = time.Now()

func TestFindUserByEmail(t *testing.T) {
	_, mock, creatingMockErr := sqlmock.New()

	if creatingMockErr != nil {
		t.Fatalf(
			"an error '%s' occurred during creating a stub database connection",
			creatingMockErr,
		)
	}

	userID := uint(1)
	userEmail := "test@email.com"
	userPassword := "test pass"
	userFullName := "test full name"

	rows := mock.NewRows([]string{
		"id",
		"created_at",
		"updated_at",
		"deleted_at",
		"email",
		"password",
		"full_name",
	}).
		AddRow(
			userID,
			testTime,
			testTime,
			nil,
			userEmail,
			userPassword,
			userFullName,
		)

	query := "SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((email = ?))"

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs("test@email.com").
		WillReturnRows(rows)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"@Dragontol123",
		"localhost",
		"3306",
		"queue",
	)

	mockedDbConn, connErr := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if connErr != nil {
		t.Fatalf(
			"an error '%s' occurred during opening a stub database connection",
			connErr,
		)
	}

	want := model.User{
		Model: gorm.Model{
			ID:        userID,
			CreatedAt: testTime,
			UpdatedAt: testTime,
			DeletedAt: gorm.DeletedAt{},
		},
		Email:    userEmail,
		Password: userPassword,
		FullName: userFullName,
	}

	got, err := NewUsersRepository(mockedDbConn).FindUserByEmail(userEmail)

	assert.NoError(t, err)
	assert.Equal(t, want, got)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindUserByEmailReturnsError(t *testing.T) {
	_, mock, creatingMockErr := sqlmock.New()

	if creatingMockErr != nil {
		t.Fatalf(
			"an error '%s' occurred during creating a stub database connection",
			creatingMockErr,
		)
	}

	query := "SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((email = ?))"

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(errors.New("test"))

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	mockedDbConn, connErr := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if connErr != nil {
		t.Fatalf(
			"an error '%s' occurred during opening a stub database connection",
			connErr,
		)
	}

	want := model.User{}

	got, err := NewUsersRepository(mockedDbConn).FindUserByEmail("test@email.com")

	assert.Error(t, err)
	assert.Equal(t, want, got)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindUserById(t *testing.T) {
	_, mock, creatingMockErr := sqlmock.New()

	if creatingMockErr != nil {
		t.Fatalf(
			"an error '%s' occurred during creating a stub database connection",
			creatingMockErr,
		)
	}

	userID := uint(1)
	userEmail := "test@email.com"
	userPassword := "test pass"
	userFullName := "test full name"

	rows := mock.NewRows([]string{
		"id",
		"created_at",
		"updated_at",
		"deleted_at",
		"email",
		"password",
		"full_name",
	}).
		AddRow(
			userID,
			testTime,
			testTime,
			nil,
			userEmail,
			userPassword,
			userFullName,
		)

	query := "SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL " +
		"AND ((id = ?)) ORDER BY `users`.`id` ASC LIMIT 1  "

	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(userID).
		WillReturnRows(rows)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	mockedDbConn, connErr := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if connErr != nil {
		t.Fatalf(
			"an error '%s' occurred during opening a stub database connection",
			connErr,
		)
	}

	want := model.User{
		Model: gorm.Model{
			ID:        userID,
			CreatedAt: testTime,
			UpdatedAt: testTime,
			DeletedAt: gorm.DeletedAt{},
		},
		Email:    userEmail,
		Password: userPassword,
		FullName: userFullName,
	}

	got := NewUsersRepository(mockedDbConn).FindUserByID(int(userID))

	assert.Equal(t, want, got)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStoreUser(t *testing.T) {
	_, mock, creatingMockErr := sqlmock.New()

	if creatingMockErr != nil {
		t.Fatalf(
			"an error '%s' occurred during creating a stub database connection",
			creatingMockErr,
		)
	}

	userID := uint(1)
	userEmail := "test@email.com"
	userPassword := "test pass"
	userFullName := "test full name"

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).
		WithArgs(
			userID,
			testTime,
			testTime,
			nil,
			userEmail,
			userPassword,
			userFullName,
		).
		WillReturnResult(sqlmock.NewResult(int64(userID), 1))

	mock.ExpectCommit()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	mockedDbConn, connErr := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if connErr != nil {
		t.Fatalf(
			"an error '%s' occurred during opening a stub database connection",
			connErr,
		)
	}

	savingErr := NewUsersRepository(mockedDbConn).StoreUser(model.User{
		Model: gorm.Model{
			ID:        userID,
			CreatedAt: testTime,
			UpdatedAt: testTime,
			DeletedAt: gorm.DeletedAt{},
		},
		Email:    userEmail,
		Password: userPassword,
		FullName: userFullName,
	})

	assert.NoError(t, savingErr)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStoreUserReturnsError(t *testing.T) {
	_, mock, creatingMockErr := sqlmock.New()

	if creatingMockErr != nil {
		t.Fatalf(
			"an error '%s' occurred during creating a stub database connection",
			creatingMockErr,
		)
	}

	userID := uint(1)
	userEmail := "test@email.com"
	userPassword := "test pass"
	userFullName := "test full name"

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).
		WithArgs(
			userID,
			testTime,
			testTime,
			nil,
			userEmail,
			userPassword,
			userFullName,
		).
		WillReturnError(errors.New("test"))

	mock.ExpectRollback()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	mockedDbConn, connErr := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if connErr != nil {
		t.Fatalf(
			"an error '%s' occurred during opening a stub database connection",
			connErr,
		)
	}

	savingErr := NewUsersRepository(mockedDbConn).StoreUser(model.User{
		Model: gorm.Model{
			ID:        userID,
			CreatedAt: testTime,
			UpdatedAt: testTime,
			DeletedAt: gorm.DeletedAt{},
		},
		Email:    userEmail,
		Password: userPassword,
		FullName: userFullName,
	})

	assert.Error(t, savingErr)

	assert.NoError(t, mock.ExpectationsWereMet())
}
