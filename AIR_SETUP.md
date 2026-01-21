# Air Configuration for ecommerce-api

## Overview

Air is installed and configured for live reloading development. Air watches your source files and automatically rebuilds and restarts your application when changes are detected.

## Installation ✅

Air v1.49.0 is already installed in your Go environment.

## Configuration

The `.air.toml` file has been created with the following settings:

- **Watch directories**: `cmd`, `internal`, `pkg`
- **Watch file extensions**: `.go`, `.tpl`, `.tmpl`, `.html`, `.yaml`, `.yml`, `.json`
- **Build output**: `./tmp/main`
- **Excluded**: `vendor`, `testdata`, `*_test.go`, and common IDE directories
- **Build delay**: 1000ms (prevents rapid rebuilds)

## Usage

### Development Mode (with live reload)

```bash
make dev
```

or

```bash
air -c .air.toml
```

### Production Mode (standard build)

```bash
make run
```

### Build Binary

```bash
make build
```

## How It Works

1. Air watches files in your project for changes
2. When a file is modified, Air waits 1000ms then triggers a rebuild
3. The application is automatically restarted with the new binary
4. Console output shows build status and application logs

## Features

- **Automatic Restart**: No manual restarts needed during development
- **Colored Output**: Different colors for build, main app, watcher, and runner
- **Smart Watching**: Only watches relevant directories and file types
- **Error Logging**: Build errors are logged to `build-errors.log`
- **Graceful Reload**: Application is properly shut down before restart

## Troubleshooting

### Common Issues

**Issue**: "Air command not found"

- **Solution**: Make sure Go bin directory is in your PATH, or use the full path to air executable

**Issue**: Application not restarting

- **Solution**: Check that you're modifying files in watched directories (`cmd`, `internal`, `pkg`)

**Issue**: Build errors not showing

- **Solution**: Check the `build-errors.log` file in the project root

## Configuration Adjustments

To modify Air's behavior, edit `.air.toml`:

- **Change rebuild delay**: Modify `delay` under `[build]` section (in milliseconds)
- **Add more watch directories**: Add to `include_dir` array
- **Exclude more files**: Add to `exclude_regex` or `exclude_dir` arrays
- **Change output binary location**: Modify `bin` and `cmd` under `[build]`

## File Structure

- `.air.toml` - Air configuration file
- `tmp/main` - Compiled binary (auto-generated, in .gitignore)
- `build-errors.log` - Build error logs (in .gitignore)

## Tips for Best Experience

1. **Monitor the logs**: Keep an eye on the console for build status
2. **Test incrementally**: Small changes trigger faster rebuilds
3. **Save frequently**: Air responds to every save
4. **Check dependencies**: If adding new imports, they will be resolved on rebuild

Enjoy fast development cycles with Air! 🚀
