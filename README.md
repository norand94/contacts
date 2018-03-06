### Установка и настройка

1. В папке с репозиторием запустить

`go get && docker-compose up -d`

2. Подключиться к mongo и создать пользователя admin, с паролем, указанным в конфиге подключения к БД.

3. Запустить `revel run`

### Роуты

Управление контактами.
                              
GET     /contact
GET     /contact/list                              
GET     /contact/show/:id                      
POST    /contact/create                         
PUT     /contact/update                        
DELETE  /contact/delete/:id                     


### Start the web server:

   revel run myapp

### Go to http://localhost:9000/ and you'll see:

    "It works"

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


