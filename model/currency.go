package model

import (
	"fmt"
	"strconv"
)

type Amount int64
type Counter uint32
type Currency uint32

const (
	CurrencyNOT Currency = 0   // Currency not used - Валюта не используется
	CurrencyAMD          = 51  // Armenian Dram - Армянский драм
	CurrencyAUD          = 36  // Australian Dollar - Австралийский доллар
	CurrencyATS          = 40  // Austrian Schilling - Австрийский шиллинг
	CurrencyAZN          = 31  // Azerbaijan Manat - Азербайджанский манат
	CurrencyBYR          = 974 // Belarusian Ruble - Белорусский рубль
	CurrencyBEF          = 56  // Belgian Franc - Бельгийский франк
	CurrencyBRL          = 986 // Brazilian Real - Бразильский реал
	CurrencyGBP          = 826 // British Pound - Британский фунт
	CurrencyCAD          = 124 // Canadian Dollar - Канадский доллар
	CurrencyCNY          = 156 // China Yuan - Китайский юань
	CurrencyCZK          = 203 // Czech Koruna - Чешская крона
	CurrencyDKK          = 208 // Danish Krone - Датская крона
	CurrencyNLG          = 528 // Dutch Guilder - Нидерландский гульден
	CurrencyEEK          = 233 // Estonian Kroon - Эстонская крона
	CurrencyEUR          = 978 // European Euro - Единая европейская валюта
	CurrencyFIM          = 246 // Finnish Mark - Финская марка
	CurrencyFRF          = 250 // French Franc - Французский франк
	CurrencyGEL          = 981 // Georgia Lari - Грузинская лари
	CurrencyDEM          = 276 // German Mark - Немецкая марка
	CurrencyGRD          = 300 // Greek Drachma - Греческая драхма
	CurrencyHKD          = 344 // Hong Kong Dollar - Гонконгский доллар
	CurrencyHUF          = 348 // Hungarian Forint - Венгерский форинт
	CurrencyINR          = 356 // Indian Rupee - Индийская рупия
	CurrencyIEP          = 372 // Irish Punt - Ирландский фунт
	CurrencyILS          = 376 // Israeli Sheqel - Израильский шекель
	CurrencyITL          = 380 // Italian Lira - Итальянская лира
	CurrencyJPY          = 392 // Japanese Yen - Японская йена
	CurrencyKZT          = 398 // Kazakhstan Tenge - Казахский тенге
	CurrencyKGS          = 417 // Kyrgyzstan Som - Киргизский сом
	CurrencyLVL          = 428 // Latvian Lat - Латвийский лат
	CurrencyLTL          = 440 // Lithuanian Lita - Литовский лит
	CurrencyMDL          = 498 // Moldovan Leu - Молдавский лей
	CurrencyMXN          = 484 // Mexican Peso - Мексиканский песо
	CurrencyNZD          = 554 // New Zealand Dollar - Новозеландский доллар
	CurrencyNOK          = 578 // Norway Krone - Норвежская крона
	CurrencyPLN          = 985 // Polish Zloty - Польский злотый
	CurrencyPTE          = 620 // Portuguese Escudo - Португальское эскудо
	CurrencyROL          = 642 // Romania Leu - Румынский лей
	CurrencyRUB          = 643 // Russian Rouble - Российский рубль (новый)
	CurrencySGD          = 702 // Singapore Dollar - Сингапурский доллар
	CurrencySKK          = 703 // Slovak Koruna - Словацкая крона
	CurrencyZAR          = 710 // South African Rand - Южноафриканский ранд
	CurrencyESP          = 724 // Spanish Peseta - Испанская песета
	CurrencySEK          = 752 // Swedish Krona - Шведская крона
	CurrencyCHF          = 756 // Swiss Franc - Швейцарский франк
	CurrencyTRY          = 949 // Turkish Lira - Турецкая лира
	CurrencyUAH          = 980 // Ukraine Hryvnia - Украинская гривна
	CurrencyUSD          = 840 // United States Dollar - Американский доллар
	CurrencyUZS          = 860 // Uzbekistan Sum - Узбекский сум
	CurrencyXAG          = 961 // Silver - Серебро
	CurrencyXAU          = 959 // Gold - Золото
	CurrencyXPD          = 964 // Palladium - Палладий
	CurrencyXPT          = 962 // Platinum - Платина
	CurrencyXXX          = 999 // Not currency (Token) - Нет валюты (жетон)
)

