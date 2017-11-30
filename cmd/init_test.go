package cmd

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestInit(t *testing.T) {
	buf := &bytes.Buffer{}
	f := &Flags{
		Logger: log.New(buf, "", 0),
		BuildN: true,
	}

	if err := Init(f); err != nil {
		t.Fatal(err)
	}

	if bytes.Compare([]byte(expectedInit), buf.Bytes()) != 0 {
		fmt.Println(expectedInit)
		fmt.Println(string(buf.Bytes()))
		t.Fatal("Output doesn't match")
	}
}

func TestBuild(t *testing.T) {
	buf := &bytes.Buffer{}
	f := &Flags{
		Threaded: false,
		Logger:   log.New(buf, "", 0),
		BuildN:   true,
	}

	if err := Build(f, []string{"gomatcha.io/matcha/examples"}); err != nil {
		t.Fatal(err)
	}

	if bytes.Compare([]byte(expectedBuild), buf.Bytes()) != 0 {
		fmt.Println(expectedBuild)
		fmt.Println(string(buf.Bytes()))
		t.Fatal("Output doesn't match")
	}
}

const expectedInit = `which go
go version
printenv GOPATH
test -d $GOPATH/pkg/matcha
rm -r -f $GOPATH/pkg/matcha
mkdir -p $GOPATH/pkg/matcha
WORK=$WORK
which xcrun
xcrun --sdk iphoneos --find clang
xcrun --sdk iphoneos --show-sdk-path
GOOS=darwin GOARCH=arm GOARM=7 CC=$CLANG_IPHONEOS CXX=$CLANG_IPHONEOS CGO_CFLAGS=-isysroot $SDK_IPHONEOS -miphoneos-version-min=6.1 -arch armv7 CGO_LDFLAGS=-isysroot $SDK_IPHONEOS -miphoneos-version-min=6.1 -arch armv7 CGO_ENABLED=1 go install -pkgdir=$GOPATH/pkg/matcha/pkg_darwin_arm std
xcrun --sdk iphoneos --find clang
xcrun --sdk iphoneos --show-sdk-path
GOOS=darwin GOARCH=arm64 CC=$CLANG_IPHONEOS CXX=$CLANG_IPHONEOS CGO_CFLAGS=-isysroot $SDK_IPHONEOS -miphoneos-version-min=6.1 -arch arm64 CGO_LDFLAGS=-isysroot $SDK_IPHONEOS -miphoneos-version-min=6.1 -arch arm64 CGO_ENABLED=1 go install -pkgdir=$GOPATH/pkg/matcha/pkg_darwin_arm64 std
xcrun --sdk iphonesimulator --find clang
xcrun --sdk iphonesimulator --show-sdk-path
GOOS=darwin GOARCH=amd64 CC=$CLANG_IPHONESIMULATOR CXX=$CLANG_IPHONESIMULATOR CGO_CFLAGS=-isysroot $SDK_IPHONESIMULATOR -mios-simulator-version-min=6.1 -arch x86_64 CGO_LDFLAGS=-isysroot $SDK_IPHONESIMULATOR -mios-simulator-version-min=6.1 -arch x86_64 CGO_ENABLED=1 go install -tags=ios -pkgdir=$GOPATH/pkg/matcha/pkg_darwin_amd64 std
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/platforms
ls $ANDROID_HOME/platforms
test -f $ANDROID_HOME/platforms/android-21/android.jar
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/ndk-bundle
which javac
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/ndk-bundle
GOOS=android GOARCH=arm CC=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang CXX=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang++ CGO_CFLAGS=-target armv7a-none-linux-androideabi -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/arm-linux-androideabi-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/arm-linux-androideabi -D__ANDROID_API__=15 CGO_CPPFLAGS=-target armv7a-none-linux-androideabi -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/arm-linux-androideabi-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/arm-linux-androideabi -D__ANDROID_API__=15 CGO_LDFLAGS=-target armv7a-none-linux-androideabi -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/arm-linux-androideabi-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/platforms/android-15/arch-arm CGO_ENABLED=1 GOARM=7 go install -pkgdir=$GOPATH/pkg/matcha/pkg_android_arm std
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/ndk-bundle
GOOS=android GOARCH=arm64 CC=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang CXX=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang++ CGO_CFLAGS=-target aarch64-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/aarch64-linux-android-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/aarch64-linux-android -D__ANDROID_API__=21 CGO_CPPFLAGS=-target aarch64-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/aarch64-linux-android-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/aarch64-linux-android -D__ANDROID_API__=21 CGO_LDFLAGS=-target aarch64-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/aarch64-linux-android-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/platforms/android-21/arch-arm64 CGO_ENABLED=1 go install -pkgdir=$GOPATH/pkg/matcha/pkg_android_arm64 std
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/ndk-bundle
GOOS=android GOARCH=386 CC=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang CXX=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang++ CGO_CFLAGS=-target i686-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/x86-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/i686-linux-android -D__ANDROID_API__=15 CGO_CPPFLAGS=-target i686-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/x86-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/i686-linux-android -D__ANDROID_API__=15 CGO_LDFLAGS=-target i686-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/x86-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/platforms/android-15/arch-x86 CGO_ENABLED=1 go install -pkgdir=$GOPATH/pkg/matcha/pkg_android_386 std
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/ndk-bundle
GOOS=android GOARCH=amd64 CC=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang CXX=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang++ CGO_CFLAGS=-target x86_64-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/x86_64-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/x86_64-linux-android -D__ANDROID_API__=21 CGO_CPPFLAGS=-target x86_64-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/x86_64-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/x86_64-linux-android -D__ANDROID_API__=21 CGO_LDFLAGS=-target x86_64-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/x86_64-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/platforms/android-21/arch-x86_64 CGO_ENABLED=1 go install -pkgdir=$GOPATH/pkg/matcha/pkg_android_amd64 std
go version
write $GOPATH/pkg/matcha/version
Matcha initialized.
rm -r -f $WORK
`

