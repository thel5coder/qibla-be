package actions

import (
	"database/sql"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) contracts.IUserRepository {
	return &UserRepository{DB: DB}
}

const selectUser = `select u."id",u."username",u."email",u."password",u."is_active",u."created_at",u."updated_at",u."odo_user_id",u."name",
                   u."mobile_phone",u."pin",u."is_admin_panel",u."fcm_device_token",r."id",r."slug",r."name",f."id",f."path",f."name" from "users" u 
                   inner join "roles" r on r."id"=u."role_id" and r."deleted_at" is null
                   left join "files" f on f."id"=u."profile_picture"`

func (repository UserRepository) BrowseNonUserAdminPanel(search, order, sort string, limit, offset int) (data []models.User, count int, err error) {
	statement := selectUser + ` where (lower(u."username") like $1 or lower(u."email") like $1) and u."deleted_at" is null and u."is_admin_panel" = false 
                 order by u.` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		fmt.Print(err)
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.User{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.UserName,
			&dataTemp.Email,
			&dataTemp.Password,
			&dataTemp.IsActive,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.OdooUserID,
			&dataTemp.Name,
			&dataTemp.MobilePhone,
			&dataTemp.PIN,
			&dataTemp.IsAdminPanel,
			&dataTemp.FcmDeviceToken,
			&dataTemp.RoleModel.ID,
			&dataTemp.RoleModel.Slug,
			&dataTemp.RoleModel.Name,
			&dataTemp.ProfilePictureID,
			&dataTemp.FileModel.Path,
			&dataTemp.FileModel.Name,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `select count("id") from "users"
                 where (lower("username") like $1 or lower("email") like $1) and "deleted_at" is null and "is_admin_panel" = false`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository UserRepository) BrowseUserAdminPanel(search, order, sort string, limit, offset int) (data []models.User, count int, err error) {
	statement := selectUser + ` where (lower(u."username") like $1 or lower(u."email") like $1 or lower(r."name") like $1) and u."deleted_at" is null and u."is_admin_panel" = true 
                order by u.` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		fmt.Print(err)
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.User{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.UserName,
			&dataTemp.Email,
			&dataTemp.Password,
			&dataTemp.IsActive,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.OdooUserID,
			&dataTemp.Name,
			&dataTemp.MobilePhone,
			&dataTemp.PIN,
			&dataTemp.IsAdminPanel,
			&dataTemp.FcmDeviceToken,
			&dataTemp.RoleModel.ID,
			&dataTemp.RoleModel.Slug,
			&dataTemp.RoleModel.Name,
			&dataTemp.ProfilePictureID,
			&dataTemp.FileModel.Path,
			&dataTemp.FileModel.Name,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `select count("users"."id") from "users"
                 inner join "roles" r on r."id"=users."role_id"
                 where (lower(users."username") like $1 or lower(users."email") like $1 or lower(r."name") like $1) and "users"."deleted_at" is null and "users"."is_admin_panel"=true`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository UserRepository) ReadBy(column, value string) (data models.User, err error) {
	statement := selectUser + ` where ` + column + `=$1 and u."deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.UserName,
		&data.Email,
		&data.Password,
		&data.IsActive,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.OdooUserID,
		&data.Name,
		&data.MobilePhone,
		&data.PIN,
		&data.IsAdminPanel,
		&data.FcmDeviceToken,
		&data.RoleModel.ID,
		&data.RoleModel.Slug,
		&data.RoleModel.Name,
		&data.ProfilePictureID,
		&data.FileModel.Path,
		&data.FileModel.Name,
	)

	return data, err
}

func (repository UserRepository) Edit(input viewmodel.UserVm, password string, tx *sql.Tx) (err error) {
	if password != "" {
		statement := `update "users" 
                     set "username"=$1, "email"=$2, "password"=$3, "is_active"=$4, "role_id"=$5, "updated_at"=$6, "name"=$7, "profile_picture"=$8, "mobile_phone"=$9,"is_admin_panel"=$10
                     where "id"=$11`
		_, err = tx.Exec(
			statement,
			input.UserName,
			input.Email,
			password,
			input.IsActive,
			input.Role.ID,
			datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
			input.Name,
			input.File.ID,
			input.MobilePhone,
			input.IsAdminPanel,
			input.ID,
		)
	} else {
		statement := `update "users" 
                      set "username"=$1, "email"=$2, "is_active"=$3, "role_id"=$4, "updated_at"=$5, "name"=$6, "profile_picture"=$7, "mobile_phone"=$8, "is_admin_panel"=$9 where "id"=$10`
		_, err = tx.Exec(
			statement,
			input.UserName,
			input.Email,
			input.IsActive,
			input.Role.ID,
			datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
			input.Name,
			input.File.ID,
			input.MobilePhone,
			input.IsAdminPanel,
			input.ID,
		)
	}

	return err
}

func (repository UserRepository) EditPIN(ID, pin, updatedAt string) (res string, err error) {
	statement := `update "users" set "pin"=$1, "updated_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, pin, datetime.StrParseToTime(updatedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository UserRepository) EditPassword(ID, password, updatedAt string) (res string, err error) {
	statement := `update "users" set "password"=$1, "updated_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, password, datetime.StrParseToTime(updatedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (UserRepository) EditUserName(ID, userName, updatedAt string, tx *sql.Tx) (err error) {
	statement := `update "users" set "username"=$1, "updated_at"=$2 where "id"=$3 returning "id"`
	_, err = tx.Exec(statement, userName, datetime.StrParseToTime(updatedAt, time.RFC3339), ID)

	return err
}

func (repository UserRepository) EditFcmDeviceToken(ID, deviceToken, updatedAt string) (res string, err error) {
	statement := `update "users" set "fcm_device_token"=$1, "updated_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, deviceToken, datetime.StrParseToTime(updatedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository UserRepository) Add(input viewmodel.UserVm, password string, tx *sql.Tx) (res string, err error) {
	fmt.Print(input.OdooUserID)
	statement := `insert into "users" 
                 ("username","name","profile_picture","email","mobile_phone","password","role_id","odo_user_id","is_active","is_admin_panel","created_at","updated_at") 
                 values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning "id"`
	err = tx.QueryRow(
		statement,
		input.UserName,
		input.Name,
		input.File.ID,
		input.Email,
		input.MobilePhone,
		password,
		input.Role.ID,
		input.OdooUserID,
		input.IsActive,
		input.IsAdminPanel,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

func (repository UserRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "users" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)

	return err
}

func (repository UserRepository) CountBy(ID, column, value string) (res int, err error) {
	fmt.Println(ID)
	if ID == "" {
		statement := `select count("id") from "users" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "users" where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
		fmt.Println(value)
		fmt.Println(statement)
		fmt.Println(res)
	}


	return res, err
}

func (repository UserRepository) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "users" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)

	return res, err
}
