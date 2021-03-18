# API Farm

A repository of APIs written in different programming languages and frameworks. Each framework's development has been unit and API tested. Also, documented with instructions for how to start up locally (addtional docker support included too).

### API Use Case

The api developed in each environment should support creation, reading, updating and deletion of a resource. The management of data can vary from in memory, file storage or a database.

Each api should be unit tested with a language supported framework where possible. Also, generic API tests should be run against each working instance of the API. Results from these tests should be collected for comparison.

To allow fair metric comparison from api tests, the api model resource will be standardised across each of the languages and frameworks. The api model resource will be a `video game`. This resource has been structured below in JSON;

```json
{
    "name": "Banjo Kazooie",
    "developers": [
        "Rare",
    ],
    "publishers": [
        "Nintendo"
    ],
    "directors": [
        "Gregg Mayles",
        "George Andreas"
    ],
    "producers": [
        "Tim Stamper",
        "Chris Stamper"
    ],
    "designers": [
        "Gregg Mayles"
    ],
    "programmers": [
        "Chris Sutherland"
    ],
    "artists": [
        "Steve Mayles",
        "John Nash",
        "Kevin Bayliss",
        "Tim Stamper"
    ],
    "composers": [
        "Grant Kirkhope"
    ],
    "platforms": [
        "Nintendo 64"
    ],
    "date_released": "29 June 1998",  
}
```

### Languages & Frameworks

Below, the various languages and frameworks are listed which are used within this repository;

| Language | API Framework | Test Framework | Storage Support | Link |
| -------- | ------------- | -------------- | --------------- | ---- |
| Ruby     | Sinatra       | Rspec          | In Memory       | TBD  |

### Project Structure

This repository is designed to clearly silo each of the api samples.

```
api_farm
|___lang-1
|   |___framework-1
|   |___framework-2
|___lang-2
    |___framework-1
```

Each framework will have its own standalone api example.