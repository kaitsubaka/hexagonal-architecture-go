## Hexagonal Architecture Go

This repository provides a template based on the principles of Hexagonal Architecture and Domain-Driven Design (DDD) for Go programming language, utilizing the Echo framework and Bun SQL ORM. 

### Overview

Hexagonal Architecture, also known as Ports and Adapters, emphasizes the separation of concerns between core application logic and external dependencies such as databases, user interfaces, and third-party services. By defining clear boundaries and interfaces, this architectural style promotes flexibility, testability, and maintainability.

Domain-Driven Design (DDD) complements Hexagonal Architecture by focusing on modeling the core business domain in a way that aligns with stakeholders' mental models. It promotes the use of ubiquitous language, aggregates, entities, value objects, and domain services to capture the essence of the problem domain effectively.

### Features

- **Clear Separation of Concerns**: The template separates the core application logic from external concerns, making it easier to understand and maintain.
  
- **Modular Design**: Components are organized into modules, each with well-defined responsibilities and interfaces, promoting encapsulation and reusability.

- **Testability**: By isolating dependencies, the code becomes more testable, enabling the use of unit tests, integration tests, and other testing techniques.

### Getting Started

To use this template:

1. Clone the repository to your local machine.
2. Explore the directory structure and familiarize yourself with the components.
3. Customize the template to fit your specific application requirements.
4. Refer to the documentation for detailed explanations and best practices.

### Documentation

#TODO

### Todo

- [ ] Add soft delete and update methods
- [ ] Add migrations with bun
- [ ] Add documentation
- [ ] Add linters
- [ ] Add unit tests

### Contribution

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

### License

This template is licensed under the [MIT License](LICENSE). Feel free to use it for both commercial and non-commercial projects.
