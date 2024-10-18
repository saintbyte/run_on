package run_on

import (
	"errors"
	"os"
)

func Detect() string {
	// Определяет по переменым среды на какой платформе запущен скрипт.
	_, exists := os.LookupEnv("VERCEL_PROJECT_PRODUCTION_URL")
	if exists {
		return Vercel
	}
	_, exists = os.LookupEnv("KOYEB_PUBLIC_DOMAIN")
	if exists {
		return Koeyb
	}
	_, exists = os.LookupEnv("RENDER_EXTERNAL_HOSTNAME")
	if exists {
		return Render
	}
	return Unknown
}

func Host() (string, error) {
	// Определяет внешний хост
	hostProvider := Detect()
	var osKey = ""
	if hostProvider == Vercel {
		osKey = "VERCEL_PROJECT_PRODUCTION_URL"
	} else if hostProvider == Koeyb {
		osKey = "KOYEB_PUBLIC_DOMAIN"
	} else if hostProvider == Render {
		osKey = "RENDER_EXTERNAL_HOSTNAME"
	} else {

	}

	if osKey != "" {
		return os.Getenv("RENDER_EXTERNAL_HOSTNAME"), nil
	}

	return "", errors.New("Cant detect cloud provider")
}
