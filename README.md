# RPG Pack Calculator

This Go application is designed to calculate the optimal number and sizes of packs to fulfill customer orders for a product with different pack sizes.

## Project Structure

```
rpg
|-- cmd
|   `-- packcalculator
|       |-- Dockerfile
|       `-- main.go
|-- internal
|   `-- packcalculator
|       |-- handlers
|       |   |-- handler.go
|       |   `-- handler_test.go
|       |-- models
|       |   `-- pack.go
|       `-- services
|           |-- mocks
|           |   |-- calculator_mocks.go
|           |   `-- graph_mocks.go
|           |-- calculator.go
|           |-- calculator_test.go
|           |-- graph.go
|           `-- graph_test.go
|-- utils
|   |-- utils.go
|   `-- utils_test.go
|-- ui
|   |-- public
|   |   `-- index.html
|   |-- src
|   |   |-- components
|   |   |   |-- OrderForm.vue
|   |   |   `-- ResultDisplay.vue
|   |   |-- App.vue
|   |   `-- main.js
|   |-- Dockerfile
|   |-- babel.config.js
|   |-- jsconfig.json
|   |-- package.json
|   |-- package-lock.json
|   `-- vue.config.js
|-- .gitignore
|-- go.mod
|-- go.sum
|-- Makefile
|-- README.md
|-- docker-compose.yml
`-- RPG Pack Calculator.postman_collection.json
```

1. `cmd/packcalculator/main.go`:
   * *Purpose*: This file serves as the entry point for your application. It initializes core components such as HTTP handlers and starts the web server.

2. `internal/packcalculator/handlers/handler.go`:
   * *Purpose*: Contains the HTTP handler code that receives requests from clients and invokes corresponding services to process the requests.

3. `internal/packcalculator/handlers/handler_test.go`:
   * *Purpose*: Test file for the HTTP handler. Used to verify the correctness of request handling and response formation.

4. `internal/packcalculator/models/pack.go`:
   * *Purpose*: Defines the data structure representing information about pack size and quantity.

5. `internal/packcalculator/services`:
   * *Purpose*: This directory contains business logic and services for pack calculations.
   * `mocks/calculator_mocks.go` and `mocks/graph_mocks.go`: Mock implementations for testing purposes.
   * `calculator.go` and `calculator_test.go`: Implement the core algorithm for calculating optimal pack combinations based on given constraints.
   * `graph.go` and `graph_test.go`: Implement the graph-related logic used in the pack calculation algorithm.

6. `utils/utils.go` and `utils/utils_test.go`:
   * *Purpose*: Contains utility functions and unit tests for them, in this case, a function to calculate the sum of integers in an array.

7. `go.mod` and `go.sum`:
   * *Purpose*: These files manage the Go module and its dependencies.

8. `README.md`:
   * *Purpose*: A documentation file providing an overview of the project structure, instructions for running tests, and details about the main algorithm used for solving the problem, etc.

9. `RPG Pack Calculator.postman_collection.json`:
   * *Purpose*: Postman collection which includes pre-configured requests for the different test cases.

## Running Tests

To run the unit tests of the `handlers` package use the following command:
```
go test ./internal/packcalculator/handlers -v
```

Result:
```
=== RUN   TestCalculateHandler_ValidRequest
--- PASS: TestCalculateHandler_ValidRequest (0.00s)
=== RUN   TestCalculateHandler_InvalidMethod
--- PASS: TestCalculateHandler_InvalidMethod (0.00s)
=== RUN   TestCalculateHandler_InvalidJSON
--- PASS: TestCalculateHandler_InvalidJSON (0.00s)
PASS
ok      rpg/internal/packcalculator/handlers    0.669s
```

To run the unit tests of the `services` package use the following command:
```
go test ./internal/packcalculator/services -v
```

