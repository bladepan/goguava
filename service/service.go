package service

type Service interface {
	// might be a blocking function
	Start() error
	Stop() error
	// TODO should i put dependencies as part of the signature of service?
}
