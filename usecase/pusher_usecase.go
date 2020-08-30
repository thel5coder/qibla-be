package usecase

type PusherUseCase struct {
	*UcContract
}

func (uc PusherUseCase) Broadcast(eventName string,body interface{}) (err error){
	uc.Pusher.Send(eventName,body)

	return nil
}
