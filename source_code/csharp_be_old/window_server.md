.NET Framework 4.8 and IIS -> run
nuget -> download thư viện
msbuild -> build source

## Install nuget and msbuild

### Download nuget

mkdir C:\tools
https://www.nuget.org/downloads/ -> C:\tools

### Download MSBuild

https://visualstudio.microsoft.com/downloads/
Tool for Visual Studio
Build Tools for Visual Studio 2026
✓ Web development build tools
✓ .NET Framework 4.8 development tools
✓ NuGet targets and build tasks

### Cài Path Nuget & MSBuild

1. Tìm `Edit the system environment variables`.
2. Chọn `Environment Variables...`
3. Trong `System variables`, chọn `Path`.
4. Nhấn `Edit` → New.
5. Thêm: `C:\tools`
6. Thêm: `C:\Program Files (x86)\Microsoft Visual Studio\18\BuildTools\MSBuild\Current\Bin`
6. Nhấn `OK` hết các cửa sổ.
7. Mở terminal mới và chạy: 
```bash
nuget help
msbuild -version
```

### Cài Path MSBuild

Open Developer PowerShell, then run:

```bash
# tải thư viện
nuget restore CSharpOldWebApi.sln -PackagesDirectory packages

# build
msbuild CSharpOldWebApi.sln /p:Configuration=Release
```

## Add web IIS

Xoá Default web nếu chưa thêm hostname

