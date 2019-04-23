// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// ToDo: Change FedWireMessage to FED...
// ToDo: Remove empty functions after all validation checks are coded

package wire

import "strings"

// FedWireMessage is a FedWire Message
type FedWireMessage struct {
	// ID
	ID string `json:"id"`
	// MessageDisposition
	MessageDisposition *MessageDisposition `json:"messageDisposition,omitempty"`
	// ReceiptTimeStamp
	ReceiptTimeStamp *ReceiptTimeStamp `json:"receiptTimeStamp,omitempty"`
	// OutputMessageAccountabilityData (OMAD)
	OutputMessageAccountabilityData *OutputMessageAccountabilityData `json:"outputMessageAccountabilityData,omitempty"`
	// ErrorWire
	ErrorWire *ErrorWire `json:"errorWire,omitempty"`
	// SenderSuppliedInformation
	SenderSupplied *SenderSupplied `json:"senderSupplied"`
	// TypeSubType
	TypeSubType *TypeSubType `json:"typeSubType"`
	// InputMessageAccountabilityData (IMAD)
	InputMessageAccountabilityData *InputMessageAccountabilityData `json:"inputMessageAccountabilityData"`
	// Amount (up to a penny less than $10 billion)
	Amount *Amount `json:"amount"`
	// SenderDepositoryInstitution
	SenderDepositoryInstitution *SenderDepositoryInstitution `json:"senderDepositoryInstitution"`
	// ReceiverDepositoryInstitution
	ReceiverDepositoryInstitution *ReceiverDepositoryInstitution `json:"receiverDepositoryInstitution"`
	// BusinessFunctionCode
	BusinessFunctionCode *BusinessFunctionCode `json:"businessFunctionCode"`
	// SenderReference
	SenderReference *SenderReference `json:"senderReference,omitempty"`
	// PreviousMessageIdentifier
	PreviousMessageIdentifier *PreviousMessageIdentifier `json:"previousMessageIdentifier,omitempty"`
	// LocalInstrument
	LocalInstrument *LocalInstrument `json:"localInstrument,omitempty"`
	// PaymentNotification
	PaymentNotification *PaymentNotification `json:"paymentNotification,omitempty"`
	// Charges
	Charges *Charges `json:"charges,omitempty"`
	// InstructedAmount
	InstructedAmount *InstructedAmount `json:"instructedAmount,omitempty"`
	// ExchangeRate
	ExchangeRate *ExchangeRate `json:"exchangeRate,omitempty"`
	// BeneficiaryIntermediaryFI
	BeneficiaryIntermediaryFI *BeneficiaryIntermediaryFI `json:"beneficiaryIntermediaryFI,omitempty"`
	// BeneficiaryFI
	BeneficiaryFI *BeneficiaryFI `json:"beneficiaryFI,omitempty"`
	// Beneficiary
	Beneficiary *Beneficiary `json:"beneficiary,omitempty"`
	// BeneficiaryReference
	BeneficiaryReference *BeneficiaryReference `json:"beneficiaryReference,omitempty"`
	// AccountDebitedDrawdown
	AccountDebitedDrawdown *AccountDebitedDrawdown `json:"accountDebitedDrawdown,omitempty"`
	// Originator
	Originator *Originator `json:"originator,omitempty"`
	// OriginatorOptionF
	OriginatorOptionF *OriginatorOptionF `json:"originatorOptionF,omitempty"`
	// OriginatorFI
	OriginatorFI *OriginatorFI `json:"originatorFI,omitempty"`
	// InstructingFI
	InstructingFI *InstructingFI `json:"instructingFI,omitempty"`
	// AccountCreditedDrawdown
	AccountCreditedDrawdown *AccountCreditedDrawdown `json:"accountCreditedDrawdown,omitempty"`
	// OriginatorToBeneficiary
	OriginatorToBeneficiary *OriginatorToBeneficiary `json:"originatorToBeneficiary,omitempty"`
	// FIReceiverFI
	FIReceiverFI *FIReceiverFI `json:"fiReceiverFI,omitempty"`
	// FIDrawdownDebitAccountAdvice
	FIDrawdownDebitAccountAdvice *FIDrawdownDebitAccountAdvice `json:"fiDrawdownDebitAccountAdvice,omitempty"`
	// FIIntermediaryFI
	FIIntermediaryFI *FIIntermediaryFI `json:"fiIntermediaryFI,omitempty"`
	// FIIntermediaryFIAdvice
	FIIntermediaryFIAdvice *FIIntermediaryFIAdvice `json:"fiIntermediaryFIAdvice,omitempty"`
	// FIBeneficiaryFI
	FIBeneficiaryFI *FIBeneficiaryFI `json:"fiBeneficiaryFI,omitempty"`
	// FIBeneficiaryFIAdvice
	FIBeneficiaryFIAdvice *FIBeneficiaryFIAdvice `json:"fiBeneficiaryFIAdvice,omitempty"`
	// FIBeneficiary

	FIBeneficiary *FIBeneficiary `json:"fiBeneficiary,omitempty"`
	// FIBeneficiaryAdvice
	FIBeneficiaryAdvice *FIBeneficiaryAdvice `json:"fiBeneficiaryAdvice,omitempty"`
	// FIPaymentMethodToBeneficiary
	FIPaymentMethodToBeneficiary *FIPaymentMethodToBeneficiary `json:"fiPaymentMethodToBeneficiary,omitempty"`
	// FIAdditionalFIToFI
	FIAdditionalFIToFI *FIAdditionalFIToFI `json:"fiAdditionalFiToFi,omitempty"`
	// CurrencyInstructedAmount
	CurrencyInstructedAmount *CurrencyInstructedAmount `json:"currencyInstructedAmount,omitempty"`
	// OrderingCustomer
	OrderingCustomer *OrderingCustomer `json:"orderingCustomer,omitempty"`
	// OrderingInstitution
	OrderingInstitution *OrderingInstitution `json:"orderingInstitution,omitempty"`
	// IntermediaryInstitution
	IntermediaryInstitution *IntermediaryInstitution `json:"intermediaryInstitution,omitempty"`
	// InstitutionAccount
	InstitutionAccount *InstitutionAccount `json:"institutionAccount,omitempty"`
	// BeneficiaryCustomer
	BeneficiaryCustomer *BeneficiaryCustomer `json:"beneficiaryCustomer,omitempty"`
	// Remittance
	Remittance *Remittance `json:"remittance,omitempty"`
	// SenderToReceiver
	SenderToReceiver *SenderToReceiver `json:"senderToReceiver,omitempty"`
	// UnstructuredAddenda
	UnstructuredAddenda *UnstructuredAddenda `json:"unstructuredAddenda,omitempty"`
	// RelatedRemittance
	RelatedRemittance *RelatedRemittance `json:"relatedRemittance,omitempty"`
	// RemittanceOriginator
	RemittanceOriginator *RemittanceOriginator `json:"remittanceOriginator,omitempty"`
	// RemittanceBeneficiary
	RemittanceBeneficiary *RemittanceBeneficiary `json:"remittanceBeneficiary,omitempty"`
	// PrimaryRemittanceDocument
	PrimaryRemittanceDocument *PrimaryRemittanceDocument `json:"primaryRemittanceDocument,omitempty"`
	// ActualAmountPaid
	ActualAmountPaid *ActualAmountPaid `json:"actualAmountPaid,omitempty"`
	// GrossAmountRemittanceDocument
	GrossAmountRemittanceDocument *GrossAmountRemittanceDocument `json:"grossAmountRemittanceDocument,omitempty"`
	// AmountNegotiatedDiscount
	AmountNegotiatedDiscount *AmountNegotiatedDiscount `json:"amountNegotiatedDiscount,omitempty"`
	// Adjustment
	Adjustment *Adjustment `json:"adjustment,omitempty"`
	// DateRemittanceDocument
	DateRemittanceDocument *DateRemittanceDocument `json:"dateRemittanceDocument,omitempty"`
	// SecondaryRemittanceDocument
	SecondaryRemittanceDocument *SecondaryRemittanceDocument `json:"secondaryRemittanceDocument,omitempty"`
	// RemittanceFreeText
	RemittanceFreeText *RemittanceFreeText `json:"remittanceFreeText,omitempty"`
	// ServiceMessage
	ServiceMessage *ServiceMessage `json:"serviceMessage,omitempty"`
}

// NewFedWireMessage returns a new FedWireMessage
func NewFedWireMessage() FedWireMessage {
	fwm := FedWireMessage{}
	return fwm
}

