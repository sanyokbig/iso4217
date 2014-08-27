package iso4217

var names = map[int]string{
	0:   "",
	8:   "ALL",
	12:  "DZD",
	32:  "ARS",
	36:  "AUD",
	44:  "BSD",
	48:  "BHD",
	50:  "BDT",
	51:  "AMD",
	52:  "BBD",
	60:  "BMD",
	64:  "BTN",
	68:  "BOB",
	72:  "BWP",
	84:  "BZD",
	90:  "SBD",
	96:  "BND",
	104: "MMK",
	108: "BIF",
	116: "KHR",
	124: "CAD",
	132: "CVE",
	136: "KYD",
	144: "LKR",
	152: "CLP",
	156: "CNY",
	170: "COP",
	174: "KMF",
	188: "CRC",
	191: "HRK",
	192: "CUP",
	203: "CZK",
	208: "DKK",
	214: "DOP",
	222: "SVC",
	230: "ETB",
	232: "ERN",
	238: "FKP",
	242: "FJD",
	262: "DJF",
	270: "GMD",
	292: "GIP",
	320: "GTQ",
	324: "GNF",
	328: "GYD",
	332: "HTG",
	340: "HNL",
	344: "HKD",
	348: "HUF",
	352: "ISK",
	356: "INR",
	360: "IDR",
	364: "IRR",
	368: "IQD",
	376: "ILS",
	388: "JMD",
	392: "JPY",
	398: "KZT",
	400: "JOD",
	404: "KES",
	408: "KPW",
	410: "KRW",
	414: "KWD",
	417: "KGS",
	418: "LAK",
	422: "LBP",
	426: "LSL",
	430: "LRD",
	434: "LYD",
	440: "LTL",
	446: "MOP",
	454: "MWK",
	458: "MYR",
	462: "MVR",
	478: "MRO",
	480: "MUR",
	484: "MXN",
	496: "MNT",
	498: "MDL",
	504: "MAD",
	512: "OMR",
	516: "NAD",
	524: "NPR",
	532: "ANG",
	533: "AWG",
	548: "VUV",
	554: "NZD",
	558: "NIO",
	566: "NGN",
	578: "NOK",
	586: "PKR",
	590: "PAB",
	598: "PGK",
	600: "PYG",
	604: "PEN",
	608: "PHP",
	634: "QAR",
	643: "RUB",
	646: "RWF",
	654: "SHP",
	678: "STD",
	682: "SAR",
	690: "SCR",
	694: "SLL",
	702: "SGD",
	704: "VND",
	706: "SOS",
	710: "ZAR",
	728: "SSP",
	748: "SZL",
	752: "SEK",
	756: "CHF",
	760: "SYP",
	764: "THB",
	776: "TOP",
	780: "TTD",
	784: "AED",
	788: "TND",
	800: "UGX",
	807: "MKD",
	818: "EGP",
	826: "GBP",
	834: "TZS",
	840: "USD",
	858: "UYU",
	860: "UZS",
	882: "WST",
	886: "YER",
	901: "TWD",
	931: "CUC",
	932: "ZWL",
	934: "TMT",
	936: "GHS",
	937: "VEF",
	938: "SDG",
	940: "UYI",
	941: "RSD",
	943: "MZN",
	944: "AZN",
	946: "RON",
	947: "CHE",
	948: "CHW",
	949: "TRY",
	950: "XAF",
	951: "XCD",
	952: "XOF",
	953: "XPF",
	955: "XBA",
	956: "XBB",
	957: "XBC",
	958: "XBD",
	959: "XAU",
	960: "XDR",
	961: "XAG",
	962: "XPT",
	963: "XTS",
	964: "XPD",
	965: "XUA",
	967: "ZMW",
	968: "SRD",
	969: "MGA",
	970: "COU",
	971: "AFN",
	972: "TJS",
	973: "AOA",
	974: "BYR",
	975: "BGN",
	976: "CDF",
	977: "BAM",
	978: "EUR",
	979: "MXV",
	980: "UAH",
	981: "GEL",
	984: "BOV",
	985: "PLN",
	986: "BRL",
	990: "CLF",
	994: "XSU",
	997: "USN",
	999: "XXX",
}

var minorUnits = map[int]int{
	0:   0,
	8:   2,
	12:  2,
	32:  2,
	36:  2,
	44:  2,
	48:  3,
	50:  2,
	51:  2,
	52:  2,
	60:  2,
	64:  2,
	68:  2,
	72:  2,
	84:  2,
	90:  2,
	96:  2,
	104: 2,
	108: 0,
	116: 2,
	124: 2,
	132: 2,
	136: 2,
	144: 2,
	152: 0,
	156: 2,
	170: 2,
	174: 0,
	188: 2,
	191: 2,
	192: 2,
	203: 2,
	208: 2,
	214: 2,
	222: 2,
	230: 2,
	232: 2,
	238: 2,
	242: 2,
	262: 0,
	270: 2,
	292: 2,
	320: 2,
	324: 0,
	328: 2,
	332: 2,
	340: 2,
	344: 2,
	348: 2,
	352: 0,
	356: 2,
	360: 2,
	364: 2,
	368: 3,
	376: 2,
	388: 2,
	392: 0,
	398: 2,
	400: 3,
	404: 2,
	408: 2,
	410: 0,
	414: 3,
	417: 2,
	418: 2,
	422: 2,
	426: 2,
	430: 2,
	434: 3,
	440: 2,
	446: 2,
	454: 2,
	458: 2,
	462: 2,
	478: 2,
	480: 2,
	484: 2,
	496: 2,
	498: 2,
	504: 2,
	512: 3,
	516: 2,
	524: 2,
	532: 2,
	533: 2,
	548: 0,
	554: 2,
	558: 2,
	566: 2,
	578: 2,
	586: 2,
	590: 2,
	598: 2,
	600: 0,
	604: 2,
	608: 2,
	634: 2,
	643: 2,
	646: 0,
	654: 2,
	678: 2,
	682: 2,
	690: 2,
	694: 2,
	702: 2,
	704: 0,
	706: 2,
	710: 2,
	728: 2,
	748: 2,
	752: 2,
	756: 2,
	760: 2,
	764: 2,
	776: 2,
	780: 2,
	784: 2,
	788: 3,
	800: 0,
	807: 2,
	818: 2,
	826: 2,
	834: 2,
	840: 2,
	858: 2,
	860: 2,
	882: 2,
	886: 2,
	901: 2,
	931: 2,
	932: 2,
	934: 2,
	936: 2,
	937: 2,
	938: 2,
	940: 0,
	941: 2,
	943: 2,
	944: 2,
	946: 2,
	947: 2,
	948: 2,
	949: 2,
	950: 0,
	951: 2,
	952: 0,
	953: 0,
	955: 0,
	956: 0,
	957: 0,
	958: 0,
	959: 0,
	960: 0,
	961: 0,
	962: 0,
	963: 0,
	964: 0,
	965: 0,
	967: 2,
	968: 2,
	969: 2,
	970: 2,
	971: 2,
	972: 2,
	973: 2,
	974: 0,
	975: 2,
	976: 2,
	977: 2,
	978: 2,
	979: 2,
	980: 2,
	981: 2,
	984: 2,
	985: 2,
	986: 2,
	990: 4,
	994: 0,
	997: 2,
	999: 0,
}

func ByCode(n int) (string, int) {
	return names[n], minorUnits[n]
}
