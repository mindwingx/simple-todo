package orm

import "gorm.io/gorm"

type (
	ISql interface {
		ISqlGeneric
		ISqlTx
	}

	ISqlGeneric interface {
		Init()
		C() *gorm.DB
		Migrate(path string)
		Seed()
		Stop()
	}

	ISqlTx interface {
		Begin()
		Commit() error
		Rollback() error
		// Resolve commit or rollback transaction by getting the error
		Resolve(dbErr error) error
	}
)
