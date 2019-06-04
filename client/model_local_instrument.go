/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// LocalInstrument is the local instrument
type LocalInstrument struct {
	// LocalInstrument  * `ANSI` - ANSI X12 format * `COVS` - Sequence B Cover Payment Structured * `GXML` - General XML format * `IXML` - ISO 20022 XML formaT * `NARR` - Narrative Text * `PROP` - Proprietary Local Instrument Code * `RMTS` - Remittance Information Structured * `RRMT` - Related Remittance Information * `S820` - STP 820 format * `SWIF` - SWIFT field 70 (Remittance Information) * `UEDI` - UN/EDIFACT format 
	LocalInstrumentCode string `json:"localInstrumentCode,omitempty"`
	// ProprietaryCode
	ProprietaryCode string `json:"proprietaryCode,omitempty"`
}