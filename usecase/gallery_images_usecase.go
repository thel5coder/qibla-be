package usecase

import (
	"database/sql"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type GalleryImagesUseCase struct {
	*UcContract
}

func (uc GalleryImagesUseCase) Browse(galleryID string) (res []viewmodel.GalleryImagesVm,err error){
	repository := actions.NewGalleryImageRepository(uc.DB)
	galleryImages,err := repository.Browse(galleryID)
	if err != nil {
		return res,err
	}

	for _, galleryImage := range galleryImages{
		res = append(res,viewmodel.GalleryImagesVm{
			ID:     galleryImage.ID,
			FileID: galleryImage.FileID,
			Path:   galleryImage.Path,
		})
	}

	return res,err
}

func (uc GalleryImagesUseCase) ReadBy(column,value string) (res viewmodel.GalleryImagesVm,err error){
	repository := actions.NewGalleryImageRepository(uc.DB)
	galleryImage,err := repository.ReadBy(column,value)
	if err != nil {
		return res,err
	}

	res = viewmodel.GalleryImagesVm{
		ID:     galleryImage.ID,
		FileID: galleryImage.FileID,
		Path:   galleryImage.Path,
	}

	return res,err
}

func (uc GalleryImagesUseCase) edit(ID,fileID,updatedAt string,tx *sql.Tx) (err error){
	repository := actions.NewGalleryImageRepository(uc.DB)
	err = repository.Edit(ID,fileID,updatedAt,tx)

	return err
}

func (uc GalleryImagesUseCase) add(galleryID,fileID,createdAt,updatedAt string,tx *sql.Tx) (err error){
	repository := actions.NewGalleryImageRepository(uc.DB)
	err = repository.Add(galleryID,fileID,createdAt,updatedAt,tx)

	return err
}

func (uc GalleryImagesUseCase) delete(ID ,updatedAt,deletedAt string,tx *sql.Tx) (err error){
	repository := actions.NewGalleryImageRepository(uc.DB)
	err = repository.Delete(ID,updatedAt,deletedAt,tx)

	return err
}

func (uc GalleryImagesUseCase) deleteByGalleryID(galleryID,updatedAt,deletedAt string, tx *sql.Tx) (err error){
	repository := actions.NewGalleryImageRepository(uc.DB)
	err = repository.DeleteByGalleryID(galleryID,updatedAt,deletedAt,tx)

	return err
}

func (uc GalleryImagesUseCase) CountBy(ID,galleryID,column,value string) (res int,err error){
	repository := actions.NewGalleryImageRepository(uc.DB)
	res,err = repository.CountBy(ID,galleryID,column,value)

	return res,err
}

func (uc GalleryImagesUseCase) Store(galleryID string,fileIDs []string,tx *sql.Tx) (err error){
	now := time.Now().UTC().Format(time.RFC3339)
	err = uc.deleteByGalleryID(galleryID,now,now,tx)
	if err != nil {
		return err
	}

	for _, fileID := range fileIDs{
		err = uc.add(galleryID,fileID,now,now,tx)
		if err != nil {
			return err
		}
	}

	return err
}
