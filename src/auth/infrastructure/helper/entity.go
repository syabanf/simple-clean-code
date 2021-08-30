package helper

import jwt "github.com/dgrijalva/jwt-go"

const (
	RequestTypeChangePassword = "change"
	RequestTypeForgotPassword = "forgot"
	RequestTypeRegistration   = "register"

	SubjectRegister       = "Verify Registration"
	SubjectForgotPassword = "Verify Forgot Password"

	ResponseEmailFormat           = "invalid mail format"
	ResponsePasswordLength        = "password minimum 8 character"
	ResponsePasswordFormat        = "password has to contain letter and number"
	ResponsePasswordDivergent     = "password is different from confirmation"
	ResponsePhoneNumberLetter     = "phone number cannot contain letter"
	ResponsePhoneNumberTooShort   = "phone number too short"
	ResponsePhoneNumberTooLong    = "phone number too long"
	ResponseNameFormat            = "name can only contain letter"
	ResponseNameTooShort          = "name too short"
	ResponseInvalidOTP            = "invalid otp"
	ResponseVerificationMailSent  = "verification mail sent"
	ResponseSessionExpired        = "session expired"
	ResponseAuthenticationInvalid = "invalid authentication"
	ResponseUpdatePasswordFailed  = "update password failed"
	ResponseInvalidSource         = "invalid source"
	ResponseInvalidType           = "invalid type"
	ResponseRegistered            = "registered"

	ResponseEmailPhoneHasExists = "email or phone number is exist"
	ResponseFailedRegister      = "failed register"

	ResponseFailedAuthFacebook = "failed to authenticate to facebook"
	ResponseFailedAuthGoogle   = "failed to authenticate to google"

	SourceTypeWebsite = "website"
	SourceTypeAndroid = "android"
	SourceTypeIOS     = "ios"

	StatusForgotPasswordChanged    = "password_changed"
	StatusForgotPasswordUnverified = "unverified"
	StatusForgotPasswordVerified   = "verified"
)

type AuthClaims struct {
	ID         string `json:"id"`
	RoleID     int64  `json:"role_id"`
	IsVerified int64  `json:"is_verified"`
	jwt.StandardClaims
}
