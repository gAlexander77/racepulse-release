# RacePulse distribution

This repository hosts public releases, user documentation, issues, changelogs, and the Hugo site. The private `racepulse` source repository builds the Windows app and transfers verified assets into an unpublished draft here. This repository does not clone or build private application source. The GitHub Pages website is the only push-triggered deploy path that should remain automatic; app packaging, API/Fly deploys, and release publication are manual.

## Handover status

Prepared 2026-09-30. The old public `.github/workflows/release.yml` was removed from main in `fa9c268` so public main no longer clones private source or publishes app releases from a tag/dispatch. Public tooling PR #3 was closed during the second review so it cannot restore a competing build path. This handover is based on public main `8725a84`, not that tooling branch. Source-owned uploads still check that the legacy workflow is absent before uploading. **0.3.5 remains on hold.** This cleanup changes neither the current stable v0.3.4 nor website version/download settings.

The public `SOURCE_REPO_PAT` Actions secret was deleted during the second review. Revoke its underlying token only after establishing that it was dedicated to this retired workflow. Public site builds need no credential for private source. The owner has stored `RELEASE_REPO_TOKEN` in private source Actions secrets; its presence was confirmed through metadata only. Actual scope/expiration/access and candidate upload acceptance remain unverified.

Private source PR #3 merged to main in `6b53855` on 2026-09-30. Source CI and its draft-build replacement are both active and accept only manual dispatch; the old private tag-triggered release is retired. The merge did not start a hosted build. Website Pages is the only automatic CI/CD path.

Active `Immutable version tags` rulesets block updates and deletion of `refs/tags/v*` in both source and distribution, with no bypass actors; new tags remain allowed. Distribution rule: [24283108](https://github.com/gAlexander77/racepulse-release/rules/24283108). Do not weaken this rule or move a version tag to repair a published release.

Credential storage is complete. Its required scope remains **`racepulse-release`, Contents read/write**, stored as **`RELEASE_REPO_TOKEN` in private source Actions secrets**, not here. Secret presence alone does not prove that scope or live access. Do not paste it into chat or reuse the broad `gh` login credential. Native Windows/live-game acceptance and the actual published-v0.3.4 upgrade to an exact candidate remain required before publication. Adding a secret alone does not build or publish anything; the private source runbook owns candidate preflight and the full acceptance matrix.

## Candidate assets

Each source build supplies:

- `RacePulse.exe`: portable Windows amd64 application, including a freshly built widget.
- `RacePulse.exe.sha256`: ASCII/UTF-8 without BOM, with the executable SHA-256 as its first field.
- `release-manifest.json`: version, approved source repository/SHA, build platform, exact tool versions, workflow run/attempt, and artifact sizes/hashes. No private source, documents, logs, or credentials.

Keep the repository name `racepulse-release`, asset names, and existing release URLs stable for installed clients. Checksums detect corrupted downloads; they are not code-signing signatures. Do not replace the portable EXE/checksum feed with installer-only assets or a different updater manifest without an explicit compatibility migration.

The source tag/SHA identifies application code. A tag in this public repository identifies distribution metadata and is not evidence of which private source was compiled. Do not move published tags or overwrite published assets.

## Review and promotion

1. Review the exact source build run and retained bundle. Inspect the draft by release ID, not just tag: tag lookup may fail for unpublished drafts.
2. Ensure the successful source upload verified all three assets byte-for-byte. Never promote partial/conflicting uploads. Retries use the original verified bundle and do not clobber assets.
3. Require native application/widget checks and a successful upgrade from the actual published v0.3.4 executable, preserving existing settings, layouts, sessions, and device identity. Its older restart behavior requires direct testing; a clean installation alone is insufficient.
4. Review public release notes and both `CHANGELOG.md` and `site/content/changelog.md`. The private build cannot assess public prose automatically.
5. Obtain owner approval, then publish the same verified draft. Stable releases must be non-draft and non-prerelease. Check anonymous EXE/checksum downloads against the manifest and confirm both GitHub's Latest result and the first eligible entry in the releases list used by installed clients.
6. Update `site/hugo.toml`'s paired `appVersion` and `downloadURL` to the verified stable release, keep both changelogs synchronized, and deploy Pages from public `main`. Inspect the live download/changelog. Publication does not itself trigger this repository's Pages workflow. Do not add push-triggered app/API/release deploy workflows here.
7. Record source SHA, public SHA, build run, release ID, asset hashes, upgrade evidence, and publication/site checks in the private source release receipt.

## Failure handling

Before publication, leave a failed draft unpublished and keep the current stable/site unchanged. Resume with identical retained assets or explicitly discard the unpublished draft after review. Never overwrite published assets to repair a version.

After publication, investigate and remove a bad version from eligible updater selection, restore the known-working website link, then publish a verified higher recovery version. Marking an older release Latest does not downgrade installed clients that already have a higher version. The private source runbook owns detailed compatibility and recovery checks.