// String returns a string explaining of the currency code
func (e Currency) String() string {
	return e.IsoCode()
}

// PlainText returns a text name of the currency
func (e Currency) PlainText() string {
	data := map[Currency]string{
		CurrencyNOT: "Currency not used",
		CurrencyAMD: "Armenian Dram",        // Armenian Dram - Армянский драм
		CurrencyAUD: "Australian Dollar",    // Australian Dollar - Австралийский доллар
		CurrencyATS: "Austrian Schilling",   // Austrian Schilling - Австрийский шиллинг
		CurrencyAZN: "Azerbaijan Manat",     // Azerbaijan Manat - Азербайджанский манат
		CurrencyBYR: "Belarusian Ruble",     // Belarusian Ruble - Белорусский рубль
		CurrencyBEF: "Belgian Franc",        // Belgian Franc - Бельгийский франк
		CurrencyBRL: "Brazilian Real",       // Brazilian Real - Бразильский реал
		CurrencyGBP: "British Pound",        // British Pound - Британский фунт
		CurrencyCAD: "Canadian Dollar",      // Canadian Dollar - Канадский доллар
		CurrencyCNY: "China Yuan",           // China Yuan - Китайский юань
		CurrencyCZK: "Czech Koruna",         // Czech Koruna - Чешская крона
		CurrencyDKK: "Danish Krone",         // Danish Krone - Датская крона
		CurrencyNLG: "Dutch Guilder",        // Dutch Guilder - Нидерландский гульден
		CurrencyEEK: "Estonian Kroon",       // Estonian Kroon - Эстонская крона
		CurrencyEUR: "European Euro",        // European Euro - Единая европейская валюта
		CurrencyFIM: "Finnish Mark",         // Finnish Mark - Финская марка
		CurrencyFRF: "French Franc",         // French Franc - Французский франк
		CurrencyGEL: "Georgia Lari",         // Georgia Lari - Грузинская лари
		CurrencyDEM: "German Mark",          // German Mark - Немецкая марка
		CurrencyGRD: "Greek Drachma",        // Greek Drachma - Греческая драхма
		CurrencyHKD: "Hong Kong Dollar",     // Hong Kong Dollar - Гонконгский доллар
		CurrencyHUF: "Hungarian Forint",     // Hungarian Forint - Венгерский форинт
		CurrencyINR: "Indian Rupee",         // Indian Rupee - Индийская рупия
		CurrencyIEP: "Irish Punt",           // Irish Punt - Ирландский фунт
		CurrencyILS: "Israeli Sheqel",       // Israeli Sheqel - Израильский шекель
		CurrencyITL: "Italian Lira",         // Italian Lira - Итальянская лира
		CurrencyJPY: "Japanese Yen",         // Japanese Yen - Японская йена
		CurrencyKZT: "Kazakhstan Tenge",     // Kazakhstan Tenge - Казахский тенге
		CurrencyKGS: "Kyrgyzstan Som",       // Kyrgyzstan Som - Киргизский сом
		CurrencyLVL: "Latvian Lat",          // Latvian Lat - Латвийский лат
		CurrencyLTL: "Lithuanian Lita",      // Lithuanian Lita - Литовский лит
		CurrencyMDL: "Moldovan Leu",         // Moldovan Leu - Молдавский лей
		CurrencyMXN: "Mexican Peso",         // Mexican Peso - Мексиканский песо
		CurrencyNZD: "New Zealand Dollar",   // New Zealand Dollar - Новозеландский доллар
		CurrencyNOK: "Norway Krone",         // Norway Krone - Норвежская крона
		CurrencyPLN: "Polish Zloty",         // Polish Zloty - Польский злотый
		CurrencyPTE: "Portuguese Escudo",    // Portuguese Escudo - Португальское эскудо
		CurrencyROL: "Romania Leu",          // Romania Leu - Румынский лей
		CurrencyRUB: "Russian Rouble",       // Russian Rouble - Российский рубль (новый)
		CurrencySGD: "Singapore Dollar",     // Singapore Dollar - Сингапурский доллар
		CurrencySKK: "Slovak Koruna",        // Slovak Koruna - Словацкая крона
		CurrencyZAR: "South African Rand",   // South African Rand - Южноафриканский ранд
		CurrencyESP: "Spanish Peseta",       // Spanish Peseta - Испанская песета
		CurrencySEK: "Swedish Krona",        // Swedish Krona - Шведская крона
		CurrencyCHF: "Swiss Franc",          // Swiss Franc - Швейцарский франк
		CurrencyTRY: "Turkish Lira",         // Turkish Lira - Турецкая лира
		CurrencyUAH: "Ukraine Hryvnia",      // Ukraine Hryvnia - Украинская гривна
		CurrencyUSD: "United States Dollar", // United States Dollar - Американский доллар
		CurrencyUZS: "Uzbekistan Sum",       // Uzbekistan Sum - Узбекский сум
		CurrencyXAG: "Silver",               // Silver - Серебро
		CurrencyXAU: "Gold",                 // Gold - Золото
		CurrencyXPD: "Palladium",            // Palladium - Палладий
		CurrencyXPT: "Platinum",             // Platinum - Платина
		CurrencyXXX: "Not currency (Token)", // Not currency (Token) - Нет валюты (жетон)
	}
	value, exists := data[e]
	if exists {
		return value
	}
	return "Unknown Currency"
}

