# 📝 12MegaBlog - Full Stack Blog Application

A modern, full-stack blogging platform built with **React** frontend and **Go** backend, featuring JWT-based authentication, rich text editing with TinyMCE, and image uploads via Cloudinary.

---

## 🎯 Project Overview

**12MegaBlog** is a complete blog application where users can:
- Register and authenticate with secure JWT tokens
- Create, read, update, and delete blog posts
- Upload featured images via Cloudinary
- Write rich, formatted content using TinyMCE editor
- Manage their own posts in a dedicated dashboard
- Browse and view all published posts from the community
- Automatic token refresh for seamless sessions

**Migration Note:** Recently migrated from Appwrite to a custom Go backend for greater control, flexibility, and performance.

---

## 🛠️ Technology Stack

### **Frontend**
- **React 18.3** - UI library
- **React Router v7** - Client-side routing
- **Redux Toolkit** - State management
- **React Hook Form** - Form handling
- **TinyMCE React** - Rich text editor
- **Tailwind CSS** - Styling
- **Vite** - Build tool

### **Backend**
- **Go 1.x** - Server language
- **Gorilla Mux** - HTTP router
- **MongoDB** - NoSQL database
- **JWT** - Token-based authentication
- **Cloudinary SDK** - Image storage and processing
- **BCrypt** - Password hashing

### **Database & Services**
- **MongoDB** - Document database for users and posts
- **Cloudinary** - Cloud-based image storage and optimization
- **JWT (JSON Web Tokens)** - Stateless, token-based authentication

### **DevOps & Tools**
- **Git** - Version control
- **npm** - Package manager (frontend)
- **Go Modules** - Package manager (backend)
- **Vite** - Fast frontend build tool

---

## � Migration from Appwrite to Go Backend

This project was **originally built with Appwrite** but has been **migrated to a custom Go backend** for the following reasons:

| Aspect | Appwrite | Go Backend |
|--------|----------|-----------|
| **Architecture** | Backend-as-a-Service (BaaS) | Self-hosted REST API |
| **Control** | Limited customization | Full control over implementation |
| **Database** | Integrated database | Flexible MongoDB integration |
| **Authentication** | Appwrite Auth SDK | Custom JWT implementation |
| **File Storage** | Appwrite Storage | Cloudinary integration |
| **Performance** | Managed service | Optimized custom code |
| **Cost** | Subscription-based | Server hosting costs |
| **Scalability** | Automatic | Manual but flexible |

### **What Changed:**
- **Frontend API Layer:** The `src/appwrite/` folder now contains REST client utilities instead of Appwrite SDK calls
- **Backend:** Custom Go REST API with JWT tokens and MongoDB
- **Database:** MongoDB replaces Appwrite's database
- **Authentication:** JWT tokens (access + refresh) instead of Appwrite Auth sessions
- **Compatibility:** Frontend code remains largely unchanged; only the API service layer was updated

---

## �📁 Project Structure

