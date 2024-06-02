package authServices

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	userModel "github.com/NabinGrz/SocialMedia/src/authentication/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

var jwtkey = []byte("N8Sns89nS2ISB09sn290bSkSHJJ2SNoiS09")

func IsValid(user userModel.User) error {
	if user.Email == "" || user.Password == "" || user.FullName == "" || user.Username == "" {

		return errors.New("email/password/username/fullname field is required")
	} else {
		return nil
	}
}

// isValidEmail checks if the email provided is a valid email format
func isValidEmail(email string) bool {
	// Define the regex pattern for a valid email address
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex pattern
	re := regexp.MustCompile(emailRegexPattern)

	// Match the email against the regex pattern
	return re.MatchString(email)
}
func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func GenerateJWT(user userModel.User) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userid": user.ID,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(user userModel.User, db *gorm.DB) (any, error) {

	emptyError := IsValid(user)
	if emptyError != nil {
		return nil, emptyError
	}
	isValid := isValidEmail(user.Email)
	if !isValid {
		return nil, errors.New("Invalid email address")
	}
	if foundUser := db.Find(&user, "email = ?", user.Email).RowsAffected; foundUser != 0 {
		return nil, errors.New("Email already registered")
	}

	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	user.ID = uuid.New()
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	m := map[string]any{
		"message": "User registered successfully",
		"user":    user,
	}
	return m, nil
}

func Login(user userModel.User, db *gorm.DB) (map[string]interface{}, error) {
	var foundUser userModel.User
	isValid := isValidEmail(user.Email)

	if user.Email == "" || user.Password == "" {
		return nil, errors.New("email/password field is required")
	}
	if !isValid {
		return nil, errors.New("invalid email address")
	}

	if err := db.First(&foundUser, "email = ?", user.Email).Error; err != nil {
		return nil, errors.New("user Not found")
	}
	match := VerifyPassword(user.Password, foundUser.Password)

	if match {
		token, _ := GenerateJWT(foundUser)
		return map[string]interface{}{"token": token}, nil
	}
	return nil, errors.New("invalid Credentials")
}

// JWT middleware function to verify token and extract user details
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the JWT token from the request headers
		header := c.GetHeader("Authorization")
		tokenString := ""
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}
		tokenString = header[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtkey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		fmt.Println(token)
		if token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				userid := claims["userid"]
				c.Set("userid", userid)
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				c.Abort()
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}