// verify checks basic valid NACHA batch rules. Assumes properly parsed records. This does not mean it is a valid batch as validity is tied to each batch type
func (fwm *FedWireMessage) verify() error {

	if err := fwm.isMandatory(); err != nil {
		return err
	}
	if err := fwm.isBusinessCodeValid(); err != nil {
		return err
	}
	if err := fwm.isAmountValid(); err != nil {
		return err
	}
	if fwm.PreviousMessageIdentifier != nil {
		if err := fwm.isPreviousMessageIdentifierValid(); err != nil {
			return err
		}
	}
	if fwm.LocalInstrument != nil {
		if err := fwm.isLocalInstrumentCodeValid(); err != nil {
			return err
		}
	}
	if fwm.PaymentNotification != nil {
		if err := fwm.isPaymentNotificationValid(); err != nil {
			return err
		}
	}
	if  fwm.Charges != nil {
		if err := fwm.isChargesValid(); err != nil {
			return err
		}
	}
	if fwm.InstructedAmount != nil {
		if err := fwm.isInstructedAmountValid(); err != nil {
			return err
		}
	}
	if fwm.ExchangeRate != nil {
		if err := fwm.isExchangeRateValid(); err != nil {
			return err
		}
	}
	if fwm.BeneficiaryIntermediaryFI != nil {
		if err := fwm.isBeneficiaryIntermediaryFIValid(); err != nil {
			return err
		}
	}
	if fwm.BeneficiaryFI != nil {
		if err := fwm.isBeneficiaryFIValid(); err != nil {
			return err
		}
	}
	if fwm.OriginatorFI != nil {
		if err := fwm.isOriginatorFIValid(); err != nil {
			return err
		}
	}
	if fwm.InstructingFI != nil {
		if err := fwm.isInstructingFIValid(); err != nil {
			return err
		}
	}
	if fwm.OriginatorToBeneficiary != nil {
		if err := fwm.isOriginatorToBeneficiaryValid(); err != nil {
			return err
		}
	}
	if fwm.FIIntermediaryFI != nil {
		if err := fwm.isFIIntermediaryFIValid(); err != nil {
			return err
		}
	}
	if fwm.FIIntermediaryFIAdvice != nil {
		if err := fwm.isFIIntermediaryFIAdviceValid(); err != nil {
			return err
		}
	}
	if fwm.FIBeneficiaryFI != nil {
		if err := fwm.isFIBeneficiaryFIValid(); err != nil {
			return err
		}
	}
	if fwm.FIBeneficiaryFIAdvice != nil {
		if err := fwm.isFIBeneficiaryFIAdviceValid(); err != nil {
			return err
		}
	}
	if fwm.FIBeneficiary != nil {
		if err := fwm.isFIBeneficiaryValid(); err != nil {
			return err
		}
	}
	if fwm.FIBeneficiaryAdvice != nil {
		if err := fwm.isFIBeneficiaryAdviceValid(); err != nil {
			return err
		}
	}
	if fwm.FIPaymentMethodToBeneficiary != nil {
		if err := fwm.isFIPaymentMethodToBeneficiaryValid(); err != nil {
			return err
		}
	}
	if fwm.UnstructuredAddenda != nil {
		if err := fwm.isUnstructuredAddendaValid(); err != nil {
			return err
		}
	}
	if fwm.RelatedRemittance != nil {
		if err := fwm.isRelatedRemittanceValid(); err != nil {
			return err
		}
	}
	if fwm.RemittanceOriginator != nil {
		if err := fwm.isRemittanceOriginatorValid(); err != nil {
			return err
		}
	}
	if fwm.RemittanceBeneficiary != nil {
		if err := fwm.isRemittanceBeneficiaryValid(); err != nil {
			return err
		}
	}
	if fwm.PrimaryRemittanceDocument != nil {
		if err := fwm.isPrimaryRemittanceDocumentValid(); err != nil {
			return err
		}
	}
	if fwm.ActualAmountPaid != nil {
		if err := fwm.isActualAmountPaidValid(); err != nil {
			return err
		}
	}
	if fwm.GrossAmountRemittanceDocument != nil {
		if err := fwm.isGrossAmountRemittanceDocumentValid(); err != nil {
			return err
		}
	}
	if fwm.Adjustment != nil {
		if err := fwm.isAdjustmentValid(); err != nil {
			return err
		}
	}
	if fwm.DateRemittanceDocument != nil {
		if err := fwm.isDateRemittanceDocumentValid(); err != nil {
			return err
		}
	}
	if fwm.RemittanceFreeText != nil {
		if err := fwm.isRemittanceFreeTextValid(); err != nil {
			return err
		}
	}
	return nil
}

// isMandatory validates mandatory tags for a FedWireMessage are defined
func (fwm *FedWireMessage) isMandatory() error {
	if fwm.SenderSupplied == nil {
		return fieldError("SenderSupplied", ErrFieldRequired)
	}
	if fwm.TypeSubType == nil {
		return fieldError("TypeSubType", ErrFieldRequired)
	}
	if fwm.InputMessageAccountabilityData == nil {
		return fieldError("InputMessageAccountabilityData", ErrFieldRequired)
	}
	if fwm.Amount == nil {
		return fieldError("Amount", ErrFieldRequired)
	}
	if fwm.SenderDepositoryInstitution == nil {
		return fieldError("SenderDepositoryInstitution", ErrFieldRequired)
	}
	if fwm.ReceiverDepositoryInstitution == nil {
		return fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}
	if fwm.BusinessFunctionCode == nil {
		return fieldError("BusinessFunctionCode", ErrFieldRequired)
	}
	return nil
}

func (fwm *FedWireMessage) isBusinessCodeValid() error {
	switch fwm.BusinessFunctionCode.BusinessFunctionCode {
	case BankTransfer:
		if err := fwm.isBankTransferValid(); err != nil {
			return err
		}
		if err := fwm.isBankTransferTags(); err != nil {
			return err
		}
		if err := fwm.isInvalidBankTransferTags(); err != nil {
			return err
		}
	case CustomerTransfer:
		if err := fwm.isCustomerTransferValid(); err != nil {
			return err
		}
		if err := fwm.isCustomerTransferTags(); err != nil {
			return err
		}
		if err := fwm.isInvalidCustomerTransferTags(); err != nil {
			return err
		}
	case CustomerTransferPlus:
		if err := fwm.isCustomerTransferPlusValid(); err != nil {
			return err
		}
		if err := fwm.isCustomerTransferPlusTags(); err != nil {
			return err
		}
		if err := fwm.isInvalidCustomerTransferPlusTags(); err != nil {
			return err
		}
	case CheckSameDaySettlement:
		if err := fwm.isCheckSameDaySettlementValid(); err != nil {
			return err
		}
		if err := fwm.isCheckSameDaySettlementTags(); err != nil {
			return err
		}
		if err := fwm.isInvalidTags(); err != nil {
			return err
		}
	case DepositSendersAccount:
		if err := fwm.isDepositSendersAccountValid(); err != nil {
			return err
		}
		if err := fwm.isDepositSendersAccountTags(); err != nil {
			return err
		}
		if err := fwm.isInvalidTags(); err != nil {
			return err
		}
	case FEDFundsReturned:
		if err := fwm.isFEDFundsReturnedValid(); err != nil {
			return err
		}
		if err := fwm.isFEDFundsReturnedTags(); err != nil {
			return err
		}
		if err := fwm.isInvalidTags(); err != nil {
			return err
		}
	case FEDFundsSold:
		if err := fwm.isFEDFundsSoldValid(); err != nil {
			return err
		}
		if err := fwm.isFEDFundsSoldTags(); err != nil {
			return err
		}
		if err := fwm.isInvalidTags(); err != nil {
			return err
		}
	case DrawdownRequest:
		if err := fwm.isDrawdownRequestValid(); err != nil {
			return err
		}
		if err := fwm.isDrawdownRequestTags(); err != nil {
			return err
		}
		if err := fwm.isInvalidTags(); err != nil {
			return err
		}
	case BankDrawdownRequest:
		if err := fwm.isBankDrawdownRequestValid(); err != nil {
			return err
		}
		if err := fwm.isBankDrawdownRequestTags(); err != nil {
			return err
		}
		if err := fwm.isInvalidTags(); err != nil {
			return err
		}
	case CustomerCorporateDrawdownRequest:
		if err := fwm.isCustomerCorporateDrawdownRequestValid(); err != nil {
			return err
		}
		if err := fwm.isCustomerCorporateDrawdownRequestTags(); err != nil {
			return err
		}
		if err := fwm.isInvalidTags(); err != nil {
			return err
		}
	case BFCServiceMessage:
		if err := fwm.isServiceMessageValid(); err != nil {
			return err
		}
		if err := fwm.isServiceMessageTags(); err != nil {
			return err
		}
		if err := fwm.isInvalidServiceMessageTags(); err != nil {
			return err
		}
	}
	return nil
}

// isBankTransferValid
func (fwm *FedWireMessage) isBankTransferValid() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	switch typeSubType {
	case
		"1000", "1002", "1008",
		"1500", "1502", "1508",
		"1600", "1602", "1608":
	default:
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// isBankTransferTags
func (fwm *FedWireMessage) isBankTransferTags() error {
	if err := fwm.isPreviousMessageIdentifierRequired(); err != nil {
		return err
	}
	return nil
}

// isInvalidBankTransferTags
func (fwm *FedWireMessage) isInvalidBankTransferTags() error {
	if strings.TrimSpace(fwm.BusinessFunctionCode.TransactionTypeCode) != "" {
		return fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, fwm.BusinessFunctionCode.TransactionTypeCode)
	}
	if fwm.LocalInstrument != nil {
		return fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument)
	}
	if fwm.PaymentNotification != nil {
		return fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification)
	}
	if fwm.Charges != nil {
		return fieldError("Charges", ErrInvalidProperty, fwm.Charges)
	}
	if fwm.InstructedAmount != nil {
		return fieldError("InstructedAmount", ErrInvalidProperty, fwm.InstructedAmount)
	}
	if fwm.ExchangeRate != nil {
		return fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate)
	}
	if fwm.Beneficiary.Personal.IdentificationCode == "T" {
		return fieldError("Beneficiary.Personal.IdentificationCode", ErrInvalidProperty, fwm.Beneficiary.Personal.IdentificationCode)
	}
	if fwm.AccountDebitedDrawdown != nil {
		return fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown)
	}
	if fwm.Originator.Personal.IdentificationCode == "T" {
		return fieldError("Originator.Personal.IdentificationCode", ErrInvalidProperty, fwm.Originator.Personal.IdentificationCode)
	}
	if fwm.OriginatorOptionF != nil {
		return fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF)
	}
	if fwm.AccountCreditedDrawdown != nil {
		return fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown)
	}
	if fwm.FIDrawdownDebitAccountAdvice != nil {
		return fieldError("FIDrawdownDebitAccountAdvice", ErrInvalidProperty, fwm.FIDrawdownDebitAccountAdvice)
	}
	if fwm.ServiceMessage != nil {
		return fieldError("BusinessFunctionCode", ErrInvalidProperty, "ServiceMessage")
	}
	if fwm.UnstructuredAddenda != nil {
		return fieldError("BusinessFunctionCode", ErrInvalidProperty, "Unstructured Addenda")
	}
	if err := fwm.invalidCoverPaymentTags(); err != nil {
		return err
	}
	if err := fwm.invalidRemittanceTags(); err != nil {
		return err
	}
	return nil
}

