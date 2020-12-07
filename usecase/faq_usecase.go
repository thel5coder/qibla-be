package usecase

import (
	"errors"
	"fmt"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type FaqUseCase struct {
	*UcContract
}

func (uc FaqUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.FaqVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewFaqRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	faqs, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, faq := range faqs {
		res = append(res, viewmodel.FaqVm{
			ID:              faq.ID,
			FaqCategoryName: faq.FaqCategoryName,
			FaqListID:       faq.FaqListID,
			Question:        faq.Question,
			Answer:          faq.Answer,
			CreatedAt:       faq.CreatedAt,
			UpdatedAt:       faq.UpdatedAt,
			DeletedAt:       faq.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc FaqUseCase) readBy(column,value string) (res []viewmodel.FaqVm,err error){
	repository := actions.NewFaqRepository(uc.DB)
	faqs,err := repository.ReadBy(column,value)

	for _, faq := range faqs {
		res = append(res, viewmodel.FaqVm{
			ID:              faq.ID,
			FaqCategoryName: faq.FaqCategoryName,
			FaqListID:       faq.FaqListID,
			Question:        faq.Question,
			Answer:          faq.Answer,
			CreatedAt:       faq.CreatedAt,
			UpdatedAt:       faq.UpdatedAt,
			DeletedAt:       faq.DeletedAt.String,
		})
	}

	return res,err
}

func (uc FaqUseCase) ReadByPk(ID string) (res []viewmodel.FaqVm, err error) {
	res,err = uc.readBy("fi.faq_id",ID)
	if err != nil {
		return res,err
	}

	return res,err
}

func (uc FaqUseCase) Edit(ID string, input *requests.FaqRequest) (err error) {
	faqListUc := FaqListUseCase{UcContract: uc.UcContract}
	var faqLists []viewmodel.FaqListVm

	transaction, err := uc.DB.Begin()
	if err != nil {
		transaction.Rollback()

		return err
	}

	for _, faqList := range input.FaqLists {
		faqLists = append(faqLists, viewmodel.FaqListVm{
			Question: faqList.Question,
			Answer:   faqList.Answer,
		})
	}

	err = faqListUc.Store(ID, faqLists, transaction)
	if err != nil {
		transaction.Rollback()

		return err
	}
	transaction.Commit()

	return nil
}

func (uc FaqUseCase) Add(input *requests.FaqRequest) (err error) {
	repository := actions.NewFaqRepository(uc.DB)
	faqListUc := FaqListUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)
	var faqLists []viewmodel.FaqListVm
	var faqID string

	isExist, err := uc.isExist("", input.FaqCategoryName)
	if err != nil {
		fmt.Println(1)
		return err
	}

	transaction, err := uc.DB.Begin()
	if err != nil {
		transaction.Rollback()

		return err
	}

	if isExist {
		faqs,err := uc.readBy("f.faq_category_name",input.FaqCategoryName)
		if err != nil {
			fmt.Println(2)
			transaction.Rollback()

			return err
		}

		if len(faqs) > 0 {
			faqID = faqs[0].ID
		}
	}else{
		body := viewmodel.FaqVm{
			FaqCategoryName: input.FaqCategoryName,
			CreatedAt: now,
			UpdatedAt: now,
		}
		faqID,err = repository.Add(body,input.WebContentCategoryID,transaction)
		if err != nil {
			fmt.Println(3)
			transaction.Rollback()

			return err
		}
	}


	for _, faqList := range input.FaqLists {
		faqLists = append(faqLists, viewmodel.FaqListVm{
			Question: faqList.Question,
			Answer:   faqList.Answer,
		})
	}

	err = faqListUc.Store(faqID, faqLists, transaction)
	if err != nil {
		fmt.Println(4)
		transaction.Rollback()

		return err
	}
	transaction.Commit()

	return nil
}

func (uc FaqUseCase) Delete(ID string) (err error){
	faqListUc := FaqListUseCase{UcContract:uc.UcContract}
	err = faqListUc.Delete(ID)
	if err != nil {
		fmt.Print(err)
		return errors.New(messages.DataNotFound)
	}

	return err
}

func (uc FaqUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewFaqRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc FaqUseCase) isExist(ID, faqCategoryName string) (res bool, err error) {
	count, err := uc.CountBy(ID, "faq_category_name", faqCategoryName)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