```
12MegaBlog/
├── 📂 src/                          # Frontend (React)
│   ├── 📂 components/               # Reusable React components
│   │   ├── AuthLayout.jsx           # Auth-protected wrapper
│   │   ├── Button.jsx               # Button component
│   │   ├── Input.jsx                # Input field component
│   │   ├── Login.jsx                # Login form
│   │   ├── Logo.jsx                 # Logo component
│   │   ├── PostCard.jsx             # Post card display
│   │   ├── RTE.jsx                  # Rich text editor
│   │   ├── Select.jsx               # Select dropdown
│   │   ├── Signup.jsx               # Signup form
│   │   ├── Header/                  # Header components
│   │   │   ├── Header.jsx           # Navigation header
│   │   │   └── LogoutBtn.jsx        # Logout button
│   │   ├── Footer/                  # Footer components
│   │   │   └── Footer.jsx           # Footer navigation
│   │   ├── container/               # Layout container
│   │   │   └── Container.jsx        # Content wrapper
│   │   └── post-form/               # Post creation/editing
│   │       └── PostForm.jsx         # Post form component
│   ├── 📂 pages/                    # Page components
│   │   ├── Home.jsx                 # Home page (all posts)
│   │   ├── AllPosts.jsx             # My posts page
│   │   ├── AddPost.jsx              # Create post page
│   │   ├── EditPost.jsx             # Edit post page
│   │   ├── Post.jsx                 # Single post view
│   │   ├── Login.jsx                # Login page
│   │   └── Signup.jsx               # Signup page
│   ├── 📂 store/                    # Redux state management
│   │   ├── store.js                 # Redux store
│   │   └── authSlice.js             # Auth state slice
   ├── 📂 appwrite/                 # Go API Service Layer (REST Client)
   │   ├── auth.js                  # Authentication API calls (register, login, logout)
   │   └── config.js                # Post API calls (CRUD operations)
│   ├── 📂 conf/                     # Configuration
│   │   └── conf.js                  # Environment config
│   ├── 📂 utils/                    # Utility functions
│   ├── App.jsx                      # Root component
│   ├── main.jsx                     # Entry point
│   ├── index.css                    # Global styles
│   └── App.css                      # App styles
├── 📂 goServer/                     # Backend (Go)
│   ├── cmd/
│   │   └── main.go                  # Server entry point
│   ├── config/
│   │   ├── database.go              # MongoDB connection
│   │   └── cloudinary.go            # Cloudinary setup
│   ├── internal/
│   │   ├── controllers/             # Business logic
│   │   │   ├── authController.go    # Auth endpoints
│   │   │   └── postController.go    # Post endpoints
│   │   ├── middleware/              # HTTP middleware
│   │   │   ├── auth.go              # JWT verification
│   │   │   ├── cors.go              # CORS handling
│   │   │   └── logger.go            # Request logging
│   │   ├── models/                  # Data models
│   │   │   ├── user.go              # User struct
│   │   │   └── post.go              # Post struct
│   │   ├── routes/                  # Route definitions
│   │   │   ├── routes.go            # Main router
│   │   │   ├── userRoute.go         # Auth routes
│   │   │   └── postRoute.go         # Post routes
│   │   ├── services/                # External services
│   │   │   └── cloudinary_service.go # Image upload
│   │   ├── server/
│   │   │   └── server.go            # Server setup
│   │   └── utils/
│   │       └── jwt.go               # JWT token generation
│   ├── go.mod                       # Go module definition
│   └── go.sum                       # Go dependencies lock
├── package.json                     # Frontend dependencies
├── vite.config.js                   # Vite configuration
├── tailwind.config.js               # Tailwind config
├── postcss.config.js                # PostCSS config
├── eslint.config.js                 # ESLint config
└── README.md                        # This file
```

---

## 🎨 Frontend Architecture

### **Component Hierarchy**

```
App (Root)
├── Header
│   ├── Logo
│   ├── Navigation Links
│   └── LogoutBtn
├── Main Content
│   ├── Home (Landing page - all posts)
│   ├── My Posts (User's posts)
│   ├── Post (Single post view)
│   ├── AddPost
│   │   └── PostForm
│   ├── EditPost
│   │   └── PostForm
│   ├── Login
│   └── Signup
└── Footer
```

### **State Management (Redux)**

```javascript
// Store Structure
{
  auth: {
    status: boolean      // Is user logged in?
    userData: {          // Current user info
      id: string
      email: string
      name: string
      createdAt: number
    }
  }
}
```

### **Authentication Flow**

1. User opens app → `App.jsx` calls `getCurrentUser()`
2. If token exists in localStorage → User auto-logged in
3. If no token → User redirected to Login page
4. On login → Tokens saved to localStorage + Redux state updated
5. Every API request includes `Authorization: Bearer {token}` header
6. Token expires after **1 day** → Auto-logout occurs

### **Routing**

```javascript
/ (Home)                    // Public - all posts
/login                      // Public - login form
/signup                     // Public - signup form
/post/:slug                 // Public - single post view
/my-posts                   // Protected - user's posts
/add-post                   // Protected - create post
/edit-post/:slug            // Protected - edit post
```

