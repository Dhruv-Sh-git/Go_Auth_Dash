import axios from 'axios';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

// Create axios instance with base configuration
const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add request interceptor to include JWT token
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Types
export interface RegisterData {
  name: string;
  email: string;
  password: string;
}

export interface LoginData {
  email: string;
  password: string;
}

export interface User {
  id: string;
  name: string;
  email: string;
}

// API Functions
export const authApi = {
  // Register a new user
  register: async (data: RegisterData): Promise<{ message: string }> => {
    const response = await api.post('/api/auth/register', data);
    return response.data;
  },

  // Login user
  login: async (data: LoginData): Promise<{ token: string }> => {
    const response = await api.post('/api/auth/login', data);
    return response.data;
  },

  // Get current user info
  getMe: async (): Promise<User> => {
    const response = await api.get('/api/user/me');
    return response.data;
  },
};

export default api;
