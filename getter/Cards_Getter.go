package getter

type Cards_Getter_Implementation struct {
}

type Cards_Getter_Interface interface {
	GetThisRandomCards(requestData []byte, requestAttributes map[string]string) (requestedCards interface{}, err error)
}

func New() Cards_Getter_Interface {
	return &Cards_Getter_Implementation{}
}
