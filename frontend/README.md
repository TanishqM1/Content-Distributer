# Content Distributer Frontend

This is a Next.js frontend application with shadcn/ui components.

## Getting Started

First, install the dependencies:

```bash
npm install
```

Then, run the development server:

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## Features

- Next.js 14 with App Router
- TypeScript
- Tailwind CSS
- shadcn/ui components
- Responsive design

## Project Structure

```
frontend/
├── app/                 # Next.js app directory
│   ├── globals.css     # Global styles
│   ├── layout.tsx      # Root layout
│   └── page.tsx        # Home page
├── components/         # Reusable components
├── lib/               # Utility functions
│   └── utils.ts       # cn utility for class merging
├── components.json    # shadcn/ui configuration
├── tailwind.config.js # Tailwind CSS configuration
├── tsconfig.json      # TypeScript configuration
└── package.json       # Dependencies and scripts
```

## Adding shadcn/ui Components

To add new shadcn/ui components, run:

```bash
npx shadcn-ui@latest add [component-name]
```

For example:
```bash
npx shadcn-ui@latest add button
npx shadcn-ui@latest add card
```