Result:
```
=== RUN   TestCalculatePacks_BaseExample
--- PASS: TestCalculatePacks_BaseExample (0.00s)
=== RUN   TestCalculatePacks_CustomSize
--- PASS: TestCalculatePacks_CustomSize (0.00s)
=== RUN   TestCalculatePacks_ZeroQuantity
--- PASS: TestCalculatePacks_ZeroQuantity (0.00s)
=== RUN   TestCalculatePacks_EmptySizes
--- PASS: TestCalculatePacks_EmptySizes (0.00s)
=== RUN   TestCalculatePacks_ZeroPackSizes
--- PASS: TestCalculatePacks_ZeroPackSizes (0.00s)
=== RUN   TestCalculatePacks_NegativePackSizes
--- PASS: TestCalculatePacks_NegativePackSizes (0.00s)
=== RUN   TestGeneratePermutations
=== RUN   TestGeneratePermutations/GeneratePermutations_with_Valid_Input
Actual result: [8 6 4 2 10 5 3 -1 9 -2 1 0 7]
=== RUN   TestGeneratePermutations/GeneratePermutations_with_Empty_Sizes
=== RUN   TestGeneratePermutations/GeneratePermutations_with_Negative_Quantity
=== RUN   TestGeneratePermutations/GeneratePermutations_with_Zero_Quantity
--- PASS: TestGeneratePermutations (0.00s)
    --- PASS: TestGeneratePermutations/GeneratePermutations_with_Valid_Input (0.00s)
    --- PASS: TestGeneratePermutations/GeneratePermutations_with_Empty_Sizes (0.00s)
    --- PASS: TestGeneratePermutations/GeneratePermutations_with_Negative_Quantity (0.00s)
    --- PASS: TestGeneratePermutations/GeneratePermutations_with_Zero_Quantity (0.00s)
=== RUN   TestClosestCandidate
=== RUN   TestClosestCandidate/NoCandidates
=== RUN   TestClosestCandidate/SingleCandidate
=== RUN   TestClosestCandidate/MultipleCandidates
--- PASS: TestClosestCandidate (0.00s)
    --- PASS: TestClosestCandidate/NoCandidates (0.00s)
    --- PASS: TestClosestCandidate/SingleCandidate (0.00s)
    --- PASS: TestClosestCandidate/MultipleCandidates (0.00s)
=== RUN   TestPruneNodes
=== RUN   TestPruneNodes/PruneNodes_with_Single_Candidate
=== RUN   TestPruneNodes/PruneNodes_with_Multiple_Candidates
--- PASS: TestPruneNodes (0.00s)
    --- PASS: TestPruneNodes/PruneNodes_with_Single_Candidate (0.00s)
    --- PASS: TestPruneNodes/PruneNodes_with_Multiple_Candidates (0.00s)
=== RUN   TestHasWeightedLine
=== RUN   TestHasWeightedLine/EmptyGraph
=== RUN   TestHasWeightedLine/SingleNodeGraph
=== RUN   TestHasWeightedLine/WeightedLinePresent
=== RUN   TestHasWeightedLine/ZeroWeightedLine
=== RUN   TestHasWeightedLine/NegativeWeightedLine
--- PASS: TestHasWeightedLine (0.00s)
    --- PASS: TestHasWeightedLine/EmptyGraph (0.00s)
    --- PASS: TestHasWeightedLine/SingleNodeGraph (0.00s)
    --- PASS: TestHasWeightedLine/WeightedLinePresent (0.00s)
    --- PASS: TestHasWeightedLine/ZeroWeightedLine (0.00s)
    --- PASS: TestHasWeightedLine/NegativeWeightedLine (0.00s)
=== RUN   TestAddWeightedLine
=== RUN   TestAddWeightedLine/AddToEmptyGraph
=== RUN   TestAddWeightedLine/AddDuplicateLine
=== RUN   TestAddWeightedLine/AddZeroWeightLine
--- PASS: TestAddWeightedLine (0.00s)
    --- PASS: TestAddWeightedLine/AddToEmptyGraph (0.00s)
    --- PASS: TestAddWeightedLine/AddDuplicateLine (0.00s)
    --- PASS: TestAddWeightedLine/AddZeroWeightLine (0.00s)
PASS
ok      rpg/internal/packcalculator/services    0.489s
```

To run the unit tests of the `utils` package use the following command:
```
go test ./utils -v
```

Result:
```
=== RUN   TestSum
=== RUN   TestSum/Positive_numbers
=== RUN   TestSum/Empty_slice
=== RUN   TestSum/Negative_numbers
=== RUN   TestSum/Mix_of_positive_and_negative_numbers
=== RUN   TestSum/Single_number
--- PASS: TestSum (0.00s)
    --- PASS: TestSum/Positive_numbers (0.00s)
    --- PASS: TestSum/Empty_slice (0.00s)
    --- PASS: TestSum/Negative_numbers (0.00s)
    --- PASS: TestSum/Mix_of_positive_and_negative_numbers (0.00s)
    --- PASS: TestSum/Single_number (0.00s)
PASS
ok      rpg/utils       0.502s
```

## Algorithm Explanation

The core algorithm is implemented inside `services` package. These services collectively deliver an effective solution aimed at identifying the most efficient combination of pack sizes to meet customer orders while satisfying predefined constraints.

1. Sorting Pack Sizes

   The algorithm begins by sorting the available pack sizes in ascending order. This ensures a systematic exploration of possibilities and simplifies subsequent calculations.


2. Order Quantity Optimization

   If the order quantity is significantly larger than the sum of available pack sizes, the algorithm takes an optimization step. It intelligently reduces the problem space by subtracting packs from the largest available size. This brings the order quantity closer to a predefined clamp value, preventing unnecessary calculations for extremely large orders.


