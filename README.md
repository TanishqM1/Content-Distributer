*this project is deprecated. It is managed privately*
# Content Distributer 

A multi-platform content distribution system designed to streamline the process of publishing content across multiple social media platforms. This application provides a unified interface for managing and distributing content, removing the need for repetitive manual posting.

<details>
  <summary>📸 Click to view screenshots</summary>

  ### Landing Page
  <img src="https://github.com/user-attachments/assets/51c2c676-a2c0-45f6-b764-3c87537fd560" width="1907" alt="Landing Page">

  ### Platform Select Page
  <img src="https://github.com/user-attachments/assets/3dcdee02-97c6-4f04-9498-9c4692dab9f9" width="1902" alt="Platform Select Page">

  ### Upload Page
  <img src="https://github.com/user-attachments/assets/ab8fab8c-5554-4b4a-b82b-596914d9d19c" width="1908" alt="Upload Page">

</details>

## Overview

The Content Distributer is a full-stack application intended to simplify social media content management for creators, marketers, and businesses. It supports automated content distribution across Instagram, Pinterest, YouTube, Reddit, and LinkedIn and more.

**Key Features:**
- **Unified Content Management**: Upload media once and distribute across multiple platforms
- **Platform-Specific Optimization**: Automatically adapts content format and metadata for each platform's requirements
- **Intelligent Validation**: Prevents incompatible platform combinations and validates content requirements
- **Concurrent Processing**: Simultaneous uploads to multiple platforms for improved efficiency
- **Dynamic Form Generation**: Platform-specific fields are dynamically generated based on user selections

The system ensures compliance with platform constraints (e.g., Pinterest's image-only requirements) and optimizes content formatting to support smooth distribution.

## Technology Stack

**Frontend:** Implemented using Next.js 14 with TypeScript for strong type safety and modern React development. React Hook Form provides form state management, Zod ensures reliable validation, and Tailwind CSS enables responsive, utility-first styling.

**Backend:** Developed in Go for high performance and reliability. The server uses Chi for lightweight HTTP routing and implements concurrent processing for optimal performance. The architecture supports scalable content distribution across multiple platforms simultaneously.

**API Integrations:** Integrates with APIs from Instagram, Pinterest, YouTube, Reddit, and LinkedIn. Platform-specific requirements and payload structures are abstracted behind a unified interface.

## Getting Started

### Prerequisites

You'll need a few things installed before we get started:
- **Node.js** (version 18 or higher)—grab it from [nodejs.org](https://nodejs.org)
- **Go** (version 1.25 or higher)—head over to [golang.org](https://golang.org)
- **Git** (because how else are you gonna clone this thing?)

### Installation

1. **Clone the repository** (obviously):
   ```bash
   git clone https://github.com/TanishqM1/Content-Distributer.git
   cd Content-Distributer
   ```

2. **Set up the backend**:
   ```bash
   cd backend
   go mod download
   ```
   
   You'll also need to create a `.env` file in the `backend/config/` directory:
   ```env
   UploadsAPI=your_upload_api_key_here
   ```

3. **Set up the frontend**:
   ```bash
   cd ../frontend
   npm install
   ```

### Running the Application

**Start the backend:**:
```bash
cd backend
go run cmd/api/main.go
```
The backend will start on `http://localhost:8000`.

**Then the frontend**:
```bash
cd frontend
npm run dev
```
The frontend will be available at `http://localhost:3000`.


## How It Actually Works

### The Upload Process

1. **Choose your platforms**—pick from Instagram, Pinterest, YouTube, Reddit, and LinkedIn. The system is smart enough to prevent you from selecting incompatible combinations (like Pinterest and YouTube together, because Pinterest doesn't do videos).

2. **Upload your media**—drag and drop works, or click to browse. The system automatically detects whether you're uploading an image or video and adjusts the available platforms accordingly.

3. **Fill out the form**—each platform has its own requirement . Required fields are clearly marked, and the form validates everything before you even try to submit.

4. **Hit submit**—Your content gets processed and sent to all the selected platforms simultaneously. No more manual posting.

### Platform-Specific Features

**Instagram:** Supports both images and videos, includes user tagging, and handles captions like a pro. The system automatically populates the image URL field when you upload media, so you don't have to think about it.

**Pinterest:** Image-only (because that's how Pinterest rolls), with automatic source type detection. The system prevents video uploads when Pinterest is selected.

**YouTube:** Full video support with title, description, tags, category, and privacy settings. Everything you need to make your content discoverable.

**Reddit:** Text posts, link posts, image posts—whatever floats your boat. Includes subreddit selection, NSFW tagging, and all the other Reddit-specific features you'd expect.

**LinkedIn:** Professional content with proper visibility settings, author attribution, and lifecycle state management

## Project Structure

```
Content-Distributer/
├── backend/                 # The Go backend 
│   ├── cmd/api/            # Main application entry point
│   ├── internal/           # Internal packages (handlers, tools, etc.)
│   ├── uploads/            # Platform-specific upload implementations
│   └── config/             # Configuration files and secrets
├── frontend/               # The Next.js frontend
│   ├── app/                # Next.js app directory
│   ├── components/         # Reusable React components
│   └── lib/                # Utilities and type definitions
└── README.md              # This file (you're reading it right now)
```

## Configuration

### Environment Variables

The backend needs a few environment variables to function properly. Create a `.env` file in `backend/config/` with the following:

```env
# Upload API configuration
UploadsAPI=your_upload_api_key_here
```

### Platform API Keys

Each platform requires its own API credentials. You'll need to:
1. Create developer accounts with each platform
2. Generate API keys and access tokens
3. Configure the appropriate credentials in your environment

or you can use upload-post :) They provide a central platform with API keys & account management.

## License

This project is licensed under the MIT License.

## Acknowledgments

Thanks to all the open-source libraries used within this project and Upload-Post.com for social media api documentation

---
