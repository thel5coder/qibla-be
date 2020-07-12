package usecase

import (
	"errors"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/hashing"
	"qibla-backend/helpers/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type UserUseCase struct {
	*UcContract
}

func (uc UserUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.UserVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewUserRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	users, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, user := range users {
		res = append(res, viewmodel.UserVm{
			ID:       user.ID,
			UserName: user.UserName,
			Email:    user.Email,
			IsActive: user.IsActive,
			Role: viewmodel.RoleVm{
				ID:        user.RoleModel.ID,
				Name:      user.RoleModel.Name,
				CreatedAt: user.RoleModel.CreatedAt,
				UpdatedAt: user.RoleModel.UpdatedAt,
				DeletedAt: user.RoleModel.DeletedAt.String,
			},
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc UserUseCase) ReadBy(column, value string) (res viewmodel.UserVm, err error) {
	repository := actions.NewUserRepository(uc.DB)

	user, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.UserVm{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		IsActive: user.IsActive,
		Role: viewmodel.RoleVm{
			ID:        user.RoleModel.ID,
			Name:      user.RoleModel.Name,
			CreatedAt: user.RoleModel.CreatedAt,
			UpdatedAt: user.RoleModel.UpdatedAt,
			DeletedAt: user.RoleModel.DeletedAt.String,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt.String,
	}

	return res, err
}

func (uc UserUseCase) ReadByPk(ID string) (res viewmodel.UserVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		return res, errors.New(messages.DataNotFound)
	}

	return res, err
}

func (uc UserUseCase) Edit(ID string,input *requests.UserRequest) (err error){
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC()
	var password string
	isEmailExist,err := uc.IsEmailExist(ID,input.Email)
	if err != nil {
		return err
	}
	if isEmailExist {
		return errors.New(messages.EmailAlreadyExist)
	}

	isUserNameExist,err := uc.IsUserNameExist(ID,input.UserName)
	if err != nil {
		return err
	}
	if isUserNameExist {
		return errors.New(messages.UserNameExist)
	}

	transaction,err := uc.DB.Begin()
	if err != nil {
		transaction.Rollback()

		return err
	}

	if input.Password != ""{
		password,_ = hashing.HashAndSalt(input.Password)
	}
	body := viewmodel.UserVm{
		ID:        ID,
		UserName:  input.UserName,
		Email:     input.Email,
		IsActive:  input.IsActive,
		Role:      viewmodel.RoleVm{
			ID:        input.RoleID,
		},
		UpdatedAt: now.Format(time.RFC3339),
	}
	err = repository.Edit(body,password,transaction)
	if err != nil {
		transaction.Rollback()

		return err
	}
	transaction.Commit()

	return err
}

func (uc UserUseCase) Add(input *requests.UserRequest) (err error){
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC()

	isEmailExist,err := uc.IsEmailExist("",input.Email)
	if err != nil {
		return err
	}
	if isEmailExist {
		return errors.New(messages.EmailAlreadyExist)
	}

	isUserNameExist,err := uc.IsUserNameExist("",input.UserName)
	if err != nil {
		return err
	}
	if isUserNameExist {
		return errors.New(messages.UserNameExist)
	}

	transaction,err := uc.DB.Begin()
	if err != nil {
		transaction.Rollback()

		return err
	}

	password,_ := hashing.HashAndSalt(input.Password)
	body := viewmodel.UserVm{
		UserName:  input.UserName,
		Email:     input.Email,
		IsActive:  input.IsActive,
		Role:      viewmodel.RoleVm{
			ID:        input.RoleID,
		},
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}
	_,err = repository.Add(body,password,transaction)
	if err != nil {
		transaction.Rollback()

		return err
	}
	transaction.Commit()

	return err
}

func (uc UserUseCase) Delete(ID string) (error error){
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC()

	count,err := uc.CountByPk(ID)
	if err != nil {
		return errors.New(messages.DataNotFound)
	}

	if count > 0 {
		transaction,err := uc.DB.Begin()
		if err != nil {
			transaction.Rollback()

			return err
		}

		err = repository.Delete(ID,now.Format(time.RFC3339),now.Format(time.RFC3339),transaction)
		if err != nil {
			transaction.Rollback()

			return err
		}
		transaction.Commit()
	}

	return err
}

func (uc UserUseCase) CountBy(ID,column,value string) (res int,err error){
	repository := actions.NewUserRepository(uc.DB)
	res,err = repository.CountBy(ID,column,value)

	return res,err
}

func (uc UserUseCase) CountByPk(ID string) (res int,err error){
	repository := actions.NewUserRepository(uc.DB)
	res,err = repository.CountByPk(ID)

	return res,err
}

func (uc UserUseCase) IsUserNameExist(ID,userName string) (res bool,err error){
	count,err := uc.CountBy(ID,"username",userName)
	if err != nil {
		return res,err
	}

	return count > 0, err
}

func (uc UserUseCase) IsEmailExist(ID, email string) (res bool,err error){
	count,err := uc.CountBy(ID,"email",email)
	if err != nil {
		return res,err
	}

	return count > 0, err
}
