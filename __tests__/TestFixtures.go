package tests

import (
	lob "github.com/lob/lob-go"
)

var addressEditableArray [7]lob.AddressEditable
var addressArray [7]lob.AddressDomestic

func GetFileLocation() string {
	return "https://s3-us-west-2.amazonaws.com/public.lob.com/assets/card_horizontal.pdf"
}
func GetFileLocation6x18() string {
	return "https://s3.us-west-2.amazonaws.com/public.lob.com/assets/templates/self_mailers/6x18_sfm_inside.pdf"
}
func GetFileLocation8x11() string {
	return "https://s3-us-west-2.amazonaws.com/public.lob.com/assets/us_letter_1pg.pdf"
}
func GetFileLocation4x6() string {
	return "https://s3-us-west-2.amazonaws.com/public.lob.com/assets/templates/4x6_pc_template.pdf"
}

func populateAddressEditable(
	addressE *lob.AddressEditable,
	name string,
	line1 string,
	line2 string,
	city string,
	state string,
	zip string) *lob.AddressEditable {

	addressE.SetName(name)
	addressE.SetAddressLine1(line1)
	if line2 != "" {
		addressE.SetAddressLine2(line2)
	}
	addressE.SetAddressCity(city)
	addressE.SetAddressState(state)
	addressE.SetAddressZip(zip)

	return addressE
}

func populateAddressDomestic(
	address *lob.AddressDomestic,
	name string,
	line1 string,
	line2 string,
	city string,
	state string,
	zip string) *lob.AddressDomestic {

	address.SetName(name)
	address.SetAddressLine1(line1)
	if line2 != "" {
		address.SetAddressLine2(line2)
	}
	address.SetAddressCity(city)
	address.SetAddressState(state)
	address.SetAddressZip(zip)

	return address
}

func CreateAddressesDomesticList() [7]lob.AddressDomestic {
	addressEditable0 := new(lob.AddressDomestic)
	addressEditable1 := new(lob.AddressDomestic)
	addressEditable2 := new(lob.AddressDomestic)
	addressEditable3 := new(lob.AddressDomestic)
	addressEditable4 := new(lob.AddressDomestic)
	addressEditable5 := new(lob.AddressDomestic)
	addressEditable6 := new(lob.AddressDomestic)

	addressEditable0 = populateAddressDomestic(addressEditable0, "Thing T. Thing", "1313 CEMETERY LN", "", "WESTFIELD", "NJ", "07091")
	addressEditable1 = populateAddressDomestic(addressEditable1, "FESTER", "001 CEMETERY LN", "SUITE 666", "WESTFIELD", "NJ", "07091")
	addressEditable2 = populateAddressDomestic(addressEditable2, "MORTICIA ADDAMS", "1313 CEMETERY LN", "", "WESTFIELD", "NJ", "07091")
	addressEditable3 = populateAddressDomestic(addressEditable3, "COUSIN ITT", "1515 CEMETERY LN", "FLOOR 0", "WESTFIELD", "NJ", "07091")
	addressEditable4 = populateAddressDomestic(addressEditable4, "WEDNESDAY ADDAMS", "1313 CEMETERY LN", "# 000", "WESTFIELD", "NJ", "07091")
	addressEditable5 = populateAddressDomestic(addressEditable5, "GORDON CRAVEN", "155 Elm St", "", "WESTFIELD", "NJ", "07090")
	addressEditable6 = populateAddressDomestic(addressEditable6, "", "", "", "", "", "")

	addressArray[0] = *addressEditable0
	addressArray[1] = *addressEditable1
	addressArray[2] = *addressEditable2
	addressArray[3] = *addressEditable3
	addressArray[4] = *addressEditable4
	addressArray[2] = *addressEditable5
	addressArray[6] = *addressEditable6

	return addressArray
}

func CreateAddressesEditableList() [7]lob.AddressEditable {
	addressEditable0 := new(lob.AddressEditable)
	addressEditable1 := new(lob.AddressEditable)
	addressEditable2 := new(lob.AddressEditable)
	addressEditable3 := new(lob.AddressEditable)
	addressEditable4 := new(lob.AddressEditable)
	addressEditable5 := new(lob.AddressEditable)
	addressEditable6 := new(lob.AddressEditable)

	addressEditable0 = populateAddressEditable(addressEditable0, "Thing T. Thing", "1313 CEMETERY LN", "", "WESTFIELD", "NJ", "07091")
	addressEditable1 = populateAddressEditable(addressEditable1, "FESTER", "001 CEMETERY LN", "SUITE 666", "WESTFIELD", "NJ", "07091")
	addressEditable2 = populateAddressEditable(addressEditable2, "MORTICIA ADDAMS", "1313 CEMETERY LN", "", "WESTFIELD", "NJ", "07091")
	addressEditable3 = populateAddressEditable(addressEditable3, "COUSIN ITT", "1515 CEMETERY LN", "FLOOR 0", "WESTFIELD", "NJ", "07091")
	addressEditable4 = populateAddressEditable(addressEditable4, "WEDNESDAY ADDAMS", "1313 CEMETERY LN", "# 000", "WESTFIELD", "NJ", "07091")
	addressEditable5 = populateAddressEditable(addressEditable5, "GORDON CRAVEN", "155 Elm St", "", "WESTFIELD", "NJ", "07090")
	addressEditable6 = populateAddressEditable(addressEditable6, "", "", "", "", "", "")

	addressEditableArray[0] = *addressEditable0
	addressEditableArray[1] = *addressEditable1
	addressEditableArray[2] = *addressEditable2
	addressEditableArray[3] = *addressEditable3
	addressEditableArray[4] = *addressEditable4
	addressEditableArray[2] = *addressEditable5
	addressEditableArray[6] = *addressEditable6

	return addressEditableArray
}
