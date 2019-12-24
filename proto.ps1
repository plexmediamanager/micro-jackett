$goPathFromEnvironment = Get-Content env:\GOPATH
$goPathSources = Join-Path $goPathFromEnvironment "src"

$currentDirectory = $PSScriptRoot

$protoSourceDirectory = Join-Path $currentDirectory "proto_src"
$protoTargetDirectory = Join-Path $currentDirectory "proto"
$protobufFiles = Get-ChildItem $protoSourceDirectory -File | Where-Object { $_.Extension -eq ".proto" }

foreach ($proto in $protobufFiles) {
    protoc -I $goPathSources --proto_path=$protoSourceDirectory --micro_out=$protoTargetDirectory --go_out=$protoTargetDirectory $proto
}