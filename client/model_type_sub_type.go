/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TypeSubType TypeSubtype is the type and sub type codes
type TypeSubType struct {
	// TypeCode:  * `10` - Funds Transfer - A funds transfer in which the sender and/or receiver may be a bank or a third party (i.e., customer of a bank). * `15` - Foreign Transfer - A funds transfer to or from a foreign central bank or government or international organization with an account at the Federal Reserve Bank of New York. * `16` - Settlement Transfer - A funds transfer between Fedwire Funds Service participants.
	TypeCode string `json:"typeCode"`
	// SubTypeCode:  * `00` - Basic Funds Transfer - A basic value funds transfer. * `01` - Request for Reversal - A non-value request for reversal of a funds transfer originated on the current business day. * `02` - Reversal of Transfer - A value reversal of a funds transfer received on the current business day.  May be used in response to a subtype code ‘01’ Request for Reversal. * `07` - Request for Reversal of a Prior Day Transfer - A non-value request for a reversal of a funds transfer originated on a prior business day. * `08` - Reversal of a Prior Day Transfer - A value reversal of a funds transfer received on a prior business day.  May be used in response to a subtype code ‘07’ Request for Reversal of a Prior Day Transfer. * `31` - Request for Credit (Drawdown) - A non-value request for the receiver to send a funds transfer to a designated party. * `32` - Funds Transfer Honoring a Request for Credit (Drawdown) -  A value funds transfer honoring a subtype 31 request for credit. * `33` -Refusal to Honor a Request for Credit (Drawdown) - A non-value message indicating refusal to honor a subtype 31 request for credit. * `90` - Service Message - A non-value message used to communicate questions and information that is not covered by a specific subtype.
	SubTypeCode string `json:"subTypeCode"`
}
