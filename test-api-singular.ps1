$baseUrl = "http://localhost:8080"
$companyIds = @("1")
$endpoints = @("/financials")

# Generate valid URLs with proper query parameters
$requests = foreach ($i in 1..10) {
    $companyId = $companyIds | Get-Random
    $endpoint = $endpoints | Get-Random
    @{
        Url = "$baseUrl$endpoint`?companyId=$companyId"  
        Name = "Request-$i"
    }
}

# Function to make API calls
$apiCall = {
    param($url)
    $start = Get-Date
    try {
        $response = Invoke-RestMethod -Uri $url -Method Get
        $result = [PSCustomObject]@{
            URL = $url
            Status = "Success"
            Response = $response
            DurationMs = [math]::Round((Get-Date $start).TotalMilliseconds, 2)
            StartTime = $start.ToString("HH:mm:ss.fff")
        }
    } catch {
        $result = [PSCustomObject]@{
            URL = $url
            Status = "Failed"
            Response = $_.Exception.Message
            DurationMs = [math]::Round((Get-Date $start).TotalMilliseconds, 2)
            StartTime = $start.ToString("HH:mm:ss.fff")
        }
    }
    return $result
}

# Run requests concurrently
$jobs = @()
foreach ($req in $requests) {
    $jobs += Start-Job -Name $req.Name -ScriptBlock $apiCall -ArgumentList $req.Url
}

# Wait for completion and show results
$results = $jobs | Wait-Job | Receive-Job
$results | Format-Table URL, Status, DurationMs, StartTime -AutoSize
$results | Select-Object -First 1 | Format-List *  # Show full details of first result

# Cleanup
$jobs | Remove-Job