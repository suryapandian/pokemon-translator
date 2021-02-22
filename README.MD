### Pokémons in a Shakespearean world

## Design

This is a simple web server with an endpoint to get the Shakespearean description of Pokémons. 

The application gets the description of Pokémons from pokeapi.co's go sdk package [pokeapi](https://pokeapi.co/) and the Shakespearean translations from [funtranslations](https://funtranslations.com/api/shakespeare). The required data from the responses are cached locally to avoid subsequent requests for the same payload to the third-party sites. The cache is a simple pluggable in-memory interface. Since the cache is a pluggable module, if needed we can replace the cache with persistent storage with minimal code change.

The app is tested and a CI pipeline is also set up through GitHub actions. The requests to third-party websites like `pokeapi` and `funtranslations` are mocked for testing purposes.

### Endpoints

# /pokemon/pikachu (GET)

 This API provides the Shakespearean description of the requested pokémon.


+ Response 200 (application/json)
```json
  {
    "string": "pikachu",
    "description": "At which hour several of these pokémon gather,  their electricity couldst buildeth and cause lightning storms."
  }
```

+ Response 404 (application/json)
```json
  {
      "code": 404,
      "message": "pokemon not found"
  }
```

## Running the application

```
 docker build -t pokemon .  

 docker run --publish 5000:5000 -i -t "pokemon"

 ```

Configurations like the log level, port number, the translation URL, etc... can also be passed as environment variables. If no environment variable is set the application uses the default configurations. 

To run the application without docker, please install go 1.15 and run `go run -mod=vendor main.go`

## Future implementations
 - Implement a generic cache proxy and use it to make requests to both pokeapi and translations.

 - Since the app is dependent on third-party websites set up a circuit breaker and a failover mechanism.

 - While invoking the internal third-party wrapper module, as of now only the API URL is passed. In the future, we might want to pass API token and other configurations for circuit breaking, etc… So, we can change the wrapper module to accept an object instead of a string to accommodate as many additional attributes as needed.

 - pokeapi's go SDK has been used so that it would be easier to extend the app and support additional endpoints. However, if there is no plan for extending the app, we could remove the dependency on the poke-api package and directly make the API call. The lesser the number of dependencies the easier the code is to maintain.

## Pre-commit 
To ensure a clean code.

   Install pre-commit [https://pre-commit.com/#install]
   Install .pre-commit-config.yaml as a pre-commit hook

  `pre-commit install`

  Go static analysis tools run automatically on pre-commit. 
  Run checks manually if needed using:

  `pre-commit run --all-files`
