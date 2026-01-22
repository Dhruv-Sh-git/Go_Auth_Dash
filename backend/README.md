# Auth Dashboard Backend - Go

A production-ready authentication backend built with Go featuring JWT-based authentication, user registration, and secure password handling.

## 🚀 Features

- ✅ User Registration with email validation
- ✅ User Login with JWT token generation
- ✅ JWT-based authentication middleware
- ✅ Protected API endpoints
- ✅ Password hashing with bcrypt
- ✅ PostgreSQL database with GORM
- ✅ CORS enabled for frontend integration
- ✅ Clean architecture with modular structure

## 📂 Project Structure

```
backend/
├── main.go                      # Application entry point
├── config/
│   ├── db.go                   # Database connection
│   └── env.go                  # Environment configuration
├── models/
│   └── user.go                 # User model and DTOs
├── controllers/
│   └── auth.controller.go      # Auth handlers (register, login, me)
├── middleware/
│   └── jwt.middleware.go       # JWT authentication middleware
├── routes/
│   └── routes.go               # API route definitions
├── utils/
│   ├── hash.go                 # Password hashing utilities
│   └── jwt.go                  # JWT token utilities
├── .env                        # Environment variables
└── go.mod                      # Go dependencies
```

## 🛠️ Tech Stack

- **Language**: Go 1.25+
- **Web Framework**: Gin
- **Authentication**: JWT (golang-jwt/jwt/v5)
- **Password Hashing**: bcrypt
- **Database**: PostgreSQL
- **ORM**: GORM
- **Environment**: godotenv

## ⚙️ Setup Instructions

### Prerequisites

- Go 1.25 or higher
- PostgreSQL database

### 1. Configure Environment Variables

Edit the `.env` file:

```env
PORT=8080
DB_URL=postgres://user:pass@localhost:5432/authdb
JWT_SECRET=supersecretkey
```

**Important**: Replace the database URL and JWT secret with your actual values.

### 2. Install Dependencies

```bash
go mod download
```

### 3. Run the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## 📡 API Endpoints

### Public Endpoints

#### Register User
```http
POST /api/auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (201 Created)**:
```json
{
  "message": "User registered successfully"
}
```

#### Login
```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (200 OK)**:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Protected Endpoints

#### Get Current User
```http
GET /api/user/me
Authorization: Bearer <your_jwt_token>
```

**Response (200 OK)**:
```json
{
  "id": "uuid-here",
  "name": "John Doe",
  "email": "john@example.com"
}
```

### Health Check
```http
GET /health
```

**Response (200 OK)**:
```json
{
  "status": "healthy",
  "message": "Auth Dashboard Backend API"
}
```

## 🔒 Security Features

- **Password Hashing**: All passwords are hashed using bcrypt with default cost
- **JWT Tokens**: 24-hour expiration, signed with HS256
- **Input Validation**: Request validation using Gin's binding
- **CORS**: Configured for frontend integration
- **Unique Email**: Database constraint ensures email uniqueness

## 🧪 Testing with Postman

1. **Register a User**: Send POST to `/api/auth/register` with user details
2. **Login**: Send POST to `/api/auth/login` to get JWT token
3. **Access Protected Route**: Send GET to `/api/user/me` with `Authorization: Bearer <token>` header

## 📦 Build for Production

```bash
go build -o auth-backend .
./auth-backend
```

## 🗄️ Database Schema

The `users` table is automatically created with the following fields:

- `id` (UUID, Primary Key)
- `name` (VARCHAR, NOT NULL)
- `email` (VARCHAR, UNIQUE, NOT NULL)
- `password` (VARCHAR, NOT NULL, hashed)
- `created_at` (TIMESTAMP)
- `updated_at` (TIMESTAMP)
- `deleted_at` (TIMESTAMP, nullable - soft delete)

## 🚨 Error Handling

The API returns appropriate HTTP status codes:

- `200` - Success
- `201` - Created
- `400` - Bad Request (validation error)
- `401` - Unauthorized (invalid/missing token)
- `404` - Not Found
- `409` - Conflict (email already exists)
- `500` - Internal Server Error

## 📝 Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `PORT` | Server port | Yes (default: 8080) |
| `DB_URL` | PostgreSQL connection string | Yes |
| `JWT_SECRET` | Secret key for JWT signing | Yes |

## 🎯 Success Criteria

✅ User registration with email validation  
✅ Secure password hashing  
✅ JWT token generation and validation  
✅ Protected routes with middleware  
✅ Clean, modular code structure  
✅ Proper error handling  
✅ CORS enabled  
✅ Database migrations  

## 📄 License

MIT License

## 👨‍💻 Development

To modify the codebase:

1. Controllers: Add new handlers in `controllers/`
2. Routes: Register routes in `routes/routes.go`
3. Models: Define new models in `models/`
4. Middleware: Add middleware in `middleware/`
5. Utilities: Add helper functions in `utils/`

---

**Built with ❤️ using Go and Gin**
