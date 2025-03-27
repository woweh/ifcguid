# ifcguid

`ifcguid` is a Go package that provides functionality for creating [IFC GUIDs](https://technical.buildingsmart.org/resources/ifcimplementationguidance/ifc-guid/),
and converting standard [UUIDs](https://en.wikipedia.org/wiki/Universally_unique_identifier) and various CAD element identifiers to IFC GUIDs and back.  
The process of converting between IFC GUIDs and standard UUIDs is sometimes also referred to as 'compressing' and 'expanding'.

[![Go Version](https://img.shields.io/github/go-mod/go-version/woweh/ifcguid)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/github/license/woweh/ifcguid)](https://github.com/woweh/ifcguid/blob/main/LICENSE)


## Table of Contents
- [Features](#features)
  - [What are IFC GUIDs?](#what-are-ifc-guids)
  - [Why convert IFC GUIDs and CAD identifiers?](#why-convert-ifc-guids-and-cad-identifiers)
  - [Microsoft GUIDs vs standard UUIDs](#microsoft-guids-vs-standard-uuids)
- [Requirements](#requirements)
- [Installation](#installation)
- [Example](#example)
- [Information about calculating ifcGUIDs](#information-about-calculating-ifcguids)
- [References and Acknowledgements](#references-and-acknowledgements)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)
- [Version](#version)


## Features
- Create new random IFC GUIDs
- Validate IFC GUIDs
- Convert between IFC GUIDs and UUIDs (aka compression/decompression)
- Validate Revit UniqueIDs
- Convert Revit UniqueIDs to IFC GUIDs
- Convert AutoCAD handles to and from IFC GUIDs
- Convert 32-bit and 64-bit integers to and from IFC GUIDs
- Convert string representations of integers to and from IFC GUIDs
- Convert arbitrary strings to and from IFC GUIDs

At the moment, this package only supports base64 encoding.

### What are IFC GUIDs?
[IFC GUIDs](https://technical.buildingsmart.org/resources/ifcimplementationguidance/ifc-guid/) (Industry Foundation Classes Globally Unique Identifiers, aka GlobalIds)
are 22-character identifiers used in Building Information Modeling (BIM) to uniquely identify elements across different software platforms and throughout a building's lifecycle.  
IFC GUIDs are effectively compressed UUIDs, developed when storage was limited.  
This package facilitates working with these identifiers in Go applications.  
While technically speaking IFC GUIDs should be universally unique, in practice you can find IFC GUIDs that are only unique _per CAD document or project_.
 
### Why convert IFC GUIDs and CAD identifiers?
Sometimes you don't have the option to persist IFC GUIDs, or keep a table with the original CAD element identifier and the associated IFC GUID.  
In those cases it's useful if the IFC GUID is based on the original CAD element identifier, so that you can get the CAD element identifier back from the IFC GUID.

### Microsoft GUIDs vs standard UUIDs
It appears that most conversions between IFC GUIDs and UUIDs (GUIDs) were initially done on Microsoft Operating Systems.  
But there is a difference between Microsoft GUIDs and standard UUIDs.   
Microsoft GUIDs are mixed endian, while standard UUIDs are sequentially encoded in big-endian.  
See [Wikipedia Universally unique identifier > Endianness](https://en.wikipedia.org/wiki/Universally_unique_identifier#Endianness).  
This package produces IFC GUIDs that are compatible with Microsoft GUIDs.  
When converting an IFC GUID to a UUID using this package, the resulting UUID will match Microsoft GUIDs.    
This approach ensures compatibility with existing CAD and BIM software.


## Requirements
This package is developed and tested with Go 1.22.3.  
It may work with earlier versions, but compatibility is not guaranteed.


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