3. Graph Generation

   The core of the algorithm involves the creation of a graph that represents the possible permutations of quantities achievable with the available pack sizes. This is accomplished through a recursive process of subtracting pack sizes from the order quantity. Each node in the graph represents a unique quantity, and edges correspond to the pack sizes used to transition between quantities.


4. Pruning Unnecessary Nodes

   After graph generation, the algorithm optimizes the graph by pruning unnecessary nodes. Nodes that don't contribute to reaching the goal or have no edges going out are removed. This pruning step significantly reduces the size of the graph, enhancing the efficiency of subsequent calculations.


5. A* Search Algorithm

   The algorithm employs the A* search algorithm to find the shortest path from the initial quantity (customer order) to the quantity closest to zero. A* uses a heuristic to guide the search, ensuring an optimal path is found efficiently.


6. Analyzing the Shortest Path

   The shortest path obtained from the A* search represents the most efficient way to fulfill the customer order. The algorithm then analyzes this path, counting the number of each pack size used.


7. Result Calculation

   The final result is a map of pack sizes and their corresponding quantities, providing the optimal combination of packs needed to fulfill the order while meeting the specified constraints.


In summary, the algorithm combines sorting, optimization, graph generation, pruning, and the A* search algorithm to systematically explore and find the optimal solution for packing items according to specific customer orders. The resulting implementation is both efficient and scalable, providing a robust solution for the given problem statement.

## Quick Launch Commands

Makefile:
```
# Build the Golang backend binary.
build-backend:
	cd ./cmd/packcalculator && go build -o ../../bin/packcalculator

# Run the Golang backend application.
run-backend:
	cd ./cmd/packcalculator && go run main.go

# Build the Vue.js frontend application.
build-frontend:
	cd ./ui && npm install && npm run build

# Run the Vue.js frontend application.
run-frontend:
	cd ./ui && npm run serve -- --port 8081
```

To quickly build and run both the Golang backend and Vue.js frontend locally on your computer, you can use the following commands from the provided Makefile:

* `make build-backend`: Build the Golang backend binary.
* `make run-backend`: Run the Golang backend application.
* `make build-frontend`: Build the Vue.js frontend application.
* `make run-frontend`: Run the Vue.js frontend application.

Execute the following commands to initiate the project within `Docker` containers:
```
docker-compose down
docker-compose build
docker-compose up -d
```

* `docker-compose down`: Stops and removes containers, networks, and volumes defined in the `docker-compose.yml`.
* `docker-compose build`: Builds or rebuilds services defined in the `docker-compose.yml`.
* `docker-compose up -d`: Starts the project in detached mode, running containers in the background.ß

## Running the Golang Backend Application

To run the Golang Backend application locally on your computer, follow these steps:

