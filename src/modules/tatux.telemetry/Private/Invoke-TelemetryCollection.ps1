function Invoke-TelemetryCollection {
    param (
        [CmdletBinding(OutputType = 'json', HelpUri = 'https://pwsh.dev.tatux.co.uk/tatux.telemetry/docs/Invoke-TelemetryCollection.html')]
        [Parameter(Mandatory = $true)]
        [array]$ExecutionContext,

        [Parameter(Mandatory = $true)]
        [switch]$Minimal
    )

    
    
}