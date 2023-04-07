package getter

func (cg *Cards_Getter_Implementation) GetThisRandomCards(requestData []byte, requestAttributes map[string]string) (requestedCards interface{}, err error) {

	dataToStr := string(requestData)

	return dataToStr, nil

}
