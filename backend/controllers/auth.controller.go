package controllers

import (
	"net/http"

	"auth-dashboard-backend/config"
	"auth-dashboard-backend/models"
	"auth-dashboard-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Register handles user registration
func Register(c *gin.Context) {
	var input models.UserRegisterInput

	// Validate request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email already exists
	config.Store.Mu.RLock()
	for _, user := range config.Store.Users {
		if user.Email == input.Email {
			config.Store.Mu.RUnlock()
			c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
			return
		}
	}
	config.Store.Mu.RUnlock()

	// Hash the password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create new user
	userID := uuid.New().String()
	user := &config.User{
		ID:       userID,
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	// Store user in memory
	config.Store.Mu.Lock()
	config.Store.Users[userID] = user
	config.Store.Mu.Unlock()

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login handles user login
func Login(c *gin.Context) {
	var input models.UserLoginInput

	// Validate request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	config.Store.Mu.RLock()
	var foundUser *config.User
	for _, user := range config.Store.Users {
		if user.Email == input.Email {
			foundUser = user
			break
		}
	}
	config.Store.Mu.RUnlock()

	if foundUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Verify password
	if !utils.CheckPasswordHash(input.Password, foundUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(foundUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetMe returns the current authenticated user's information
func GetMe(c *gin.Context) {
	// Get user ID from context (set by JWT middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Fetch user from memory
	config.Store.Mu.RLock()
	user, found := config.Store.Users[userID.(string)]
	config.Store.Mu.RUnlock()

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}