// isCustomerTransferValid
func (fwm *FedWireMessage) isCustomerTransferValid() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	switch typeSubType {
	case
		FundsTransfer + BasicFundsTransfer,
		FundsTransfer + ReversalTransfer,
		FundsTransfer + ReversalPriorDayTransfer,
		ForeignTransfer + BasicFundsTransfer,
		ForeignTransfer + ReversalTransfer,
		ForeignTransfer + ReversalPriorDayTransfer,
		SettlementTransfer + BasicFundsTransfer,
		SettlementTransfer + ReversalTransfer,
		SettlementTransfer + ReversalPriorDayTransfer:
	default:
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// isCustomerTransferTags
func (fwm *FedWireMessage) isCustomerTransferTags() error {
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	if fwm.Originator == nil && fwm.OriginatorFI == nil {
		return fieldError("Originator or OriginatorFI", ErrFieldRequired)
	}
	if err := fwm.isPreviousMessageIdentifierRequired(); err != nil {
		return err
	}
	return nil
}

// isInvalidCustomerTransferTags
func (fwm *FedWireMessage) isInvalidCustomerTransferTags() error {
	// This covers the edit requirement
	if fwm.BusinessFunctionCode.TransactionTypeCode == "COV" {
		return fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, fwm.BusinessFunctionCode.TransactionTypeCode)
	}
	if fwm.LocalInstrument != nil {
		return fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument)
	}
	if fwm.PaymentNotification != nil {
		return fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification)
	}
	if fwm.AccountDebitedDrawdown != nil {
		return fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown)
	}
	if fwm.OriginatorOptionF != nil {
		return fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF)
	}
	if fwm.AccountCreditedDrawdown != nil {
		return fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown)
	}
	if fwm.FIDrawdownDebitAccountAdvice != nil {
		return fieldError("FIDrawdownDebitAccountAdvice", ErrInvalidProperty, fwm.FIDrawdownDebitAccountAdvice)
	}
	if fwm.ServiceMessage != nil {
		return fieldError("BusinessFunctionCode", ErrInvalidProperty, "ServiceMessage")
	}
	if fwm.UnstructuredAddenda != nil {
		return fieldError("BusinessFunctionCode", ErrInvalidProperty, "Unstructured Addenda")
	}
	if err := fwm.invalidCoverPaymentTags(); err != nil {
		return err
	}
	if err := fwm.invalidRemittanceTags(); err != nil {
		return err
	}
	return nil
}

// isCustomerTransferPlusValid
func (fwm *FedWireMessage) isCustomerTransferPlusValid() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	switch typeSubType {
	case
		FundsTransfer + BasicFundsTransfer,
		FundsTransfer + RequestReversal,
		FundsTransfer + ReversalTransfer,
		FundsTransfer + RequestReversalPriorDayTransfer,
		FundsTransfer + ReversalPriorDayTransfer,
		ForeignTransfer + BasicFundsTransfer,
		ForeignTransfer + RequestReversal,
		ForeignTransfer + ReversalTransfer,
		ForeignTransfer + RequestReversalPriorDayTransfer,
		ForeignTransfer + ReversalPriorDayTransfer,
		SettlementTransfer + BasicFundsTransfer,
		SettlementTransfer + RequestReversal,
		SettlementTransfer + ReversalTransfer,
		SettlementTransfer + RequestReversalPriorDayTransfer,
		SettlementTransfer + ReversalPriorDayTransfer:
	default:
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// isCustomerTransferPlusTags
func (fwm *FedWireMessage) isCustomerTransferPlusTags() error {
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	if fwm.Originator == nil {
		return fieldError("Originator", ErrFieldRequired)
	}
	if err := fwm.isPreviousMessageIdentifierRequired(); err != nil {
		return err
	}
	switch fwm.LocalInstrument.LocalInstrumentCode {
	case SequenceBCoverPaymentStructured:
		if fwm.BeneficiaryReference == nil {
			return fieldError("Beneficiary Reference", ErrFieldRequired)
		}
		if fwm.OrderingCustomer == nil {
			return fieldError("Ordering Customer", ErrFieldRequired)
		}
		if fwm.BeneficiaryCustomer == nil {
			return fieldError("BeneficiaryCustomer", ErrFieldRequired)
		}
	case ANSIX12format, GeneralXMLformat, ISO20022XMLformat,
		NarrativeText, STP820format, SWIFTfield70, UNEDIFACTformat:
		if fwm.UnstructuredAddenda == nil {
			return fieldError("UnstructuredAddenda", ErrFieldRequired)
		}
	case RelatedRemittanceInformation:
		if fwm.RelatedRemittance == nil {
			return fieldError("RelatedRemittance", ErrFieldRequired)
		}
	case RemittanceInformationStructured:
		if fwm.RemittanceOriginator == nil {
			return fieldError("RemittanceOriginator", ErrFieldRequired)
		}
		if fwm.RemittanceBeneficiary == nil {
			return fieldError("RemittanceBeneficiary", ErrFieldRequired)
		}
		if fwm.PrimaryRemittanceDocument == nil {
			return fieldError("PrimaryRemittanceDocument", ErrFieldRequired)
		}
		if fwm.ActualAmountPaid == nil {
			return fieldError("ActualAmountPaid", ErrFieldRequired)
		}
	case ProprietaryLocalInstrumentCode:
		if fwm.LocalInstrument.ProprietaryCode == "" {
			return fieldError("ProprietaryCode", ErrFieldRequired)
		}
	}
	if fwm.LocalInstrument.LocalInstrumentCode != "COVS" {
		if err := fwm.invalidCoverPaymentTags(); err != nil {
			return err
		}
	}
	return nil
}

// isInvalidCustomerTransferPlusTags
func (fwm *FedWireMessage) isInvalidCustomerTransferPlusTags() error {
	if strings.TrimSpace(fwm.BusinessFunctionCode.TransactionTypeCode) != "" {
		return fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, fwm.BusinessFunctionCode.TransactionTypeCode)
	}
	if fwm.AccountDebitedDrawdown != nil {
		return fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown)
	}
	if fwm.AccountCreditedDrawdown != nil {
		return fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown)
	}
	if fwm.FIReceiverFI != nil {
		return fieldError("FIReceiverFI", ErrInvalidProperty, fwm.FIReceiverFI)
	}
	if fwm.LocalInstrument.LocalInstrumentCode == SequenceBCoverPaymentStructured {
		if fwm.Charges != nil {
			return fieldError("Charges", ErrInvalidProperty, fwm.Charges)
		}
		if fwm.InstructedAmount != nil {
			return fieldError("InstructedAmount", ErrInvalidProperty, fwm.InstructedAmount)
		}
		if fwm.ExchangeRate != nil {
			return fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate)
		}
	}
	// ToDo: From the spec - Certain {7xxx} tags & {8xxx} tags may not be permitted depending upon value of {3610}.  I'm not sure how to code this yet
	return nil
}

