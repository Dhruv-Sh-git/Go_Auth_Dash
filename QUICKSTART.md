# Go Auth Dashboard - Quick Start Guide

This project consists of a Go backend with a Next.js frontend for secure user authentication.

## 🚀 Quick Start

### 1. Start the Backend

```bash
cd backend
go run main.go
```

The backend will run on **http://localhost:8080**

### 2. Start the Frontend

In a new terminal:

```bash
cd frontend
npm run dev
```

The frontend will run on **http://localhost:3000**

### 3. Test the Application

1. Open **http://localhost:3000** in your browser
2. Click "Get Started" to register a new account
3. Fill in the registration form
4. You'll be automatically logged in and redirected to the dashboard
5. View your profile information on the dashboard
6. Click "Logout" to sign out

## 📁 Project Structure

```
go-auth-dashboard/
├── backend/                 # Go backend with Gin
│   ├── config/             # Database and environment config
│   ├── controllers/        # Request handlers
│   ├── middleware/         # JWT authentication middleware
│   ├── models/             # Data models
│   ├── routes/             # API routes
│   ├── utils/              # Helper functions (hash, JWT)
│   ├── .env                # Environment variables
│   └── main.go             # Entry point
│
└── frontend/               # Next.js frontend
    ├── app/                # App router pages
    │   ├── dashboard/      # Protected dashboard
    │   ├── login/          # Login page
    │   └── register/       # Registration page
    ├── context/            # React Context for auth
    ├── lib/                # API service layer
    └── .env.local          # Frontend environment variables
```

## 🔌 API Endpoints

### Public Endpoints
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login user

### Protected Endpoints (Requires JWT Token)
- `GET /api/user/me` - Get current user info

## 🔐 Authentication Flow

1. **Register**: User creates account → Auto-login → Dashboard
2. **Login**: User enters credentials → Receives JWT token → Dashboard
3. **Protected Routes**: Token stored in localStorage → Sent with API requests
4. **Logout**: Clear token → Redirect to login

## 🛠️ Tech Stack

### Backend
- **Language**: Go
- **Framework**: Gin
- **Auth**: JWT + bcrypt
- **Storage**: In-memory (can be replaced with database)

### Frontend
- **Framework**: Next.js 15
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **HTTP Client**: Axios
- **State**: React Context API

## 📝 Environment Variables

### Backend (.env)
```env
PORT=8080
JWT_SECRET=your_super_secret_jwt_key_change_this_in_production
```

### Frontend (.env.local)
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## 🎯 Features

✅ User registration with validation  
✅ User login with JWT tokens  
✅ Protected dashboard route  
✅ Persistent authentication (localStorage)  
✅ Automatic token refresh  
✅ Logout functionality  
✅ Beautiful, responsive UI  
✅ Dark mode support  
✅ Password hashing with bcrypt  
✅ CORS enabled  
✅ Type-safe with TypeScript  

## 🔧 Development

### Backend Commands
```bash
cd backend

# Run in development
go run main.go

# Build executable
go build -o auth-backend.exe .

# Run executable
./auth-backend.exe
```

### Frontend Commands
```bash
cd frontend

# Install dependencies
npm install

# Development server
npm run dev

# Production build
npm run build

# Run production server
npm start
```

## 📚 Testing with Postman

A Postman collection is included in `backend/postman_collection.json`. Import it to test the API endpoints directly.

## 🐛 Troubleshooting

### Backend won't start
- Ensure port 8080 is not in use
- Check `.env` file exists with correct values
- Run `go mod tidy` to sync dependencies

### Frontend can't connect to backend
- Verify backend is running on port 8080
- Check `.env.local` has correct API URL
- Clear browser localStorage and try again

### Authentication errors
- Clear browser localStorage
- Check JWT_SECRET is set in backend `.env`
- Verify CORS is enabled in backend

## 🚀 Next Steps

- Add database integration (PostgreSQL, MongoDB, etc.)
- Implement password reset functionality
- Add email verification
- Implement refresh tokens
- Add user profile editing
- Add role-based access control (RBAC)
- Deploy to production

## 📄 License

This project is open source and available for learning purposes.
