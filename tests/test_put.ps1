$url = "http://localhost:8080/users/2"
$headers = @{ "Content-Type" = "application/json" }
$body = @{
    name = "New Name"
    age = 33
} | ConvertTo-Json

Invoke-RestMethod -Uri $url -Method PUT -Headers $headers -Body $body
