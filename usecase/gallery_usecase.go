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

type GalleryUseCase struct {
	*UcContract
}

func (uc GalleryUseCase) Browse(search,order,sort string,page, limit int) (res []viewmodel.GalleryVm,pagination viewmodel.PaginationVm,err error){
	repository := actions.NewGalleryRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	galleries, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _,gallery := range galleries{
		res = append(res,viewmodel.GalleryVm{
			ID:          gallery.ID,
			GalleryName: gallery.GalleryName,
			CreatedAt: gallery.CreatedAt,
			UpdatedAt: gallery.UpdatedAt,
			DeletedAt: gallery.DeletedAt.String,
		})
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc GalleryUseCase) readBy(ID,column,value string) (res viewmodel.GalleryDetailVm,err error){
	repository := actions.NewGalleryRepository(uc.DB)
	var galleryImages []viewmodel.GalleryImagesVm
	gallery,err := repository.ReadBy(column,value)
	if err != nil {
		return res,err
	}

	if ID != ""{
		galleryImageUc := GalleryImagesUseCase{UcContract:uc.UcContract}
		galleryImages,err = galleryImageUc.Browse(ID)
		if err != nil {
			return res,err
		}
	}

	res = viewmodel.GalleryDetailVm{
		ID:          gallery.ID,
		GalleryName: gallery.GalleryName,
		CreatedAt:   gallery.CreatedAt,
		UpdatedAt:   gallery.UpdatedAt,
		DeletedAt:   gallery.DeletedAt.String,
		Images:      galleryImages,
	}

	return res,err
}

func (uc GalleryUseCase) ReadByPk(ID string) (res viewmodel.GalleryDetailVm,err error){
	res,err = uc.readBy(ID,"id",ID)
	if err != nil {
		return res,err
	}

	return res,err
}

func (uc GalleryUseCase) Edit(ID string, input *requests.GalleryRequest) (err error){
	repository := actions.NewGalleryRepository(uc.DB)
	galleryImageUc := GalleryImagesUseCase{UcContract:uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)

	isExist,err := uc.isExist(ID,input.GalleryName)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	transaction,err := uc.DB.Begin()
	if err != nil {
		fmt.Println(3)
		transaction.Rollback()
		return err
	}

	err = galleryImageUc.Store(ID,input.ImagesID,transaction)
	if err != nil {
		fmt.Println(4)
		transaction.Rollback()

		return err
	}

	err = repository.Edit(ID,input.GalleryName,now,transaction)
	if err != nil {
		fmt.Println(5)
		transaction.Rollback()

		return err
	}
	transaction.Commit()

	return nil
}

func (uc GalleryUseCase) Add(input *requests.GalleryRequest) (err error){
	repository := actions.NewGalleryRepository(uc.DB)
	galleryImageUc := GalleryImagesUseCase{uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)

	isExist,err := uc.isExist("",input.GalleryName)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	transaction,err := uc.DB.Begin()
	if err != nil {
		transaction.Rollback()
		return err
	}

	galleryID,err := repository.Add(input.WebContentCategoryID,input.GalleryName,now,now,transaction)
	if err != nil {
		transaction.Rollback()

		return err
	}

	err = galleryImageUc.Store(galleryID,input.ImagesID,transaction)
	if err != nil {
		transaction.Rollback()

		return err
	}
	transaction.Commit()

	return nil
}

func (uc GalleryUseCase) Delete(ID string) (err error){
	repository := actions.NewGalleryRepository(uc.DB)
	galleryImageUc := GalleryImagesUseCase{UcContract:uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)

	count,err := uc.countBy("","id",ID)
	if err != nil {
		return errors.New(messages.DataNotFound)
	}

	if count > 0 {
		transaction,err := uc.DB.Begin()
		if err != nil {
			transaction.Rollback()

			return err
		}

		err = galleryImageUc.deleteByGalleryID(ID,now,now,transaction)
		if err != nil {
			transaction.Rollback()

			return err
		}

		err = repository.Delete(ID,now,now,transaction)
		if err != nil {
			transaction.Rollback()

			return err
		}
		transaction.Commit()
	}

	return err
}

func (uc GalleryUseCase) isExist(ID,name string) (res bool,err error){
	count,err := uc.countBy(ID,"gallery_name",name)
	if err != nil {
		return res,err
	}

	return count > 0,err
}

func (uc GalleryUseCase) countBy(ID,column,value string) (res int,err error){
	repository := actions.NewGalleryRepository(uc.DB)
	res,err = repository.CountBy(ID,column,value)
	if err != nil {
		return res,err
	}

	return res,err
}
