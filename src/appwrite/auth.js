import conf from '../conf/conf.js';

// Authentication service for Go backend (formerly Appwrite, now using Go REST API)
export class AuthService {
    constructor() {
        this.apiUrl = conf.goServerUrl;
    }

    // Store tokens in localStorage
    setTokens(accessToken, refreshToken) {
        if (accessToken) localStorage.setItem('accessToken', accessToken);
        if (refreshToken) localStorage.setItem('refreshToken', refreshToken);
    }

    // Get access token
    getAccessToken() {
        return localStorage.getItem('accessToken');
    }

    // Get refresh token
    getRefreshToken() {
        return localStorage.getItem('refreshToken');
    }

    // Clear tokens
    clearTokens() {
        localStorage.removeItem('accessToken');
        localStorage.removeItem('refreshToken');
    }

    // Make API request with auth header
    async apiCall(endpoint, options = {}) {
        const accessToken = this.getAccessToken();
        const headers = {
            'Content-Type': 'application/json',
            ...options.headers,
        };

        if (accessToken) {
            headers['Authorization'] = `Bearer ${accessToken}`;
        }

        const response = await fetch(`${this.apiUrl}/api${endpoint}`, {
            ...options,
            headers,
        });

        if (!response.ok) {
            const error = await response.text();
            throw new Error(error || `HTTP ${response.status}`);
        }

        return response.json();
    }

    async createAccount({ email, password, name }) {
        try {
            const response = await this.apiCall('/register', {
                method: 'POST',
                body: JSON.stringify({ email, password, name }),
            });

            if (response.accessToken && response.refreshToken) {
                this.setTokens(response.accessToken, response.refreshToken);
                // Return user data from getCurrentUser after successful registration
                return this.getCurrentUser();
            }
            return response;
        } catch (error) {
            throw new Error(`Registration failed: ${error.message}`);
        }
    }

    async login({ email, password }) {
        try {
            const response = await this.apiCall('/login', {
                method: 'POST',
                body: JSON.stringify({ email, password }),
            });

            if (response.accessToken && response.refreshToken) {
                this.setTokens(response.accessToken, response.refreshToken);
                return response;
            }
            return response;
        } catch (error) {
            throw new Error(`Login failed: ${error.message}`);
        }
    }

    async getCurrentUser() {
        try {
            const user = await this.apiCall('/me', {
                method: 'GET',
            });
            return user;
        } catch (error) {
            // Silent fail - 401 is expected when user is not logged in
            this.clearTokens();
            return null;
        }
    }

    async logout() {
        try {
            await this.apiCall('/logout', {
                method: 'POST',
            });
            this.clearTokens();
        } catch (error) {
            // Silent fail - logout errors are expected
            this.clearTokens();
        }
    }

    async refreshAccessToken() {
        try {
            const refreshToken = this.getRefreshToken();
            if (!refreshToken) {
                throw new Error('No refresh token available');
            }

            const response = await this.apiCall('/refresh', {
                method: 'POST',
                body: JSON.stringify({ refreshToken }),
            });

            if (response.accessToken) {
                localStorage.setItem('accessToken', response.accessToken);
                return response.accessToken;
            }
            return null;
        } catch (error) {
            // Silent fail - token refresh errors are expected
            this.clearTokens();
            return null;
        }
    }
}

const authService = new AuthService();

export default authService