### **Key Features**
- ✅ Redux for global auth state
- ✅ Protected routes with AuthLayout
- ✅ React Hook Form for form management
- ✅ TinyMCE for rich text editing
- ✅ Responsive design with Tailwind CSS

---

## 🔌 Frontend API Service Layer

The `src/appwrite/` folder contains the REST API client that communicates with the Go backend:

### **AuthService** (`src/appwrite/auth.js`)
Handles all authentication operations:
- `createAccount(email, password, name)` - Register new user
- `login(email, password)` - Authenticate user
- `logout()` - Clear tokens and logout
- `getCurrentUser()` - Get current authenticated user
- `apiCall(endpoint, options)` - Helper method for authenticated API calls

### **Service** (`src/appwrite/config.js`)
Handles all post/blog operations:
- `createPost(title, slug, content, status, image)` - Create post
- `getPosts()` - Get all active posts
- `getPost(slug)` - Get single post
- `updatePost(slug, title, content, status, image)` - Update post
- `deletePost(slug)` - Delete post
- `authenticatedFetch(endpoint, options)` - Helper for authenticated requests

### **Token Management**
- Tokens stored in `localStorage`
- Access Token: 15-minute expiry (short-lived)
- Refresh Token: 7-day expiry (long-lived)
- Every API request auto-adds `Authorization: Bearer {accessToken}` header

---

## ⚙️ Backend Architecture

### **Server Setup Flow**

```
1. main.go starts server
2. config/database.go → Connect to MongoDB
3. config/cloudinary.go → Initialize Cloudinary
4. internal/routes/routes.go → Setup HTTP router
5. Middleware stack applied:
   - CORS handling
   - Request logging
6. Routes registered:
   - /api/register, /api/login, /api/refresh, /api/logout, /api/me
   - /api/posts (GET, POST, PUT, DELETE)
7. Server listens on :8080
```

### **Data Models**

#### **User Model**
```go
type User struct {
  ID           primitive.ObjectID  // MongoDB _id
  Name         string              // User full name
  Email        string              // User email (unique)
  Password     string              // Bcrypt hashed password
  RefreshToken string              // JWT refresh token
  CreatedAt    int64               // Timestamp
}
```

#### **Post Model**
```go
type Post struct {
  ID        primitive.ObjectID  // MongoDB _id
  Title     string              // Post title
  Slug      string              // URL-friendly identifier (unique)
  Content   string              // HTML content (from TinyMCE)
  ImageURL  string              // Cloudinary URL
  ImageID   string              // Cloudinary public ID
  Status    string              // "active" or "inactive"
  UserID    string              // Post creator's email
  CreatedAt int64               // Timestamp
}
```

### **Middleware Stack**

1. **CORS Middleware** - Allows requests from `http://localhost:5173`
2. **Logger Middleware** - Logs all HTTP requests with method and path
3. **Auth Middleware** - Verifies JWT tokens for protected routes
4. **Options Handler** - Handles CORS preflight requests

### **Service Architecture**

```
HTTP Request
    ↓
Router (Gorilla Mux)
    ↓
Middleware (CORS, Logger, Auth)
    ↓
Controller (Business Logic)
    ↓
Services (Cloudinary, JWT)
    ↓
Models (Data structures)
    ↓
MongoDB
```

---

## 🔌 API Endpoints

### **Base URL**
```
http://localhost:8080/api
```

### **Authentication Endpoints**

#### `POST /register` - Create Account
```json
Request:
{
  "name": "Satyam Tripathi",
  "email": "satyam@gmail.com",
  "password": "securepass123"
}

Response (200):
{
  "accessToken": "eyJhbGc...",
  "refreshToken": "eyJhbGc..."
}
```

#### `POST /login` - Login User
```json
Request:
{
  "email": "satyam@gmail.com",
  "password": "securepass123"
}

Response (200):
{
  "accessToken": "eyJhbGc...",
  "refreshToken": "eyJhbGc..."
}
```

