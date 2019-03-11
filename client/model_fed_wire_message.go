/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// FEDWireMessage
type FedWireMessage struct {
	SenderSuppliedInfo SenderSuppliedInfo `json:"senderSuppliedInfo"`
	TypeSubType TypeSubType `json:"typeSubType"`
	InputMessageAccountabilityData InputMessageAccountabilityData `json:"inputMessageAccountabilityData"`
	Amount Amount `json:"amount"`
	SenderDepositoryInstitution SenderDepositoryInstitution `json:"senderDepositoryInstitution"`
	ReceiverDepositoryInstitution ReceiverDepositoryInstitution `json:"receiverDepositoryInstitution"`
	BusinessFunctionCode BusinessFunctionCode `json:"businessFunctionCode"`
	IntermediaryFinancialInstitution FinancialInstitution `json:"intermediaryFinancialInstitution,omitempty"`
	BeneficiaryFinancialInstitution FinancialInstitution `json:"beneficiaryFinancialInstitution,omitempty"`
	Beneficiary Personal `json:"beneficiary,omitempty"`
	// ReferenceForBeneficiary {4320}
	ReferenceForBeneficiary string `json:"referenceForBeneficiary,omitempty"`
	AccountDebitedDrawdown AccountDebitedDrawdown `json:"accountDebitedDrawdown,omitempty"`
	Originator Personal `json:"originator,omitempty"`
	// OriginatorOptionF {5010}
	OriginatorOptionF map[string]interface{} `json:"originatorOptionF,omitempty"`
	OriginatorFinancialInstitution FinancialInstitution `json:"originatorFinancialInstitution,omitempty"`
	InstructingFinancialInstitution FinancialInstitution `json:"instructingFinancialInstitution,omitempty"`
	AccountCreditedDrawdown AccountCreditedDrawdown `json:"accountCreditedDrawdown,omitempty"`
	OriginatorToBeneficiaryInfo OriginatorToBeneficiaryInfo `json:"originatorToBeneficiaryInfo,omitempty"`
	ReceiverFinancialInstitutionInfo FiToFiInfo `json:"receiverFinancialInstitutionInfo,omitempty"`
	DrawdownDebitAccountAdviceInfo AdviceInfo `json:"drawdownDebitAccountAdviceInfo,omitempty"`
	IntermediaryFinancialInstitutionInfo FiToFiInfo `json:"intermediaryFinancialInstitutionInfo,omitempty"`
	IntermediaryFinacialInstitutionAdviceInfo AdviceInfo `json:"intermediaryFinacialInstitutionAdviceInfo,omitempty"`
	BeneficiaryFinancialInstitutionInfo FiToFiInfo `json:"beneficiaryFinancialInstitutionInfo,omitempty"`
	BeneficiaryFinancialInstitutionAdviceInfo AdviceInfo `json:"beneficiaryFinancialInstitutionAdviceInfo,omitempty"`
	BeneficiaryInfo FiToFiInfo `json:"beneficiaryInfo,omitempty"`
	BeneficiaryAdviceInfo AdviceInfo `json:"beneficiaryAdviceInfo,omitempty"`
	// PaymentMethodToBeneficiary
	PaymentMethodToBeneficiary map[string]interface{} `json:"paymentMethodToBeneficiary,omitempty"`
	// AdditionalFIToFIInfo
	AdditionalFIToFIInfo map[string]interface{} `json:"additionalFIToFIInfo,omitempty"`
}
