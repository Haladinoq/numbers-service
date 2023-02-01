package data

import (
	"github.com/golobby/container/v3"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/repo"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/sql"
	"gorm.io/gorm"
)

//RunConfig ..
func RunConfig(container *container.Container) {
	container.Singleton(func(db *gorm.DB) repo.INumbersRepo {
		return sql.NewNumbers(db)
	})
}
