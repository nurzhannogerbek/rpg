# RPG Pack Calculator

This Go application is designed to calculate the optimal number and sizes of packs to fulfill customer orders for a product with different pack sizes.

## Project Structure

```
rpg
|-- cmd
|   `-- packcalculator
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
|-- .gitignore
|-- go.mod
|-- go.sum
|-- README.md
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

## Running the Project Locally

To run the project locally, follow these steps:

1. Download the project dependencies using the following command:
```
go mod download
```

2. Run the project with the following command:
```
go run cmd/packcalculator/main.go
```

The server will start on port `8080`.

## API Testing

You can use `curl` utility or a tool like `Postman` to test the API with the following scenarios:

### 1. Single Item Order
```
curl -X POST -H "Content-Type: application/json" -d '{
    "order": 1,
    "pack_sizes": [250, 500, 1000, 2000, 5000]
}' http://localhost:8080/calculate
```

![1. Single Item Order](screenshots/1.%20Single%20Item%20Order.png)

### 2. Order Matching a Single Pack
```
curl -X POST -H "Content-Type: application/json" -d '{
    "order": 250,
    "pack_sizes": [250, 500, 1000, 2000, 5000]
}' http://localhost:8080/calculate
```

![2. Order Matching a Single Pack](screenshots/2.%20Order%20Matching%20a%20Single%20Pack.png)

### 3. Order Just Above Single Pack Size
```
curl -X POST -H "Content-Type: application/json" -d '{
    "order": 251,
    "pack_sizes": [250, 500, 1000, 2000, 5000]
}' http://localhost:8080/calculate
```

![3. Order Just Above Single Pack Size](screenshots/3.%20Order%20Just%20Above%20Single%20Pack%20Size.png)

### 4. Order Requiring Multiple Packs
```
curl -X POST -H "Content-Type: application/json" -d '{
    "order": 501,
    "pack_sizes": [250, 500, 1000, 2000, 5000]
}' http://localhost:8080/calculate
```

![4. Order Requiring Multiple Packs](screenshots/4.%20Order%20Requiring%20Multiple%20Packs.png)

### 5. Large Order
```
curl -X POST -H "Content-Type: application/json" -d '{
    "order": 12001,
    "pack_sizes": [250, 500, 1000, 2000, 5000]
}' http://localhost:8080/calculate
```

![5. Large Order](screenshots/5.%20Large%20Order.png)

### 6. Custom Pack Sizes
```
curl -X POST -H "Content-Type: application/json" -d '{
    "order": 263,
    "pack_sizes": [23, 31, 53]
}' http://localhost:8080/calculate
```

![Custom Pack Sizes](screenshots/6.%20Custom%20Pack%20Sizes.png)

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