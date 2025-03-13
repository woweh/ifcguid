# ifcguid

`ifcguid` is a Go package that provides functionality for converting between various identifier formats used in CAD and BIM systems, including IFC GUIDs, UUIDs, Revit Unique IDs, AutoCAD handles, and integer representations.

[![Go Version](https://img.shields.io/github/go-mod/go-version/woweh/ifcguid)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/github/license/woweh/ifcguid)](https://github.com/woweh/ifcguid/blob/main/LICENSE)

## Table of Contents
- [Features](#features)
- [What are IFC GUIDs?](#what-are-ifc-guids)
- [Requirements](#requirements)
- [Installation](#installation)
- [Example](#example)
- [Information about calculating ifcGUIDs](#information-about-calculating-ifcguids)
- [References and Acknowledgements](#references-and-acknowledgements)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Features
- Convert between IFC GUIDs and UUIDs
- Convert Revit Unique IDs to IFC GUIDs
- Convert AutoCAD handles to and from IFC GUIDs
- Convert 32-bit and 64-bit integers to and from IFC GUIDs
- Convert string representations of integers to and from IFC GUIDs
- Validate Revit Unique IDs

## What are IFC GUIDs?
IFC GUIDs (Industry Foundation Classes Globally Unique Identifiers) are special 22-character identifiers used in Building Information Modeling (BIM) to uniquely identify elements across different software platforms and throughout a building's lifecycle.
This package facilitates working with these identifiers in Go applications.

## Requirements
This package is developed and tested with Go 1.22.3. It may work with earlier versions, but compatibility is not guaranteed.

## Installation
To install the `ifcguid` package, use the following command:
```shell
go get github.com/woweh/ifcguid
```

## Example
An example of how to use this package is provided in the `example_test.go` file in the root of the repository.
You can run this example using:
```shell
go test -v -run=Example
```

## Information about calculating ifcGUIDs
The calculation of ifcGUIDs is based on the following resources:
- https://technical.buildingsmart.org/resources/ifcimplementationguidance/ifc-guid/
- https://forums.buildingsmart.org/t/ifcgloballyuniqueids-spec-description-is-incorrect-proposal-to-simplify/1083/7
- https://github.com/buildingSMART/NextGen-IFC/issues/8

## References and Acknowledgements
This package was developed with reference to and inspiration from several existing implementations across different programming languages.
We acknowledge the following projects that have contributed to our understanding and approach:
- [IfcOpenShell's Python GUID implementation](https://github.com/IfcOpenShell/IfcOpenShell/blob/master/src/ifcopenshell-python/ifcopenshell/guid.py#L38)
- [XbimTeam's C# implementation in XbimEssentials](https://github.com/xBimTeam/XbimEssentials/blob/f9562fc2bdd6f34ec667de70a3e4d19daa6986ef/Xbim.Ifc2x3/UtilityResource/IfcGloballyUniqueIdPartial.cs)
- [Håkon Clausen's C# IfcGuid implementation](https://github.com/hakonhc/IfcGuid/blob/master/IfcGuid/IfcGuid.cs)
- [Jonathon Broughton's JavaScript implementation](https://github.com/jsdbroughton/ifc-guid/blob/master/Guid.js)
- [Devon Sparks' ifcidc project](https://github.com/devonsparks/ifcidc)

These projects have been valuable resources in understanding the intricacies of IFC GUID handling across different platforms and languages.

## Testing
To run the tests for this package, use the following command in the root directory of the project:
```shell
go test ./...
```

## Contributing
Contributions to the ifcguid package are welcome. Please feel free to submit issues, fork the repository and send pull requests!

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Version
For the latest version and release notes, please check the [releases page](https://github.com/woweh/ifcguid/releases).