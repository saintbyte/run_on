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
		return Koyeb
	}
	_, exists = os.LookupEnv("RENDER_EXTERNAL_HOSTNAME")
	if exists {
		return Render
	}
	_, exists = os.LookupEnv("X_YCF_WORKER_ID")
	if exists {
		return YandexCloudFunction
	}
	return Unknown
}

func Host() (string, error) {
	// Определяет внешний хост
	hostProvider := Detect()
	var osKey = ""
	if hostProvider == Vercel {
		osKey = "VERCEL_PROJECT_PRODUCTION_URL"
	} else if hostProvider == Koyeb {
		osKey = "KOYEB_PUBLIC_DOMAIN"
	} else if hostProvider == Render {
		osKey = "RENDER_EXTERNAL_HOSTNAME"
	} else if hostProvider == YandexCloudFunction {
		return "", errors.New("Yandex Cloud Function has no external hostname env var")
	}
	if osKey != "" {
		return os.Getenv(osKey), nil
	}
	return "", errors.New("cant detect cloud provider")
}
