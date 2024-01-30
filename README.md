  # URL Shortener
  ---------------------

  ## Technologies & Language
  - Golang
  - Chi
  - MYSQL DB

  ## API
  API will have 2 endpoints:
  - URL Shortener endpoint
  - Redirect to original URL endpoint
  
  ## Design
  
  ### Databse Driver
  Database Driver will be intialized once and set into a universal variable
  and used by functions

  ### Structs
  There will be only one struct "URL" that will be used to get the URLs, 
  Shorten URLs, store and retrieve from database through interfaces

  ### Interfaces
  The Project will have only 2 interfaces
  - Database Operations interface
    - ```Store();```
    - ```Retrieve();```

  - Shorten Interface
    - ```Shorten();```

  ### Design Overview
  The Service Recieves the data in the URL struct and based on the request we will know what to do through our interfaces, this design keep things simple and not dependable and if we wanted to extend new features it'll be easier through the interfaces.