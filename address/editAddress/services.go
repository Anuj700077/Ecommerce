package address

import (
	"Ecommerce/address"
	"errors"
)

func UpdateAddressService(addr *address.Address) error {

	if addr.ID == 0 {
		return errors.New("address id is required")
	}

	if addr.Address1 == "" || addr.City == "" {
		return errors.New("missing required fields")
	}

	return address.UpdateAddress(addr)
}
