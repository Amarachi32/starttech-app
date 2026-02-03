import axios from 'axios';

//const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'https://d3qh39lmclj4j6.cloudfront.net/api';
export const apiClient = axios.create({
    baseURL: API_BASE_URL,
    withCredentials: true, // Crucial for httpOnly cookies
});
