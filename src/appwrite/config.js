import conf from '../conf/conf.js';

// Service layer for Go backend API calls (formerly Appwrite, now using Go REST API)
export class Service {
    constructor() {
        this.apiUrl = conf.goServerUrl;
    }

    // Helper method to get auth token
    getAuthToken() {
        return localStorage.getItem('accessToken');
    }

    // Helper method to make authenticated requests
    async authenticatedFetch(endpoint, options = {}) {
        const token = this.getAuthToken();
        const headers = {
            ...options.headers,
        };

        if (token) {
            headers['Authorization'] = `Bearer ${token}`;
        }

        // Only set Content-Type if not FormData (FormData sets its own boundary)
        if (!(options.body instanceof FormData)) {
            headers['Content-Type'] = 'application/json';
        }

        const response = await fetch(`${this.apiUrl}/api${endpoint}`, {
            ...options,
            headers,
        });

        if (!response.ok) {
            const error = await response.text();
            throw new Error(error || `HTTP ${response.status}`);
        }

        // Always try to parse JSON for successful responses
        try {
            const data = await response.json();
            return data;
        } catch (e) {
            // If not JSON, return the response object
            return response;
        }
    }

    // ================= CREATE POST =================
    async createPost({ title, slug, content, status, userId, image }) {
        try {
            const formData = new FormData();
            formData.append('title', title);
            formData.append('slug', slug);
            formData.append('content', content);
            formData.append('status', status);

            // Append image file if provided
            if (image && image[0]) {
                formData.append('image', image[0]);
            } else if (image instanceof File) {
                formData.append('image', image);
            }

            const response = await this.authenticatedFetch('/posts', {
                method: 'POST',
                body: formData,
                // Don't set Content-Type header; browser will set it with boundary
                headers: {},
            });

            return response;
        } catch (error) {
            console.error("Service :: createPost :: error", error);
            throw error;
        }
    }

    // ================= UPDATE POST =================
    async updatePost(slug, { title, content, status, image }) {
        try {
            const formData = new FormData();
            
            if (title) formData.append('title', title);
            if (content) formData.append('content', content);
            if (status) formData.append('status', status);

            // Append image file if provided
            if (image && image[0]) {
                formData.append('image', image[0]);
            } else if (image instanceof File) {
                formData.append('image', image);
            }

            const response = await this.authenticatedFetch(`/posts/${slug}`, {
                method: 'PUT',
                body: formData,
                // Don't set Content-Type header; browser will set it with boundary
                headers: {},
            });

            return response;
        } catch (error) {
            console.error("Service :: updatePost :: error", error);
            throw error;
        }
    }

    // ================= DELETE POST =================
    async deletePost(slug) {
        try {
            const response = await this.authenticatedFetch(`/posts/${slug}`, {
                method: 'DELETE',
            });

            return true;
        } catch (error) {
            console.error("Service :: deletePost :: error", error);
            return false;
        }
    }

    // ================= GET SINGLE POST =================
    async getPost(slug) {
        try {
            const response = await this.authenticatedFetch(`/posts/${slug}`, {
                method: 'GET',
            });

            return response;
        } catch (error) {
            console.error("Service :: getPost :: error", error);
            throw error;
        }
    }

    // ================= GET ALL POSTS =================
    async getPosts(queries = []) {
        try {
            // Go backend returns array of posts directly
            const response = await this.authenticatedFetch('/posts', {
                method: 'GET',
            });

            
            // Ensure we always return an array
            if (Array.isArray(response)) {
                return response;
            } else if (response && response.documents) {
                return response.documents;
            }
            return [];
        } catch (error) {
            console.error("Service :: getPosts :: error", error);
            return [];
        }
    }

    // ================= GET FILE PREVIEW =================
    getFilePreview(imageUrl) {
        if (!imageUrl) return null;
        // If it's already a full URL (http/https), return as is
        if (imageUrl.startsWith('http://') || imageUrl.startsWith('https://')) {
            return imageUrl;
        }
        // Otherwise prepend the server URL
        return `${this.apiUrl}${imageUrl}`;
    }
}

const service = new Service();
export default service;
