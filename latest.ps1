$goModContent = Get-Content .\go.mod
$packageSource = ""
$shouldBeUpdated = New-Object System.Collections.Generic.List[String]

foreach ($line in $goModContent) {
    $currentIndex = $goModContent.IndexOf($line)
    if ($currentIndex -eq "0" -and $packageSource -eq "") {
        $moduleName = $line.Replace("module", "").Trim()
        $packageSource = $moduleName.Substring(0, $moduleName.lastIndexOf('/'))
    }
    if ($packageSource -ne "" -and $currentIndex -ne "0") {
        if ($line -match $packageSource -and $line -inotmatch "replace") {
            $shouldBeUpdated.Add($line.Substring(0, $line.indexOf(' ')).Trim())
        }
    }
}

foreach ($module in $shouldBeUpdated) {
    & go get $module.Trim()@master
}