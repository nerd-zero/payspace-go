#!/usr/bin/env bash
#
# version.sh — Bump CalVer version, commit, and tag.
#
# Format: 0.YYYYMM.MICRO (MICRO is zero-padded to 3 digits, starts at 001)
#
# Usage:
#   ./scripts/version.sh             # bump, commit, tag, push to GitHub
#   ./scripts/version.sh --dry-run   # preview without making changes

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
VERSION_FILE="$REPO_ROOT/VERSION"

# Parse flags.
DRY_RUN=false
for arg in "$@"; do
    case "$arg" in
        --dry-run)  DRY_RUN=true ;;
        -h|--help)
            echo "Usage: $0 [--dry-run]"
            echo ""
            echo "  --dry-run   Preview the next version without making changes"
            exit 0
            ;;
        *)
            echo "Unknown flag: $arg" >&2
            exit 1
            ;;
    esac
done

# Read current version.
if [[ ! -f "$VERSION_FILE" ]]; then
    echo "Error: VERSION file not found at $VERSION_FILE" >&2
    exit 1
fi
CURRENT="$(tr -d '[:space:]' < "$VERSION_FILE")"

# Parse current version: 0.YYYYMM.MICRO
CURRENT_YYYYMM="$(echo "$CURRENT" | cut -d. -f2)"
CURRENT_MICRO="$(echo "$CURRENT" | cut -d. -f3)"

# Compute new version based on today's date.
NOW_YYYYMM="$(date +%Y%m)"

if [[ "$NOW_YYYYMM" == "$CURRENT_YYYYMM" ]]; then
    # Same month — increment MICRO.
    NEXT_MICRO="$(printf '%03d' $(( 10#$CURRENT_MICRO + 1 )))"
    NEXT="0.${NOW_YYYYMM}.${NEXT_MICRO}"
else
    # New month — reset MICRO to 001.
    NEXT="0.${NOW_YYYYMM}.001"
fi

TAG="v${NEXT}"

echo "Current version: $CURRENT"
echo "Next version:    $NEXT"
echo "Tag:             $TAG"

if $DRY_RUN; then
    echo ""
    echo "(dry run — no changes made)"
    exit 0
fi

# Ensure we are on the main branch.
BRANCH="$(git -C "$REPO_ROOT" rev-parse --abbrev-ref HEAD)"
if [[ "$BRANCH" != "main" ]]; then
    echo "Error: releases must be created from the main branch (currently on '$BRANCH')." >&2
    exit 1
fi

# Ensure working tree is clean (except VERSION).
if [[ -n "$(git -C "$REPO_ROOT" diff --name-only -- ':!VERSION')" ]]; then
    echo "Error: working tree has uncommitted changes. Commit or stash them first." >&2
    exit 1
fi

# Write new version.
echo "$NEXT" > "$VERSION_FILE"

# Commit and tag.
git -C "$REPO_ROOT" add VERSION
git -C "$REPO_ROOT" commit -m "chore: bump version to $NEXT"
git -C "$REPO_ROOT" tag -a "$TAG" -m "Release $TAG"

echo ""
echo "Created commit and tag $TAG"

git -C "$REPO_ROOT" push --follow-tags origin main
echo "Pushed to remote."
