$url = "http://localhost:8080/users/3"
$headers = @{ "Content-Type" = "application/json" }
$body = @{
    age = 34
} | ConvertTo-Json

Invoke-RestMethod -Uri $url -Method PATCH -Headers $headers -Body $body