// isCheckSameDaySettlementValid
func (fwm *FedWireMessage) isCheckSameDaySettlementValid() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	switch typeSubType {
	case
		SettlementTransfer + BasicFundsTransfer,
		SettlementTransfer + ReversalTransfer,
		SettlementTransfer + ReversalPriorDayTransfer:
	default:
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// isCheckSameDaySettlementTags
func (fwm *FedWireMessage) isCheckSameDaySettlementTags() error {
	return nil
}

// isDepositSendersAccountValid
func (fwm *FedWireMessage) isDepositSendersAccountValid() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	switch typeSubType {
	case
		SettlementTransfer + BasicFundsTransfer,
		SettlementTransfer + ReversalTransfer,
		SettlementTransfer + ReversalPriorDayTransfer:
	default:
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// isDepositSendersAccountTags
func (fwm *FedWireMessage) isDepositSendersAccountTags() error {
	return nil
}

// isFEDFundsReturnedValid
func (fwm *FedWireMessage) isFEDFundsReturnedValid() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	switch typeSubType {
	case
		SettlementTransfer + BasicFundsTransfer,
		SettlementTransfer + ReversalTransfer,
		SettlementTransfer + ReversalPriorDayTransfer:
	default:
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// isFEDFundsReturnedTag
func (fwm *FedWireMessage) isFEDFundsReturnedTags() error {
	return nil
}

// isFEDFundsSoldValid
func (fwm *FedWireMessage) isFEDFundsSoldValid() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	switch typeSubType {
	case
		SettlementTransfer + BasicFundsTransfer,
		SettlementTransfer + ReversalTransfer,
		SettlementTransfer + ReversalPriorDayTransfer:
	default:
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// isFEDFundsSoldTags
func (fwm *FedWireMessage) isFEDFundsSoldTags() error {
	return nil
}

// isDrawdownRequestValid
func (fwm *FedWireMessage) isDrawdownRequestValid() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	switch typeSubType {
	case
		FundsTransfer + FundsTransferRequestCredit,
		SettlementTransfer + FundsTransferRequestCredit:
	default:
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// isDrawdownRequestTags
func (fwm *FedWireMessage) isDrawdownRequestTags() error {
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	if fwm.Originator == nil {
		return fieldError("Originator", ErrFieldRequired)
	}
	return nil
}

// isBankDrawdownRequestValid
func (fwm *FedWireMessage) isBankDrawdownRequestValid() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	switch typeSubType {
	case
		SettlementTransfer + RequestCredit,
		SettlementTransfer + RefusalRequestCredit:
	default:
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// isBankDrawdownRequestTags
func (fwm *FedWireMessage) isBankDrawdownRequestTags() error {
	if fwm.AccountDebitedDrawdown == nil {
		return fieldError("AccountDebitedDrawdown", ErrFieldRequired)
	}
	if fwm.AccountCreditedDrawdown == nil {
		return fieldError("AccountCreditedDrawdown", ErrFieldRequired)
	}
	return nil
}

// isCustomerCorporateDrawdownRequestValid
func (fwm *FedWireMessage) isCustomerCorporateDrawdownRequestValid() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	switch typeSubType {
	case
		FundsTransfer + RequestCredit,
		FundsTransfer + RefusalRequestCredit:
	default:
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// isCustomerCorporateDrawdownRequestTags
func (fwm *FedWireMessage) isCustomerCorporateDrawdownRequestTags() error {
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	if fwm.AccountDebitedDrawdown == nil {
		return fieldError("AccountDebitedDrawdown", ErrFieldRequired)
	}
	if fwm.AccountCreditedDrawdown == nil {
		return fieldError("AccountCreditedDrawdown", ErrFieldRequired)
	}
	return nil
}

// isServiceMessageValid
func (fwm *FedWireMessage) isServiceMessageValid() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	switch typeSubType {
	case
		FundsTransfer + RequestReversal,
		FundsTransfer + RequestReversalPriorDayTransfer,
		FundsTransfer + RefusalRequestCredit,
		FundsTransfer + SSIServiceMessage,
		ForeignTransfer + RequestReversal,
		ForeignTransfer + RequestReversalPriorDayTransfer,
		ForeignTransfer + SSIServiceMessage,
		SettlementTransfer + RequestReversal,
		SettlementTransfer + RequestReversalPriorDayTransfer,
		SettlementTransfer + RefusalRequestCredit,
		SettlementTransfer + SSIServiceMessage:
	default:
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// isServiceMessageTags
func (fwm *FedWireMessage) isServiceMessageTags() error {
	return nil
}

// isInvalidServiceMessageTags
func (fwm *FedWireMessage) isInvalidServiceMessageTags() error {
	// BusinessFunctionCode.TransactionTypeCode (Element 02) is invalid
	if strings.TrimSpace(fwm.BusinessFunctionCode.TransactionTypeCode) != "" {
		return fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, fwm.BusinessFunctionCode.TransactionTypeCode)
	}
	if fwm.LocalInstrument != nil {
		return fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument)
	}
	if fwm.PaymentNotification != nil {
		return fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification)
	}
	if fwm.Charges != nil {
		return fieldError("Charges", ErrInvalidProperty, fwm.Charges)
	}
	if fwm.InstructedAmount != nil {
		return fieldError("InstructedAmount", ErrInvalidProperty, fwm.InstructedAmount)
	}
	if fwm.ExchangeRate != nil {
		return fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate)
	}
	if fwm.Beneficiary.Personal.IdentificationCode == "T" {
		return fieldError("Beneficiary.Personal.IdentificationCode", ErrInvalidProperty, fwm.Beneficiary.Personal.IdentificationCode)
	}
	if fwm.Originator.Personal.IdentificationCode == "T" {
		return fieldError("Originator.Personal.IdentificationCode", ErrInvalidProperty, fwm.Originator.Personal.IdentificationCode)
	}
	if fwm.OriginatorOptionF != nil {
		return fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF)
	}
	if fwm.UnstructuredAddenda != nil {
		return fieldError("BusinessFunctionCode", ErrInvalidProperty, "Unstructured Addenda")
	}
	if err := fwm.invalidCoverPaymentTags(); err != nil {
		return err
	}
	if err := fwm.invalidRemittanceTags(); err != nil {
		return err
	}
	return nil
}

// isPreviousMessageIdentifierRequired
func (fwm *FedWireMessage) isPreviousMessageIdentifierRequired() error {
	switch fwm.TypeSubType.SubTypeCode {
	case ReversalTransfer, ReversalPriorDayTransfer:
		if fwm.PreviousMessageIdentifier == nil {
			return fieldError("PreviousMessageIdentifier", ErrFieldRequired)
		}
	}
	return nil
}

// ToDo:  May revisit this create separate functions for each of the case BusinessFunctionCode statements as I did with
//  Valid and Tags or vice versa

// isInvalidTags
// isInvalidTags uses case logic for BusinessFunctionCodes that have the same invalid tags.  If this were to change per
// BusinessFunctionCode, create function isInvalidBusinessFunctionCodeTag() with the specific invalid tags for that
// BusinessFunctionCode (e.g. isInvalidBankTransferTags)
func (fwm *FedWireMessage) isInvalidTags() error {
	switch fwm.BusinessFunctionCode.BusinessFunctionCode {
	case CheckSameDaySettlement, DepositSendersAccount, FEDFundsReturned, FEDFundsSold:
		if strings.TrimSpace(fwm.BusinessFunctionCode.TransactionTypeCode) != "" {
			return fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, fwm.BusinessFunctionCode.TransactionTypeCode)
		}
		if fwm.LocalInstrument != nil {
			return fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument)
		}
		if fwm.PaymentNotification != nil {
			return fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification)
		}
		if fwm.Charges != nil {
			return fieldError("Charges", ErrInvalidProperty, fwm.Charges)
		}
		if fwm.InstructedAmount != nil {
			return fieldError("InstructedAmount", ErrInvalidProperty, fwm.InstructedAmount)
		}
		if fwm.ExchangeRate != nil {
			return fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate)
		}
		if fwm.Beneficiary.Personal.IdentificationCode == "T" {
			return fieldError("Beneficiary.Personal.IdentificationCode", ErrInvalidProperty, fwm.Beneficiary.Personal.IdentificationCode)
		}
		if fwm.AccountDebitedDrawdown != nil {
			return fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown)
		}
		if fwm.Originator.Personal.IdentificationCode == "T" {
			return fieldError("Originator.Personal.IdentificationCode", ErrInvalidProperty, fwm.Originator.Personal.IdentificationCode)
		}
		if fwm.OriginatorOptionF != nil {
			return fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF)
		}
		if fwm.AccountCreditedDrawdown != nil {
			return fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown)
		}
		if fwm.FIDrawdownDebitAccountAdvice != nil {
			return fieldError("FIDrawdownDebitAccountAdvice", ErrInvalidProperty, fwm.FIDrawdownDebitAccountAdvice)
		}
		if fwm.ServiceMessage != nil {
			return fieldError("BusinessFunctionCode", ErrInvalidProperty, "ServiceMessage")
		}
		if fwm.UnstructuredAddenda != nil {
			return fieldError("BusinessFunctionCode", ErrInvalidProperty, "Unstructured Addenda")
		}
		if err := fwm.invalidCoverPaymentTags(); err != nil {
			return err
		}

		if err := fwm.invalidRemittanceTags(); err != nil {
			return err
		}
	case DrawdownRequest, BankDrawdownRequest, CustomerCorporateDrawdownRequest:
		if strings.TrimSpace(fwm.BusinessFunctionCode.TransactionTypeCode) != "" {
			return fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, fwm.BusinessFunctionCode.TransactionTypeCode)
		}
		if fwm.LocalInstrument != nil {
			return fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument)
		}
		if fwm.PaymentNotification != nil {
			return fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification)
		}
		if fwm.Charges != nil {
			return fieldError("Charges", ErrInvalidProperty, fwm.Charges)
		}
		if fwm.InstructedAmount != nil {
			return fieldError("InstructedAmount", ErrInvalidProperty, fwm.InstructedAmount)
		}
		if fwm.ExchangeRate != nil {
			return fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate)
		}
		if fwm.Beneficiary.Personal.IdentificationCode == "T" {
			return fieldError("Beneficiary.Personal.IdentificationCode", ErrInvalidProperty, fwm.Beneficiary.Personal.IdentificationCode)
		}
		if fwm.Originator.Personal.IdentificationCode == "T" {
			return fieldError("Originator.Personal.IdentificationCode", ErrInvalidProperty, fwm.Originator.Personal.IdentificationCode)
		}
		if fwm.OriginatorOptionF != nil {
			return fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF)
		}
		if fwm.ServiceMessage != nil {
			return fieldError("BusinessFunctionCode", ErrInvalidProperty, "ServiceMessage")
		}
		if fwm.UnstructuredAddenda != nil {
			return fieldError("BusinessFunctionCode", ErrInvalidProperty, "Unstructured Addenda")
		}
		if err := fwm.invalidCoverPaymentTags(); err != nil {
			return err
		}
		if err := fwm.invalidRemittanceTags(); err != nil {
			return err
		}
	}
	return nil
}