#### `GET /me` - Get Current User
```
Headers: Authorization: Bearer {accessToken}

Response (200):
{
  "id": "69c565e7df49418e252f8b62",
  "name": "Satyam Tripathi",
  "email": "satyam@gmail.com",
  "createdAt": 1774544359
}
```

#### `POST /refresh` - Refresh Token
```json
Request:
{
  "refreshToken": "eyJhbGc..."
}

Response (200):
{
  "accessToken": "eyJhbGc..."
}
```

#### `POST /logout` - Logout User
```
Headers: Authorization: Bearer {accessToken}

Response (200):
{
  "message": "Logged out successfully"
}
```

---

### **Post Endpoints**

#### `POST /posts` - Create Post
```
Headers: Authorization: Bearer {accessToken}
Content-Type: multipart/form-data

Form Data:
- title: "My First Post"
- slug: "my-first-post"
- content: "<p>Post content here</p>"
- status: "active"
- image: [File object]

Response (200):
{
  "id": "69c69b9a1b32c64a93bebac0",
  "title": "My First Post",
  "slug": "my-first-post",
  "content": "<p>Post content here</p>",
  "imageUrl": "https://res.cloudinary.com/.../image.png",
  "imageId": "blog/xyz123",
  "status": "active",
  "userId": "satyam@gmail.com",
  "createdAt": 1774625254
}
```

#### `GET /posts` - Get All Active Posts
```
Response (200):
[
  {
    "id": "69c69b9a1b32c64a93bebac0",
    "title": "My First Post",
    "slug": "my-first-post",
    "content": "<p>Post content</p>",
    "imageUrl": "https://res.cloudinary.com/.../image.png",
    "imageId": "blog/xyz123",
    "status": "active",
    "userId": "satyam@gmail.com",
    "createdAt": 1774625254
  }
]
```

#### `GET /posts/{slug}` - Get Single Post
```
Response (200):
{
  "id": "69c69b9a1b32c64a93bebac0",
  "title": "My First Post",
  "slug": "my-first-post",
  "content": "<p>Post content</p>",
  "imageUrl": "https://res.cloudinary.com/.../image.png",
  "imageId": "blog/xyz123",
  "status": "active",
  "userId": "satyam@gmail.com",
  "createdAt": 1774625254
}
```

#### `PUT /posts/{slug}` - Update Post
```
Headers: Authorization: Bearer {accessToken}
Content-Type: multipart/form-data

Form Data:
- title: "Updated Title" (optional)
- content: "<p>Updated content</p>" (optional)
- status: "active" (optional)
- image: [File object] (optional)

Response (200):
{
  "id": "69c69b9a1b32c64a93bebac0",
  "title": "Updated Title",
  "slug": "my-first-post",
  "content": "<p>Updated content</p>",
  "imageUrl": "https://res.cloudinary.com/.../new-image.png",
  "imageId": "blog/abc789",
  "status": "active",
  "userId": "satyam@gmail.com",
  "createdAt": 1774625254
}
```

#### `DELETE /posts/{slug}` - Delete Post
```
Headers: Authorization: Bearer {accessToken}

Response (200):
{
  "message": "Post deleted"
}
```

---

## 🚀 Setup & Installation

### **Prerequisites**
- Node.js 16+ (Frontend)
- Go 1.18+ (Backend)
- MongoDB (Local or Atlas)
- Cloudinary Account

### **Environment Variables**

#### **Frontend (.env)**
```env
VITE_GO_SERVER_URL=http://localhost:8080
```

#### **Backend (goServer/config/.env)**
```env
PORT=8080
MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/12megablog
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret
JWT_ACCESS_SECRET=your_strong_access_secret_here
JWT_REFRESH_SECRET=your_strong_refresh_secret_here
```

⚠️ **Security Note:** The JWT secrets in `goServer/internal/utils/jwt.go` and `goServer/internal/middleware/auth.go` are currently hardcoded. For production, update these files to read from environment variables (`os.Getenv("JWT_ACCESS_SECRET")` and `os.Getenv("JWT_REFRESH_SECRET")`).

### **Installation Steps**

#### **1. Clone Repository**
```bash
git clone <repo-url>
cd 12MegaBlog
```

