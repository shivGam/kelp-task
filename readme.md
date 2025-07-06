# Kelp Task 

## Prerequisites to run the code

- Go 1.24.1 or higher
- PowerShell (for running test scripts on Windows)

## Installation

1. **Clone the repository** (if not already done):
   ```bash
   git clone <https://github.com/shivGam/kelp-task.git>
   cd kelp-task
   ```

2. **Install Go dependencies**:
   ```bash
   go mod download
   ```

## Configuration

### Company ID Range

The companyId parameter in API calls must be within the configured range (1-3 by default).
To change the number of companies, modify the `main.go` file:

```go
// In main.go, line 10
db.InitDB(3)  // Change this number to desired company count
```

**Note**: 
## Running the Application

1. **Start the server**:
   ```bash
   go run main.go
   ```

2. **Server will start on**: `http://localhost:8080`


## API Endpoints

| Endpoint | Method | Description | Query Parameter |
|----------|--------|-------------|-----------------|
| `/companies` | GET | Get all companies | None |
| `/financials` | GET | Get financial data | `companyId` (1-3) |
| `/sales` | GET | Get sales data | `companyId` (1-3) |
| `/employees` | GET | Get employee data | `companyId` (1-3) |

### Example API Calls

```bash
# Get all companies
curl http://localhost:8080/companies

# Get financial data for company ID 1
curl http://localhost:8080/financials?companyId=1

# Get sales data for company ID 2
curl http://localhost:8080/sales?companyId=2

# Get employee data for company ID 3
curl http://localhost:8080/employees?companyId=3
```

## Testing

### PowerShell Test Scripts

The project includes two PowerShell test scripts:

1. **`test-api.ps1`** - Tests all endpoints with random company IDs (1-3)
2. **`test-api-singular.ps1`** - Tests financial endpoint with company ID 1 only

### Running Tests

1. **Make sure the server is running** (see Running the Application above)

2. **Run comprehensive tests**:
   ```powershell
   .\test-api.ps1
   ```

3. **Run singular endpoint test**:
   ```powershell
   .\test-api-singular.ps1
   ```

4. **Test again**: Delete `companies.db` and run the tests again.

## Project Structure

```
kelp-task/
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── db/
│   └── db.go              # Database initialization and setup
├── handlers/
│   ├── companies.go        # Company data handlers
│   ├── employee.go         # Employee data handlers
│   ├── financial.go        # Financial data handlers
│   └── sale.go            # Sales data handlers
├── models/
│   ├── company.go          # Company data model
│   ├── employee-model.go   # Employee data model
│   ├── financial-model.go  # Financial data model
│   └── sale-model.go       # Sales data model
├── test-api.ps1           # Comprehensive API test script
└── test-api-singular.ps1  # Single endpoint test script
```

## Dependencies

- **Gin**: Web framework for HTTP routing
- **SQLite**: Database driver (modernc.org/sqlite)
- **Standard Go libraries**: database/sql, math, rand, etc.

