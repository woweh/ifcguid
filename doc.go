// Package ifcguid provides functions for working with IFC GUIDs (Globally Unique Identifiers, aka GlobalIDs).
//
// IFC GUIDs are 22-character strings used in the Industry Foundation Classes (IFC) schema
// for uniquely identifying building elements and other objects in building information models.
// IFC GUIDs are effectively compressed UUIDs, developed when storage was limited.
//
// This package offers a comprehensive set of functions to generate, validate, and convert IFC GUIDs
// to and from various formats, including UUIDs, Revit UniqueIds, AutoCAD handles, and integer representations.
//
// Key features:
//   - Generate new random IFC GUIDs
//   - Validate IFC GUIDs
//   - Convert between IFC GUIDs and UUIDs (aka compression/decompression)
//   - Validate Revit UniqueIDs
//   - Convert Revit UniqueIDs to IFC GUIDs
//   - Convert between IFC GUIDs and AutoCAD handles
//   - Convert between IFC GUIDs and integer representations
//   - Convert arbitrary strings to and from IFC GUIDs
//
// Usage:
//
//	import "github.com/woweh/ifcguid"
//
//	// Generate a new IFC GUID
//	guid, err := ifcguid.New()
//
//	// Convert UUID to IFC GUID
//	uuid := uuid.New()
//	ifcGuid, err := ifcguid.FromUuid(uuid)
//
//	// Validate an IFC GUID
//	err = ifcguid.IsValid(ifcGuid)
//
//	// Convert IFC GUID to UUID
//	uuid, err = ifcguid.ToUuid(ifcGuid)
//
// For more detailed information on each function, refer to the individual function documentation.
package ifcguid
