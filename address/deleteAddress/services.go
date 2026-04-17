package address

import (
	"Ecommerce/address"
	"errors"
)

func DeleteAddressService(id uint, userID uint) error {

	if id == 0 {
		return errors.New("invalid address id")
	}

	return address.DeleteAddress(id, userID)
}
