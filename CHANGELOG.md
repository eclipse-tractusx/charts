# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres
to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- This `CHANGELOG.md` file
- GitHub Workflow - PullRequest linting

### Changed

#### helmRepoIndex Go Code

- remove obsolete check for existing `gh-pages` branch in `func main`
- remove check `StatusCode == 200` from previous branch check
- func `downloadProductHelmRepoIndex`: Download `index.yaml` via GH RAW URL, error handling, add log output
- func `buildHelmRepoIndex` add log output

## [0.0.4] - 2022-11-28

### Fixed

- Path to workflow artefact (`git add ...`)

## [0.0.3] - 2022-11-28

### Fixed

- Use correct argument names for binary

## [0.0.2] - 2022-11-25

### Fixed

- Fix GH page download issues

## [0.0.1] 2022-11-23

### Added

- `DEPENDENCIES.md` file
- Basic Helm repository web page
- GitHub Workflow to build Helm repository
- GitHub Workflow to build Go binary
- initial `helmRepoIndex` Golang code
- `Authors.md` file
- `CODE_OF_CONDICT.md` file
- `CONTRIBUTING.md` file
- `LICENSE` file
- `Notice.md` file
- `Security.md` file

### Changed

- `README.md` file

[unreleased]: https://github.com/eclipse-tractusx/charts/compare/v0.0.4...HEAD
[0.0.4]: https://github.com/eclipse-tractusx/charts/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/eclipse-tractusx/charts/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/eclipse-tractusx/charts/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/eclipse-tractusx/charts/releases/tag/v0.0.1
