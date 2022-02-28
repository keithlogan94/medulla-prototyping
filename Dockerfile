#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY ["DatabaseControllerKubeOps/DatabaseControllerKubeOps.csproj", "DatabaseControllerKubeOps/"]
RUN dotnet restore "DatabaseControllerKubeOps/DatabaseControllerKubeOps.csproj"
COPY . .
WORKDIR "/src/DatabaseControllerKubeOps"
RUN dotnet build "DatabaseControllerKubeOps.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "DatabaseControllerKubeOps.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "DatabaseControllerKubeOps.dll"]