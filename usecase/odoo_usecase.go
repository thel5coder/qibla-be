package usecase

import (
	"github.com/skilld-labs/go-odoo"
)

type OdooUseCase struct {
	*UcContract
}

func (uc OdooUseCase) GetField(objectName string) (res map[string]interface{}, err error) {
	var odooOption *odoo.Options
	res, err = uc.Odoo.FieldsGet(objectName, odooOption)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc OdooUseCase) Browse(objectName, search string, limit, offset int, res interface{}) (err error) {
	var odooOption *odoo.Options
	var odoCriteria *odoo.Criteria
	err = uc.Odoo.SearchRead(objectName, odoCriteria, odooOption, res)
	if err != nil {
		return err
	}

	return nil
}

func (uc OdooUseCase) Read(objectName string, id int64, res interface{}) (err error) {
	var odooOption *odoo.Options
	err = uc.Odoo.Read(objectName, []int64{id}, odooOption, res)
	if err != nil {
		return err
	}

	return nil
}
