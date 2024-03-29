/*
Lob

The Lob API is organized around REST. Our API is designed to have predictable, resource-oriented URLs and uses HTTP response codes to indicate any API errors. <p> Looking for our [previous documentation](https://lob.github.io/legacy-docs/)? 

API version: 1.3.0
Contact: lob-openapi@lob.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lob

import (
	"encoding/json"
	
	"fmt"
)

// CountryExtendedExpanded Full name of country
type CountryExtendedExpanded string

// List of country_extended_expanded
const (
	COUNTRYEXTENDEDEXPANDED_EMPTY CountryExtendedExpanded = ""
	COUNTRYEXTENDEDEXPANDED_AFGHANISTAN CountryExtendedExpanded = "AFGHANISTAN"
	COUNTRYEXTENDEDEXPANDED_ALBANIA CountryExtendedExpanded = "ALBANIA"
	COUNTRYEXTENDEDEXPANDED_ALGERIA CountryExtendedExpanded = "ALGERIA"
	COUNTRYEXTENDEDEXPANDED_AMERICAN_SAMOA CountryExtendedExpanded = "AMERICAN SAMOA"
	COUNTRYEXTENDEDEXPANDED_ANDORRA CountryExtendedExpanded = "ANDORRA"
	COUNTRYEXTENDEDEXPANDED_ANGOLA CountryExtendedExpanded = "ANGOLA"
	COUNTRYEXTENDEDEXPANDED_ANGUILLA CountryExtendedExpanded = "ANGUILLA"
	COUNTRYEXTENDEDEXPANDED_ANTARCTICA CountryExtendedExpanded = "ANTARCTICA"
	COUNTRYEXTENDEDEXPANDED_ANTIGUA_AND_BARBUDA CountryExtendedExpanded = "ANTIGUA AND BARBUDA"
	COUNTRYEXTENDEDEXPANDED_ARGENTINA CountryExtendedExpanded = "ARGENTINA"
	COUNTRYEXTENDEDEXPANDED_ARUBA CountryExtendedExpanded = "ARUBA"
	COUNTRYEXTENDEDEXPANDED_AUSTRALIA CountryExtendedExpanded = "AUSTRALIA"
	COUNTRYEXTENDEDEXPANDED_AUSTRIA CountryExtendedExpanded = "AUSTRIA"
	COUNTRYEXTENDEDEXPANDED_AZERBAIJAN CountryExtendedExpanded = "AZERBAIJAN"
	COUNTRYEXTENDEDEXPANDED_BAHRAIN CountryExtendedExpanded = "BAHRAIN"
	COUNTRYEXTENDEDEXPANDED_BANGLADESH CountryExtendedExpanded = "BANGLADESH"
	COUNTRYEXTENDEDEXPANDED_BARBADOS CountryExtendedExpanded = "BARBADOS"
	COUNTRYEXTENDEDEXPANDED_BELARUS CountryExtendedExpanded = "BELARUS"
	COUNTRYEXTENDEDEXPANDED_BELGIUM CountryExtendedExpanded = "BELGIUM"
	COUNTRYEXTENDEDEXPANDED_BELIZE CountryExtendedExpanded = "BELIZE"
	COUNTRYEXTENDEDEXPANDED_BENIN CountryExtendedExpanded = "BENIN"
	COUNTRYEXTENDEDEXPANDED_BERMUDA CountryExtendedExpanded = "BERMUDA"
	COUNTRYEXTENDEDEXPANDED_BHUTAN CountryExtendedExpanded = "BHUTAN"
	COUNTRYEXTENDEDEXPANDED_BOLIVIA__PLURINATIONAL_STATE_OF CountryExtendedExpanded = "BOLIVIA (PLURINATIONAL STATE OF)"
	COUNTRYEXTENDEDEXPANDED_BONAIRE_SAINT_EUSTATIUS_AND_SABA CountryExtendedExpanded = "BONAIRE, SAINT EUSTATIUS AND SABA"
	COUNTRYEXTENDEDEXPANDED_BOSNIA_AND_HERZEGOVINA CountryExtendedExpanded = "BOSNIA AND HERZEGOVINA"
	COUNTRYEXTENDEDEXPANDED_BOTSWANA CountryExtendedExpanded = "BOTSWANA"
	COUNTRYEXTENDEDEXPANDED_BRAZIL CountryExtendedExpanded = "BRAZIL"
	COUNTRYEXTENDEDEXPANDED_BRITISH_INDIAN_OCEAN_TERRITORY CountryExtendedExpanded = "BRITISH INDIAN OCEAN TERRITORY"
	COUNTRYEXTENDEDEXPANDED_BRITISH_VIRGIN_ISLANDS CountryExtendedExpanded = "BRITISH VIRGIN ISLANDS"
	COUNTRYEXTENDEDEXPANDED_BRUNEI_DARUSSALAM CountryExtendedExpanded = "BRUNEI DARUSSALAM"
	COUNTRYEXTENDEDEXPANDED_BULGARIA CountryExtendedExpanded = "BULGARIA"
	COUNTRYEXTENDEDEXPANDED_BURKINA_FASO CountryExtendedExpanded = "BURKINA FASO"
	COUNTRYEXTENDEDEXPANDED_BURUNDI CountryExtendedExpanded = "BURUNDI"
	COUNTRYEXTENDEDEXPANDED_CABO_VERDE CountryExtendedExpanded = "CABO VERDE"
	COUNTRYEXTENDEDEXPANDED_CAMBODIA CountryExtendedExpanded = "CAMBODIA"
	COUNTRYEXTENDEDEXPANDED_CAMEROON CountryExtendedExpanded = "CAMEROON"
	COUNTRYEXTENDEDEXPANDED_CANADA CountryExtendedExpanded = "CANADA"
	COUNTRYEXTENDEDEXPANDED_CAYMAN_ISLANDS CountryExtendedExpanded = "CAYMAN ISLANDS"
	COUNTRYEXTENDEDEXPANDED_CENTRAL_AFRICAN_REPUBLIC CountryExtendedExpanded = "CENTRAL AFRICAN REPUBLIC"
	COUNTRYEXTENDEDEXPANDED_CHAD CountryExtendedExpanded = "CHAD"
	COUNTRYEXTENDEDEXPANDED_CHILE CountryExtendedExpanded = "CHILE"
	COUNTRYEXTENDEDEXPANDED_CHINA CountryExtendedExpanded = "CHINA"
	COUNTRYEXTENDEDEXPANDED_COLOMBIA CountryExtendedExpanded = "COLOMBIA"
	COUNTRYEXTENDEDEXPANDED_COMOROS CountryExtendedExpanded = "COMOROS"
	COUNTRYEXTENDEDEXPANDED_CONGO CountryExtendedExpanded = "CONGO"
	COUNTRYEXTENDEDEXPANDED_CONGO_DEMOCRATIC_REPUBLIC_OF_THE CountryExtendedExpanded = "CONGO, DEMOCRATIC REPUBLIC OF THE"
	COUNTRYEXTENDEDEXPANDED_COOK_ISLANDS CountryExtendedExpanded = "COOK ISLANDS"
	COUNTRYEXTENDEDEXPANDED_COSTA_RICA CountryExtendedExpanded = "COSTA RICA"
	COUNTRYEXTENDEDEXPANDED_CTE_DIVOIRE CountryExtendedExpanded = "CÔTE D`IVOIRE"
	COUNTRYEXTENDEDEXPANDED_CROATIA CountryExtendedExpanded = "CROATIA"
	COUNTRYEXTENDEDEXPANDED_CUBA CountryExtendedExpanded = "CUBA"
	COUNTRYEXTENDEDEXPANDED_CURAAO CountryExtendedExpanded = "CURAÇAO"
	COUNTRYEXTENDEDEXPANDED_CYPRUS CountryExtendedExpanded = "CYPRUS"
	COUNTRYEXTENDEDEXPANDED_CZECH_REPUBLIC CountryExtendedExpanded = "CZECH REPUBLIC"
	COUNTRYEXTENDEDEXPANDED_DENMARK CountryExtendedExpanded = "DENMARK"
	COUNTRYEXTENDEDEXPANDED_DJIBOUTI CountryExtendedExpanded = "DJIBOUTI"
	COUNTRYEXTENDEDEXPANDED_DOMINICA CountryExtendedExpanded = "DOMINICA"
	COUNTRYEXTENDEDEXPANDED_DOMINICAN_REPUBLIC CountryExtendedExpanded = "DOMINICAN REPUBLIC"
	COUNTRYEXTENDEDEXPANDED_ECUADOR CountryExtendedExpanded = "ECUADOR"
	COUNTRYEXTENDEDEXPANDED_EGYPT CountryExtendedExpanded = "EGYPT"
	COUNTRYEXTENDEDEXPANDED_EL_SALVADOR CountryExtendedExpanded = "EL SALVADOR"
	COUNTRYEXTENDEDEXPANDED_EQUATORIAL_GUINEA CountryExtendedExpanded = "EQUATORIAL GUINEA"
	COUNTRYEXTENDEDEXPANDED_ERITREA CountryExtendedExpanded = "ERITREA"
	COUNTRYEXTENDEDEXPANDED_ESTONIA CountryExtendedExpanded = "ESTONIA"
	COUNTRYEXTENDEDEXPANDED_ESWATINI CountryExtendedExpanded = "ESWATINI"
	COUNTRYEXTENDEDEXPANDED_ETHIOPIA CountryExtendedExpanded = "ETHIOPIA"
	COUNTRYEXTENDEDEXPANDED_FALKLAND_ISLANDS__MALVINAS CountryExtendedExpanded = "FALKLAND ISLANDS (MALVINAS)"
	COUNTRYEXTENDEDEXPANDED_FAROE_ISLANDS CountryExtendedExpanded = "FAROE ISLANDS"
	COUNTRYEXTENDEDEXPANDED_FIJI CountryExtendedExpanded = "FIJI"
	COUNTRYEXTENDEDEXPANDED_FINLAND CountryExtendedExpanded = "FINLAND"
	COUNTRYEXTENDEDEXPANDED_FRANCE CountryExtendedExpanded = "FRANCE"
	COUNTRYEXTENDEDEXPANDED_GABON CountryExtendedExpanded = "GABON"
	COUNTRYEXTENDEDEXPANDED_GAMBIA CountryExtendedExpanded = "GAMBIA"
	COUNTRYEXTENDEDEXPANDED_GEORGIA CountryExtendedExpanded = "GEORGIA"
	COUNTRYEXTENDEDEXPANDED_GERMANY CountryExtendedExpanded = "GERMANY"
	COUNTRYEXTENDEDEXPANDED_GHANA CountryExtendedExpanded = "GHANA"
	COUNTRYEXTENDEDEXPANDED_GIBRALTAR CountryExtendedExpanded = "GIBRALTAR"
	COUNTRYEXTENDEDEXPANDED_GREECE CountryExtendedExpanded = "GREECE"
	COUNTRYEXTENDEDEXPANDED_GREENLAND CountryExtendedExpanded = "GREENLAND"
	COUNTRYEXTENDEDEXPANDED_GRENADA CountryExtendedExpanded = "GRENADA"
	COUNTRYEXTENDEDEXPANDED_GUATEMALA CountryExtendedExpanded = "GUATEMALA"
	COUNTRYEXTENDEDEXPANDED_GUINEA CountryExtendedExpanded = "GUINEA"
	COUNTRYEXTENDEDEXPANDED_GUINEA_BISSAU CountryExtendedExpanded = "GUINEA-BISSAU"
	COUNTRYEXTENDEDEXPANDED_GUYANA CountryExtendedExpanded = "GUYANA"
	COUNTRYEXTENDEDEXPANDED_HAITI CountryExtendedExpanded = "HAITI"
	COUNTRYEXTENDEDEXPANDED_HOLY_SEE CountryExtendedExpanded = "HOLY SEE"
	COUNTRYEXTENDEDEXPANDED_HONDURAS CountryExtendedExpanded = "HONDURAS"
	COUNTRYEXTENDEDEXPANDED_HONG_KONG CountryExtendedExpanded = "HONG KONG"
	COUNTRYEXTENDEDEXPANDED_HUNGARY CountryExtendedExpanded = "HUNGARY"
	COUNTRYEXTENDEDEXPANDED_ICELAND CountryExtendedExpanded = "ICELAND"
	COUNTRYEXTENDEDEXPANDED_INDIA CountryExtendedExpanded = "INDIA"
	COUNTRYEXTENDEDEXPANDED_INDONESIA CountryExtendedExpanded = "INDONESIA"
	COUNTRYEXTENDEDEXPANDED_IRAN__ISLAMIC_REPUBLIC_OF CountryExtendedExpanded = "IRAN (ISLAMIC REPUBLIC OF)"
	COUNTRYEXTENDEDEXPANDED_IRAQ CountryExtendedExpanded = "IRAQ"
	COUNTRYEXTENDEDEXPANDED_IRELAND CountryExtendedExpanded = "IRELAND"
	COUNTRYEXTENDEDEXPANDED_ISRAEL CountryExtendedExpanded = "ISRAEL"
	COUNTRYEXTENDEDEXPANDED_ITALY CountryExtendedExpanded = "ITALY"
	COUNTRYEXTENDEDEXPANDED_JAMAICA CountryExtendedExpanded = "JAMAICA"
	COUNTRYEXTENDEDEXPANDED_JAPAN CountryExtendedExpanded = "JAPAN"
	COUNTRYEXTENDEDEXPANDED_JORDAN CountryExtendedExpanded = "JORDAN"
	COUNTRYEXTENDEDEXPANDED_KAZAKHSTAN CountryExtendedExpanded = "KAZAKHSTAN"
	COUNTRYEXTENDEDEXPANDED_KENYA CountryExtendedExpanded = "KENYA"
	COUNTRYEXTENDEDEXPANDED_KIRIBATI CountryExtendedExpanded = "KIRIBATI"
	COUNTRYEXTENDEDEXPANDED_KOREA__DEMOCRATIC_PEOPLES_REPUBLIC_OF CountryExtendedExpanded = "KOREA (DEMOCRATIC PEOPLE’S REPUBLIC OF)"
	COUNTRYEXTENDEDEXPANDED_KOREA_REPUBLIC_OF CountryExtendedExpanded = "KOREA, REPUBLIC OF"
	COUNTRYEXTENDEDEXPANDED_KUWAIT CountryExtendedExpanded = "KUWAIT"
	COUNTRYEXTENDEDEXPANDED_KYRGYZSTAN CountryExtendedExpanded = "KYRGYZSTAN"
	COUNTRYEXTENDEDEXPANDED_LAO_PEOPLES_DEMOCRATIC_REPUBLIC CountryExtendedExpanded = "LAO PEOPLE’S DEMOCRATIC REPUBLIC"
	COUNTRYEXTENDEDEXPANDED_LATVIA CountryExtendedExpanded = "LATVIA"
	COUNTRYEXTENDEDEXPANDED_LEBANON CountryExtendedExpanded = "LEBANON"
	COUNTRYEXTENDEDEXPANDED_LESOTHO CountryExtendedExpanded = "LESOTHO"
	COUNTRYEXTENDEDEXPANDED_LIBERIA CountryExtendedExpanded = "LIBERIA"
	COUNTRYEXTENDEDEXPANDED_LIBYA CountryExtendedExpanded = "LIBYA"
	COUNTRYEXTENDEDEXPANDED_LIECHTENSTEIN CountryExtendedExpanded = "LIECHTENSTEIN"
	COUNTRYEXTENDEDEXPANDED_LITHUANIA CountryExtendedExpanded = "LITHUANIA"
	COUNTRYEXTENDEDEXPANDED_LUXEMBOURG CountryExtendedExpanded = "LUXEMBOURG"
	COUNTRYEXTENDEDEXPANDED_MACAO CountryExtendedExpanded = "MACAO"
	COUNTRYEXTENDEDEXPANDED_MACEDONIA CountryExtendedExpanded = "MACEDONIA"
	COUNTRYEXTENDEDEXPANDED_MADAGASCAR CountryExtendedExpanded = "MADAGASCAR"
	COUNTRYEXTENDEDEXPANDED_MALAWI CountryExtendedExpanded = "MALAWI"
	COUNTRYEXTENDEDEXPANDED_MALAYSIA CountryExtendedExpanded = "MALAYSIA"
	COUNTRYEXTENDEDEXPANDED_MALDIVES CountryExtendedExpanded = "MALDIVES"
	COUNTRYEXTENDEDEXPANDED_MALI CountryExtendedExpanded = "MALI"
	COUNTRYEXTENDEDEXPANDED_MALTA CountryExtendedExpanded = "MALTA"
	COUNTRYEXTENDEDEXPANDED_MAURITANIA CountryExtendedExpanded = "MAURITANIA"
	COUNTRYEXTENDEDEXPANDED_MAURITIUS CountryExtendedExpanded = "MAURITIUS"
	COUNTRYEXTENDEDEXPANDED_MEXICO CountryExtendedExpanded = "MEXICO"
	COUNTRYEXTENDEDEXPANDED_MOLDOVA_REPUBLIC_OF CountryExtendedExpanded = "MOLDOVA, REPUBLIC OF"
	COUNTRYEXTENDEDEXPANDED_MONACO CountryExtendedExpanded = "MONACO"
	COUNTRYEXTENDEDEXPANDED_MONGOLIA CountryExtendedExpanded = "MONGOLIA"
	COUNTRYEXTENDEDEXPANDED_MONTENEGRO CountryExtendedExpanded = "MONTENEGRO"
	COUNTRYEXTENDEDEXPANDED_MONTSERRAT CountryExtendedExpanded = "MONTSERRAT"
	COUNTRYEXTENDEDEXPANDED_MOROCCO CountryExtendedExpanded = "MOROCCO"
	COUNTRYEXTENDEDEXPANDED_MOZAMBIQUE CountryExtendedExpanded = "MOZAMBIQUE"
	COUNTRYEXTENDEDEXPANDED_MYANMAR CountryExtendedExpanded = "MYANMAR"
	COUNTRYEXTENDEDEXPANDED_NAMIBIA CountryExtendedExpanded = "NAMIBIA"
	COUNTRYEXTENDEDEXPANDED_NAURU CountryExtendedExpanded = "NAURU"
	COUNTRYEXTENDEDEXPANDED_NEPAL CountryExtendedExpanded = "NEPAL"
	COUNTRYEXTENDEDEXPANDED_NETHERLAND_ANTILLES CountryExtendedExpanded = "NETHERLAND ANTILLES"
	COUNTRYEXTENDEDEXPANDED_NETHERLANDS CountryExtendedExpanded = "NETHERLANDS"
	COUNTRYEXTENDEDEXPANDED_NEW_ZEALAND CountryExtendedExpanded = "NEW ZEALAND"
	COUNTRYEXTENDEDEXPANDED_NICARAGUA CountryExtendedExpanded = "NICARAGUA"
	COUNTRYEXTENDEDEXPANDED_NIGER CountryExtendedExpanded = "NIGER"
	COUNTRYEXTENDEDEXPANDED_NIGERIA CountryExtendedExpanded = "NIGERIA"
	COUNTRYEXTENDEDEXPANDED_NIUE CountryExtendedExpanded = "NIUE"
	COUNTRYEXTENDEDEXPANDED_NORFOLK_ISLAND CountryExtendedExpanded = "NORFOLK ISLAND"
	COUNTRYEXTENDEDEXPANDED_NORWAY CountryExtendedExpanded = "NORWAY"
	COUNTRYEXTENDEDEXPANDED_OMAN CountryExtendedExpanded = "OMAN"
	COUNTRYEXTENDEDEXPANDED_PAKISTAN CountryExtendedExpanded = "PAKISTAN"
	COUNTRYEXTENDEDEXPANDED_PANAMA CountryExtendedExpanded = "PANAMA"
	COUNTRYEXTENDEDEXPANDED_PAPUA_NEW_GUINEA CountryExtendedExpanded = "PAPUA NEW GUINEA"
	COUNTRYEXTENDEDEXPANDED_PARAGUAY CountryExtendedExpanded = "PARAGUAY"
	COUNTRYEXTENDEDEXPANDED_PERU CountryExtendedExpanded = "PERU"
	COUNTRYEXTENDEDEXPANDED_PHILIPPINES CountryExtendedExpanded = "PHILIPPINES"
	COUNTRYEXTENDEDEXPANDED_PITCAIRN CountryExtendedExpanded = "PITCAIRN"
	COUNTRYEXTENDEDEXPANDED_POLAND CountryExtendedExpanded = "POLAND"
	COUNTRYEXTENDEDEXPANDED_PORTUGAL CountryExtendedExpanded = "PORTUGAL"
	COUNTRYEXTENDEDEXPANDED_QATAR CountryExtendedExpanded = "QATAR"
	COUNTRYEXTENDEDEXPANDED_ROMANIA CountryExtendedExpanded = "ROMANIA"
	COUNTRYEXTENDEDEXPANDED_RUSSIAN_FEDERATION CountryExtendedExpanded = "RUSSIAN FEDERATION"
	COUNTRYEXTENDEDEXPANDED_RWANDA CountryExtendedExpanded = "RWANDA"
	COUNTRYEXTENDEDEXPANDED_SAINT_HELENA CountryExtendedExpanded = "SAINT HELENA"
	COUNTRYEXTENDEDEXPANDED_SAINT_KITTS_AND_NEVIS CountryExtendedExpanded = "SAINT KITTS AND NEVIS"
	COUNTRYEXTENDEDEXPANDED_SAINT_LUCIA CountryExtendedExpanded = "SAINT LUCIA"
	COUNTRYEXTENDEDEXPANDED_SAINT_VINCENT_AND_THE_GRENADINES CountryExtendedExpanded = "SAINT VINCENT AND THE GRENADINES"
	COUNTRYEXTENDEDEXPANDED_SAMOA CountryExtendedExpanded = "SAMOA"
	COUNTRYEXTENDEDEXPANDED_SAN_MARINO CountryExtendedExpanded = "SAN MARINO"
	COUNTRYEXTENDEDEXPANDED_SAO_TOME_AND_PRINCIPE CountryExtendedExpanded = "SAO TOME AND PRINCIPE"
	COUNTRYEXTENDEDEXPANDED_SAUDI_ARABIA CountryExtendedExpanded = "SAUDI ARABIA"
	COUNTRYEXTENDEDEXPANDED_SENEGAL CountryExtendedExpanded = "SENEGAL"
	COUNTRYEXTENDEDEXPANDED_SERBIA CountryExtendedExpanded = "SERBIA"
	COUNTRYEXTENDEDEXPANDED_SEYCHELLES CountryExtendedExpanded = "SEYCHELLES"
	COUNTRYEXTENDEDEXPANDED_SIERRA_LEONE CountryExtendedExpanded = "SIERRA LEONE"
	COUNTRYEXTENDEDEXPANDED_SINGAPORE CountryExtendedExpanded = "SINGAPORE"
	COUNTRYEXTENDEDEXPANDED_SINT_MAARTEN CountryExtendedExpanded = "SINT MAARTEN"
	COUNTRYEXTENDEDEXPANDED_SLOVAKIA CountryExtendedExpanded = "SLOVAKIA"
	COUNTRYEXTENDEDEXPANDED_SLOVENIA CountryExtendedExpanded = "SLOVENIA"
	COUNTRYEXTENDEDEXPANDED_SOLOMON_ISLANDS CountryExtendedExpanded = "SOLOMON ISLANDS"
	COUNTRYEXTENDEDEXPANDED_SOMALIA CountryExtendedExpanded = "SOMALIA"
	COUNTRYEXTENDEDEXPANDED_SOUTH_AFRICA CountryExtendedExpanded = "SOUTH AFRICA"
	COUNTRYEXTENDEDEXPANDED_SOUTH_GEORGIA_AND_THE_SOUTH_SANDWICH_ISLANDS CountryExtendedExpanded = "SOUTH GEORGIA AND THE SOUTH SANDWICH ISLANDS"
	COUNTRYEXTENDEDEXPANDED_SOUTH_SUDAN CountryExtendedExpanded = "SOUTH SUDAN"
	COUNTRYEXTENDEDEXPANDED_SPAIN CountryExtendedExpanded = "SPAIN"
	COUNTRYEXTENDEDEXPANDED_SRI_LANKA CountryExtendedExpanded = "SRI LANKA"
	COUNTRYEXTENDEDEXPANDED_SUDAN CountryExtendedExpanded = "SUDAN"
	COUNTRYEXTENDEDEXPANDED_SURINAME CountryExtendedExpanded = "SURINAME"
	COUNTRYEXTENDEDEXPANDED_SWEDEN CountryExtendedExpanded = "SWEDEN"
	COUNTRYEXTENDEDEXPANDED_SWITZERLAND CountryExtendedExpanded = "SWITZERLAND"
	COUNTRYEXTENDEDEXPANDED_SYRIAN_ARAB_REPUBLIC CountryExtendedExpanded = "SYRIAN ARAB REPUBLIC"
	COUNTRYEXTENDEDEXPANDED_TAIWAN CountryExtendedExpanded = "TAIWAN"
	COUNTRYEXTENDEDEXPANDED_TAJIKISTAN CountryExtendedExpanded = "TAJIKISTAN"
	COUNTRYEXTENDEDEXPANDED_TANZANIA CountryExtendedExpanded = "TANZANIA"
	COUNTRYEXTENDEDEXPANDED_THAILAND CountryExtendedExpanded = "THAILAND"
	COUNTRYEXTENDEDEXPANDED_THE_BAHAMAS CountryExtendedExpanded = "THE BAHAMAS"
	COUNTRYEXTENDEDEXPANDED_TIMOR_LESTE CountryExtendedExpanded = "TIMOR-LESTE"
	COUNTRYEXTENDEDEXPANDED_TOGO CountryExtendedExpanded = "TOGO"
	COUNTRYEXTENDEDEXPANDED_TOKELAU CountryExtendedExpanded = "TOKELAU"
	COUNTRYEXTENDEDEXPANDED_TONGA CountryExtendedExpanded = "TONGA"
	COUNTRYEXTENDEDEXPANDED_TRINIDAD_AND_TOBAGO CountryExtendedExpanded = "TRINIDAD AND TOBAGO"
	COUNTRYEXTENDEDEXPANDED_TUNISIA CountryExtendedExpanded = "TUNISIA"
	COUNTRYEXTENDEDEXPANDED_TURKEY CountryExtendedExpanded = "TURKEY"
	COUNTRYEXTENDEDEXPANDED_TURKMENISTAN CountryExtendedExpanded = "TURKMENISTAN"
	COUNTRYEXTENDEDEXPANDED_TURKS_AND_CAICOS_ISLANDS CountryExtendedExpanded = "TURKS AND CAICOS ISLANDS"
	COUNTRYEXTENDEDEXPANDED_TUVALU CountryExtendedExpanded = "TUVALU"
	COUNTRYEXTENDEDEXPANDED_UGANDA CountryExtendedExpanded = "UGANDA"
	COUNTRYEXTENDEDEXPANDED_UKRAINE CountryExtendedExpanded = "UKRAINE"
	COUNTRYEXTENDEDEXPANDED_UNITED_ARAB_EMIRATES CountryExtendedExpanded = "UNITED ARAB EMIRATES"
	COUNTRYEXTENDEDEXPANDED_UNITED_KINGDOM CountryExtendedExpanded = "UNITED KINGDOM"
	COUNTRYEXTENDEDEXPANDED_UNITED_STATES CountryExtendedExpanded = "UNITED STATES"
	COUNTRYEXTENDEDEXPANDED_URUGUAY CountryExtendedExpanded = "URUGUAY"
	COUNTRYEXTENDEDEXPANDED_UZBEKISTAN CountryExtendedExpanded = "UZBEKISTAN"
	COUNTRYEXTENDEDEXPANDED_VANUATU CountryExtendedExpanded = "VANUATU"
	COUNTRYEXTENDEDEXPANDED_VENEZUELA CountryExtendedExpanded = "VENEZUELA"
	COUNTRYEXTENDEDEXPANDED_VIET_NAM CountryExtendedExpanded = "VIET NAM"
	COUNTRYEXTENDEDEXPANDED_WESTERN_SAHARA CountryExtendedExpanded = "WESTERN SAHARA"
	COUNTRYEXTENDEDEXPANDED_YEMEN CountryExtendedExpanded = "YEMEN"
	COUNTRYEXTENDEDEXPANDED_ZAMBIA CountryExtendedExpanded = "ZAMBIA"
	COUNTRYEXTENDEDEXPANDED_ZIMBABWE CountryExtendedExpanded = "ZIMBABWE"
)

// All allowed values of CountryExtendedExpanded enum
var AllowedCountryExtendedExpandedEnumValues = []CountryExtendedExpanded{
	"",
	"AFGHANISTAN",
	"ALBANIA",
	"ALGERIA",
	"AMERICAN SAMOA",
	"ANDORRA",
	"ANGOLA",
	"ANGUILLA",
	"ANTARCTICA",
	"ANTIGUA AND BARBUDA",
	"ARGENTINA",
	"ARUBA",
	"AUSTRALIA",
	"AUSTRIA",
	"AZERBAIJAN",
	"BAHRAIN",
	"BANGLADESH",
	"BARBADOS",
	"BELARUS",
	"BELGIUM",
	"BELIZE",
	"BENIN",
	"BERMUDA",
	"BHUTAN",
	"BOLIVIA (PLURINATIONAL STATE OF)",
	"BONAIRE, SAINT EUSTATIUS AND SABA",
	"BOSNIA AND HERZEGOVINA",
	"BOTSWANA",
	"BRAZIL",
	"BRITISH INDIAN OCEAN TERRITORY",
	"BRITISH VIRGIN ISLANDS",
	"BRUNEI DARUSSALAM",
	"BULGARIA",
	"BURKINA FASO",
	"BURUNDI",
	"CABO VERDE",
	"CAMBODIA",
	"CAMEROON",
	"CANADA",
	"CAYMAN ISLANDS",
	"CENTRAL AFRICAN REPUBLIC",
	"CHAD",
	"CHILE",
	"CHINA",
	"COLOMBIA",
	"COMOROS",
	"CONGO",
	"CONGO, DEMOCRATIC REPUBLIC OF THE",
	"COOK ISLANDS",
	"COSTA RICA",
	"CÔTE D`IVOIRE",
	"CROATIA",
	"CUBA",
	"CURAÇAO",
	"CYPRUS",
	"CZECH REPUBLIC",
	"DENMARK",
	"DJIBOUTI",
	"DOMINICA",
	"DOMINICAN REPUBLIC",
	"ECUADOR",
	"EGYPT",
	"EL SALVADOR",
	"EQUATORIAL GUINEA",
	"ERITREA",
	"ESTONIA",
	"ESWATINI",
	"ETHIOPIA",
	"FALKLAND ISLANDS (MALVINAS)",
	"FAROE ISLANDS",
	"FIJI",
	"FINLAND",
	"FRANCE",
	"GABON",
	"GAMBIA",
	"GEORGIA",
	"GERMANY",
	"GHANA",
	"GIBRALTAR",
	"GREECE",
	"GREENLAND",
	"GRENADA",
	"GUATEMALA",
	"GUINEA",
	"GUINEA-BISSAU",
	"GUYANA",
	"HAITI",
	"HOLY SEE",
	"HONDURAS",
	"HONG KONG",
	"HUNGARY",
	"ICELAND",
	"INDIA",
	"INDONESIA",
	"IRAN (ISLAMIC REPUBLIC OF)",
	"IRAQ",
	"IRELAND",
	"ISRAEL",
	"ITALY",
	"JAMAICA",
	"JAPAN",
	"JORDAN",
	"KAZAKHSTAN",
	"KENYA",
	"KIRIBATI",
	"KOREA (DEMOCRATIC PEOPLE’S REPUBLIC OF)",
	"KOREA, REPUBLIC OF",
	"KUWAIT",
	"KYRGYZSTAN",
	"LAO PEOPLE’S DEMOCRATIC REPUBLIC",
	"LATVIA",
	"LEBANON",
	"LESOTHO",
	"LIBERIA",
	"LIBYA",
	"LIECHTENSTEIN",
	"LITHUANIA",
	"LUXEMBOURG",
	"MACAO",
	"MACEDONIA",
	"MADAGASCAR",
	"MALAWI",
	"MALAYSIA",
	"MALDIVES",
	"MALI",
	"MALTA",
	"MAURITANIA",
	"MAURITIUS",
	"MEXICO",
	"MOLDOVA, REPUBLIC OF",
	"MONACO",
	"MONGOLIA",
	"MONTENEGRO",
	"MONTSERRAT",
	"MOROCCO",
	"MOZAMBIQUE",
	"MYANMAR",
	"NAMIBIA",
	"NAURU",
	"NEPAL",
	"NETHERLAND ANTILLES",
	"NETHERLANDS",
	"NEW ZEALAND",
	"NICARAGUA",
	"NIGER",
	"NIGERIA",
	"NIUE",
	"NORFOLK ISLAND",
	"NORWAY",
	"OMAN",
	"PAKISTAN",
	"PANAMA",
	"PAPUA NEW GUINEA",
	"PARAGUAY",
	"PERU",
	"PHILIPPINES",
	"PITCAIRN",
	"POLAND",
	"PORTUGAL",
	"QATAR",
	"ROMANIA",
	"RUSSIAN FEDERATION",
	"RWANDA",
	"SAINT HELENA",
	"SAINT KITTS AND NEVIS",
	"SAINT LUCIA",
	"SAINT VINCENT AND THE GRENADINES",
	"SAMOA",
	"SAN MARINO",
	"SAO TOME AND PRINCIPE",
	"SAUDI ARABIA",
	"SENEGAL",
	"SERBIA",
	"SEYCHELLES",
	"SIERRA LEONE",
	"SINGAPORE",
	"SINT MAARTEN",
	"SLOVAKIA",
	"SLOVENIA",
	"SOLOMON ISLANDS",
	"SOMALIA",
	"SOUTH AFRICA",
	"SOUTH GEORGIA AND THE SOUTH SANDWICH ISLANDS",
	"SOUTH SUDAN",
	"SPAIN",
	"SRI LANKA",
	"SUDAN",
	"SURINAME",
	"SWEDEN",
	"SWITZERLAND",
	"SYRIAN ARAB REPUBLIC",
	"TAIWAN",
	"TAJIKISTAN",
	"TANZANIA",
	"THAILAND",
	"THE BAHAMAS",
	"TIMOR-LESTE",
	"TOGO",
	"TOKELAU",
	"TONGA",
	"TRINIDAD AND TOBAGO",
	"TUNISIA",
	"TURKEY",
	"TURKMENISTAN",
	"TURKS AND CAICOS ISLANDS",
	"TUVALU",
	"UGANDA",
	"UKRAINE",
	"UNITED ARAB EMIRATES",
	"UNITED KINGDOM",
	"UNITED STATES",
	"URUGUAY",
	"UZBEKISTAN",
	"VANUATU",
	"VENEZUELA",
	"VIET NAM",
	"WESTERN SAHARA",
	"YEMEN",
	"ZAMBIA",
	"ZIMBABWE",
}

func (v *CountryExtendedExpanded) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CountryExtendedExpanded(value)
	for _, existing := range AllowedCountryExtendedExpandedEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CountryExtendedExpanded", value)
}

// NewCountryExtendedExpandedFromValue returns a pointer to a valid CountryExtendedExpanded
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCountryExtendedExpandedFromValue(v string) (*CountryExtendedExpanded, error) {
	ev := CountryExtendedExpanded(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CountryExtendedExpanded: valid values are %v", v, AllowedCountryExtendedExpandedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CountryExtendedExpanded) IsValid() bool {
	for _, existing := range AllowedCountryExtendedExpandedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to country_extended_expanded value
func (v CountryExtendedExpanded) Ptr() *CountryExtendedExpanded {
	return &v
}

type NullableCountryExtendedExpanded struct {
	value *CountryExtendedExpanded
	isSet bool
}

func (v NullableCountryExtendedExpanded) Get() *CountryExtendedExpanded {
	return v.value
}

func (v *NullableCountryExtendedExpanded) Set(val *CountryExtendedExpanded) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryExtendedExpanded) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryExtendedExpanded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryExtendedExpanded(val *CountryExtendedExpanded) *NullableCountryExtendedExpanded {
	return &NullableCountryExtendedExpanded{value: val, isSet: true}
}

func (v NullableCountryExtendedExpanded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryExtendedExpanded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

