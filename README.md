# GoTimeZoneMapper
[![Maintainability](https://api.codeclimate.com/v1/badges/85a98cfb0a056bbc4d4e/maintainability)](https://codeclimate.com/github/SiddhantAgarwal/GoTimezoneMapper/maintainability)
[![Coverage Status](https://coveralls.io/repos/github/SiddhantAgarwal/GoTimezoneMapper/badge.svg?branch=master)](https://coveralls.io/github/SiddhantAgarwal/GoTimezoneMapper?branch=master)
![CI](https://github.com/SiddhantAgarwal/GoTimezoneMapper/workflows/CI/badge.svg?branch=master)

A library to lookup [Country-Code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) (ISO 3166-1 alpha-2) (ISO ) from [Timezone](https://en.wikipedia.org/wiki/Tz_database) (tz database)

Example Lookup
```
America/New_York => US
```

How to use
```
code, err := mapper.FindCountryCodeForTimezone("America/New_York")
if err != nil {
	fmt.Println(err)
}
fmt.Println(code)
```