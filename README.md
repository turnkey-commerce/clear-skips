# clear-skips

This is a simple Go app to remove skip statements from Exercism C# tests. It is useful when mentoring C# exercises.

Skipped tests are defined in the xUnit framework as a parameter on the Fact attribute on the test:

`[Fact(Skip = "Remove to run test")]`

After running clear-skips it will remove the parameter for all of the test attributes, changing them to:

`[Fact]`

## Usage

`clear-skips FilenameTest.cs`