func (fwm *FedWireMessage) invalidRemittanceTags() error {
	if fwm.RelatedRemittance != nil {
		return fieldError("RelatedRemittance", ErrInvalidProperty, "RelatedRemittance")
	}
	if fwm.RemittanceOriginator != nil {
		return fieldError("RemittanceOriginator", ErrInvalidProperty, "RemittanceOriginator")
	}
	if fwm.RemittanceBeneficiary != nil {
		return fieldError("RemittanceBeneficiary", ErrInvalidProperty, "RemittanceBeneficiary")
	}
	if fwm.PrimaryRemittanceDocument != nil {
		return fieldError("PrimaryRemittanceDocument", ErrInvalidProperty, "PrimaryRemittanceDocument")
	}
	if fwm.ActualAmountPaid != nil {
		return fieldError("ActualAmountPaid", ErrInvalidProperty, "ActualAmountPaid")
	}
	if fwm.GrossAmountRemittanceDocument != nil {
		return fieldError("GrossAmountRemittanceDocument", ErrInvalidProperty, "GrossAmountRemittanceDocument")
	}
	if fwm.AmountNegotiatedDiscount != nil {
		return fieldError("AmountNegotiatedDiscount", ErrInvalidProperty, "AmountNegotiatedDiscount")
	}
	if fwm.Adjustment != nil {
		return fieldError("Adjustment", ErrInvalidProperty, "Adjustment")
	}
	if fwm.DateRemittanceDocument != nil {
		return fieldError("DateRemittanceDocument", ErrInvalidProperty, "DateRemittanceDocument")
	}
	if fwm.SecondaryRemittanceDocument != nil {
		return fieldError("SecondaryRemittanceDocument", ErrInvalidProperty, "SecondaryRemittanceDocument")
	}
	if fwm.RemittanceFreeText != nil {
		return fieldError("RemittanceFreeText", ErrInvalidProperty, "RemittanceFreeText")
	}
	return nil
}

func (fwm *FedWireMessage) invalidCoverPaymentTags() error {
	if fwm.CurrencyInstructedAmount != nil {
		return fieldError("CurrencyInstructedAmount ", ErrInvalidProperty, fwm.CurrencyInstructedAmount)
	}
	if fwm.OrderingCustomer != nil {
		return fieldError("OrderingCustomer", ErrInvalidProperty, fwm.OrderingCustomer)
	}
	if fwm.OrderingInstitution != nil {
		return fieldError("OrderingInstitution", ErrInvalidProperty, fwm.OrderingInstitution)
	}
	if fwm.IntermediaryInstitution != nil {
		return fieldError("IntermediaryInstitution", ErrInvalidProperty, fwm.IntermediaryInstitution)
	}
	if fwm.InstitutionAccount != nil {
		return fieldError("InstitutionAccount", ErrInvalidProperty, fwm.InstitutionAccount)
	}
	if fwm.BeneficiaryCustomer != nil {
		return fieldError("BeneficiaryCustomer", ErrInvalidProperty, fwm.BeneficiaryCustomer)
	}
	if fwm.Remittance != nil {
		return fieldError("Remittance", ErrInvalidProperty, fwm.Remittance)
	}
	if fwm.SenderToReceiver != nil {
		return fieldError("SenderToReceiver", ErrInvalidProperty, fwm.SenderToReceiver)
	}
	return nil
}

// SetSenderSupplied appends a SenderSupplied to the FedWireMessage
func (fwm *FedWireMessage) SetSenderSupplied(ss *SenderSupplied) {
	fwm.SenderSupplied = ss
}

// GetSenderSupplied returns the current SenderSupplied
func (fwm *FedWireMessage) GetSenderSupplied() *SenderSupplied {
	return fwm.SenderSupplied
}

// SetTypeSubType appends a TypeSubType to the FedWireMessage
func (fwm *FedWireMessage) SetTypeSubType(tst *TypeSubType) {
	fwm.TypeSubType = tst
}

// GetTypeSubType returns the current TypeSubType
func (fwm *FedWireMessage) GetTypeSubType() *TypeSubType {
	return fwm.TypeSubType
}

// SetInputMessageAccountabilityData appends a InputMessageAccountabilityData to the FedWireMessage
func (fwm *FedWireMessage) SetInputMessageAccountabilityData(imad *InputMessageAccountabilityData) {
	fwm.InputMessageAccountabilityData = imad
}

// GetInputMessageAccountabilityData returns the current InputMessageAccountabilityData
func (fwm *FedWireMessage) GetInputMessageAccountabilityData() *InputMessageAccountabilityData {
	return fwm.InputMessageAccountabilityData
}

// SetAmount appends a Amount to the FedWireMessage
func (fwm *FedWireMessage) SetAmount(amt *Amount) {
	fwm.Amount = amt
}

// GetAmount returns the current Amount
func (fwm *FedWireMessage) GetAmount() *Amount {
	return fwm.Amount
}

// SetSenderDepositoryInstitution appends a SenderDepositoryInstitution to the FedWireMessage
func (fwm *FedWireMessage) SetSenderDepositoryInstitution(sdi *SenderDepositoryInstitution) {
	fwm.SenderDepositoryInstitution = sdi
}

// GetSenderDepositoryInstitution returns the current SenderDepositoryInstitution
func (fwm *FedWireMessage) GetSenderDepositoryInstitution() *SenderDepositoryInstitution {
	return fwm.SenderDepositoryInstitution
}

// SetReceiverDepositoryInstitution appends a ReceiverDepositoryInstitution to the FedWireMessage
func (fwm *FedWireMessage) SetReceiverDepositoryInstitution(rdi *ReceiverDepositoryInstitution) {
	fwm.ReceiverDepositoryInstitution = rdi
}

// GetReceiverDepositoryInstitution returns the current ReceiverDepositoryInstitution
func (fwm *FedWireMessage) GetReceiverDepositoryInstitution() *ReceiverDepositoryInstitution {
	return fwm.ReceiverDepositoryInstitution
}

// SetBusinessFunctionCode appends a BusinessFunctionCode to the FedWireMessage
func (fwm *FedWireMessage) SetBusinessFunctionCode(bfc *BusinessFunctionCode) {
	fwm.BusinessFunctionCode = bfc
}

// GetBusinessFunctionCode returns the current BusinessFunctionCode
func (fwm *FedWireMessage) GetBusinessFunctionCode() *BusinessFunctionCode {
	return fwm.BusinessFunctionCode
}

// SetSenderReference appends a SenderReference to the FedWireMessage
func (fwm *FedWireMessage) SetSenderReference(sr *SenderReference) {
	fwm.SenderReference = sr
}

// GetSenderReference returns the current SenderReference
func (fwm *FedWireMessage) GetSenderReference() *SenderReference {
	return fwm.SenderReference
}

// SetPreviousMessageIdentifier appends a PreviousMessageIdentifier to the FedWireMessage
func (fwm *FedWireMessage) SetPreviousMessageIdentifier(pmi *PreviousMessageIdentifier) {
	fwm.PreviousMessageIdentifier = pmi
}

// GetPreviousMessageIdentifier returns the current PreviousMessageIdentifier
func (fwm *FedWireMessage) GetPreviousMessageIdentifier() *PreviousMessageIdentifier {
	return fwm.PreviousMessageIdentifier
}

// SetLocalInstrument appends a LocalInstrument to the FedWireMessage
func (fwm *FedWireMessage) SetLocalInstrument(li *LocalInstrument) {
	fwm.LocalInstrument = li
}

// GetLocalInstrument returns the current LocalInstrument
func (fwm *FedWireMessage) GetLocalInstrument() *LocalInstrument {
	return fwm.LocalInstrument
}

// SetPaymentNotification appends a PaymentNotification to the FedWireMessage
func (fwm *FedWireMessage) SetPaymentNotification(pn *PaymentNotification) {
	fwm.PaymentNotification = pn
}

// GetPaymentNotification returns the current PaymentNotification
func (fwm *FedWireMessage) GetPaymentNotification() *PaymentNotification {
	return fwm.PaymentNotification
}

// SetCharges appends a Charges to the FedWireMessage
func (fwm *FedWireMessage) SetCharges(c *Charges) {
	fwm.Charges = c
}

// GetCharges returns the current Charges
func (fwm *FedWireMessage) GetCharges() *Charges {
	return fwm.Charges
}

// SetInstructedAmount appends a InstructedAmount to the FedWireMessage
func (fwm *FedWireMessage) SetInstructedAmount(ia *InstructedAmount) {
	fwm.InstructedAmount = ia
}

// GetInstructedAmount returns the current Instructed Amount
func (fwm *FedWireMessage) GetInstructedAmount() *InstructedAmount {
	return fwm.InstructedAmount
}

// SetExchangeRate appends a ExchangeRate to the FedWireMessage
func (fwm *FedWireMessage) SetExchangeRate(er *ExchangeRate) {
	fwm.ExchangeRate = er
}

// GetExchangeRate returns the current ExchangeRate
func (fwm *FedWireMessage) GetExchangeRate() *ExchangeRate {
	return fwm.ExchangeRate
}

// SetBeneficiaryIntermediaryFI appends a BeneficiaryIntermediaryFI to the FedWireMessage
func (fwm *FedWireMessage) SetBeneficiaryIntermediaryFI(bifi *BeneficiaryIntermediaryFI) {
	fwm.BeneficiaryIntermediaryFI = bifi
}

// GetBeneficiaryIntermediaryFI returns the current BeneficiaryIntermediaryFI
func (fwm *FedWireMessage) GetBeneficiaryIntermediaryFI() *BeneficiaryIntermediaryFI {
	return fwm.BeneficiaryIntermediaryFI
}

// SetBeneficiaryFI appends a BeneficiaryFI to the FedWireMessage
func (fwm *FedWireMessage) SetBeneficiaryFI(bfi *BeneficiaryFI) {
	fwm.BeneficiaryFI = bfi
}

