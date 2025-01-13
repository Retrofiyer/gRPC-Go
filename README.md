# Go - gRPC

## About The Project

The "gRPC-UserManagement" project demonstrates how to use gRPC in Golang to create a service that allows users to be added and listed. This project is ideal for learning about gRPC's powerful features, including protocol buffer serialization, service definition, and streamlined communication.

## Table of Contents

<ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#overview">Overview</a></li>
        <li><a href="#features">Features</a></li>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#configuration">Configuration</a></li>
        <li><a href="#running-the-service">Running the service</a></li>
    </li>
    <li>
      <a href="#contributing">Contributing</a>
    </li>
 </ol>

## Overview

This project is built using Golang and gRPC to create a microservice for user management. It includes functionality to add and list users, leveraging gRPC's efficient and structured communication.
## Features

<div> <ul> <li> <b>gRPC Communication:</b> Uses gRPC for efficient client-server interaction.</li> <li> <b>Protocol Buffers:</b> Defines services and messages using `.proto` files.</li> <li> <b>Simple Setup:</b> Easy to run and extend for more complex use cases.</li> </ul> </div>


## Built With

[![Go][go.dev]][go-url]

<!-- GETTING STARTED -->
## Getting Started

## Prerequisites

Before you begin, make sure you have the following tools installed on your machine:

- **Golang 1.23.0 or higher** - [Download Golang](https://go.dev/dl/)

If you don't have any of these tools installed, follow the provided links to install them.

## Installation

1.- Clone the repository
   ```sh
   git clone https://github.com/Retrofiyer/gRPC-Go.git
   cd gRPC-Go
   ```
2.- Install Dependencies
 ```sh
   go mod tidy
   ```

## Running the service

  ```sh
    go run main.go
   ```

Open any browser and type 

  ```sh
    localhost:50051
   ```

## Contributing
I would like you to contribute to this project. Whether it's fixing a bug, adding a new feature or improving the documentation, your help is always welcome. Please email me at `sebitas5225@gmail.com` with all the details for improvement.

<!-- LINKS & IMAGES -->

[docker.com]: https://img.shields.io/badge/Docker-black?style=for-the-badge&logo=docker&logoColor=white
[docker-url]: https://www.docker.com/
[go.dev]: https://img.shields.io/badge/Go-black?style=for-the-badge&logo=go&logoColor=white
[go-url]: https://go.dev/