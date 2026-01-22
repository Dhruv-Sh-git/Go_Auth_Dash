# Auth Dashboard Frontend

A modern authentication dashboard built with Next.js 15, TypeScript, and Tailwind CSS, connected to a Go backend API.

## Features

- 🔐 **User Authentication**: Register and login with JWT tokens
- 🎨 **Modern UI**: Beautiful, responsive design with Tailwind CSS
- 🌙 **Dark Mode**: Full dark mode support
- 🔒 **Protected Routes**: Dashboard accessible only to authenticated users
- ⚡ **Fast Performance**: Built with Next.js 15 and React 19
- 🎯 **Type Safety**: Full TypeScript support

## Project Structure

```
frontend/
├── app/
│   ├── dashboard/         # Protected dashboard page
│   │   └── page.tsx
│   ├── login/            # Login page
│   │   └── page.tsx
│   ├── register/         # Registration page
│   │   └── page.tsx
│   ├── globals.css       # Global styles
│   ├── layout.tsx        # Root layout with AuthProvider
│   └── page.tsx          # Home/landing page
├── context/
│   └── AuthContext.tsx   # Authentication context and hooks
├── lib/
│   └── api.ts           # API service layer with axios
└── .env.local           # Environment variables
```

## Getting Started

### Prerequisites

- Node.js 18+ installed
- Backend API running on `http://localhost:8080`

### Installation

1. Install dependencies:
```bash
npm install
```

2. Configure environment variables:
The `.env.local` file is already configured with:
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

3. Run the development server:

```bash
npm run dev
```

4. Open [http://localhost:3000](http://localhost:3000) in your browser

## API Integration

The frontend connects to the following backend endpoints:

### Authentication Endpoints
- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login and receive JWT token

### User Endpoints (Protected)
- `GET /api/user/me` - Get current user information

## Pages

### Home Page (`/`)
- Landing page with feature overview
- Navigation to login/register pages
- Tech stack information

### Register Page (`/register`)
- User registration form
- Name, email, and password fields
- Password confirmation validation
- Auto-redirect to dashboard after successful registration

### Login Page (`/login`)
- User login form
- Email and password fields
- Token stored in localStorage
- Redirect to dashboard on success

### Dashboard Page (`/dashboard`)
- Protected route (requires authentication)
- Displays user information
- Profile card with name, email, and ID
- Status cards showing account info
- Logout functionality

## Authentication Flow

1. **Registration**:
   - User fills registration form
   - Frontend sends POST to `/api/auth/register`
   - Auto-login after successful registration
   - Redirect to dashboard

2. **Login**:
   - User enters credentials
   - Frontend sends POST to `/api/auth/login`
   - JWT token stored in localStorage
   - Fetch user data with token
   - Redirect to dashboard

3. **Protected Routes**:
   - Dashboard checks for authentication
   - Token sent with all API requests via axios interceptor
   - Redirect to login if not authenticated

4. **Logout**:
   - Clear token from localStorage
   - Clear user state
   - Redirect to login page

## Tech Stack

- **Framework**: Next.js 15 (App Router)
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **HTTP Client**: Axios
- **Authentication**: JWT with localStorage
- **State Management**: React Context API

## Available Scripts

```bash
# Development server
npm run dev

# Production build
npm run build

# Start production server
npm start

# Lint code
npm run lint
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `NEXT_PUBLIC_API_URL` | Backend API base URL | `http://localhost:8080` |

## Notes

- JWT tokens are stored in localStorage
- Axios interceptor automatically adds Bearer token to requests
- Authentication state persists across page reloads
- Protected routes redirect to login when not authenticated

## Backend Integration

Ensure your Go backend is running with CORS enabled for the frontend origin. The backend should be accessible at the URL specified in `.env.local`.

## Troubleshooting

### Cannot connect to backend
- Verify backend is running on port 8080
- Check `.env.local` has correct `NEXT_PUBLIC_API_URL`
- Ensure backend has CORS properly configured

### Token issues
- Clear localStorage and login again
- Check JWT_SECRET matches between frontend expectations and backend

### Build errors
- Run `npm install` to ensure all dependencies are installed
- Delete `.next` folder and rebuild

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.
