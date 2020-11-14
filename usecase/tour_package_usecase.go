package usecase

import "qibla-backend/usecase/viewmodel"

type TourPackageUseCase struct {
	*UcContract
}

func (uc TourPackageUseCase) BrowseByPartner(partnerID string) (res []viewmodel.TourPackageVm,err error){

}