package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/messages"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type FaqListUseCase struct {
	*UcContract
}

func (uc FaqListUseCase) Browse(faqID string) (res []viewmodel.FaqListVm,err error){
	repository := actions.NewFaqListRepository(uc.DB)
	faqLists,err := repository.Browse(faqID)
	if err != nil {
		return res,err
	}

	for _,faqList := range faqLists{
		res = append(res,viewmodel.FaqListVm{
			ID:       faqList.ID,
			Question: faqList.Question,
			Answer:   faqList.Answer,
			CreatedAt: faqList.CreatedAt,
			UpdatedAt: faqList.UpdatedAt,
			DeletedAt: faqList.DeletedAt.String,
		})
	}

	return res,err
}

func (uc FaqListUseCase) add(faqID,question,answer string, tx *sql.Tx) (err error){
	repository := actions.NewFaqListRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	err = repository.Add(faqID,question,answer,now,now,tx)

	return err
}

func (uc FaqListUseCase) deleteByFaqID(faqID,updatedAt,deletedAt string,tx *sql.Tx)(err error){
	repository := actions.NewFaqListRepository(uc.DB)
	err = repository.DeleteByFaqID(faqID,updatedAt,deletedAt,tx)

	return err
}

func (uc FaqListUseCase) Delete(ID string) (err error){
	repository := actions.NewFaqListRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	_,err = repository.Delete(ID,now,now)
	if err != nil {
		fmt.Print(err)
		return errors.New(messages.DataNotFound)
	}

	return nil
}

func (uc FaqListUseCase) Store(faqID string,faqLists []viewmodel.FaqListVm,tx *sql.Tx) (err error){
	now := time.Now().UTC().Format(time.RFC3339)
	err = uc.deleteByFaqID(faqID,now,now,tx)
	if err != nil {
		return err
	}

	for _,faqList := range faqLists{
		err = uc.add(faqID,faqList.Question,faqList.Answer,tx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc FaqListUseCase) CountBy(column,value string) (res int,err error){
	repository := actions.NewFaqListRepository(uc.DB)
	res,err = repository.CountBy(column,value)

	return res,err
}