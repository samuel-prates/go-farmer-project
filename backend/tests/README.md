# API Tests

This directory contains documentation about the testing approach for the Farm Project API.

## Test Structure

The tests are organized to match the structure of the codebase:

- **Handler Tests**: Test the HTTP handlers that process API requests and generate responses.
- **Route Tests**: Verify that all API endpoints are correctly registered.
- **Main Tests**: Basic smoke tests to ensure the server initialization doesn't panic.

## Test Files

- `/internal/api/handlers/farmer_handler_test.go`: Tests for farmer-related endpoints
- `/internal/api/handlers/dashboard_handler_test.go`: Tests for dashboard-related endpoints
- `/internal/api/routes/routes_test.go`: Tests for route registration
- `/cmd/api/main_test.go`: Tests for server initialization

## Testing Approach

The tests use the following approach:

1. **Unit Tests**: Each handler is tested in isolation using mocks for dependencies.
2. **Mock Services**: Service layer is mocked to avoid database dependencies.
3. **HTTP Testing**: Uses Go's `httptest` package to simulate HTTP requests and responses.
4. **Table-Driven Tests**: Tests are structured as tables of test cases to make them more maintainable.

## Running the Tests

To run all tests:

```bash
cd backend
go test ./...
```

To run tests for a specific package:

```bash
go test ./internal/api/handlers
```

To run a specific test:

```bash
go test ./internal/api/handlers -run TestFarmerHandler_Create
```

To run tests with verbose output:

```bash
go test -v ./...
```

## Test Coverage

To generate a test coverage report:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

This will open a browser window showing which lines of code are covered by tests.

## Adding New Tests

When adding new endpoints or modifying existing ones:

1. Add corresponding test cases to the appropriate test file
2. Ensure both success and error scenarios are tested
3. Mock any external dependencies
4. Verify the HTTP status codes and response bodies

## Mocking Strategy

The tests use a simple mocking approach:

1. Define mock structs that implement the same interfaces as the real dependencies
2. Configure the mock's behavior using function fields
3. Inject the mocks into the handlers during test setup

This approach allows for flexible control over the behavior of dependencies during testing.