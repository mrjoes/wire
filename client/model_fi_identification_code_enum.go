/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// FiIdentificationCodeEnum : IdentificationCode:  * `B` - SWIFT Bank Identifier Code (BIC) * `C` - CHIPS Participant * `D` - Demand Deposit Account (DDA) Number * `F` - Fed Routing Number * `T` - SWIFT BIC or Bank Entity Identifier (BEI) and Account Number * `U` - CHIPS Identifier 
type FiIdentificationCodeEnum string

// List of FIIdentificationCodeEnum
const (
	B FiIdentificationCodeEnum = "B"
	C FiIdentificationCodeEnum = "C"
	D FiIdentificationCodeEnum = "D"
	F FiIdentificationCodeEnum = "F"
	T FiIdentificationCodeEnum = "T"
	U FiIdentificationCodeEnum = "U"
)