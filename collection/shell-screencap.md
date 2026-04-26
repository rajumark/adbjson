# shell screencap

Capture device screenshot.

## Command
```bash
adbjson shell screencap [options]
```

## Description
Executes `adb shell screencap` and outputs the result as structured JSON. Captures the device screen and returns image data or file information.

## Sample Output
```json
{
  "capture_time": "2024-01-15T10:30:45Z",
  "width": 1080,
  "height": 2400,
  "density": 420,
  "format": "png",
  "size_bytes": 245760,
  "file_path": "/storage/emulated/0/Pictures/screenshot.png",
  "base64_data": "iVBORw0KGgoAAAANSUhEUgAA..."
}
```

## Options
- `-p` - Output PNG data to stdout
- `file_path` - Save to specific file path

## Examples
```bash
# Capture screenshot and get base64 data
adbjson shell screencap -p

# Save screenshot to file
adbjson shell screencap /storage/emulated/0/Pictures/screenshot.png

# Capture with timestamp filename
adbjson shell screencap /storage/emulated/0/Pictures/screenshot_$(date +%Y%m%d_%H%M%S).png
```

## Output Formats
- **Base64** - When using `-p` flag, returns base64 encoded image data
- **File Info** - When saving to file, returns file information and metadata

## Screen Information
The JSON output includes:
- `width` - Screen width in pixels
- `height` - Screen height in pixels  
- `density` - Screen density (DPI)
- `format` - Image format (PNG)
- `size_bytes` - File size in bytes
- `capture_time` - When screenshot was taken

## Flags
- `--pretty` - Pretty print JSON (default: true)
- `--compact` - Compact JSON output
- `--base64` - Return base64 encoded image data
- `--metadata-only` - Return only metadata, no image data
