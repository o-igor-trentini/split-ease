package sepostgresmigration

import (
	"gorm.io/gorm"
	"reflect"
	"se-backend/config/selog"
	"se-backend/model/repository/entity"
)

type migrationRun struct {
	Name string
	Run  func(*gorm.DB) error
}

// ExecMigrations realiza o versionamento do banco de dados.
func ExecMigrations(db *gorm.DB) {
	var migrations []interface{}

	tx := db.Begin()

	for _, migration := range migrations {
		f := reflect.ValueOf(migration)
		result := f.Call(nil)
		function := result[0].Interface().(migrationRun)

		current := entity.Migration{
			Name: function.Name,
		}

		if err := db.Where("mi_name", current.Name).Find(&current).Error; err != nil {
			tx.Rollback()
			selog.Info("[MIGRATION] Erro ao executar migration ["+current.Name+"]", selog.MigrationTag)
			return
		}

		if current.ID > 0 {
			selog.Info("[MIGRATION] " + current.Name + " já executada")
			continue
		}

		if err := function.Run(db); err != nil {
			tx.Rollback()
			selog.Fatal("Erro ao executar migration ["+current.Name+"]", err, selog.MigrationTag)
			return
		}

		db.Create(&current)
	}

	tx.Commit()
}