// IsoCode returns the text ISO code of the currency
func (e Currency) IsoCode() string {
	data := map[Currency]string{
		CurrencyNOT: "NOT",
		CurrencyAMD: "AMD", // Armenian Dram - Армянский драм
		CurrencyAUD: "AUD", // Australian Dollar - Австралийский доллар
		CurrencyATS: "ATS", // Austrian Schilling - Австрийский шиллинг
		CurrencyAZN: "AZN", // Azerbaijan Manat - Азербайджанский манат
		CurrencyBYR: "BYR", // Belarusian Ruble - Белорусский рубль
		CurrencyBEF: "BEF", // Belgian Franc - Бельгийский франк
		CurrencyBRL: "BRL", // Brazilian Real - Бразильский реал
		CurrencyGBP: "GBP", // British Pound - Британский фунт
		CurrencyCAD: "CAD", // Canadian Dollar - Канадский доллар
		CurrencyCNY: "CNY", // China Yuan - Китайский юань
		CurrencyCZK: "CZK", // Czech Koruna - Чешская крона
		CurrencyDKK: "DKK", // Danish Krone - Датская крона
		CurrencyNLG: "NLG", // Dutch Guilder - Нидерландский гульден
		CurrencyEEK: "EEK", // Estonian Kroon - Эстонская крона
		CurrencyEUR: "EUR", // European Euro - Единая европейская валюта
		CurrencyFIM: "FIM", // Finnish Mark - Финская марка
		CurrencyFRF: "FRF", // French Franc - Французский франк
		CurrencyGEL: "GEL", // Georgia Lari - Грузинская лари
		CurrencyDEM: "DEM", // German Mark - Немецкая марка
		CurrencyGRD: "GRD", // Greek Drachma - Греческая драхма
		CurrencyHKD: "HKD", // Hong Kong Dollar - Гонконгский доллар
		CurrencyHUF: "HUF", // Hungarian Forint - Венгерский форинт
		CurrencyINR: "INR", // Indian Rupee - Индийская рупия
		CurrencyIEP: "IEP", // Irish Punt - Ирландский фунт
		CurrencyILS: "ILS", // Israeli Sheqel - Израильский шекель
		CurrencyITL: "ITL", // Italian Lira - Итальянская лира
		CurrencyJPY: "JPY", // Japanese Yen - Японская йена
		CurrencyKZT: "KZT", // Kazakhstan Tenge - Казахский тенге
		CurrencyKGS: "KGS", // Kyrgyzstan Som - Киргизский сом
		CurrencyLVL: "LVL", // Latvian Lat - Латвийский лат
		CurrencyLTL: "LTL", // Lithuanian Lita - Литовский лит
		CurrencyMDL: "MDL", // Moldovan Leu - Молдавский лей
		CurrencyMXN: "MXN", // Mexican Peso - Мексиканский песо
		CurrencyNZD: "NZD", // New Zealand Dollar - Новозеландский доллар
		CurrencyNOK: "NOK", // Norway Krone - Норвежская крона
		CurrencyPLN: "PLN", // Polish Zloty - Польский злотый
		CurrencyPTE: "PTE", // Portuguese Escudo - Португальское эскудо
		CurrencyROL: "ROL", // Romania Leu - Румынский лей
		CurrencyRUB: "RUB", // Russian Rouble - Российский рубль (новый)
		CurrencySGD: "SGD", // Singapore Dollar - Сингапурский доллар
		CurrencySKK: "SKK", // Slovak Koruna - Словацкая крона
		CurrencyZAR: "ZAR", // South African Rand - Южноафриканский ранд
		CurrencyESP: "ESP", // Spanish Peseta - Испанская песета
		CurrencySEK: "SEK", // Swedish Krona - Шведская крона
		CurrencyCHF: "CHF", // Swiss Franc - Швейцарский франк
		CurrencyTRY: "TRY", // Turkish Lira - Турецкая лира
		CurrencyUAH: "UAH", // Ukraine Hryvnia - Украинская гривна
		CurrencyUSD: "USD", // United States Dollar - Американский доллар
		CurrencyUZS: "UZS", // Uzbekistan Sum - Узбекский сум
		CurrencyXAG: "XAG", // Silver - Серебро
		CurrencyXAU: "XAU", // Gold - Золото
		CurrencyXPD: "XPD", // Palladium - Палладий
		CurrencyXPT: "XPT", // Platinum - Платина
		CurrencyXXX: "XXX", // Not currency (Token) - Нет валюты (жетон)
	}
	value, exists := data[e]
	if exists {
		return value
	}
	return "???"
}

