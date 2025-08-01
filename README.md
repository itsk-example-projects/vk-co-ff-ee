# VK CO FF EE Encryption/Decryption Service

A lightweight Go web server that provides a simple web interface to encrypt and decrypt text using the “VK CO FF EE” cipher algorithm (as seen in the VK Coffee app)

[vk-co-ff-ee.itskoshkin.site](https://vk-co-ff-ee.itskoshkin.site)

## Features

- Encrypt plain text → `VK CO FF EE … VK CO FF EE` cipher
- Decrypt cipher text → original plain text
- Support for custom key (derivation via AES-ECB + “pepper”)

## Prerequisites

- Go 1.18+ 

## Installation & Run

```bash
git clone https://github.com/itsk-example-projects/vk-co-ff-ee.git
cd vk-co-ff-ee
go mod tidy
go run cmd/main.go
```
