import axios from 'axios';

//const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'https://starttech-alb-1968314094.us-east-1.elb.amazonaws.com/swagger/index.html';
export const apiClient = axios.create({
    baseURL: API_BASE_URL,
    withCredentials: true, // Crucial for httpOnly cookies
});
