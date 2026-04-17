package address

import (
	"Ecommerce/database"
	"errors"
)

func CreateAddress(addr *Address) error {

	query := `
	INSERT INTO addresses (user_id, address1, pincode, city, country)
	VALUES ($1, $2, $3, $4, $5)
	`

	return database.DB.Exec(query,
		addr.UserID,
		addr.Address1,
		addr.Pincode,
		addr.City,
		addr.Country,
	).Error
}

func GetAddressByUserID(userID uint) ([]Address, error) {
	var addresses []Address

	query := `
	SELECT id, user_id, address1, pincode, city, country, created_at
	FROM addresses
	WHERE user_id =$1
	ORDER BY created_at DESC
	`

	rows, err := database.SQLDB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var addr Address
		err := rows.Scan(
			&addr.ID, &addr.UserID, &addr.Address1, &addr.Pincode, &addr.City,
			&addr.Country, &addr.CreatedAt,
		)
		if err != nil {
			return nil, err

		}
		addresses = append(addresses, addr)

	}
	return addresses, nil
}

func UpdateAddress(addr *Address) error {

	query := `
	UPDATE addresses
	SET address1=$1, pincode=$2, city=$3, country=$4
	WHERE id=$5 AND user_id=$6
	`

	result := database.DB.Exec(
		query,
		addr.Address1,
		addr.Pincode,
		addr.City,
		addr.Country,
		addr.ID,
		addr.UserID,
	)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("address not found or unauthorized")
	}

	return nil
}



func DeleteAddress(id uint, userID uint) error {

	query := `
	DELETE FROM addresses
	WHERE id=$1 AND user_id=$2
	`

	result := database.DB.Exec(query, id, userID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("address not found or unauthorized")
	}

	return nil
}


