# URL Shortener

A simple yet functional URL shortener built with Go, featuring a clean web interface and SQLite database storage.

## Features

- 🔗 Shorten long URLs to 5-character codes
- 🌐 Web interface for creating and viewing shortened URLs
- 📊 View all shortened URLs in one place
- 🗄️ SQLite database for persistent storage
- ↩️ Automatic redirection from short codes to original URLs
- 🎨 Custom styling with CSS

## Tech Stack

- **Go 1.24.5** - Backend server and routing
- **GORM** - ORM for database operations
- **SQLite** - Lightweight database
- **HTML Templates** - Server-side rendering
- **Standard Library HTTP** - Built-in Go HTTP server

## Project Structure

```
url-shortener/
├── main.go                 # Application entry point and HTTP handlers
├── internal/
│   ├── short-url.go       # URL shortening logic and models
│   └── db/
│       ├── db.go          # Database connection
│       └── short-url.db   # SQLite database file
└── web/
    ├── templates/         # HTML templates
    │   ├── index.html    # Home page with URL list
    │   ├── shorten.html  # Short URL result page
    │   └── 404.html      # Not found page
    └── styles/
        └── styles.css     # Application styles
```

## Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/cjosue15/url-shortener.git
   cd url-shortener
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the application**
   ```bash
   go run main.go
   ```

4. **Access the application**
   
   Open your browser and navigate to `http://localhost:8080`

## Usage

### Web Interface

1. **Shorten a URL**
   - Enter a URL in the input field on the home page
   - Click submit
   - Copy your shortened URL

2. **Access shortened URLs**
   - Navigate to `http://localhost:8080/{code}`
   - You'll be automatically redirected to the original URL

3. **View all URLs**
   - The home page displays all shortened URLs in your database

### API Endpoint

**POST** `/api/shorten`
- Form field: `url` - The URL to shorten
- Redirects to `/shorten` page with the result

## How It Works

1. When you submit a URL, the application generates a random 5-character code
2. The original URL and short code are stored in the SQLite database
3. When someone visits `/{code}`, the app looks up the original URL and redirects
4. If a code doesn't exist, users are redirected to a 404 page

## Dependencies

- `gorm.io/gorm` - ORM library
- `gorm.io/driver/sqlite` - SQLite driver for GORM
- `github.com/mattn/go-sqlite3` - SQLite3 driver

## Development

To run in development mode with auto-reload, you can use tools like [Air](https://github.com/air-verse/air):

```bash
go install github.com/air-verse/air@latest
air
```

## License

This project is open source and available under the MIT License.

## Author

Carlos Josue ([@cjosue15](https://github.com/cjosue15))