#### **2. Frontend Setup**
```bash
# Install dependencies
npm install

# Start development server (runs on port 5173)
npm run dev
```

#### **3. Backend Setup**
```bash
cd goServer

# Download Go dependencies
go mod download

# Run the server (runs on port 8080)
go run cmd/main.go
```

### **Verify Setup**
```bash
# Check Frontend
curl http://localhost:5173

# Check Backend
curl http://localhost:8080/health
# Expected: {"status": "healthy"}
```

---

## 📋 Features

### ✅ **Authentication**
- User registration with email validation
- Secure login with JWT tokens
- Token refresh mechanism
- Automatic logout on token expiry
- Session persistence

### ✅ **Post Management**
- Create posts with rich text editor (TinyMCE)
- Upload featured images to Cloudinary
- Edit own posts
- Delete own posts
- View all published posts
- View single post with full content

### ✅ **User Interface**
- Responsive design (Mobile & Desktop)
- Tailwind CSS styling
- Navigation between pages
- Protected routes
- Beautiful post cards
- Rich text preview

### ✅ **Security**
- Password hashing (bcrypt)
- JWT authentication
- CORS protection
- Authorization checks
- Protected API endpoints

### ✅ **Data Persistence**
- MongoDB database
- File upload to Cloudinary
- User data storage
- Post data storage

---

## 🔐 Authentication & Authorization

### **JWT Tokens**

#### **Access Token** (24 hours)
- Used for API requests
- Sent in `Authorization: Bearer {token}` header
- Stored in localStorage

#### **Refresh Token** (7 days)
- Used to get new access token
- Stored in MongoDB (user document)
- Stored in localStorage (client)

### **Protected Routes**

Routes requiring authentication:
```
/my-posts        - View user's own posts
/add-post        - Create new post
/edit-post/:slug - Edit own post
```

Public routes:
```
/                - Home (all posts)
/post/:slug      - View single post
/login           - Login form
/signup          - Signup form
```

---

## 📊 Database Schema

### **Users Collection**
```javascript
{
  _id: ObjectId,
  name: String,
  email: String (unique),
  password: String (bcrypt hash),
  refreshToken: String,
  createdAt: Number (timestamp)
}
```

### **Posts Collection**
```javascript
{
  _id: ObjectId,
  title: String,
  slug: String (unique),
  content: String (HTML),
  imageUrl: String (Cloudinary URL),
  imageId: String (Cloudinary public_id),
  status: String ("active" | "inactive"),
  userId: String (user email),
  createdAt: Number (timestamp)
}
```

---

## 🔄 Data Flow

### **Creating a Post**
```
1. User fills form in PostForm component
2. Submit → appwriteService.createPost()
3. Frontend sends FormData with image
4. Backend receives FormData
5. Upload image to Cloudinary → get URL
6. Save post to MongoDB with imageUrl
7. Return created post
8. Frontend navigates to post view
```

### **Fetching Posts**
```
1. Home component mounts
2. Call appwriteService.getPosts()
3. Backend queries MongoDB (status: "active")
4. Return all active posts
5. Frontend displays PostCard components
6. User can click to view full post
```

### **Updating a Post**
```
1. User edits PostForm on edit page
2. Submit → appwriteService.updatePost()
3. Backend checks if user is author (userId === email)
4. If image provided: upload to Cloudinary, delete old image
5. Update post in MongoDB
6. Return updated post
7. Frontend navigates to updated post
```

---

## 🐛 Troubleshooting

### **CORS Errors**
- Ensure Go backend is running on port 8080
- Check `VITE_GO_SERVER_URL` in frontend config
- Verify CORS middleware is enabled

### **MongoDB Connection Failed**
- Check MongoDB URI in `.env`
- Verify MongoDB is running (or Atlas connection)
- Check username/password

### **Image Upload Fails**
- Verify Cloudinary credentials
- Check image file size (max 5MB recommended)
- Ensure .env has correct Cloudinary config

### **Login Issues**
- Clear localStorage and try again
- Check MongoDB for user record
- Verify password is correct

