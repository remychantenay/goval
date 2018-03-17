package goval

const ArgTypeString = "string"
const ArgTypeNumber = "number"
const ArgTypeEmail = "email"
const ArgTypeUuid = "uuid"
const ArgTypeCountryCode = "country_code" // ISO-Alpha-2
const ArgTypeCurrency = "currency"        // ISO 4217
const ArgTypeEnum = "enum"

const ArgConstraintMax = "max="
const ArgConstraintMin = "min="
const ArgConstraintRequired = "required="
const ArgConstraintDomain = "domain=" // For email addresses (e.g. @google.com)
const ArgConstraintExcludeEu = "exclude_eu="
const ArgConstraintExclude = "exclude=" // Exclusion parameters must be separated by a pipe | e.g. exclude=EUR|GBP
const ArgConstraintValues = "values="