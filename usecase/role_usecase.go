package usecase

import (
	"errors"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/skilld-labs/go-odoo"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type RoleUseCase struct {
	*UcContract
}

func (uc RoleUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.RoleVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewRoleRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	roles, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, role := range roles {
		res = append(res, viewmodel.RoleVm{
			ID:        role.ID,
			Name:      role.Name,
			Slug:      role.Slug,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
			DeletedAt: role.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc RoleUseCase) ReadBy(column, value string) (res viewmodel.RoleVm, err error) {
	repository := actions.NewRoleRepository(uc.DB)
	role, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.RoleVm{
		ID:        role.ID,
		Name:      role.Name,
		Slug:      role.Slug,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: role.DeletedAt.String,
	}

	return res, err
}

func (uc RoleUseCase) ReadByPk(ID string) (res viewmodel.RoleVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		fmt.Println(err)
		return res, errors.New(messages.DataNotFound)
	}

	return res, err
}

func (uc RoleUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewRoleRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc RoleUseCase) Edit(ID string, input *requests.RoleRequest) (err error) {
	repository := actions.NewRoleRepository(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsNameExist(ID, input.Name)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.RoleVm{
		ID:        ID,
		Name:      input.Name,
		Slug:      slug.Make(input.Name),
		UpdatedAt: now.Format(time.RFC3339),
	}
	_, err = repository.Edit(body)

	return err
}

func (uc RoleUseCase) Add(input *requests.RoleRequest) (error error) {
	repository := actions.NewRoleRepository(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsNameExist("", input.Name)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.RoleVm{
		Name:      input.Name,
		Slug:      slug.Make(input.Name),
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}
	_, err = repository.Add(body)

	return err
}

func (uc RoleUseCase) Delete(ID string) (err error) {
	repository := actions.NewRoleRepository(uc.DB)
	now := time.Now().UTC()

	count, err := uc.CountByPk(ID)
	if err != nil {
		return errors.New(messages.DataNotFound)
	}

	if count > 0 {
		_, err = repository.Delete(ID, now.Format(time.RFC3339), now.Format(time.RFC3339))
	}

	return err
}

func (uc RoleUseCase) CountByPk(ID string) (res int, err error) {
	repository := actions.NewRoleRepository(uc.DB)
	res, err = repository.CountByPk(ID)

	return res, err
}

func (uc RoleUseCase) IsNameExist(ID, name string) (res bool, err error) {
	count, err := uc.CountBy(ID, "name", name)
	if err != nil {
		return res, err
	}

	return count > 0, err
}

func (uc RoleUseCase) GetRes() (res map[string]interface{}, err error) {
	res = make(map[string]interface{})
	err = uc.Read("travel.package", odoo.NewCriteria().Add("display_name", "like", "paket"), odoo.NewOptions(), &res)
	if err != nil {
		return res, err
	}

	return res, err
}
