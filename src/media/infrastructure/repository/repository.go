package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"sagara-test/src/auth/infrastructure/helper"
	"sagara-test/src/common/infrastructure/db"
	"sagara-test/src/common/utility"
	"sagara-test/src/media/domain/entity"
	"sagara-test/src/media/domain/interfaces"
	"strconv"
	"strings"
)

var (
	URLImages = os.Getenv("URLImages")
)

type MediaRepository struct {
	db *db.ConnectTo
	interfaces.IMediaRepository
}

// MediaRepository
func NewMediaRepository(db *db.ConnectTo) MediaRepository {
	return MediaRepository{db: db}
}

func (a *MediaRepository) GetMedia(ctx context.Context, data entity.StructQuery) (response []entity.ModelMedia, countAll int64, err error) {
	query := `SELECT * FROM Medias WHERE deleted_at IS NULL`

	// Condition Query
	if data.Keys != "" && data.Keyword != "" {
		query += " AND " + data.Keys + " LIKE '%" + data.Keyword + "%' "
	}

	if data.Sort != "" {
		query += " ORDER BY " + data.Sort + " " + data.Order
	}

	// Count All
	queryCount := strings.ReplaceAll(query, "*", "count(*)")
	if err = a.db.PGRead.Get(&countAll, queryCount); err != nil {
		log.Println("Error get data Media: ", err)
	}

	// Limit offset condition for pagination
	if data.Limit == 0 {
		data.Limit = 5
	}
	if data.Page == 0 {
		data.Page = (1 * data.Limit) - data.Limit
	} else {
		data.Page = (data.Page * data.Limit) - data.Limit
	}

	query += " LIMIT " + strconv.Itoa(data.Limit) + " OFFSET " + strconv.Itoa(data.Page)
	err = a.db.PGRead.Select(&response, query)
	if err != nil {
		err = fmt.Errorf("Error get data Media: ", err)
		log.Println(err)
		return
	}

	return

}

func (a *MediaRepository) InsertMedia(ctx context.Context, data entity.ModelMedia) (response entity.ModelMedia, err error) {
	tx := a.db.PGRead.MustBegin()
	// Generate UUID
	data.GUID = utility.GenerateGoogleUUID()
	query := `INSERT INTO Medias `
	var strField strings.Builder
	var strValue strings.Builder

	fieldItem := helper.GetNamedStruct(data)
	for _, v := range fieldItem {
		if v != "remember_token" {
			strField.WriteString(v + ",")
			strValue.WriteString(":" + v + ",")
		}
	}

	query += "(" + strings.TrimSuffix(strField.String(), ",") + ")" + " VALUES(" + strings.TrimSuffix(strValue.String(), ",") + ")"

	_, err = tx.NamedExec(query, data)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}
	tx.Commit()

	response = data
	return
}

func (a *MediaRepository) UpdateMedia(ctx context.Context, data entity.ModelMedia) (response entity.ModelMedia, err error) {
	tx := a.db.PGRead.MustBegin()

	query := `UPDATE Medias SET `
	var strField strings.Builder

	fieldItem := helper.GetNamedStruct(data)
	for _, v := range fieldItem {
		if v == "id" || v == "created_at" || v == "updated_at" || v == "email" || v == "deleted_at" || v == "last_logout_at" || v == "last_login_at" || v == "password" {
			continue
		}

		strField.WriteString(v + "=:" + v + ", ")
	}

	query += strings.TrimSuffix(strField.String(), ", ") + " WHERE id = '" + data.GUID + "'"
	log.Println(query)
	_, err = tx.NamedExec(query, data)
	if err != nil {
		tx.Rollback()
		err = fmt.Errorf("Error update FundraisingDonation: ", err)
		log.Println(err)
		return
	}
	tx.Commit()
	response = data
	return
}

func (a *MediaRepository) DeleteMedia(ctx context.Context, data string) (response entity.ModelMedia, err error) {
	tx := a.db.PGRead.MustBegin()
	// Generate UUID

	queryDelete := `DELETE FROM Medias WHERE Medias_id = '` + data + `'`
	_, err = tx.Exec(queryDelete)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}
	tx.Commit()
	return
}
