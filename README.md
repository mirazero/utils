# utils

1. instantclient-sdk-windows.x64-12.1.0.2.0   instantclient-basic-windows.x64-12.1.0.2.0 파일 설치
2. gcc compiler인 tdm64-gcc-10.3.0-2.exe 설치
3. pkg-config_0.26-1_win32, gettext-runtime_0.18.1.1-2_win32, glib_2.28.8-1_win32 를 위 gcc 컴파일러 bin 폴더에 복사
4. oci8.pc 파일 생성하고, D:\pkg-config\bin 에 복사
~~
ora=D:/instantclient_12_1
gcc=D:/gcc/TDM-GCC-64

oralib=${ora}/sdk/lib/msvc
orainclude=${ora}/sdk/include

gcclib=${gcc}/lib
gccinclude=${gcc}/include

glib_genmarshal=glib-genmarshal
gobject_query=gobject-query
glib_mkenums=glib-mkenums

Name: OCI
Description: Oracle database engine
Version: 12.1
Libs: -L${oralib} -L${gcclib} -loci
Libs.private: 
Cflags: -I${orainclude} -I${gccinclude}
~~
5. 사용자 변수
  - path에 D:\instantclient_12_1, D:\pkg-config\bin, D:\gcc\TDM-GCC-64\bin 추가
  - PKG_CONFIG_PATH = D:\pkg-config\bin 추가
6. go get github.com/mattn/go-oci8
7. pkg-config --cflags --libs oci8 해서 오류없나 확인
