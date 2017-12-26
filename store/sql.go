package store

// GetDestinationByID sql to get a destination by id
const GetDestinationByID = `SELECT 
	id,
	budget,
	city,
	country,
	currency
	FROM destination WHERE id = ?`

// GetDestinations sql to get all destinations
const GetDestinations = `SELECT 
	id,
	budget,
	city,
	country,
	currency 
	FROM destination`

// AddDestination sql to add a destination
const AddDestination = `INSERT INTO destination ('budget', 'country', 'city', 'currency') Values (?, ?, ?, ?)`

// UpdateDestination sql to update a destination
const UpdateDestination = `UPDATE destination set budget = ?, country = ?, city = ?, currency = ? WHERE id = ?`

// DeleteDestination sql to delete a destination
const DeleteDestination = `DELETE FROM destination WHERE id = ?`
