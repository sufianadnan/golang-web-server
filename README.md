# My First Golang Web Server Project

Welcome to my first Golang web server project! In this project, I have created a simple HTTP web server using Golang to serve static web pages and handle form submissions. Let's take a closer look at what I've built.

## Project Overview

I am Sufian Adnan, a 3rd-year IT student at Seneca College in Ontario, Canada. This project is part of my journey into Golang development.

### Features

- Serving static web pages (index.html and form.html)
- Handling form submissions
- Minimalistic and user-friendly interface

## How to Run

Follow these simple steps to run my Golang web server project:

1. Clone the project to your local machine:
   ```bash
   git clone https://github.com/sufianadnan/golang-web-server.git
   ```
2. Navigate to the project directory:
   ```bash
   cd golang-web-server
   ```
3. Build and run the Golang web server:
   ```bash
   go run ./cmd/web
   ```
4. Open your web browser and visit http://localhost:8080 to explore the project.

## Project Structure

Here's a brief overview of the project files:

- index.html: A simple HTML page served by the web server. It displays a greeting message and a Golang logo.

- form.html: Another HTML page containing a form. Users can enter their name and address and submit the form.

- main.go: The heart of the project, this Go code defines the web server and handles requests. It serves static files, processes form submissions, and provides a simple "hello" response.

## Endpoints

1. `/` - Serves the index.html page, the landing page of the project.

2. `/form` - Handles form submissions from the form.html page. It parses and displays the submitted data.

3. `/hello` - Responds with a "hello" message when you access this endpoint.

## Author

This project was created with ❤️ in Toronto by Sufian Adnan, a student at Seneca College.
Thank you for checking out my first Golang web server project. I hope you find it interesting and inspiring as I continue my journey in Golang development.
