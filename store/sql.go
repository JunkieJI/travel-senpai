package store

// GetDestinationByID sql to get a destination by id
const GetDestinationByID = `SELECT id FROM destination WHERE id = ?`

// GetDestinations sql to get all destinations
const GetDestinations = `SELECT id FROM destination`
