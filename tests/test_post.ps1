$url = "http://localhost:8080/users"
$headers = @{ "Content-Type" = "application/json" }
$body = @{
    name = "John Doe"
    age = 30
} | ConvertTo-Json

Invoke-RestMethod -Uri $url -Method POST -Headers $headers -Body $body
