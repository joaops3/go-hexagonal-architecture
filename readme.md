### hexagonal architecture implementation

This project demonstrates the **Hexagonal Architecture (Ports and Adapters)**

## What is Hexagonal Architecture?

Hexagonal Architecture, also known as **Ports and Adapters**, is a software architecture pattern that aims to isolate the **application core** (business logic) from **external systems** (like databases, UIs, APIs). This separation enables:

- Easy testing of business logic in isolation
- Flexibility to change external systems without affecting the core
- Clear boundaries between layers

```bash
.
├── adapters            # External interfaces (Adapters)
│   ├── cli             # Command-line interface (input)
│   ├── db              # Database interaction (output)
│   ├── dto             # Data transfer objects
│   └── web             # Web API server
├── application         # Core business logic (Use Cases, Interfaces)
├── cmd                 # Application entry points
├── main.go             # Main entry point
├── docker-compose.yml  # Docker setup

```
