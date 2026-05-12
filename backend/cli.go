package main

import (
	"flag"
	"fmt"
	"os"
)

type cliFlags struct {
	ConfigPath   string
	DataDir      string
	AuthUserPass string
	BasePath     string
	APIBaseURL   string
	S3Endpoint   string
	S3Region     string
	APIAdminKey  string
	CookieSecure string
	Host         string
	Port         string
}

func parseFlags() *cliFlags {
	var f cliFlags

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Garage WebUI\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\nCLI flags take precedence over environment variables.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Boolean flags accept: true, TRUE, 1, false, FALSE, 0.\n")
	}

	flag.StringVar(&f.ConfigPath, "config", "", "Path to garage.toml (env: CONFIG_PATH, default: /etc/garage.toml)")
	flag.StringVar(&f.DataDir, "data-dir", "", "Enable multi-user SQLite mode (env: DATA_DIR)")
	flag.StringVar(&f.AuthUserPass, "auth", "", "Credentials as user:bcrypt_hash (env: AUTH_USER_PASS)")
	flag.StringVar(&f.BasePath, "base-path", "", "URL prefix for reverse proxy (env: BASE_PATH)")
	flag.StringVar(&f.APIBaseURL, "api-url", "", "Garage Admin API URL (env: API_BASE_URL)")
	flag.StringVar(&f.S3Endpoint, "s3-url", "", "Garage S3 API URL (env: S3_ENDPOINT_URL)")
	flag.StringVar(&f.S3Region, "s3-region", "", "S3 region (env: S3_REGION, default: garage)")
	flag.StringVar(&f.APIAdminKey, "admin-key", "", "Garage Admin API bearer token (env: API_ADMIN_KEY)")
	flag.StringVar(&f.CookieSecure, "cookie-secure", "", "Set session cookie Secure flag (env: COOKIE_SECURE, default: true)")
	flag.StringVar(&f.Host, "host", "", "Listen address (env: HOST, default: 0.0.0.0)")
	flag.StringVar(&f.Port, "port", "", "Listen port (env: PORT, default: 3909)")

	flag.Parse()
	return &f
}

func (f *cliFlags) applyToEnv() {
	setEnv("CONFIG_PATH", f.ConfigPath)
	setEnv("DATA_DIR", f.DataDir)
	setEnv("AUTH_USER_PASS", f.AuthUserPass)
	setEnv("BASE_PATH", f.BasePath)
	setEnv("API_BASE_URL", f.APIBaseURL)
	setEnv("S3_ENDPOINT_URL", f.S3Endpoint)
	setEnv("S3_REGION", f.S3Region)
	setEnv("API_ADMIN_KEY", f.APIAdminKey)
	setEnv("COOKIE_SECURE", f.CookieSecure)
	setEnv("HOST", f.Host)
	setEnv("PORT", f.Port)
}

func setEnv(key, value string) {
	if value != "" {
		os.Setenv(key, value)
	}
}
