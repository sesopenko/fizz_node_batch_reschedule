# sesopenko/fizz_node_batch_reschedule

Reschedules a batch of keyframe directives used by [ComfyUI FizzNodes](https://github.com/FizzleDorf/ComfyUI_FizzNodes).

When skipping frames in an animation, the key frames have to be rescheduled based on the number of key frames skipped.

## Requirements

* [go 1.21](https://go.dev/doc/install)

## Usage (Windows):

From windows terminal:

```bat
REM build project
go build
REM copy all example files
prep.bat
REM run program and skip 120 frames
batch_reschedule 120
```

### Linux support

Theoretically this should work on linux if you build it with similar steps. See what [prep.bat](prep.bat) is doing.

## License

Licenced under Apache version 2.0. The license is included in [LICENSE-2.0.txt](LICENSE-2.0.txt)

## Copyright

Copyright (c) Sean Esopenko 2023