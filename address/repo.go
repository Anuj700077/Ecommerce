package address

import "Ecommerce/database"

func CreateAddress(addr *Address) error {

	query := `
	INSERT INTO addresses (user_id, address1 , pincode, city, country, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	return database.DB.Exec(query,
		addr.UserID,
		addr.Address1,
		addr.Pincode,
		addr.City,
		addr.Country,
		addr.CreatedAt,
	).Error
}