// Precision returns an amount of minor units in the currency unit
func (e Currency) Precision() int {
	data := map[Currency]int{
		CurrencyAMD: 100,  // Armenian Dram - Армянский драм
		CurrencyAUD: 100,  // Australian Dollar - Австралийский доллар
		CurrencyATS: 100,  // Austrian Schilling - Австрийский шиллинг
		CurrencyAZN: 100,  // Azerbaijan Manat - Азербайджанский манат
		CurrencyBYR: 100,  // Belarusian Ruble - Белорусский рубль
		CurrencyBEF: 100,  // Belgian Franc - Бельгийский франк
		CurrencyBRL: 100,  // Brazilian Real - Бразильский реал
		CurrencyGBP: 100,  // British Pound - Британский фунт
		CurrencyCAD: 100,  // Canadian Dollar - Канадский доллар
		CurrencyCNY: 10,   // China Yuan - Китайский юань
		CurrencyCZK: 100,  // Czech Koruna - Чешская крона
		CurrencyDKK: 100,  // Danish Krone - Датская крона
		CurrencyNLG: 100,  // Dutch Guilder - Нидерландский гульден
		CurrencyEEK: 100,  // Estonian Kroon - Эстонская крона
		CurrencyEUR: 100,  // European Euro - Единая европейская валюта
		CurrencyFIM: 100,  // Finnish Mark - Финская марка
		CurrencyFRF: 100,  // French Franc - Французский франк
		CurrencyGEL: 100,  // Georgia Lari - Грузинская лари
		CurrencyDEM: 100,  // German Mark - Немецкая марка
		CurrencyGRD: 100,  // Greek Drachma - Греческая драхма
		CurrencyHKD: 100,  // Hong Kong Dollar - Гонконгский доллар
		CurrencyHUF: 100,  // Hungarian Forint - Венгерский форинт
		CurrencyINR: 100,  // Indian Rupee - Индийская рупия
		CurrencyIEP: 100,  // Irish Punt - Ирландский фунт
		CurrencyILS: 100,  // Israeli Sheqel - Израильский шекель
		CurrencyITL: 100,  // Italian Lira - Итальянская лира
		CurrencyJPY: 100,  // Japanese Yen - Японская йена
		CurrencyKZT: 100,  // Kazakhstan Tenge - Казахский тенге
		CurrencyKGS: 100,  // Kyrgyzstan Som - Киргизский сом
		CurrencyLVL: 100,  // Latvian Lat - Латвийский лат
		CurrencyLTL: 100,  // Lithuanian Lita - Литовский лит
		CurrencyMDL: 100,  // Moldovan Leu - Молдавский лей
		CurrencyMXN: 100,  // Mexican Peso - Мексиканский песо
		CurrencyNZD: 100,  // New Zealand Dollar - Новозеландский доллар
		CurrencyNOK: 100,  // Norway Krone - Норвежская крона
		CurrencyPLN: 100,  // Polish Zloty - Польский злотый
		CurrencyPTE: 100,  // Portuguese Escudo - Португальское эскудо
		CurrencyROL: 100,  // Romania Leu - Румынский лей
		CurrencyRUB: 100,  // Russian Rouble - Российский рубль (новый)
		CurrencySGD: 100,  // Singapore Dollar - Сингапурский доллар
		CurrencySKK: 100,  // Slovak Koruna - Словацкая крона
		CurrencyZAR: 100,  // South African Rand - Южноафриканский ранд
		CurrencyESP: 100,  // Spanish Peseta - Испанская песета
		CurrencySEK: 100,  // Swedish Krona - Шведская крона
		CurrencyCHF: 100,  // Swiss Franc - Швейцарский франк
		CurrencyTRY: 100,  // Turkish Lira - Турецкая лира
		CurrencyUAH: 100,  // Ukraine Hryvnia - Украинская гривна
		CurrencyUSD: 100,  // United States Dollar - Американский доллар
		CurrencyUZS: 100,  // Uzbekistan Sum - Узбекский сум
		CurrencyXAG: 1000, // Silver - Серебро
		CurrencyXAU: 1000, // Gold - Золото
		CurrencyXPD: 1000, // Palladium - Палладий
		CurrencyXPT: 1000, // Platinum - Платина
		CurrencyXXX: 1,    // Not currency (Token) - Нет валюты (жетон)
	}
	value, exists := data[e]
	if exists {
		return value
	}
	return 1
}

func AmountText(value Amount, curr Currency) string {
	precision := curr.Precision()
	if precision == 0 || precision == 1 {
		return strconv.Itoa(int(value))
	}
	maxUnit := int(value) / precision
	minUnit := int(value) % precision
	switch precision {
	case 10:
		return fmt.Sprintf("%d.%d", maxUnit, minUnit)
	case 100:
		return fmt.Sprintf("%d.%2d", maxUnit, minUnit)
	case 1000:
		return fmt.Sprintf("%d.%3d", maxUnit, minUnit)
	default:
		return strconv.Itoa(int(value))
	}
}
