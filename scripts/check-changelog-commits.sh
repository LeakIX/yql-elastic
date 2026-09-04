#!/usr/bin/env bash
#
# Verify that the commit references in CHANGELOG.md are real.
#
# A changelog entry is only useful if its link takes you to the change. Ours are
# reference-style:
#
#     - Some change ([abc1234])
#     [abc1234]: https://github.com/LeakIX/yql-elastic/commit/abc1234
#
# and the hash is written by hand, so it rots in three ways. This checks all
# three:
#
#   1. The referenced commit exists in this repository. A hash typed from memory,
#      or one left behind after a commit was amended or rebased (which changes the
#      hash), points at nothing.
#   2. Every reference has a link definition, so it renders as a link rather than
#      as literal "[abc1234]" text.
#   3. The link URL points at the hash it is keyed by. Copy-pasting a previous
#      link and editing only the key gives you an entry that silently sends the
#      reader to the wrong commit.
#
# Usage: scripts/check-changelog-commits.sh [path-to-changelog]

set -euo pipefail

changelog="${1:-CHANGELOG.md}"

if [ ! -f "$changelog" ]; then
    echo "error: $changelog not found" >&2
    exit 1
fi

errors=0

# A short hash is 7 to 40 hex characters. Issue references ([#42]) do not match,
# and neither does prose, because the brackets and the hex-only body are required.
hash_re='[0-9a-f]{7,40}'

# --- 1. every referenced commit exists ---------------------------------------

referenced=$(grep -oE "\[${hash_re}\]" "$changelog" | tr -d '[]' | sort -u || true)

for key in $referenced; do
    if ! git cat-file -e "${key}^{commit}" 2>/dev/null; then
        echo "error: $changelog references commit $key, which does not exist here."
        echo "       If the commit was amended or rebased its hash changed; update the entry."
        errors=$((errors + 1))
    fi
done

# --- 2. every reference has a link definition --------------------------------

defined=$(grep -oE "^\[${hash_re}\]:" "$changelog" | tr -d '[]:' | sort -u || true)

for key in $referenced; do
    if ! echo "$defined" | grep -qx "$key"; then
        echo "error: $changelog references commit $key but never defines a link for it."
        echo "       Add: [$key]: https://github.com/LeakIX/yql-elastic/commit/$key"
        errors=$((errors + 1))
    fi
done

# --- 3. every link definition points at its own key ---------------------------

while IFS= read -r line; do
    [ -z "$line" ] && continue
    key=$(echo "$line" | grep -oE "^\[${hash_re}\]" | tr -d '[]')
    url_hash=$(echo "$line" | grep -oE "commit/${hash_re}" | sed 's|commit/||')
    if [ -n "$key" ] && [ -n "$url_hash" ] && [ "$key" != "$url_hash" ]; then
        echo "error: $changelog link [$key] points at commit/$url_hash."
        echo "       The link and its key must be the same commit."
        errors=$((errors + 1))
    fi
done < <(grep -E "^\[${hash_re}\]: https?://.*/commit/${hash_re}" "$changelog" || true)

# --- summary -----------------------------------------------------------------

if [ "$errors" -gt 0 ]; then
    echo
    echo "Found $errors changelog error(s)."
    exit 1
fi

count=$(echo "$referenced" | grep -c . || true)
echo "CHANGELOG.md: $count commit reference(s), all present and correctly linked."
