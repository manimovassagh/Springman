# 🧰 Springman

Springman is a lightweight and fast CLI tool written in Go for bootstrapping, running, and managing Spring Boot projects.

## 🚀 Features

- Generate new Spring Boot apps with Maven or Gradle
- Run projects using wrapper scripts (`./mvnw`, `./gradlew`)
- Add Maven dependencies to `pom.xml`
- Remove existing dependencies from `pom.xml`
- Clean XML formatting and prevents duplicates
- Works offline after the initial setup

---

## 📦 Installation

Follow these steps to build and install Springman:

### 1. Clone the repository

```bash
git clone https://github.com/yourname/springman.git
cd springman
```

### 2. Build the CLI

Make sure you have Go installed (version 1.21 or higher):

```bash
go build -o springman
```

### 3. Move the binary to your system path

This allows you to use `springman` from anywhere:

```bash
sudo mv springman /usr/local/bin/
```

> 🛠️ You may need to enter your system password

### 4. Verify installation

```bash
springman --help
```

You should see the available commands and options.

---

## 🛠 Commands

### `new`

Create a new Spring Boot project.

```bash
springman new myapp --build maven
```

### `run`

Run the specified Spring Boot project.

```bash
springman run myapp
```

### `add`

Add a dependency to `pom.xml`.

```bash
springman add myapp org.springframework.boot:spring-boot-starter-data-jpa
springman add myapp org.springframework.boot:spring-boot-starter-web:3.3.0
```

### `remove`

Remove a dependency from `pom.xml`.

```bash
springman remove myapp org.springframework.boot:spring-boot-starter-web
```

---

## 🧪 Requirements

- Go 1.21+
- Java 17+
- Internet connection for downloading starter ZIP

---

## 📁 Project Structure

- `cmd/` – contains CLI command implementations (`new`, `run`, `add`, `remove`)
- `main.go` – entry point of the CLI
- `go.mod` – Go module definition

---

## 📜 License

MIT License

---

## 🙌 Author

Made with ❤️ by [Mani Movassagh](https://github.com/manimovassagh)
