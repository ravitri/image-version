# image-version

This repository is to hose files for a basic script/program to fetch a version from a Openshift's release image.

The `image-version.sh` is the simplest way to fetch the version of an image by referring to the `io.openshift.release` label in the image. An equivalent Golang program is written in `image-version.go`.

## Usage

### Running shell script
```bash
bash image-version.sh ${IMAGE_DIGEST}
```

### Running golang program
```bash
go run image-version.go
```

The script/program is very basic and definitely can be improved further as per best practices. 
