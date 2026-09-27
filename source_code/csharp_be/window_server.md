ASP.NET Core (.NET 10 Hosting Bundle) + IIS -> run
.NET 10 SDK -> tải thư viện và build



# Download and Install
https://dotnet.microsoft.com/en-us/download/dotnet/10.0

ASP.NET Core (.NET 10 Hosting Bundle)
-> Bên phải, phần ASP.NET Core Runtime 10.0.xx → Windows → Hosting Bundle

.NET 10 SDK 
-> Bên trái  -> Windows -> x64

# reset IIS
```bash
iisreset
```

# đóng cửa sổ Terminal hiện tại, mở Terminal mới rồi kiểm tra
```bash
dotnet --version
dotnet --list-runtimes

# 10.0.xx
# Microsoft.AspNetCore.App 10.0.xx
# Microsoft.NETCore.App 10.0.xx
```

# Tải và build source
```bash
# restore → build → publish
dotnet publish -c Release -o .\publish
```