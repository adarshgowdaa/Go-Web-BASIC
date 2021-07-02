package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

const Path = "file://migration"

func runMigrations(config *AppConfig, logger *logrus.Logger) error  {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	m , err := migrate.New(Path,dsn)
	if err != nil {
		logger.WithError(err).Error("Error Creating Migrations!")
	}

	defer m.Close()

	err = m.Up()

	if errors.Is(err,migrate.ErrNoChange)	{
		logger.Info("No New Migrations!")
		return nil
	}

	if err != nil {
		logger.WithError(err).Error("Error Running Migrations!")
		return err
	}

	logger.Info("Successfully Executed Migrations!")
	return nil
}
