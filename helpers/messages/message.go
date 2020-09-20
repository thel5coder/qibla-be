package messages

var (
	DataNotFound           = "data not found"
	DeleteFailed           = "delete data failed"
	AddFailed              = "add data failed"
	EditFailed             = "edit data failed"
	DataAlreadyExist       = "data already exist"
	PhoneAlreadyExist      = "nomor handphone sudah ada"
	EmailAlreadyExist      = "email already exist"
	UserNameExist          = "username already exist"
	AddUserFailed          = "add user failed"
	UserNotFound           = "user not found"
	UserIsNotActive        = "user has not been activated"
	UserIsBlocked          = "user has been blocked"
	CredentialDoNotMatch   = "kredensial tidak sesuai"
	PasswordNotMatch       = "password not match"
	OtpStoreToRedis        = "otp failed store to redis"
	InvalidOtpStoreToRedis = "invalid otp failed store to redis"
	PushQueueFailed        = "push queue failed"
	InvalidToken           = "invalid token" //
	AuthHeaderNotPresent   = "authorization header not present"
	ExpiredToken           = "expired token"
	FailedLoadPayload      = "failed load payload"
	InvalidSession         = "invalid session"
	InternalServer         = "internal server"
	MaxRetryKey            = "max retry key"
	InvalidKey             = "invalid key"
	PhoneNotFound          = "phone not found"
	Unauthorized           = "unauthorized"
	ExpiredOtp             = "expired otp"
	IncompleteProfile      = "incomplete profile"
	DateInvalid            = "invalid date"
	RoleNotFound           = "role not found"
	OtpNotMatch            = "otp not match"
	PaymentFailed          = "payment failed"
	SameMethod             = "same method"
)
