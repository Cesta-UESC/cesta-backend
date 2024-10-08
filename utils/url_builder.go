package utils

import (
    "fmt"
    "os"
)

func ConnectionURLBuilder(n string) (string, error) {
    var url string

    switch n {
    case "mysql":
        url = fmt.Sprintf(
            "%s:%s@tcp(%s:%s)/%s",
            os.Getenv("DB_USER"),
            os.Getenv("DB_PASSWORD"),
            os.Getenv("DB_HOST"),
            os.Getenv("DB_PORT"),
            os.Getenv("DB_NAME"),
        )
    /* case "redis":
        url = fmt.Sprintf(
            "%s:%s",
            os.Getenv("REDIS_HOST"),
            os.Getenv("REDIS_PORT"),
        ) */
    default:
        return "", fmt.Errorf("connection name '%v' is not supported", n)
    }

    return url, nil
}
