package helper

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegisterValidation(validate *validator.Validate) {
	validate.RegisterValidation("account_type", validateAccountType)
	validate.RegisterValidation("discount_proposal_confirmation_status_type", validateDiscountProposalConfirmationStatusType)
	validate.RegisterValidation("discount_proposal_confirmation_status_unit", validateDiscountProposalConfirmationStatusUnit)
	validate.RegisterValidation("discount_proposal_transfer_type", validateDiscountProposalTransferType)
	validate.RegisterValidation("gender", validateGender)
	validate.RegisterValidation("ktp", validateKtp)
	validate.RegisterValidation("npwp", validateNpwp)
	validate.RegisterValidation("period_day", validatePeriodDay)
	validate.RegisterValidation("period_month", validatePeriodMonth)
	validate.RegisterValidation("religion", validateReligion)
}

func validateNpwp(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	if len(value) != 20 {
		return false
	}

	regex, _ := regexp.Compile("[0-9]{2}[.][0-9]{3}[.][0-9]{3}[.][0-9]{1}[-][0-9]{3}[.][0-9]{3}")
	result := regex.MatchString(value)
	return result
}

func validateKtp(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	if value == "" {
		return true
	}

	if len(value) != 16 {
		return false
	}

	regex, _ := regexp.Compile("(1[1-9]|21|[37][1-6]|5[1-3]|6[1-5]|[89][12])\\d{2}\\d{2}([04][1-9]|[1256][0-9]|[37][01])(0[1-9]|1[0-2])\\d{2}\\d{4}")
	result := regex.MatchString(value)
	return result
}

func validateAccountType(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	if value == "" {
		return true
	}

	regex, _ := regexp.Compile("(RKC|RBO)")
	result := regex.MatchString(value)
	return result
}

func validateGender(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	regex, _ := regexp.Compile("(Laki-laki|Perempuan)")
	result := regex.MatchString(value)
	return result
}

func validateReligion(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	regex, _ := regexp.Compile("(Islam|Katolik|Protestan|Hindu|Buddha|Konghucu)")
	result := regex.MatchString(value)
	return result
}

func validatePeriodDay(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	if len(value) != 8 {
		return false
	}

	regex, _ := regexp.Compile("\\d{4}(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])")
	result := regex.MatchString(value)
	return result
}

func validateDiscountProposalConfirmationStatusUnit(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	regex, _ := regexp.Compile("(Days|Minutes|Hours)")
	result := regex.MatchString(value)
	return result
}

func validateDiscountProposalConfirmationStatusType(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	regex, _ := regexp.Compile("(SKI|DPF)")
	result := regex.MatchString(value)
	return result
}

func validateDiscountProposalTransferType(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	regex, _ := regexp.Compile("(BO|504)")
	result := regex.MatchString(value)
	return result
}

func validatePeriodMonth(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	if len(value) != 6 {
		return false
	}

	regex, _ := regexp.Compile("\\d{4}(0[1-9]|1[012])")
	result := regex.MatchString(value)
	return result
}
