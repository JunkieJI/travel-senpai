package store

func (s *store) GetDestinationByID(destinationID int64) (*Destination, error) {
	var destination Destination
	err := s.db.Get(&destination, GetDestinationByID, destinationID)
	if err != nil {
		return nil, err
	}
	return &destination, nil
}

func (s *store) GetDestinations() (*[]Destination, error) {
	var destinations []Destination
	err := s.db.Select(&destinations, GetDestinations)
	if err != nil {
		return nil, err
	}
	return &destinations, nil
}
