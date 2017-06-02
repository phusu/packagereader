# Linux dpkg status file reader and server

In Debian/Ubuntu Linux system there exists a file /var/lib/dpkg/status which lists information about each installed package in the system. 
This application reads the file, parses the contents and starts a web server listening to port 8080. User can browse the package information with web browser
from localhost:8080. Index page shows each package and the user can click on a package to see more information about the package, such as description and dependencies. 
An example status file is included.

## Goals of the project
- Learn Go