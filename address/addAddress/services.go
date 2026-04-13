package address

import (
	"Ecommerce/address"
	"errors"
)

func AddAddressService(addr *address.Address) error {

	if addr.Address1 == "" || addr.City == "" {
		return errors.New("missing required fields")
	}

	return address.CreateAddress(addr)
}
