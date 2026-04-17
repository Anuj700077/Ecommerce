package address

import (
	"Ecommerce/address"
	"errors"
)

func GetAddressService(userID uint) ([]address.Address, error) {
	addresses, err := address.GetAddressByUserID(userID)
	if err != nil {
		return nil, err
	}
	if len(addresses) == 0 {
		return nil, errors.New("No address found")
	}
	return addresses, nil
}
