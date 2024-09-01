$baseUrl = "http://localhost:8080"

function Get-ApiResponse {
    param (
        [string]$endpoint
    )

    $url = "$baseUrl$endpoint"
    Write-Host "Executing GET to $url"

    try {
        $response = Invoke-WebRequest -Uri $url -Method Get
        Write-Host "Status code: $($response.StatusCode)"
        Write-Host "Response:"
        $content = $response.Content 
        Write-Output ($content | Out-String)
    } catch {
        Write-Host "Error in the request:"
        Write-Output $_.Exception.Message
    }
}

# Execute endpoints
Get-ApiResponse "/"
Start-Sleep -Seconds 1
Get-ApiResponse "/users"
Start-Sleep -Seconds 1
Get-ApiResponse "/users/1"
Start-Sleep -Seconds 1
Get-ApiResponse "/users/2"
