package bootstrap

import (
	"qibla-backend/usecase"
	"time"

	"github.com/robfig/cron/v3"
)

// RegisterCronjob ...
func (boot *Bootstrap) RegisterCronjob() {
	location, err := time.LoadLocation(usecase.DefaultLocation)
	if err != nil {
		panic(err)
	}

	cronjobUc := usecase.CronjobUseCase{UcContract: &boot.UseCaseContract}

	c := cron.New(cron.WithLocation(location))

	// Test
	// c.AddFunc("* * * * *", cronjobUc.Test)

	c.AddFunc("* * * * *", cronjobUc.DisbursementMutation)

	c.Run()
}
