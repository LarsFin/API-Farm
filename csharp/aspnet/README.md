# C# > Asp&#46;Net Core

ASP&#46;Net Core is a development framework for building web applications. With Visual Studio, it is possible to generate a functioning Controller at the initation of a project.
Much of the challenge when developing the API Farm use case, came in customising the middleware. The development of the `VideoGameController` and its service dependencies was much more straightforward.

## Dependencies

&#46;Net Core uses Nuget to manage project dependencies. Nuget is included with Visual Studio, so importing dependencies is quick and easy.
The structure of a &#46;Net Core code base is made up of projects which are referenced by a solution. The solution is what is opened with the Visual Studio IDE.
Each project has its own dependencies which are listed within `csproj` files. These are the files which define the projects themselves.
Additionally, a project can have a dependency on another project. For example; the test project `ApiFarm.Test` has a dependency on `ApiFarm`.
This means the unit tests can reference any dependencies of the source code.
However, it is not possible to have a cyclic dependency; `ApiFarm` could not have a dependency on `ApiFarm.Test`.

Note; each dependency has target &#46;Net frameworks. So, when installing dependencies in your local environment you may have to amend your &#46;Net Core verion (alternatively, use Docker).

## Running

You can choose to run the `ApiFarm` project either through the Visual Studio IDE or alternatively via the dotnet executable via the command line.
Below, is the command to run the project via dotnet;

```shell
dotnet run --project ApiFarm
```

### Configuration

ASP&#46;Net Core uses the environment variable `ASPNETCORE_ENVIRONMENT` to determine which environment to run under.
Within the `ApiFarm` project are `appsettings.<env>.json` files with custom settings for running the `ApiFarm` project.
If the environment is invalid; the default `appsettings.json` configuration will be used.

## Testing

You can run the unit test suite from Visual Studio using the Test Explorer.
However, for coverage data you should use the dotnet executable with some custom parameters to ensure coverage reporting.

```shell
dotnet test /p:Exclude=\"[*]ApiFarm.Scaffolding.*,[*]ApiFarm.Models.*,[*]ApiFarm.Utils.*\" /p:CollectCoverage=true /p:Threshold=100
```

The tool used for coverage assessment is Coverlet. The parameters above, are various instructions for Coverlet.

1. **/p:Exclude** - Excludes the specified assembly namespaces from coverage collection. Without excluding `Scaffolding`; code found within
`Startup.cs` and `Program.cs` files would be considered necessary for testing. The `Models` namespace contains model classes with getters and
setters, these are considered testworthy by Coverlet; so model files have been excluded from coverage. `Utils` contains wrapper function and 
other utilities for writing testable code so has been excluded. However, some tests have been written for utility classes which bear logic.
2. **/p:CollectCoverage=true** - Instructs Coverlet to build a coverage report to `ApiFarm.Test/coverage.json`.
3. **/p:Threshold=100** - The threshold required to achieve a 0 exit code after tests have run. It is also possible to provide separate thresholds 
for different coverage metrics like 'lines', 'branches' or 'methods'.

## Linting

By default, Intellisense in Visual Studio will provide warnings as part of builds where obvious code smells are present. StyleCop has been used to 
extend scrutiny to ensure a code standard is conformed to across the solution. StyleCop will append warnings to the Error List in Visual Studio
after performing a build. The rules used to lint the code base have been defined within `StyleCop.Analyzers.ruleset` files. There is a file for both
`ApiFarm` and `ApiFarm.Test`, this is because some standard concerns didn't seem fitting for unit test code. The ruleset files themselves are 
enforced by being referenced as a `<PropertyGroup>` in each `.csproj` file.

To run the linter you can simply perform a build in visual studio. Alternatively, you can build via the dotnet executable;

```shell
dotnet build
```

As part of the CI pipeline we append the switch `/warnaserror` so warnings caught from linting result in a non zero exit code.

## Docker

This API is supported with Docker. You can check out how to run it by following the instructions in the root README [here](https://github.com/LarsFin/API-Farm).

## Resource Documentation

- **ASP&#46;Net Core:** https://docs.microsoft.com/en-us/aspnet/core/?view=aspnetcore-5.0
- **Coverlet:** https://github.com/coverlet-coverage/coverlet/blob/master/README.md
- **Moq:** https://documentation.help/Moq/
- **NUnit:** https://docs.nunit.org/articles/nunit/intro.html
- **Shouldly:** https://github.com/shouldly/shouldly
- **StyleCop:** https://documentation.help/StyleCop/