package store

import (
	"context"
)

func (s *store) GetDestinationByID(ctx context.Context, destinationID int64) (*Destination, error) {
	var destination Destination
	err := s.db.Get(&destination, GetDestinationByID, destinationID)
	if err != nil {
		return nil, err
	}
	return &destination, nil
}

func (s *store) GetDestinations(ctx context.Context) (*[]Destination, error) {
	var destinations []Destination
	err := s.db.Select(&destinations, GetDestinations)
	if err != nil {
		return nil, err
	}
	return &destinations, nil
}

func (s *store) CreateDestination(ctx context.Context, dest *Destination) (*Destination, error) {
	res, err := s.db.Exec(CreateDestination, dest.Budget, dest.Country, dest.City, dest.Currency)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return s.GetDestinationByID(ctx, id)
}

func (s *store) UpdateDestination(ctx context.Context, dest *Destination) (*Destination, error) {
	_, err := s.db.Exec(UpdateDestination, dest.Budget, dest.Country, dest.City, dest.Currency, dest.ID)
	if err != nil {
		return nil, err
	}

	return s.GetDestinationByID(ctx, dest.ID)
}

func (s *store) DeleteDestination(ctx context.Context, id int64) error {
	_, err := s.db.Exec(DeleteDestination, id)
	if err != nil {
		return err
	}

	return nil
}
