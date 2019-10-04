/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CurrencyInstructedAmount CurrencyInstructedAmount
type CurrencyInstructedAmount struct {
	// SwiftFieldTag
	SwiftFieldTag string `json:"swiftFieldTag,omitempty"`
	// Amount
	Amount string `json:"amount,omitempty"`
}
