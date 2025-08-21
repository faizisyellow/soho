## soHo

![soho logo](/assets/logo.png)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/faizisyellow/soho)

soHo is a CLI tool for generating CRUD files into [falcon](https://github.com/faizisyellow/falcon) Go REST APIs project.<br>
It assists you to generate CRUD files to get faster when developing REST APIs.

## Features

- Generate new repository files.
- Generate new service files.
- Generate new handler files.
- Generate new resource which is generate all repository, service, handler and route files.


## Installation

### Via ``` go install ```

With go 1.24 or higher:

```
go install github.com/faizisyellow/soho@latest
```

## Usage

1. Initialize soHo CLI into your project.
   
   ```
   soho init 
   ``` 

2. To generate resource files.

   ```
   soho generate resource {NAME}
   ```

3. To generate repository file.

   ```
   soho generate repository {NAME}
   ```

See more commands in :

 ``` 
 soho --help
 ```  
