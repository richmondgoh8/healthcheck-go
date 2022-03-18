Sample Health Checkpoint which utilizes the Gin Framework and Simple Validation

## Local Installations
```
make run
curl --location --request GET 'http://localhost:8080/health?isListening=1'
```

To update current mocking packages
```
make gen-mocks
```

To check for test coverage
```
make coverage
```


## Libraries
- Gin Framework
- Go Swagger
- Viper for Reading Config
- Gen Mocks