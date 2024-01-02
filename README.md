# URL Shortener

## Overview

Welcome to the URL Shortener project! This application provides a simple and efficient way to shorten URLs using a base-62 encoding method. Inspired by the principles outlined in "System Design Interview: An Insider’s Guide" by Alex Xu, this URL shortener is designed to be scalable and follows best practices for system design.

## Features

- URL shortening using base-62 encoding
- Three main endpoints:
  - `GET /v1/short/{encoded_string}`: Retrieve the original URL for the given encoded string.
  - `POST /v1/short`: Generate a short URL.
  - `DELETE /v1/short`: Delete a short URL.

## Getting Started

Follow these steps to get started with the URL Shortener:

1. Clone the repository to your local machine.
2. Ensure you have Go installed.
3. Navigate to the project directory.
4. Run the application by executing `go run main.go`.
5. The server will be accessible at `http://127.0.0.1:8000`.

## Usage

1. **Shorten a URL:**
   ```bash
   curl -X POST -d "long_url=https://www.example.com" http://127.0.0.1:8000/v1/short
   ```

2. **Retrieve the original URL:**
   ```bash
   curl http://127.0.0.1:8000/v1/short/{encoded_string}
   ```

3. **Delete a short URL:**
   ```bash
   curl -X DELETE http://127.0.0.1:8000/v1/short
   ```

## Inspiration

This project draws inspiration from the valuable insights shared in "System Design Interview: An Insider’s Guide" by Alex Xu. It aims to demonstrate how to architect scalable systems with best practices in mind.

## Contributions

Contributions are welcome! Feel free to open issues or submit pull requests to enhance the URL Shortener. Your contributions will help make it even more efficient and robust.

## Acknowledgments

Special thanks to Alex Xu for the guidance provided in "System Design Interview: An Insider’s Guide."

## License

This repository is licensed under the [MIT License](LICENSE). Use, modify, and distribute the code as needed, adhering to the license terms.

Happy URL shortening!
