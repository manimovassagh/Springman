# Springman 🧰

Springman is a lightweight CLI tool written in Go that helps you quickly generate and run Spring Boot projects using either Maven or Gradle.

## 🚀 Features

- Generate new Spring Boot projects via [start.spring.io](https://start.spring.io)
- Supports both Maven and Gradle
- Automatically detects and runs the correct wrapper (`mvnw` or `gradlew`)
- Easy to compile and use as a global CLI tool

## 📦 Installation

Build the CLI locally:

```bash
go build -o springman
```

Then move it to a global location:

```bash
sudo mv springman /usr/local/bin/
```

## 🛠 Usage

### Create a new project

```bash
springman new myapp --build maven
```

Or with Gradle:

```bash
springman new myapp --build gradle
```

### Run a project

```bash
springman run myapp
```

## 💡 Example

```bash
springman new blogapp --build gradle
cd blogapp
springman run .
```

## 📁 Project Structure

- `cmd/new.go` — command for project creation
- `cmd/run.go` — command to run the generated Spring Boot project

## 🧪 Requirements

- Go 1.21+
- Internet connection (to fetch Spring Boot ZIP)
- Java 17+ installed

## 📜 License

MIT
