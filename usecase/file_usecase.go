package usecase

import (
	"errors"
	"mime/multipart"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/messages"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type FileUseCase struct {
	*UcContract
}

func (uc FileUseCase) ReadBy(column, value string) (res viewmodel.FileVm, err error) {
	repository := actions.NewFileRepository(uc.DB)
	count, err := uc.CountBy(column, value)
	if err != nil {
		return res, errors.New(messages.DataNotFound)
	}
	if count > 0 {
		file, err := repository.ReadBy(column, value)
		if err != nil {
			return res, errors.New(messages.DataNotFound)
		}
		res = viewmodel.FileVm{
			ID:        file.ID,
			Name:      file.Name.String,
			Path:      file.Path.String,
			CreatedAt: file.CreatedAt,
			UpdatedAt: file.UpdatedAt,
			DeletedAt: file.DeletedAt.String,
		}
	}

	return res, err
}

func (uc FileUseCase) ReadByPk(ID string) (res viewmodel.FileVm, err error) {
	file, err := uc.ReadBy("id", ID)
	if err != nil {
		return res, err
	}

	urlFile,err := uc.AWSS3.GetURL(file.Path)
	if err != nil {
		return res,err
	}
	file.Path = urlFile

	return file, err
}

func (uc FileUseCase) CountBy(column, value string) (res int, err error) {
	repository := actions.NewFileRepository(uc.DB)
	res, err = repository.CountBy(column, value)

	return res, err
}

func (uc FileUseCase) Add(fileName, path string) (res viewmodel.FileVm, err error) {
	repository := actions.NewFileRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	body := viewmodel.FileVm{
		Name:      fileName,
		Path:      path,
		CreatedAt: now,
		UpdatedAt: now,
	}
	body.ID, err = repository.Add(body)
	if err != nil {
		return res, err
	}

	return body, err
}

func (uc FileUseCase) UploadFile(fileHeader *multipart.FileHeader) (fileVm viewmodel.FileVm, err error) {
	path, fileName, err := uc.AWSS3.UploadManager(fileHeader)
	if err != nil {
		return fileVm, err
	}

	fileVm, err = uc.Add(fileName, path)
	urlFile,err := uc.AWSS3.GetURL(fileVm.Path)
	if err != nil {
		return fileVm,err
	}
	fileVm.Path = urlFile

	return fileVm, err
}
