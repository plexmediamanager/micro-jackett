param(
    [string] $CustomTag
)

# Prompt user for the commit message
$COMMIT_MESSAGE = Read-Host "Please provide commit message"

& .\latest

# Get previous version string
$LATEST_COMMIT_HASH = git rev-list --tags --max-count=1
$VERSION = git describe --tags "$($LATEST_COMMIT_HASH)"
$VERSION_BITS = $VERSION.split('.')

Write-Host "Previous version tag: $($VERSION)"

if ($CustomTag.Length -lt 5) {
    $MAJOR_VERSION = $VERSION_BITS[0]
    $MINOR_VERSION = $VERSION_BITS[1]
    $PATCH_VERSION = $VERSION_BITS[2]

    $COUNT_OF_COMMIT_MSG_HAVE_SEMVER_MAJOR = (git log -1 --pretty=%B | Select-String -Pattern '\+semver:\s?(breaking|major)').Matches.Count
    $COUNT_OF_COMMIT_MSG_HAVE_SEMVER_MINOR = (git log -1 --pretty=%B | Select-String -Pattern '\+semver:\s?(feature|minor)').Matches.Count

    if ($COUNT_OF_COMMIT_MSG_HAVE_SEMVER_MAJOR > 0) {
        $MAJOR_VERSION += 1
    }
    if ($COUNT_OF_COMMIT_MSG_HAVE_SEMVER_MINOR > 0) {
        $MINOR_VERSION += 1
    }
    $conversionResult = [int]::TryParse($PATCH_VERSION, [ref]$PATCH_VERSION)
    $PATCH_VERSION += 1


    $NEW_RELEASE_TAG = "$($MAJOR_VERSION).$($MINOR_VERSION).$($PATCH_VERSION)"
} else {
    $NEW_RELEASE_TAG = $CustomTag
}

# Prepare files to be pushed
& git add .

$GIT_COMMIT_COUNT = $(git rev-list --count HEAD)
Write-Host "This is commit #$($GIT_COMMIT_COUNT)"

& git commit -m "$($COMMIT_MESSAGE)"
& git tag "$($NEW_RELEASE_TAG)"

Write-Host "Updating $($VERSION) to $($NEW_RELEASE_TAG)"

& git push --tags
& git push
