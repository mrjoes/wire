/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// OriginatorOptionF OptionF is originator option F
type OriginatorOptionF struct {
	// PartyIdentifier  Must be one of the following two formats: 1. /Account Number (slash followed by at least one valid non-space character:  e.g., /123456)  2. Unique Identifier/ (4 character code followed by a slash and at least one valid non-space character:    e.g., SOSE/123-456-789) ARNU: Alien Registration Number CCPT: Passport Number CUST: Customer Identification Number  DRLC/    Driver’s License Number  EMPL/    Employer Number NIDN: National Identify Number  SOSE/    Social Security Number TXID: Tax Identification Number
	PartyIdentifier string `json:"partyIdentifier,omitempty"`
	// Name  Format:  Must begin with Line Code 1 followed by a slash and at least one valid non-space character: e.g., 1/SMITH JOHN.
	Name string `json:"name,omitempty"`
	// LineOne  Format: Must begin with one of the following Line Codes followed by a slash and at least one valid non-space character. 1 Name 2 Address 3 Country and Town 4 Date of Birth 5 Place of Birth 6 Customer Identification Number 7 National Identity Number 8 Additional Information  For example: 2/123 MAIN STREET 3/US/NEW YORK, NY 10000 7/111-22-3456
	LineOne string `json:"lineOne,omitempty"`
	// LineTwo  Format: Must begin with one of the following Line Codes followed by a slash and at least one valid non-space character. 1 Name 2 Address 3 Country and Town 4 Date of Birth 5 Place of Birth 6 Customer Identification Number 7 National Identity Number 8 Additional Information  For example: 2/123 MAIN STREET 3/US/NEW YORK, NY 10000 7/111-22-3456
	LineTwo string `json:"lineTwo,omitempty"`
	// LineThree  Format: Must begin with one of the following Line Codes followed by a slash and at least one valid non-space character. 1 Name 2 Address 3 Country and Town 4 Date of Birth 5 Place of Birth 6 Customer Identification Number 7 National Identity Number 8 Additional Information  For example: 2/123 MAIN STREET 3/US/NEW YORK, NY 10000 7/111-22-3456
	LineThree string `json:"lineThree,omitempty"`
}
