package ifcguid_test

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/woweh/ifcguid"
)

func Example() {
	fmt.Println("ifcguid version:", ifcguid.Version)

	// Example UUID
	id := uuid.MustParse("01cf62c8-e9bc-bf88-0000-000000000005")

	// Convert UUID to IFC GUID
	ifcGuid, _ := ifcguid.FromUuid(id)
	fmt.Println("UUID to IFC GUID:", id, "->", ifcGuid)

	// Convert IFC GUID back to UUID
	backToUuid, _ := ifcguid.ToUuid(ifcGuid)
	fmt.Println("IFC GUID to UUID:", ifcGuid, "->", backToUuid)

	// Convert Revit Unique ID to IFC GUID
	revitUniqueId := "8d814f39-b6ea-4766-9a4f-8ac3de3501b2-00007c0e"
	revitGuid, _ := ifcguid.FromRevitUniqueId(revitUniqueId)
	fmt.Println("Revit Unique ID to IFC GUID:", revitUniqueId, "->", revitGuid)

	// Convert AutoCAD handle to IFC GUID
	acadHandle := "1A"
	acadGuid, _ := ifcguid.FromAutoCadHandle(acadHandle)
	fmt.Println("AutoCAD handle to IFC GUID:", acadHandle, "->", acadGuid)

	// Convert integer to IFC GUID
	intValue := int64(123456789)
	intGuid, _ := ifcguid.FromInt64(intValue)
	fmt.Println("Int64 to IFC GUID:", intValue, "->", intGuid)

	// Output:
	// ifcguid version: 1.0.0
	// UUID to IFC GUID: 01cf62c8-e9bc-bf88-0000-000000000005 -> 01psB8wRo$Y00000000005
	// IFC GUID to UUID: 01psB8wRo$Y00000000005 -> 01cf62c8-e9bc-bf88-0000-000000000005
	// Revit Unique ID to IFC GUID: 8d814f39-b6ea-4766-9a4f-8ac3de3501b2-00007c0e -> 2DWKyvjkf7PffFYiFUDNsy
	// AutoCAD handle to IFC GUID: 1A -> 000000000000000000000Q
	// Int64 to IFC GUID: 123456789 -> 000000000000000007MyqL
}
