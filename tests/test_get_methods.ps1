# Definir la URL base de la API
$baseUrl = "http://localhost:8080"

# Función para realizar una solicitud GET y mostrar la respuesta
function Get-ApiResponse {
    param (
        [string]$endpoint
    )

    $url = "$baseUrl$endpoint"
    Write-Host "Haciendo solicitud GET a $url"

    try {
        $response = Invoke-RestMethod -Uri $url -Method Get
        Write-Host "Respuesta:"
        Write-Output $response
    } catch {
        Write-Host "Error:"
        Write-Output $_.Exception.Message
    }
}

# Ejecutar solicitudes GET a los endpoints deseados
Get-ApiResponse "/"
Get-ApiResponse "/users"