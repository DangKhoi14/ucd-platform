# UCD Platform  
_Universal Code Documentation Platform_

---

## 🌐 Overview  
**UCD Platform** is a web-based application designed to analyze and generate documentation for codebases.  
The initial version focuses on **COBOL**, providing developers with a modern tool to understand, maintain, and document legacy COBOL applications.  
The platform is designed to be extensible, allowing future support for multiple programming languages.

---

## 🚀 Key Features
- 📄 **Upload COBOL source files** for automatic parsing and documentation generation.
- 🗃️ **Temporary storage** of uploaded files for secure processing.
- 🔍 **COBOL Code Analyzer** with potential for expansion into other languages.
- ⚙️ **Built with Golang Fiber framework** for fast and lightweight backend processing.
- 🔒 Optional **HTTPS support** for secure communication.

---

## 📂 Folder Structure
```plaintext
ucd-platform/
├── cmd/                 # Main application entry point
├── internal/            # Internal application logic
│   ├── handler/         # HTTP request handlers
│   ├── analyzer/        # Code analyzers
│   │   └── cobol/       # COBOL-specific analysis logic
│   └── util/            # Utility functions
├── tmp/                 # Temporary file storage
├── certs/               # SSL certificates (excluded from version control)
├── go.mod               # Go module definition
├── go.sum               # Go module checksum
└── README.md            # This file
```

---

## 🛠️ Tech Stack
Golang with Fiber for backend

COBOL code parsing and documentation tools

HTTPS (self-signed for local, production-ready with real certificates)

---

## 🏗️ How to Run

Clone the repository
```
git clone https://github.com/your-username/ucd-platform.git
cd ucd-platform
```

Install dependencies
```
go mod tidy
```

Run the server (HTTP)
```
go run cmd/main.go
```

OR run with HTTPS (self-signed)
```
go run cmd/main.go --tls
```

---

## 🌍 Future Enhancements

- Add support for other programming languages (Java, Python, etc.)

- User authentication and project management

- Enhanced code visualization and reporting

- Dockerization for easy deployment

---

## 📢 Contributing
Feel free to contribute! Fork the repo, create a branch, and submit a pull request.
