# Social Media Project

A project inspired by early versions of Twitter written in Go and Echo framework.


## Technologies
- Architecture: Clean
- Delivery: REST API
- Database: Postgresql
- Authentication: JWT
- Object Storage: MinIO
- Deployment: Docker

## Key Improvements

**Input Sanitization**  
   Input data should be validated and sanitized properly to prevent security issues.

**Error Handling**  
   Implementing **custom richerror** package for better and more structured custom error management.

**Configuration Management**  
   Loading configuration variables from env for flexibility and security.

**Custom Middlewares**  
   Writting custom middlewares to dive deeper into authentication & authorization based on project needs.

**Database Migrations**  
   Adding DB migrations to handle schema updates and maintain consistency.

**Essential Endpoints**  
   Implementing crucial endpoints like **forgot-password** for better user experience.

**API Documentation**  
   Integrate **Swagger** for clean, interactive API documentation.

**Pagination & Sorting**  
   Add pagination and sorting for each users' feeds.
