/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/
import axios from 'axios';


const instance = axios.create({
    baseURL: import.meta.env.VITE_API_URL || 'http://localhost:3000/api/v1',
});

export default instance;