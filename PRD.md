Product Requirements Document (PRD)

Title
    • FTP File Downloader Web Application using Supabase Integration

Document Control
    • Author: David Bowman
    • Date: March 20, 2025 Thursday
    • Version: 1.0

Overview
    • This web application will securely download files from multiple FTP servers and store them in Supabase Storage and store download info in a database
    • The app will also use Supabase Database for managing file metadata (with an already defined schema) and Supabase Authentication to handle user logins
    • The primary goal is to provide a streamlined, secure solution that enables authenticated users to trigger FTP downloads and manage stored files

Objectives & Goals
    • Store FTP login information: Store FTP server login information in Supabase table called logins
    • Store radio programs metadata information in the table called programs
    • Store radio show files download information the table called shows
    • Seamless FTP Integration: Automatically connect to and download files from a specified FTP server
    • Centralized Storage: Store downloaded files in Supabase Storage with corresponding metadata in the database
    • Secure Authentication: Leverage Supabase Authentication to manage user sessions and access control
    • User-friendly Interface: Provide an intuitive interface for users to log in, initiate downloads, and view file statuses
    • Scalability & Security: Ensure that the system is scalable and secure, following best practices in data handling and user management

Target Audience
    • Internal Users/Administrators: Individuals responsible for managing file imports and verifying successful transfers
    • End Users: Authenticated users who need to trigger file downloads or access stored files for further processing

Features

1. User Authentication & Management
    • Login/Logout: Secure login and logout functionality using Supabase Authentication
    • User Roles: Define roles (e.g., Admin, Standard User) for different levels of access and control
    • Session Management: Maintain session security and implement proper timeout and re-authentication mechanisms

2. FTP Integration
    • FTP Connectivity: Connect to a specified FTP server using secure credentials
    • File Download: Trigger downloads manually or schedule automated file pulls
    • Error Handling: Log and notify users of FTP connection issues or download errors

3. File Storage & Metadata Management
    • Supabase Storage Integration: Store downloaded files securely
    • Database Updates: Insert or update file metadata in the pre-defined Supabase database schema
    • File Management UI: Provide an interface where users can see a list of downloaded files, their statuses, and relevant metadata

4. User Interface & Experience
    • Dashboard: A central dashboard displaying file download history, current FTP connection status, and any alerts or error messages
    • Responsive Design: Ensure the application is accessible on both desktop and mobile devices
    • Notifications: In-app notifications (and optional email alerts) for critical actions like successful downloads or error occurrences

Functional Requirements
    • Authentication: Utilize Supabase's built-in authentication methods to secure the application
    • FTP Client Module: Develop or integrate a robust FTP client that can handle secure file transfers
    • File Handling: Implement functions to process downloaded files (e.g., renaming, metadata extraction) before storing them
    • Data Persistence: Use the pre-defined database structure to persist file details such as filename, size, download timestamp, and FTP source
    • API Endpoints: Create backend endpoints (or serverless functions) to facilitate FTP operations, file storage, and user management

Non-Functional Requirements
    • Security: Ensure data encryption in transit and at rest; follow best practices for user authentication and authorization
    • Performance: Optimize file download and storage processes to handle large files and multiple simultaneous requests
    • Scalability: Architect the solution to handle growth in user base and volume of FTP transactions
    • Reliability: Implement error logging, monitoring, and retry mechanisms for failed file downloads
    • Maintainability: Use modular code design and proper documentation to ensure ease of maintenance and future enhancements

Technical Architecture
    • Frontend: A web interface built with html/htmx/tailwind served from Go Echo, that communicates with backend services
    • Backend/Serverless: Go Lang Functions that handle FTP connections, file processing, and interactions with Supabase (authentication, database, and storage)
    • Template System:
        • Technology: Go's native `html/template` package
        • Template Syntax: Go template syntax (not Django/Jinja2)
        • Template Structure:
            - Main template definition: `{{define "template_name"}}`
            - Block definitions: `{{define "block_name"}}`
            - Block inclusion: `{{template "block_name" .}}`
        • Template Organization:
            - Each page has its own template file
            - Common blocks (title, nav, content) defined in each template
            - Templates stored in `web/templates/` directory
        • Template Features:
            - Safe HTML rendering
            - Template inheritance through blocks
            - Data passing using dot notation
            - Custom template functions support
    • Supabase Integration:
        • Authentication: Managing user sessions
        • Database: Storing file metadata
        • Storage: Saving downloaded files
        • FTP Client: A dedicated module (or third-party library) integrated into the backend to manage FTP operations
    • Environment Variables:
        • All sensitive data must be stored in environment variables
        • Required environment variables:
            - SUPABASE_URL: Supabase project URL
            - SUPABASE_KEY: Supabase anon/public key
            - SUPABASE_JWT_KEY: Supabase JWT secret key
            - PORT: Server port (optional, defaults to 8080)
            - ENVIRONMENT: Application environment (development/production)
        • Environment files:
            - Use .env file for local development
            - .env file must be added to .gitignore
            - Production environments should use secure environment variable management
        • Security considerations:
            - Never commit sensitive data to version control
            - Use different keys for development and production
            - Rotate keys regularly
            - Use secure key management services in production

    • API Testing Format:
        • Curl Commands:
            - Always use single quotes for data parameters to prevent special character issues
            - Example format for login:
              ```
              curl -v -b cookies.txt -c cookies.txt \
                -H 'Content-Type: application/x-www-form-urlencoded' \
                -d '_csrf=TOKEN&email=user@example.com&password=pass' \
                http://localhost:8080/api/auth/login
              ```
            - Double quotes may cause issues with special characters in passwords
            - Store cookies in a file for session persistence
            - Include CSRF token from previous requests

Database Relationships and Workflow

1. FTP Server Management
    • The `logins` table stores credentials for multiple FTP servers
    • Each FTP server can host multiple radio programs
    • Credentials are securely stored and accessed only when needed for downloads

2. Program Configuration
    • The `programs` table links to FTP servers via `loginid` (foreign key to `logins.loginid`)
    • Each program has its own subdirectory on the FTP server
    • Programs are configured with specific parsing rules and metadata

3. File Download Process
    • When a program is selected for download:
        • System retrieves FTP credentials from `logins` table using the program's `loginid`
        • Connects to the FTP server using these credentials
        • Navigates to the program's specific subdirectory
        • Compares available files against the `shows` table to identify new files
        • Downloads only files that haven't been previously downloaded
        • Creates new records in the `shows` table for each downloaded file

4. File Processing and Parsing
    • The `parsers` table contains rules for processing downloaded files
    • Each parser defines patterns for extracting:
        • Year information
        • Month information
        • Date information
        • Topic information
    • Parsing rules are referenced by programs and applied to downloaded files
    • The parsing process is configurable and can be updated as needed

5. Download History Tracking
    • The `shows` table maintains a complete history of downloaded files
    • Records include:
        • Original filename
        • New filename (after processing)
        • Show date and title
        • Access information
        • Program reference
    • This history prevents duplicate downloads and enables tracking of file processing status

This workflow ensures:
    • Efficient file downloads by avoiding duplicates
    • Proper organization of files by program
    • Secure credential management
    • Flexible file processing through configurable parsers
    • Complete tracking of all downloaded and processed files