const expectedBuild = `go findpackage gomatcha.io/matcha
which go
go version
WORK=$WORK
printenv GOPATH
test -d $GOPATH/pkg/matcha
read $GOPATH/pkg/matcha/version
go version
pwd
go importall $CWD gomatcha.io/matcha/examples
go findpackage gomatcha.io/matcha/bridge
which xcrun
mkdir -p $WORK/matcha-ios
mkdir -p $WORK/matcha-ios/MatchaBridge/MatchaBridge
write $WORK/src/iosbin/main.go
xcrun --sdk iphoneos --find clang
xcrun --sdk iphoneos --show-sdk-path
xcrun --sdk iphoneos --find clang
xcrun --sdk iphoneos --show-sdk-path
xcrun --sdk iphonesimulator --find clang
xcrun --sdk iphonesimulator --show-sdk-path
printenv GOPATH
test -d $GOPATH/pkg/matcha/pkg_darwin_arm
GOOS=darwin GOARCH=arm GOARM=7 CC=$CLANG_IPHONEOS CXX=$CLANG_IPHONEOS CGO_CFLAGS=-isysroot $SDK_IPHONEOS -miphoneos-version-min=6.1 -arch armv7 CGO_LDFLAGS=-isysroot $SDK_IPHONEOS -miphoneos-version-min=6.1 -arch armv7 CGO_ENABLED=1 GOPATH=$WORK/IOS-GOPATH:$GOPATH go build -pkgdir=$GOPATH/pkg/matcha/pkg_darwin_arm -tags matcha ios -buildmode=c-archive -o $WORK/matcha-arm.a $WORK/src/iosbin/main.go
printenv GOPATH
test -d $GOPATH/pkg/matcha/pkg_darwin_arm64
GOOS=darwin GOARCH=arm64 CC=$CLANG_IPHONEOS CXX=$CLANG_IPHONEOS CGO_CFLAGS=-isysroot $SDK_IPHONEOS -miphoneos-version-min=6.1 -arch arm64 CGO_LDFLAGS=-isysroot $SDK_IPHONEOS -miphoneos-version-min=6.1 -arch arm64 CGO_ENABLED=1 GOPATH=$WORK/IOS-GOPATH:$GOPATH go build -pkgdir=$GOPATH/pkg/matcha/pkg_darwin_arm64 -tags matcha ios -buildmode=c-archive -o $WORK/matcha-arm64.a $WORK/src/iosbin/main.go
printenv GOPATH
test -d $GOPATH/pkg/matcha/pkg_darwin_amd64
GOOS=darwin GOARCH=amd64 CC=$CLANG_IPHONESIMULATOR CXX=$CLANG_IPHONESIMULATOR CGO_CFLAGS=-isysroot $SDK_IPHONESIMULATOR -mios-simulator-version-min=6.1 -arch x86_64 CGO_LDFLAGS=-isysroot $SDK_IPHONESIMULATOR -mios-simulator-version-min=6.1 -arch x86_64 CGO_ENABLED=1 GOPATH=$WORK/IOS-GOPATH:$GOPATH go build -pkgdir=$GOPATH/pkg/matcha/pkg_darwin_amd64 -tags matcha ios -buildmode=c-archive -o $WORK/matcha-amd64.a $WORK/src/iosbin/main.go
xcrun lipo -create -arch armv7 $WORK/matcha-arm.a -arch arm64 $WORK/matcha-arm64.a -arch x86_64 $WORK/matcha-amd64.a -o $WORK/matcha-ios/MatchaBridge/MatchaBridge/MatchaBridge.a
cp $WORK/matcha-ios/MatchaBridge/MatchaBridge/MatchaBridge.a $GOPATH/src/gomatcha.io/matcha/ios/MatchaBridge/MatchaBridge/MatchaBridge.a
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/platforms
ls $ANDROID_HOME/platforms
test -f $ANDROID_HOME/platforms/android-21/android.jar
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/ndk-bundle
which javac
write $WORK/androidlib/main.go
mkdir -p $WORK/android/src/main/java/io/gomatcha/bridge
cp $GOPATH/src/gomatcha.io/matcha/bridge/java-GoValue.java $WORK/android/src/main/java/io/gomatcha/bridge/GoValue.java
cp $GOPATH/src/gomatcha.io/matcha/bridge/java-Bridge.java $WORK/android/src/main/java/io/gomatcha/bridge/Bridge.java
cp $GOPATH/src/gomatcha.io/matcha/bridge/java-Tracker.java $WORK/android/src/main/java/io/gomatcha/bridge/Tracker.java
mkdir -p $WORK/matcha-android
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/ndk-bundle
printenv GOPATH
test -d $GOPATH/pkg/matcha/pkg_android_arm
GOOS=android GOARCH=arm CC=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang CXX=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang++ CGO_CFLAGS=-target armv7a-none-linux-androideabi -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/arm-linux-androideabi-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/arm-linux-androideabi -D__ANDROID_API__=15 CGO_CPPFLAGS=-target armv7a-none-linux-androideabi -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/arm-linux-androideabi-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/arm-linux-androideabi -D__ANDROID_API__=15 CGO_LDFLAGS=-target armv7a-none-linux-androideabi -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/arm-linux-androideabi-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/platforms/android-15/arch-arm CGO_ENABLED=1 GOARM=7 GOPATH=$WORK/ANDROID-GOPATH:$GOPATH go build -pkgdir=$GOPATH/pkg/matcha/pkg_android_arm -tags matcha -buildmode=c-shared -o=$WORK/android/src/main/jniLibs/armeabi-v7a/libgojni.so $WORK/androidlib/main.go
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/ndk-bundle
printenv GOPATH
test -d $GOPATH/pkg/matcha/pkg_android_arm64
GOOS=android GOARCH=arm64 CC=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang CXX=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang++ CGO_CFLAGS=-target aarch64-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/aarch64-linux-android-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/aarch64-linux-android -D__ANDROID_API__=21 CGO_CPPFLAGS=-target aarch64-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/aarch64-linux-android-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/aarch64-linux-android -D__ANDROID_API__=21 CGO_LDFLAGS=-target aarch64-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/aarch64-linux-android-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/platforms/android-21/arch-arm64 CGO_ENABLED=1 GOPATH=$WORK/ANDROID-GOPATH:$GOPATH go build -pkgdir=$GOPATH/pkg/matcha/pkg_android_arm64 -tags matcha -buildmode=c-shared -o=$WORK/android/src/main/jniLibs/arm64-v8a/libgojni.so $WORK/androidlib/main.go
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/ndk-bundle
printenv GOPATH
test -d $GOPATH/pkg/matcha/pkg_android_386
GOOS=android GOARCH=386 CC=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang CXX=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang++ CGO_CFLAGS=-target i686-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/x86-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/i686-linux-android -D__ANDROID_API__=15 CGO_CPPFLAGS=-target i686-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/x86-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/i686-linux-android -D__ANDROID_API__=15 CGO_LDFLAGS=-target i686-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/x86-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/platforms/android-15/arch-x86 CGO_ENABLED=1 GOPATH=$WORK/ANDROID-GOPATH:$GOPATH go build -pkgdir=$GOPATH/pkg/matcha/pkg_android_386 -tags matcha -buildmode=c-shared -o=$WORK/android/src/main/jniLibs/x86/libgojni.so $WORK/androidlib/main.go
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/ndk-bundle
printenv GOPATH
test -d $GOPATH/pkg/matcha/pkg_android_amd64
GOOS=android GOARCH=amd64 CC=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang CXX=$ANDROID_HOME/ndk-bundle/toolchains/llvm/prebuilt/darwin-x86_64/bin/clang++ CGO_CFLAGS=-target x86_64-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/x86_64-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/x86_64-linux-android -D__ANDROID_API__=21 CGO_CPPFLAGS=-target x86_64-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/x86_64-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/sysroot -isystem $ANDROID_HOME/ndk-bundle/sysroot/usr/include/x86_64-linux-android -D__ANDROID_API__=21 CGO_LDFLAGS=-target x86_64-none-linux-android -gcc-toolchain $ANDROID_HOME/ndk-bundle/toolchains/x86_64-4.9/prebuilt/darwin-x86_64 --sysroot $ANDROID_HOME/ndk-bundle/platforms/android-21/arch-x86_64 CGO_ENABLED=1 GOPATH=$WORK/ANDROID-GOPATH:$GOPATH go build -pkgdir=$GOPATH/pkg/matcha/pkg_android_amd64 -tags matcha -buildmode=c-shared -o=$WORK/android/src/main/jniLibs/x86_64/libgojni.so $WORK/androidlib/main.go
mkdir -p $WORK/javac-output
printenv ANDROID_HOME
test -d $ANDROID_HOME
test -d $ANDROID_HOME/platforms
ls $ANDROID_HOME/platforms
test -f $ANDROID_HOME/platforms/android-21/android.jar
PWD=$WORK/android/src/main/java javac -d $WORK/javac-output -source 1.7 -target 1.7 -bootclasspath $ANDROID_HOME/platforms/android-21/android.jar *.java
write $WORK/matcha-android/classes.jar
cp $WORK/matcha-android/classes.jar $GOPATH/src/gomatcha.io/matcha/android/Matcha/Matcha/libs/classes.jar
rm -r -f $GOPATH/src/gomatcha.io/matcha/android/Matcha/Matcha/jniLibs
cp $WORK/android/src/main/jniLibs/armeabi-v7a/libgojni.so $GOPATH/src/gomatcha.io/matcha/android/Matcha/Matcha/jniLibs/armeabi-v7a/libgojni.so
cp $WORK/android/src/main/jniLibs/arm64-v8a/libgojni.so $GOPATH/src/gomatcha.io/matcha/android/Matcha/Matcha/jniLibs/arm64-v8a/libgojni.so
cp $WORK/android/src/main/jniLibs/x86/libgojni.so $GOPATH/src/gomatcha.io/matcha/android/Matcha/Matcha/jniLibs/x86/libgojni.so
cp $WORK/android/src/main/jniLibs/x86_64/libgojni.so $GOPATH/src/gomatcha.io/matcha/android/Matcha/Matcha/jniLibs/x86_64/libgojni.so
rm -r -f $WORK
`