// GetBeneficiaryFI returns the current BeneficiaryFI
func (fwm *FedWireMessage) GetBeneficiaryFI() *BeneficiaryFI {
	return fwm.BeneficiaryFI
}

// SetBeneficiary appends a Beneficiary to the FedWireMessage
func (fwm *FedWireMessage) SetBeneficiary(ben *Beneficiary) {
	fwm.Beneficiary = ben
}

// GetBeneficiary returns the current Beneficiary
func (fwm *FedWireMessage) GetBeneficiary() *Beneficiary {
	return fwm.Beneficiary
}

// SetBeneficiaryReference appends a BeneficiaryReference to the FedWireMessage
func (fwm *FedWireMessage) SetBeneficiaryReference(br *BeneficiaryReference) {
	fwm.BeneficiaryReference = br
}

// GetBeneficiaryReference returns the current BeneficiaryReference
func (fwm *FedWireMessage) GetBeneficiaryReference() *BeneficiaryReference {
	return fwm.BeneficiaryReference
}

// SetAccountDebitedDrawdown appends a AccountDebitedDrawdown to the FedWireMessage
func (fwm *FedWireMessage) SetAccountDebitedDrawdown(debitDD *AccountDebitedDrawdown) {
	fwm.AccountDebitedDrawdown = debitDD
}

// GetAccountDebitedDrawdown returns the current AccountDebitedDrawdown
func (fwm *FedWireMessage) GetAccountDebitedDrawdown() *AccountDebitedDrawdown {
	return fwm.AccountDebitedDrawdown
}

// SetOriginator appends a Originator to the FedWireMessage
func (fwm *FedWireMessage) SetOriginator(o *Originator) {
	fwm.Originator = o
}

// GetOriginator returns the current Originator
func (fwm *FedWireMessage) GetOriginator() *Originator {
	return fwm.Originator
}

// SetOriginatorOptionF appends a OriginatorOptionF to the FedWireMessage
func (fwm *FedWireMessage) SetOriginatorOptionF(oof *OriginatorOptionF) {
	fwm.OriginatorOptionF = oof
}

// GetOriginatorOptionF returns the current OriginatorOptionF
func (fwm *FedWireMessage) GetOriginatorOptionF() *OriginatorOptionF {
	return fwm.OriginatorOptionF
}

// SetOriginatorFI appends a OriginatorFI to the FedWireMessage
func (fwm *FedWireMessage) SetOriginatorFI(ofi *OriginatorFI) {
	fwm.OriginatorFI = ofi
}

// GetOriginatorFI returns the current OriginatorFI
func (fwm *FedWireMessage) GetOriginatorFI() *OriginatorFI {
	return fwm.OriginatorFI
}

// SetInstructingFI appends a InstructingFI to the FedWireMessage
func (fwm *FedWireMessage) SetInstructingFI(ifi *InstructingFI) {
	fwm.InstructingFI = ifi
}

// GetInstructingFI returns the current InstructingFI
func (fwm *FedWireMessage) GetInstructingFI() *InstructingFI {
	return fwm.InstructingFI
}

// SetAccountCreditedDrawdown appends a AccountCreditedDrawdown to the FedWireMessage
func (fwm *FedWireMessage) SetAccountCreditedDrawdown(creditDD *AccountCreditedDrawdown) {
	fwm.AccountCreditedDrawdown = creditDD
}

// GetAccountCreditedDrawdown returns the current AccountCreditedDrawdown
func (fwm *FedWireMessage) GetAccountCreditedDrawdown() *AccountCreditedDrawdown {
	return fwm.AccountCreditedDrawdown
}

// SetOriginatorToBeneficiary appends a OriginatorToBeneficiary to the FedWireMessage
func (fwm *FedWireMessage) SetOriginatorToBeneficiary(ob *OriginatorToBeneficiary) {
	fwm.OriginatorToBeneficiary = ob
}

// GetOriginatorToBeneficiary returns the current OriginatorToBeneficiary
func (fwm *FedWireMessage) GetOriginatorToBeneficiary() *OriginatorToBeneficiary {
	return fwm.OriginatorToBeneficiary
}

// SetFIReceiverFI appends a FIReceiverFI to the FedWireMessage
func (fwm *FedWireMessage) SetFIReceiverFI(firfi *FIReceiverFI) {
	fwm.FIReceiverFI = firfi
}

// GetFIReceiverFI returns the current FIReceiverFI
func (fwm *FedWireMessage) GetFIReceiverFI() *FIReceiverFI {
	return fwm.FIReceiverFI
}

// SetFIDrawdownDebitAccountAdvice appends a FIDrawdownDebitAccountAdvice to the FedWireMessage
func (fwm *FedWireMessage) SetFIDrawdownDebitAccountAdvice(debitDDAdvice *FIDrawdownDebitAccountAdvice) {
	fwm.FIDrawdownDebitAccountAdvice = debitDDAdvice
}

// GetFIDrawdownDebitAccountAdvice returns the current FIDrawdownDebitAccountAdvice
func (fwm *FedWireMessage) GetFIDrawdownDebitAccountAdvice() *FIDrawdownDebitAccountAdvice {
	return fwm.FIDrawdownDebitAccountAdvice
}

// SetFIIntermediaryFI appends a FIIntermediaryFI to the FedWireMessage
func (fwm *FedWireMessage) SetFIIntermediaryFI(fiifi *FIIntermediaryFI) {
	fwm.FIIntermediaryFI = fiifi
}

// GetFIIntermediaryFI returns the current FIIntermediaryFI
func (fwm *FedWireMessage) GetFIIntermediaryFI() *FIIntermediaryFI {
	return fwm.FIIntermediaryFI
}

// SetFIIntermediaryFIAdvice appends a FIIntermediaryFIAdvice to the FedWireMessage
func (fwm *FedWireMessage) SetFIIntermediaryFIAdvice(fiifia *FIIntermediaryFIAdvice) {
	fwm.FIIntermediaryFIAdvice = fiifia
}

// GetFIIntermediaryFIAdvice returns the current FIIntermediaryFIAdvice
func (fwm *FedWireMessage) GetFIIntermediaryFIAdvice() *FIIntermediaryFIAdvice {
	return fwm.FIIntermediaryFIAdvice
}

// SetFIBeneficiaryFI appends a FIBeneficiaryFI to the FedWireMessage
func (fwm *FedWireMessage) SetFIBeneficiaryFI(fibfi *FIBeneficiaryFI) {
	fwm.FIBeneficiaryFI = fibfi
}

// GetFIBeneficiaryFI returns the current FIBeneficiaryFI
func (fwm *FedWireMessage) GetFIBeneficiaryFI() *FIBeneficiaryFI {
	return fwm.FIBeneficiaryFI
}

// SetFIBeneficiaryFIAdvice appends a FIBeneficiaryFIAdvice to the FedWireMessage
func (fwm *FedWireMessage) SetFIBeneficiaryFIAdvice(fibfia *FIBeneficiaryFIAdvice) {
	fwm.FIBeneficiaryFIAdvice = fibfia
}

// GetFIBeneficiaryFIAdvice returns the current FIBeneficiaryFIAdvice
func (fwm *FedWireMessage) GetFIBeneficiaryFIAdvice() *FIBeneficiaryFIAdvice {
	return fwm.FIBeneficiaryFIAdvice
}

// SetFIBeneficiary appends a FIBeneficiary to the FedWireMessage
func (fwm *FedWireMessage) SetFIBeneficiary(fib *FIBeneficiary) {
	fwm.FIBeneficiary = fib
}

// GetFIBeneficiary returns the current FIBeneficiary
func (fwm *FedWireMessage) GetFIBeneficiary() *FIBeneficiary {
	return fwm.FIBeneficiary
}

// SetFIBeneficiaryAdvice appends a FIBeneficiaryAdviceto the FedWireMessage
func (fwm *FedWireMessage) SetFIBeneficiaryAdvice(fiba *FIBeneficiaryAdvice) {
	fwm.FIBeneficiaryAdvice = fiba
}

// GetFIBeneficiaryAdvice returns the current FIBeneficiaryAdvice
func (fwm *FedWireMessage) GetFIBeneficiaryAdvice() *FIBeneficiaryAdvice {
	return fwm.FIBeneficiaryAdvice
}

// SetFIPaymentMethodToBeneficiary appends a FIPaymentMethodToBeneficiary to the FedWireMessage
func (fwm *FedWireMessage) SetFIPaymentMethodToBeneficiary(pm *FIPaymentMethodToBeneficiary) {
	fwm.FIPaymentMethodToBeneficiary = pm
}

// GetFIPaymentMethodToBeneficiary returns the current FIPaymentMethodToBeneficiary
func (fwm *FedWireMessage) GetFIPaymentMethodToBeneficiary() *FIPaymentMethodToBeneficiary {
	return fwm.FIPaymentMethodToBeneficiary
}

// SetFIAdditionalFIToFI appends a FIAdditionalFIToFI to the FedWireMessage
func (fwm *FedWireMessage) SetFIAdditionalFIToFI(fifi *FIAdditionalFIToFI) {
	fwm.FIAdditionalFIToFI = fifi
}

// GetFIAdditionalFIToFI returns the current FIAdditionalFIToFI
func (fwm *FedWireMessage) GetFIAdditionalFIToFI() *FIAdditionalFIToFI {
	return fwm.FIAdditionalFIToFI
}

// SetCurrencyInstructedAmount appends a CurrencyInstructedAmount to the FedWireMessage
func (fwm *FedWireMessage) SetCurrencyInstructedAmount(cia *CurrencyInstructedAmount) {
	fwm.CurrencyInstructedAmount = cia
}

