# Content Distributer 🚀

A comprehensive multi-platform content distribution system that streamlines the process of publishing content across multiple social media platforms simultaneously. This application eliminates the need for manual, repetitive posting by providing a unified interface for content management and distribution.

## Overview

The Content Distributer is a full-stack application designed to simplify social media content management for creators, marketers, and businesses. The platform supports content distribution across five major social media networks: Instagram, Pinterest, YouTube, Reddit, and LinkedIn.

**Key Features:**
- **Unified Content Management**: Upload media once and distribute across multiple platforms
- **Platform-Specific Optimization**: Automatically adapts content format and metadata for each platform's requirements
- **Intelligent Validation**: Prevents incompatible platform combinations and validates content requirements
- **Concurrent Processing**: Simultaneous uploads to multiple platforms for improved efficiency
- **Dynamic Form Generation**: Platform-specific fields are dynamically generated based on user selections

The system intelligently handles platform-specific requirements, such as Pinterest's image-only policy or Instagram's support for both images and videos, ensuring optimal content delivery across all selected platforms.

## Technology Stack

**Frontend:** Built with Next.js 14 and TypeScript for robust type safety and modern React development. The application utilizes React Hook Form for efficient form management, Zod for comprehensive data validation, and Tailwind CSS for responsive, utility-first styling.

**Backend:** Developed in Go for high performance and reliability. The server uses Chi for lightweight HTTP routing and implements concurrent processing for optimal performance. The architecture supports scalable content distribution across multiple platforms simultaneously.

**API Integrations:** The system integrates with five major social media platforms: Instagram, Pinterest, YouTube, Reddit, and LinkedIn. Each platform's unique API requirements and data formats are abstracted into a unified interface, simplifying the content distribution process.

## Getting Started (The Not-So-Fun Part)

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
   (Don't ask us where to get the API key—that's between you and the upload service provider.)

3. **Set up the frontend**:
   ```bash
   cd ../frontend
   npm install
   ```

### Running the Application

**Backend first** (because the frontend needs something to talk to):
```bash
cd backend
go run cmd/api/main.go
```
The backend will start on `http://localhost:8000`. You'll know it's working when you see some logs that don't look like error messages.

**Then the frontend** (in a new terminal, because we're not savages):
```bash
cd frontend
npm run dev
```
The frontend will be available at `http://localhost:3000`. Open it up in your browser and marvel at what we've built.

## How It Actually Works

### The Upload Process

1. **Choose your platforms**—pick from Instagram, Pinterest, YouTube, Reddit, and LinkedIn. The system is smart enough to prevent you from selecting incompatible combinations (like Pinterest and YouTube together, because Pinterest doesn't do videos).

2. **Upload your media**—drag and drop works, or click to browse. The system automatically detects whether you're uploading an image or video and adjusts the available platforms accordingly.

3. **Fill out the form**—each platform has its own requirements, but we've made it as painless as possible. Required fields are clearly marked, and the form validates everything before you even try to submit.

4. **Hit submit**—the magic happens. Your content gets processed and sent to all the selected platforms simultaneously. No more manual posting.

### Platform-Specific Features

**Instagram:** Supports both images and videos, includes user tagging, and handles captions like a pro. The system automatically populates the image URL field when you upload media, so you don't have to think about it.

**Pinterest:** Image-only (because that's how Pinterest rolls), with automatic source type detection. The system prevents video uploads when Pinterest is selected, because we learned that lesson the hard way.

**YouTube:** Full video support with title, description, tags, category, and privacy settings. Everything you need to make your content discoverable.

**Reddit:** Text posts, link posts, image posts—whatever floats your boat. Includes subreddit selection, NSFW tagging, and all the other Reddit-specific features you'd expect.

**LinkedIn:** Professional content with proper visibility settings, author attribution, and lifecycle state management. Because your professional network deserves better than a hastily written post.

## Project Structure

```
Content-Distributer/
├── backend/                 # The Go backend (where the magic happens)
│   ├── cmd/api/            # Main application entry point
│   ├── internal/           # Internal packages (handlers, tools, etc.)
│   ├── uploads/            # Platform-specific upload implementations
│   └── config/             # Configuration files and secrets
├── frontend/               # The Next.js frontend (the pretty part)
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

# Add other environment variables as needed
```

### Platform API Keys

Each platform requires its own API credentials. You'll need to:
1. Create developer accounts with each platform
2. Generate API keys and access tokens
3. Configure the appropriate credentials in your environment

(We'd give you the exact steps, but platform APIs change more often than we change our socks, so it's best to check their current documentation.)

## Contributing

Found a bug? Have a feature request? Think we're doing something completely wrong? We'd love to hear from you!

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Just please, for the love of all that is holy, test your changes before submitting. We've seen enough broken code to last a lifetime.

## Known Issues (Because Nothing's Perfect)

- The Pinterest API can be... finicky. Sometimes it works perfectly, sometimes it throws errors that make no sense. We're working on it.
- File upload validation could be more robust (but honestly, what couldn't be?).
- The error messages could be more user-friendly (we're working on that too).

## Roadmap

- **Better error handling**—because nobody likes cryptic error messages
- **Scheduling functionality**—because sometimes you want to post at 3 AM without actually being awake at 3 AM
- **Analytics integration**—because numbers are fun
- **More platforms**—because why stop at five?

## License

This project is licensed under the MIT License—use it, modify it, make it better. Just don't blame us if something goes wrong.

## Acknowledgments

Thanks to all the open-source libraries that make this project possible. Thanks to the platform APIs (even when they're being difficult). And thanks to anyone who's ever had to manually post the same content to multiple social media platforms—you inspired this project.

## Support

Having issues? Check the GitHub issues page first. Someone else might have had the same problem. If not, create a new issue and we'll get back to you... eventually. (We're not Amazon, but we try to be reasonably responsive.)

---

*Built with ❤️ and way too much coffee by people who got tired of manually posting to social media.*
