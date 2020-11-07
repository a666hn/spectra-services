package model

import "time"

type SpectraConfig struct {
	AppName       string
	GRPCTimeout   time.Duration
	CacheExpiry   time.Duration
	CacheCleanup  time.Duration
	ClientURI     string
	FileURI       string
	AllowedOrigin []string
	SmtpConfig    SmtpConfig
}

type SmtpConfig struct {
	SmptpHOST      string
	SmtpPORT       int16
	EmailSender    string
	PasswordSender string
}