### **Token Expired**
- App automatically refreshes token
- If not working, clear localStorage and re-login
- Check refresh token expiry (7 days)

---

## 📈 Performance Optimizations

### **Frontend**
- Code splitting with React Router
- Image optimization via Cloudinary
- Lazy loading of components
- Redux for efficient state management

### **Backend**
- MongoDB indexing on slug and userId
- Request logging for debugging
- Connection pooling (MongoDB driver)
- Cloudinary caching

---

## 🚦 Status & Token Expiry

| Item | Value |
|------|-------|
| Access Token | 1 day (24 hours) |
| Refresh Token | 7 days |
| Frontend Port | 5173 |
| Backend Port | 8080 |
| CORS Origin | http://localhost:5173 |

---

## 📝 Key Code Examples

### **API Request in Frontend**
```javascript
// src/appwrite/config.js
async authenticatedFetch(endpoint, options = {}) {
  const token = this.getAuthToken();
  const headers = {
    'Authorization': `Bearer ${token}`,
    ...options.headers,
  };
  
  const response = await fetch(`${this.apiUrl}/api${endpoint}`, {
    ...options,
    headers,
  });
  
  return response.json();
}
```

### **Protected Route in Backend**
```go
// All POST requests require auth middleware
router.Handle("/posts", 
  middleware.Auth(http.HandlerFunc(controllers.CreatePost))
).Methods("POST")
```

### **Redux Login**
```javascript
// src/store/authSlice.js
const authSlice = createSlice({
  name: "auth",
  initialState: { status: false, userData: null },
  reducers: {
    login: (state, action) => {
      state.status = true;
      state.userData = action.payload.userData;
    },
  }
})
```

---

## 💡 Best Practices

### **Frontend**
- Always check `getCurrentUser()` before accessing protected routes
- Store tokens in localStorage; never expose them in code
- Use Redux for auth state to avoid prop drilling
- Validate form inputs before sending to API
- Handle errors gracefully with user-friendly messages
- Use React.lazy() for code splitting on large routes

### **Backend**
- ✅ Validate all input data before processing
- ✅ Use prepared MongoDB queries to prevent injection
- ✅ Add rate limiting to prevent brute force attacks
- ✅ Log all sensitive operations (login, post deletion)
- ✅ Use HTTPS in production
- ✅ Regularly rotate JWT secrets
- ✅ Implement request timeouts to prevent hanging requests
- ⚠️ **TODO:** Move JWT secrets from hardcoded values to environment variables

### **General**
- Keep `.env` files out of version control
- Test API endpoints before frontend integration
- Document any new endpoints in README
- Use meaningful commit messages
- Keep dependencies updated for security patches

---

## 🔄 Migration Guide for Appwrite Users

If you're migrating from Appwrite:

### **What Changed**
| Feature | Appwrite | Go Backend |
|---------|----------|-----------|
| User Registration | `account.create()` | `POST /api/register` |
| Authentication | Session-based | JWT-based (access + refresh) |
| Post Storage | Appwrite Database | MongoDB |
| File Upload | Appwrite Storage | Cloudinary |
| Client Library | Appwrite SDK | Vanilla fetch API |

### **Frontend Changes**
1. **Old (Appwrite SDK)**
   ```javascript
   import { Client, Account, Databases } from "appwrite";
   const client = new Client(); // Initialize Appwrite
   ```

2. **New (Go REST API)**
   ```javascript
   import { AuthService } from './appwrite/auth';
   const authService = new AuthService(); // REST client
   ```

### **Code Migration**
- Replace `account.create()` with `AuthService.createAccount()`
- Replace `databases.createDocument()` with `Service.createPost()`
- Replace session handling with localStorage JWT tokens
- Update error handling to work with HTTP status codes instead of SDK exceptions

---

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

---

## 📜 License

This project is open source and available under the MIT License.

---

## 👤 Author

Created with ❤️ by the development team

---

## 📞 Support

For issues and questions, please open an issue on GitHub.

---

**Happy Blogging! 🚀**