1. **Install Golang**

   * Ensure that Golang is installed on your system. You can download and install Golang from the [official website](https://go.dev/doc/install).
   * After installation, make sure the `GOPATH` environment variable is set and added to the `PATH` variable.


2. **Download Dependencies**
   * Navigate to the project root directory:
   ```
   cd rpg
   ```
   * Execute the following command to download and install the project dependencies specified in `go.mod`:
   ```
   go mod download
   ```

3. **Run the Golang Application**

   Execute the following command to run the Golang application:
   ```
   go run cmd/packcalculator/main.go
   ```

4. **Verification**:
   * The Golang application uses the `RPG_BACKEND_PORT` environment variable to determine the port on which the server should listen. If the variable is not set, the application defaults to port `8080`.
   * After successful startup, you should see a log message indicating the server starting on a specific port, for example:
   ```
   Server starting on port 8080...
   ```
## Running the Vue.js Frontend Application

To run the Vue.js Frontend application locally on your computer, follow these steps:

1. **Install Node.js and npm**

   Ensure that Node.js and npm are installed on your system. You can download and install them from the [official Node.js website](https://nodejs.org/en).


2. **Navigate to the UI Directory**
   ```
   cd rpg/ui
   ```

3. **Install Dependencies**

   Execute the following command to install project dependencies:
   ```
   npm install
   ```

4. **Run the Vue.js Application**

   Run the Vue.js application using the command. Specify a port that does not conflict with the Golang Backend application (e.g., `8081`):
   ```
   npm run serve -- --port 8081
   ```

5. **Verification**
   * After successful startup, you should see a message in the console indicating the URL of the local server, for example, `http://localhost:8081/`.
   * Open this URL in your web browser, and you will see your Vue.js Frontend application.

## API Testing

You can use `curl` utility, a tool like `Postman` or `Vue.js Frontend Application` with simple `UI` to test the API with the following scenarios:

### 1. Single Item Order
```
curl -X POST -H "Content-Type: application/json" -d '{
    "order": 1,
    "pack_sizes": [250, 500, 1000, 2000, 5000]
}' http://localhost:8080/calculate
```

![Single Item Order](screenshots/1.%20Single%20Item%20Order.png)

![Single Item Order (UI)](screenshots/1.%20Single%20Item%20Order%20(UI).png)

### 2. Order Matching a Single Pack
```
curl -X POST -H "Content-Type: application/json" -d '{
    "order": 250,
    "pack_sizes": [250, 500, 1000, 2000, 5000]
}' http://localhost:8080/calculate
```

![Order Matching a Single Pack](screenshots/2.%20Order%20Matching%20a%20Single%20Pack.png)

![Order Matching a Single Pack (UI)](screenshots/2.%20Order%20Matching%20a%20Single%20Pack%20(UI).png)

### 3. Order Just Above Single Pack Size
```
curl -X POST -H "Content-Type: application/json" -d '{
    "order": 251,
    "pack_sizes": [250, 500, 1000, 2000, 5000]
}' http://localhost:8080/calculate
```

![Order Just Above Single Pack Size](screenshots/3.%20Order%20Just%20Above%20Single%20Pack%20Size.png)

![Order Just Above Single Pack Size (UI)](screenshots/3.%20Order%20Just%20Above%20Single%20Pack%20Size%20(UI).png)

### 4. Order Requiring Multiple Packs
```
curl -X POST -H "Content-Type: application/json" -d '{
    "order": 501,
    "pack_sizes": [250, 500, 1000, 2000, 5000]
}' http://localhost:8080/calculate
```

![Order Requiring Multiple Packs](screenshots/4.%20Order%20Requiring%20Multiple%20Packs.png)

![Order Requiring Multiple Packs (UI)](screenshots/4.%20Order%20Requiring%20Multiple%20Packs%20(UI).png)

### 5. Large Order
```
curl -X POST -H "Content-Type: application/json" -d '{
    "order": 12001,
    "pack_sizes": [250, 500, 1000, 2000, 5000]
}' http://localhost:8080/calculate
```

![Large Order](screenshots/5.%20Large%20Order.png)

![Large Order (UI)](screenshots/5.%20Large%20Order%20(UI).png)

### 6. Custom Pack Sizes
```
curl -X POST -H "Content-Type: application/json" -d '{
    "order": 263,
    "pack_sizes": [23, 31, 53]
}' http://localhost:8080/calculate
```

![Custom Pack Sizes](screenshots/6.%20Custom%20Pack%20Sizes.png)

![Custom Pack Sizes (UI)](screenshots/6.%20Custom%20Pack%20Sizes%20(UI).png)

To test the API, you can use `Postman` and import the provided Postman collection called `RPG Pack Calculator.postman_collection.json` located in the root directory. This collection includes pre-configured requests for the following test cases:
1. Single Item Order | Order: 1, Pack Sizes: [250, 500, 1000, 2000, 5000]
2. Order Matching a Single Pack | Order: 250, Pack Sizes: [250, 500, 1000, 2000, 5000]
3. Order Just Above Single Pack Size | Order: 251, Pack Sizes: [250, 500, 1000, 2000, 5000]
4. Order Requiring Multiple Packs | Order: 501, Pack Sizes: [250, 500, 1000, 2000, 5000]
5. Large Order | Order: 12001, Pack Sizes: [250, 500, 1000, 2000, 5000]
6. Custom Pack Sizes | Order: 263, Pack Sizes: [23, 31, 53]

Follow these steps to use the Postman collection:
1. Open Postman.
2. Click on "Import" in the top left corner.
3. Select "Choose File" and navigate to the project's root directory.
4. Select the file RPG Pack Calculator.postman_collection.json and click "Open" to import the collection.
5. The collection will now be visible in the left sidebar of Postman.

You can then execute each request from the collection to test the API for different scenarios. Adjust the request payload or parameters as needed.

## User Interface Validation

The Vue.js Frontend Application includes a simple yet effective input validation in `App.vue`. The validation is applied to the order quantity and pack sizes fields:

* **Order Quantity Validation**:
  * Validates that the order quantity is a positive integer.
  * Displays an error message if the input is invalid.
* **Pack Sizes Validation**:
  * Validates that pack sizes are valid numbers separated by commas.
  * Displays an error message if the input is invalid.
* **Calculate Button State**:
  * The "Calculate Packs" button is disabled until both the order quantity and pack sizes are valid.
  * Prevents the user from triggering calculations with invalid input.
* **Error Messages**:
  * Error messages are displayed below the respective input fields when validation fails.
  * Provides clear feedback to the user about the nature of the validation error.

This validation ensures that users provide correct and reasonable input before triggering the pack calculation, enhancing the overall user experience.