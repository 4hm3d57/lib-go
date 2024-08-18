# Library Management System

This is a library management system built with Golang and the Gin framework. 
The system allows administrators to manage books, students, and book recommendations.
The application uses MongoDB as the database and serves HTML templates for the user interface.

## Features

- **User Authentication**: Users can sign up and log in to the system.
- **Admin Dashboard**: Admins can manage students, books, and book recommendations.
- **Book Management**: Add, edit, and view books available in the library.
- **Student Management**: Manage student records and their borrowed books.
- **Recommendations**: Admins can manage book recommendations for students.

  
## Future Enhancements
- Complete Borrowing/Returning Functionality: Add the ability for students to borrow and return books.
- Enhanced User Interface: Improve the UI/UX for better usability.
- Search Functionality: Implement a search feature to easily find books.

## Prerequisites

Make sure you have the following installed on your system:

- [Golang](https://golang.org/doc/install)
- [MongoDB](https://docs.mongodb.com/manual/installation/)
- [Git](https://git-scm.com/)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/4hm3d57/lib-go.git
   cd lib-go
   
   # install dependencies
   go mod tidy
   
   # run the application
   go build .
   ```
   - Open your browser and go to  `http://localhost:5050`

   
## Contributing

Feel free to fork this repository and contribute by submitting a pull request. For major changes, please open an issue first to discuss what you would like to change.




