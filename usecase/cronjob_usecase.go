package usecase

import (
	"fmt"
	"time"

	timepkg "qibla-backend/helpers/time"
)

type CronjobUseCase struct {
	*UcContract
}

func (uc CronjobUseCase) Test() {
	now := time.Now().UTC()
	date := timepkg.InFormatNoErr(now, DefaultLocation, "2006-01-02")

	fmt.Println(date)
}
