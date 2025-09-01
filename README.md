# VK CO FF EE Encryption/Decryption Service

A lightweight Go web server that provides a simple web interface to encrypt and decrypt text using the “VK CO FF EE” cipher algorithm (as seen in the VK Coffee app)

[vk-co-ff-ee.itskoshkin.site](https://vk-co-ff-ee.itskoshkin.site)

## Features

- Encrypt plain text → `VK CO FF EE … VK CO FF EE` cipher
- Decrypt cipher text → original plain text
- Support for custom key (derivation via AES-ECB + “pepper”)

## Getting Started

### Prerequisites

- Go 1.23+ 

### Running
1. Plain
    1. Clone the repository
        ```bash
        git clone https://github.com/itsk-example-projects/vk-co-ff-ee.git
        cd vk-co-ff-ee
        ```
    2. Install dependencies
        ```bash
        go mod tidy
        ```
    3. Run the application
        ```bash
        go run cmd/main.go
        ```
        *(Use `screen` to run app in a separate session so your terminal stays free)*
2. Docker
    1. Clone the repository
        ```bash
        git clone https://github.com/itsk-example-projects/vk-co-ff-ee.git
        cd vk-co-ff-ee
        ```
    2. Build Docker image
        ```bash
        docker build -t vk-co-ff-ee-example:latest .
        ```
    3. Run the container
        ```bash
        docker run -dit --name vk-co-ff-ee-example -p 8080:8080 vk-co-ff-ee-example:latest
        ```
        *App now available at http://localhost:8080*