package usecase

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/messages"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type FileUseCase struct {
	*UcContract
}

func (uc FileUseCase) ReadBy(column,value string) (res viewmodel.FileVm,err error){
	repository := actions.NewFileRepository(uc.DB)
	count,err := uc.CountBy(column,value)
	if err != nil {
		return res,errors.New(messages.DataNotFound)
	}
	if count > 0 {
		file,err := repository.ReadBy(column,value)
		if err != nil {
			return res,errors.New(messages.DataNotFound)
		}
		res = viewmodel.FileVm{
			ID:        file.ID,
			Name:      file.Name,
			Path:      file.Path,
			CreatedAt: file.CreatedAt,
			UpdatedAt: file.UpdatedAt,
			DeletedAt: file.DeletedAt.String,
		}
	}

	return res,err
}

func (uc FileUseCase) ReadByPk(ID string) (res string,err error){
	file,err := uc.ReadBy("id",ID)
	if err != nil {
		return res,err
	}

	return file.Path,err
}

func (uc FileUseCase) CountBy(column,value string) (res int,err error){
	repository := actions.NewFileRepository(uc.DB)
	res,err = repository.CountBy(column,value)

	return res,err
}

func (uc FileUseCase) Add(file *multipart.FileHeader) (res viewmodel.FileVm,err error){
	repository := actions.NewFileRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	src,err := file.Open()
	if err != nil {
		return res,err
	}
	//defer src.Close()

	//size := file.Size
	//buffer := make([]byte, size)
	//src.Read(buffer)
	//fileName := bson.NewObjectId().Hex() + filepath.Ext(file.Filename)
	path := StaticBaseUrl+"/"+ file.Filename
	dst,err := os.Create("../server/statics/"+file.Filename)
	if err != nil {
		return res,err
	}
	defer dst.Close()

	_,err = io.Copy(dst,src)
	if err != nil {
		return res,err
	}

	body := viewmodel.FileVm{
		Name:      file.Filename,
		Path:      path,
		CreatedAt: now,
		UpdatedAt: now,
	}
	body.ID,err = repository.Add(body)
	if err != nil {
		return res,err
	}

	return body,err
}
