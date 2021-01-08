package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) contracts.IUserRepository {
	userSelectParams = []interface{}{}
	userUpdateStatement = ``
	userUpdateParams = []interface{}{}

	return &UserRepository{DB: DB}
}

const (
	selectUser = `select u."id",u."username",u."email",u."password",u."is_active",u."created_at",u."updated_at",u."odo_user_id",u."name",
                  u."mobile_phone",u."pin",u."is_admin_panel",u."fcm_device_token",u.is_fingerprint_set,r."id",r."slug",r."name",f."id",f."path",f."name",
                  array_to_string(array_agg(mu."id" || ':' || mu."menu_id"),','),array_to_string(array_agg(mup."menu_id" || ':' || mup."menu_permission_id"),',')`
	joinUser = `inner join "roles" r on r."id"=u."role_id" and r."deleted_at" is null
                left join "files" f on f."id"=u."profile_picture" and f."deleted_at" is null
                left join "menu_users" mu on mu."user_id"=u."id"
                left join "menu_user_permissions" mup on mup."menu_id"=mu."menu_id"`
	groupByUser = `group by u."id",r."id",f."id"`
)

var (
	whereStatementUser  = `where (lower(u."username") like $1 or lower(u."email") like $1) and u."deleted_at" is null and u."is_admin_panel"=$2`
	userSelectParams    = []interface{}{}
	userUpdateStatement = ``
	userUpdateParams    = []interface{}{}
)

func (repository UserRepository) scanRow(row *sql.Row) (res models.User, err error) {
	err = row.Scan(&res.ID, &res.UserName, &res.Email, &res.Password, &res.IsActive, &res.CreatedAt, &res.UpdatedAt, &res.OdooUserID, &res.Name, &res.MobilePhone, &res.PIN,
		&res.IsAdminPanel, &res.FcmDeviceToken, &res.IsFingerprintSet, &res.RoleModel.ID, &res.RoleModel.Slug, &res.RoleModel.Name, &res.ProfilePictureID, &res.FileModel.Path, &res.FileModel.Name,
		&res.MenuUser, &res.MenuPermissionUser)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository UserRepository) scanRows(rows *sql.Rows) (res models.User, err error) {
	err = rows.Scan(&res.ID, &res.UserName, &res.Email, &res.Password, &res.IsActive, &res.CreatedAt, &res.UpdatedAt, &res.OdooUserID, &res.Name, &res.MobilePhone, &res.PIN,
		&res.IsAdminPanel, &res.FcmDeviceToken, &res.IsFingerprintSet, &res.RoleModel.ID, &res.RoleModel.Slug, &res.RoleModel.Name, &res.ProfilePictureID, &res.FileModel.Path, &res.FileModel.Name,
		&res.MenuUser, &res.MenuPermissionUser)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository UserRepository) Browse(isAdminPanel bool, search, order, sort string, limit, offset int) (data []models.User, count int, err error) {
	userSelectParams = append(userSelectParams, []interface{}{"%" + strings.ToLower(search) + "%", isAdminPanel, limit, offset}...)
	statement := selectUser + ` from "users" u ` + joinUser + ` ` + whereStatementUser + ` ` + groupByUser + ` order by u.` + order + ` ` + sort + ` limit $3 offset $4`
	rows, err := repository.DB.Query(statement, userSelectParams...)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, count, err
		}
		data = append(data, temp)
	}

	statement = `select distinct count(u."id") from "users" u ` + joinUser + ` ` + whereStatementUser + ` ` + groupByUser
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%", isAdminPanel).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository UserRepository) ReadBy(column, value string) (data models.User, err error) {
	statement := selectUser + ` from "users" u ` + joinUser + ` where  ` + column + `=$1 and u."deleted_at" is null ` + groupByUser
	row := repository.DB.QueryRow(statement, value)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository UserRepository) Edit(input viewmodel.UserVm, password string, tx *sql.Tx) (err error) {
	userUpdateParams = append(userUpdateParams, []interface{}{input.UserName, input.Email, input.IsActive, input.Role.ID, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.Name, input.File.ID, input.MobilePhone, input.IsAdminPanel, input.ID}...)
	userUpdateStatement = `set "username"=$1, "email"=$2, "is_active"=$3, "role_id"=$4, "updated_at"=$5, "name"=$6, "profile_picture"=$7, "mobile_phone"=$8, "is_admin_panel"=$9 where "id"=$10`
	if password != "" {
		userUpdateParams = append(userUpdateParams, []interface{}{password}...)
		userUpdateStatement = `set "username"=$1, "email"=$2, "is_active"=$3, "role_id"=$4, "updated_at"=$5, "name"=$6, "profile_picture"=$7, "mobile_phone"=$8, "is_admin_panel"=$9, 
                              "password"=$11 where "id"=$10`
	}
	statement := `update "users" ` + userUpdateStatement
	_, err = tx.Exec(statement, userUpdateParams...)
	if err != nil {
		return err
	}

	return nil
}

func (repository UserRepository) EditPIN(ID, pin, updatedAt string) (res string, err error) {
	statement := `update "users" set "pin"=$1, "updated_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, pin, datetime.StrParseToTime(updatedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository UserRepository) EditFingerPrintStatus(ID, updatedAt string, status bool) (res string, err error) {
	statement := `update "users" set "is_fingerprint_set"=$1, "updated_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, status, datetime.StrParseToTime(updatedAt, time.RFC3339), ID).Scan(&res)

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

func (repository UserRepository) EditIsActiveStatus(email, updatedAt string, status bool) (res string, err error) {
	statement := `update "users" set "is_active"=$1, "updated_at"=$2 where "email"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, status, datetime.StrParseToTime(updatedAt, time.RFC3339), email).Scan(&res)

	return res, err
}

func (repository UserRepository) Add(input viewmodel.UserVm, password string, tx *sql.Tx) (res string, err error) {
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
	if ID == "" {
		statement := `select count("id") from "users" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "users" where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}

func (repository UserRepository) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "users" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)

	return res, err
}
