package utils

import (
    "golang.org/x/crypto/bcrypt"
    "log"
)

// HashPassword toma una contraseña en texto claro y devuelve su hash
func HashPassword(password string) (string, error) {
    // Generar el hash de la contraseña
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Println("Error hashing password:", err)
        return "", err
    }
    return string(bytes), nil
}

// VerifyPassword compara un hash y una contraseña en texto claro
func VerifyPassword(hashedPassword, password string) bool {
    // Comparar el hash con la contraseña
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}