// GetCurrencyInstructedAmount returns the current CurrencyInstructedAmount
func (fwm *FedWireMessage) GetCurrencyInstructedAmount() *CurrencyInstructedAmount {
	return fwm.CurrencyInstructedAmount
}

// SetOrderingCustomer appends a OrderingCustomer to the FedWireMessage
func (fwm *FedWireMessage) SetOrderingCustomer(oc *OrderingCustomer) {
	fwm.OrderingCustomer = oc
}

// GetOrderingCustomer returns the current OrderingCustomer
func (fwm *FedWireMessage) GetOrderingCustomer() *OrderingCustomer {
	return fwm.OrderingCustomer
}

// SetOrderingInstitution appends a OrderingInstitution to the FedWireMessage
func (fwm *FedWireMessage) SetOrderingInstitution(oi *OrderingInstitution) {
	fwm.OrderingInstitution = oi
}

// GetOrderingInstitution returns the current OrderingInstitution
func (fwm *FedWireMessage) GetOrderingInstitution() *OrderingInstitution {
	return fwm.OrderingInstitution
}

// SetIntermediaryInstitution appends a IntermediaryInstitution to the FedWireMessage
func (fwm *FedWireMessage) SetIntermediaryInstitution(ii *IntermediaryInstitution) {
	fwm.IntermediaryInstitution = ii
}

// GetIntermediaryInstitution returns the current IntermediaryInstitution
func (fwm *FedWireMessage) GetIntermediaryInstitution() *IntermediaryInstitution {
	return fwm.IntermediaryInstitution
}

// SetInstitutionAccount appends a InstitutionAccount to the FedWireMessage
func (fwm *FedWireMessage) SetInstitutionAccount(iAccount *InstitutionAccount) {
	fwm.InstitutionAccount = iAccount
}

// GetInstitutionAccount returns the current InstitutionAccount
func (fwm *FedWireMessage) GetInstitutionAccount() *InstitutionAccount {
	return fwm.InstitutionAccount
}

// SetBeneficiaryCustomer appends a BeneficiaryCustomer to the FedWireMessage
func (fwm *FedWireMessage) SetBeneficiaryCustomer(bc *BeneficiaryCustomer) {
	fwm.BeneficiaryCustomer = bc
}

// GetBeneficiaryCustomer returns the current BeneficiaryCustomer
func (fwm *FedWireMessage) GetBeneficiaryCustomer() *BeneficiaryCustomer {
	return fwm.BeneficiaryCustomer
}

// SetRemittance appends a Remittance to the FedWireMessage
func (fwm *FedWireMessage) SetRemittance(ri *Remittance) {
	fwm.Remittance = ri
}

// GetRemittance returns the current Remittance
func (fwm *FedWireMessage) GetRemittance() *Remittance {
	return fwm.Remittance
}

// SetSenderToReceiver appends a SenderToReceiver to the FedWireMessage
func (fwm *FedWireMessage) SetSenderToReceiver(str *SenderToReceiver) {
	fwm.SenderToReceiver = str
}

// GetSenderToReceiver  returns the current SenderToReceiver
func (fwm *FedWireMessage) GetSenderToReceiver() *SenderToReceiver {
	return fwm.SenderToReceiver
}

// SetUnstructuredAddenda appends a UnstructuredAddenda to the FedWireMessage
func (fwm *FedWireMessage) SetUnstructuredAddenda(ua *UnstructuredAddenda) {
	fwm.UnstructuredAddenda = ua
}

// GetUnstructuredAddenda returns the current UnstructuredAddenda
func (fwm *FedWireMessage) GetUnstructuredAddenda() *UnstructuredAddenda {
	return fwm.UnstructuredAddenda
}

// SetRelatedRemittance appends a RelatedRemittance to the FedWireMessage
func (fwm *FedWireMessage) SetRelatedRemittance(rr *RelatedRemittance) {
	fwm.RelatedRemittance = rr
}

// GetRelatedRemittance returns the current RelatedRemittance
func (fwm *FedWireMessage) GetRelatedRemittance() *RelatedRemittance {
	return fwm.RelatedRemittance
}

// SetRemittanceOriginator appends a RemittanceOriginator to the FedWireMessage
func (fwm *FedWireMessage) SetRemittanceOriginator(ro *RemittanceOriginator) {
	fwm.RemittanceOriginator = ro
}

// GetRemittanceOriginator returns the current RemittanceOriginator
func (fwm *FedWireMessage) GetRemittanceOriginator() *RemittanceOriginator {
	return fwm.RemittanceOriginator
}

// SetRemittanceBeneficiary appends a RemittanceBeneficiary to the FedWireMessage
func (fwm *FedWireMessage) SetRemittanceBeneficiary(rb *RemittanceBeneficiary) {
	fwm.RemittanceBeneficiary = rb
}

// GetRemittanceBeneficiary returns the current RemittanceBeneficiary
func (fwm *FedWireMessage) GetRemittanceBeneficiary() *RemittanceBeneficiary {
	return fwm.RemittanceBeneficiary
}

// SetPrimaryRemittanceDocument appends a PrimaryRemittanceDocument to the FedWireMessage
func (fwm *FedWireMessage) SetPrimaryRemittanceDocument(prd *PrimaryRemittanceDocument) {
	fwm.PrimaryRemittanceDocument = prd
}

// GetPrimaryRemittanceDocument returns the current PrimaryRemittanceDocument
func (fwm *FedWireMessage) GetPrimaryRemittanceDocument() *PrimaryRemittanceDocument {
	return fwm.PrimaryRemittanceDocument
}

// SetActualAmountPaid appends a ActualAmountPaid to the FedWireMessage
func (fwm *FedWireMessage) SetActualAmountPaid(aap *ActualAmountPaid) {
	fwm.ActualAmountPaid = aap
}

// GetActualAmountPaid returns the current ActualAmountPaid
func (fwm *FedWireMessage) GetActualAmountPaid() *ActualAmountPaid {
	return fwm.ActualAmountPaid
}

// SetGrossAmountRemittanceDocument appends a GrossAmountRemittanceDocument to the FedWireMessage
func (fwm *FedWireMessage) SetGrossAmountRemittanceDocument(gard *GrossAmountRemittanceDocument) {
	fwm.GrossAmountRemittanceDocument = gard
}

// GetGrossAmountRemittanceDocument returns the current GrossAmountRemittanceDocument
func (fwm *FedWireMessage) GetGrossAmountRemittanceDocument() *GrossAmountRemittanceDocument {
	return fwm.GrossAmountRemittanceDocument
}

// SetAmountNegotiatedDiscount appends a AmountNegotiatedDiscount to the FedWireMessage
func (fwm *FedWireMessage) SetAmountNegotiatedDiscount(nd *AmountNegotiatedDiscount) {
	fwm.AmountNegotiatedDiscount = nd
}

// GetAmountNegotiatedDiscount returns the current AmountNegotiatedDiscount
func (fwm *FedWireMessage) GetAmountNegotiatedDiscount() *AmountNegotiatedDiscount {
	return fwm.AmountNegotiatedDiscount
}

// SetAdjustment appends a Adjustment to the FedWireMessage
func (fwm *FedWireMessage) SetAdjustment(adj *Adjustment) {
	fwm.Adjustment = adj
}

// GetAdjustment returns the current Adjustment
func (fwm *FedWireMessage) GetAdjustment() *Adjustment {
	return fwm.Adjustment
}

// SetDateRemittanceDocument appends a DateRemittanceDocument to the FedWireMessage
func (fwm *FedWireMessage) SetDateRemittanceDocument(drd *DateRemittanceDocument) {
	fwm.DateRemittanceDocument = drd
}

// GetDateRemittanceDocument returns the current DateRemittanceDocument
func (fwm *FedWireMessage) GetDateRemittanceDocument() *DateRemittanceDocument {
	return fwm.DateRemittanceDocument
}

// SetSecondaryRemittanceDocument appends a SecondaryRemittanceDocument to the FedWireMessage
func (fwm *FedWireMessage) SetSecondaryRemittanceDocument(srd *SecondaryRemittanceDocument) {
	fwm.SecondaryRemittanceDocument = srd
}

// GetSecondaryRemittanceDocument returns the current SecondaryRemittanceDocument
func (fwm *FedWireMessage) GetSecondaryRemittanceDocument() *SecondaryRemittanceDocument {
	return fwm.SecondaryRemittanceDocument
}

// SetRemittanceFreeText appends a RemittanceFreeText to the FedWireMessage
func (fwm *FedWireMessage) SetRemittanceFreeText(rft *RemittanceFreeText) {
	fwm.RemittanceFreeText = rft
}

// GetRemittanceFreeText returns the current RemittanceFreeText
func (fwm *FedWireMessage) GetRemittanceFreeText() *RemittanceFreeText {
	return fwm.RemittanceFreeText
}

// SetServiceMessage appends a ServiceMessage to the FedWireMessage
func (fwm *FedWireMessage) SetServiceMessage(sm *ServiceMessage) {
	fwm.ServiceMessage = sm
}

// GetServiceMessage returns the current ServiceMessage
func (fwm *FedWireMessage) GetServiceMessage() *ServiceMessage {
	return fwm.ServiceMessage
}

func (fwm *FedWireMessage) isAmountValid() error {
	if  fwm.TypeSubType.SubTypeCode != "90" && fwm.Amount.Amount == "000000000000" {
		return NewErrInvalidPropertyForProperty("Amount", fwm.Amount.Amount,
			"SubTypeCode", fwm.TypeSubType.SubTypeCode)
	}
	return nil
}

