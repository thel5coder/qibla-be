package usecase

// FcmUseCase ...
type FcmUseCase struct {
	*UcContract
}

// Send ...
func (uc FcmUseCase) Send(to []string, title, body string, data map[string]interface{}) (res string, err error) {
	var recipients []string
	for _, t := range to {
		if t != "" {
			recipients = append(recipients, t)
		}
	}

	res, err = uc.Fcm.SendAndroid(recipients, title, body, data)
	if err != nil {
		return res, err
	}

	return res, err
}
