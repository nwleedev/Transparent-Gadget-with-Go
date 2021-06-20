# The Simple Windows Gadget with Transparent Background

## Used Language and Module

- Go Programming Language
- Walk, Go GUI Module

## How to Use

1. Store the image and icons.

- Icons are divided into two extensions.
- PNG and ICO

2. Specify Exact Image Name in main.go

- Image
- Icon
- Window Name

3. Run These Commands

```
go get github.com/akavel/rsrc
rsrc -manifest main.exe.manifest -ico="./assets/YOUR_ICON_NAME.ico" -o rsrc.syso
```

4. Compile Go File

```
go build -o YOUR_APP_NAME.exe -ldflags="-H windowsgui"
- OR -
go run .
```

5. Now you can execute the exe file generated.

## The Screen Shot Example

You can see the result image in result folder.

**The cat is carring multiple icons. 🤣**
