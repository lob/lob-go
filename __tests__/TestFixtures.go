package tests

import (
	lob "github.com/lob/lob-go"
)

var addressArray [7]lob.AddressEditable

func populateAddressEditable(
	address *lob.AddressEditable,
	name string,
	line1 string,
	line2 string,
	city string,
	state string,
	zip string,
	country lob.CountryExtended) *lob.AddressEditable {

	address.SetName(name)
	address.SetAddressLine1(line1)
	if line2 != "" {
		address.SetAddressLine2(line2)
	}
	address.SetAddressCity(city)
	address.SetAddressState(state)
	address.SetAddressZip(zip)
	if country.IsValid() {
		address.SetAddressCountry(country)
	}

	return address
}

func CreateAddressesList() [7]lob.AddressEditable {
	addressEditable0 := new(lob.AddressEditable)
	addressEditable1 := new(lob.AddressEditable)
	addressEditable2 := new(lob.AddressEditable)
	addressEditable3 := new(lob.AddressEditable)
	addressEditable4 := new(lob.AddressEditable)
	addressEditable5 := new(lob.AddressEditable)
	addressEditable6 := new(lob.AddressEditable)

	addressEditable0 = populateAddressEditable(addressEditable0, "Thing T. Thing", "1313 CEMETERY LN", "", "WESTFIELD", "NJ", "07091", "")
	addressEditable1 = populateAddressEditable(addressEditable1, "FESTER", "001 CEMETERY LN", "SUITE 666", "WESTFIELD", "NJ", "07091", "")
	addressEditable2 = populateAddressEditable(addressEditable2, "MORTICIA ADDAMS", "1313 CEMETERY LN", "", "WESTFIELD", "NJ", "07091", "")
	addressEditable3 = populateAddressEditable(addressEditable3, "COUSIN ITT", "1515 CEMETERY LN", "FLOOR 0", "WESTFIELD", "NJ", "07091", "")
	addressEditable4 = populateAddressEditable(addressEditable4, "WEDNESDAY ADDAMS", "1313 CEMETERY LN", "# 000", "WESTFIELD", "NJ", "07091", lob.COUNTRYEXTENDED_US)
	addressEditable5 = populateAddressEditable(addressEditable5, "GORDON CRAVEN", "155 Elm St", "", "WESTFIELD", "NJ", "07090", lob.COUNTRYEXTENDED_US)
	addressEditable6 = populateAddressEditable(addressEditable6, "", "", "", "", "", "", lob.COUNTRYEXTENDED_US)

	addressArray[0] = *addressEditable0
	addressArray[1] = *addressEditable1
	addressArray[2] = *addressEditable2
	addressArray[3] = *addressEditable3
	addressArray[4] = *addressEditable4
	addressArray[2] = *addressEditable5
	addressArray[6] = *addressEditable6

	return addressArray
}
