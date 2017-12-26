package store

// Destination model
type Destination struct {
	ID       int64  `db:"id"`
	Budget   int64  `db:"budget"`
	City     string `db:"city"`
	Country  string `db:"country"`
	Currency string `db:"currency"`
}

// Transport model
type Transport struct {
	ID            int64  `db:"id"`
	Type          string `db:"type"`
	Cost          int64  `db:"cost"`
	Provider      string `db:"privider"`
	DestinationID int64  `db:"destination_id"`
	OutboundID    int64  `db:"outbound_id"`
	ReturnID      *int64 `db:"return_id"`
}

// TransportDetails model
type TransportDetails struct {
	DepartureDate string `db:"departure_date"`
	ArrivalDate   string `db:"arrival_date"`
	Layover       *int64 `db:"layover"`
}
