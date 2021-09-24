# minderaWeatherService!

![weatherAppFlow](https://user-images.githubusercontent.com/17809047/134525202-404272ba-cff0-48f5-9ad2-63a12bdb141e.jpg)

1. Weather Service Build And Run
    * Install Golang
    * Install necessay golang packages 1. gin-gonic : Used as a http web framework; testify : Used to test golang project
        * sudo apt-get update -y
        * go get -u github.com/gin-gonic/gin 
            --- [used as router instead of native http library]
        * sudo apt-get install golang-github-stretchr-testify-dev 
            --- [used to run test cases]
        * go get github.com/patrickmn/go-cache 
            --- [used to store weather-info in cache]
        * go get github.com/gregjones/httpcache/diskcache 
            --- [to store weather info permanently in local disk as key-value pair]
2. Compile and Run the services
    * Checkout source code
       * https://github.com/hariprathap-hp/minderaWeatherService.git

3. Go to the folder minderaWeatherService
4. Modify the config file at ./config/config.go with both weatherstack api key and openweather api key (The keys can also be set as environment variables)
5. Run the application "go run main.go"
6. Send the api request from postman/browser "http://localhost:8080/v1/weather?city=melbourne"

Errors To Take care of while using library:
    * While running tests, if th error "Use 'mock' flag to tell package rest that you would like to use mockups." is observed, do the below configuration change
        --- go to "github.com/mercadolibre/golang-restclient/rest/mockup.go"
        --- Inside the func init(), comment the line "flag.Parse()"