func (fwm *FedWireMessage) isPreviousMessageIdentifierValid() error {
	if fwm.TypeSubType.SubTypeCode == "02" || fwm.TypeSubType.SubTypeCode == "08" {
		switch fwm.BusinessFunctionCode.BusinessFunctionCode {
		case BankTransfer, CustomerTransfer, CustomerTransferPlus:
			if fwm.PreviousMessageIdentifier == nil {
				return fieldError("PreviousMessageIdentifier", ErrFieldRequired)
			}
		}
	}
	return nil
}

func (fwm *FedWireMessage) isLocalInstrumentCodeValid() error {
/*	if fwm.BusinessFunctionCode.BusinessFunctionCode != "CTP" {
		return NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode,
			"BusinessFunctionCode.BusinessFunctionCode", fwm.BusinessFunctionCode.BusinessFunctionCode)
	}*/
	if fwm.LocalInstrument.ProprietaryCode != "" && fwm.LocalInstrument.LocalInstrumentCode != "PROP" {
		return NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode,
			"LocalInstrument.ProprietaryCode", fwm.LocalInstrument.ProprietaryCode)
	}
	if fwm.LocalInstrument.LocalInstrumentCode == "COVS" {
		if fwm.BeneficiaryReference == nil {
			return fieldError("BeneficiaryReference", ErrFieldRequired)
		}
	}
	return nil
}

func (fwm *FedWireMessage) isPaymentNotificationValid() error {
	// ToDo: I'm not sure of anyway to indicate this from a code stand point
	//  Payment Notification Indicator is mandatory.
	//   Indicators 0 through 6 – Reserved for market practice conventions.
	//   Indicators 7 through 9 – Reserved for bilateral agreements between Fedwire senders and receivers.
	return nil
}

func (fwm *FedWireMessage) isChargesValid() error {
	if fwm.LocalInstrument != nil {
		if fwm.LocalInstrument.LocalInstrumentCode == "COVS" {
			return NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode,
				"Charges", "Charges Defined")
		}
	}
	return nil
}

func (fwm *FedWireMessage) isInstructedAmountValid() error {
	if fwm.LocalInstrument != nil {
		if fwm.LocalInstrument.LocalInstrumentCode == "COVS" {
			return NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode,
				"Instructed Amount", "Instructed Amount")
		}
	}
	return nil
}

func (fwm *FedWireMessage) isExchangeRateValid() error {
	if fwm.InstructedAmount == nil {
		return fieldError("InstructedAmount", ErrFieldRequired)
	}
	if fwm.LocalInstrument != nil {
		if fwm.LocalInstrument.LocalInstrumentCode == "COVS" {
			return NewErrInvalidPropertyForProperty("ExchangeRate", fwm.ExchangeRate.ExchangeRate,
				"Instructed Amount", "Instructed Amount")
		}
	}
	return nil
}

func (fwm *FedWireMessage) isBeneficiaryIntermediaryFIValid() error {
	if fwm.BeneficiaryFI == nil {
		return fieldError("BeneficiaryFI", ErrFieldRequired)
	}
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	return nil
}

func (fwm *FedWireMessage) isBeneficiaryFIValid() error {
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	return nil
}

func (fwm *FedWireMessage) isOriginatorFIValid() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus {
		if fwm.Originator == nil && fwm.OriginatorOptionF == nil {
			return fieldError("Originator or Originator Option F", ErrFieldRequired)
		}
	} else {
		if fwm.Originator == nil {
			return fieldError("Originator", ErrFieldRequired)
		}
	}
	return nil
}

func (fwm *FedWireMessage) isInstructingFIValid() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus {
		if fwm.Originator == nil && fwm.OriginatorOptionF == nil {
			return fieldError("Originator or Originator Option F", ErrFieldRequired)
		}
	} else {
		if fwm.Originator == nil {
			return fieldError("Originator", ErrFieldRequired)
		}
	}
	return nil
}

func (fwm *FedWireMessage) isOriginatorToBeneficiaryValid() error {
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus {
		if fwm.Originator == nil && fwm.OriginatorOptionF == nil {
			return fieldError("Originator or Originator Option F", ErrFieldRequired)
		}
	} else {
		if fwm.Originator == nil {
			return fieldError("Originator", ErrFieldRequired)
		}
	}
	return nil
}

func (fwm *FedWireMessage) isFIIntermediaryFIValid() error {
	if fwm.BeneficiaryIntermediaryFI == nil {
		return fieldError("BeneficiaryIntermediaryFI", ErrFieldRequired)
	}
	if fwm.BeneficiaryFI == nil {
		return fieldError("BeneficiaryFI", ErrFieldRequired)
	}
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	return nil
}

func (fwm *FedWireMessage) isFIIntermediaryFIAdviceValid() error {
	if fwm.BeneficiaryIntermediaryFI == nil {
		return fieldError("BeneficiaryIntermediaryFI", ErrFieldRequired)
	}
	if fwm.BeneficiaryFI == nil {
		return fieldError("BeneficiaryFI", ErrFieldRequired)
	}
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	return nil
}

func (fwm *FedWireMessage) isFIBeneficiaryFIValid() error {
	if fwm.BeneficiaryFI == nil {
		return fieldError("BeneficiaryFI", ErrFieldRequired)
	}
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	return nil
}

func (fwm *FedWireMessage) isFIBeneficiaryFIAdviceValid() error {
	if fwm.BeneficiaryFI == nil {
		return fieldError("BeneficiaryFI", ErrFieldRequired)
	}
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	return nil
}

func (fwm *FedWireMessage) isFIBeneficiaryValid() error {
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	return nil
}
func (fwm *FedWireMessage) isFIBeneficiaryAdviceValid() error {
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	return nil
}

func (fwm *FedWireMessage) isFIPaymentMethodToBeneficiaryValid() error {
	if fwm.FIBeneficiary == nil {
		return fieldError("FIBeneficiary", ErrFieldRequired)
	}
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	return nil
}

func (fwm *FedWireMessage) isUnstructuredAddendaValid() error {
	switch fwm.LocalInstrument.LocalInstrumentCode {
	case
		SequenceBCoverPaymentStructured,
		ProprietaryLocalInstrumentCode,
		RemittanceInformationStructured,
		RelatedRemittanceInformation:
			// ToDo may want to also check Addenda Length
		if len(fwm.UnstructuredAddenda.Addenda) > 1 {
			return NewErrInvalidPropertyForProperty("UnstructuredAddenda", fwm.UnstructuredAddenda.Addenda,
				"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode)
		}
	}
	return nil
}

func (fwm *FedWireMessage) isRelatedRemittanceValid() error {
	if fwm.LocalInstrument.LocalInstrumentCode != RelatedRemittanceInformation {
		return NewErrInvalidPropertyForProperty("RelatedRemittance", "RelatedRemittance",
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode)
	}
	return nil
}

func (fwm *FedWireMessage) isRemittanceOriginatorValid() error {
	if fwm.LocalInstrument.LocalInstrumentCode != RemittanceInformationStructured {
		return NewErrInvalidPropertyForProperty("RemittanceOriginator", "RemittanceOriginator",
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode)
	}
	return nil
}

func (fwm *FedWireMessage) isRemittanceBeneficiaryValid() error {
	if fwm.LocalInstrument.LocalInstrumentCode != RemittanceInformationStructured {
		return NewErrInvalidPropertyForProperty("RemittanceBeneficiary", "RemittanceBeneficiary",
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode)
	}
	return nil
}

func (fwm *FedWireMessage) isPrimaryRemittanceDocumentValid() error {
	if fwm.LocalInstrument.LocalInstrumentCode != RemittanceInformationStructured {
		return NewErrInvalidPropertyForProperty("PrimaryRemittanceDocument", "PrimaryRemittanceDocument",
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode)
	}
	return nil
}

func (fwm *FedWireMessage) isActualAmountPaidValid() error {
	if fwm.LocalInstrument.LocalInstrumentCode != RemittanceInformationStructured {
		return NewErrInvalidPropertyForProperty("ActualAmountPaid", "ActualAmountPaid",
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode)
	}
	return nil
}

func (fwm *FedWireMessage) isGrossAmountRemittanceDocumentValid() error {
	if fwm.LocalInstrument.LocalInstrumentCode != RemittanceInformationStructured {
		return NewErrInvalidPropertyForProperty("GrossAmountRemittanceDocument", "GrossAmountRemittanceDocument",
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode)
	}
	return nil
}

func (fwm *FedWireMessage) isAdjustmentValid() error {
	if fwm.LocalInstrument.LocalInstrumentCode != RemittanceInformationStructured {
		return NewErrInvalidPropertyForProperty("Adjustment", "Adjustment",
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode)
	}
	return nil
}

func (fwm *FedWireMessage) isDateRemittanceDocumentValid() error {
	if fwm.LocalInstrument.LocalInstrumentCode != RemittanceInformationStructured {
		return NewErrInvalidPropertyForProperty("DateRemittanceDocument", "DateRemittanceDocument",
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode)
	}
	return nil
}

func (fwm *FedWireMessage) isSecondaryRemittanceDocumentValid() error {
	if fwm.LocalInstrument.LocalInstrumentCode != RemittanceInformationStructured {
		return NewErrInvalidPropertyForProperty("SecondaryRemittanceDocument", "SecondaryRemittanceDocument",
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode)
	}
	return nil
}

func (fwm *FedWireMessage) isRemittanceFreeTextValid() error {
	if fwm.LocalInstrument.LocalInstrumentCode != RemittanceInformationStructured {
		return NewErrInvalidPropertyForProperty("RemittanceFreeText", "RemittanceFreeText",
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode)
	}
	return nil
}