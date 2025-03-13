package ifcguid

import (
	"encoding/binary"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

const (
	// _conversionTable is used during the ifcGuid <-> uuid conversions.
	_conversionTable = `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_$`

	// _int64base is the base for formatting int64 to a string during the int64 <-> uuid conversions.
	// This is chosen so that the biggest int64 value can still be formatted as a 16char string.
	_int64base = 17
)

// FromRevitUniqueId converts a Revit 'unique identifier' to an IFC GUID.
//
//	Note that there is no reverse conversion.
//
// For a reverse conversion, you would need the Revit element ID,
// which is part of the Revit element 'unique identifier'.
func FromRevitUniqueId(uniqueId string) (string, error) {
	guid, err := revitUniqueIdToUuid(uniqueId)
	if err != nil {
		return "", err
	}

	return FromUuid(guid)
}

// FromAutoCadHandle converts an AutoCAD handle to an IFC GUID.
func FromAutoCadHandle(handle string) (string, error) {
	guid, err := autoCadHandleToUuid(handle)
	if err != nil {
		return "", err
	}

	return FromUuid(guid)
}

// ToAutoCadHandle converts an IFC GUID to an AutoCAD handle.
func ToAutoCadHandle(ifcGuid string) (string, error) {
	guid, err := ToUuid(ifcGuid)
	if err != nil {
		return "", err
	}

	return uuidToAutoCadHandle(guid)
}

// FromInt32 converts a 32-bit integer to an IFC GUID.
// Revit element IDs are 32-bit integers.
func FromInt32(value int32) (string, error) {
	return FromInt64(int64(value))
}

// ToInt32 converts an IFC GUID to a 32-bit integer.
// Revit element IDs are 32-bit integers.
func ToInt32(ifcGuid string) (int32, error) {
	i64, err := ToInt64(ifcGuid)
	if err != nil {
		return 0, err
	}

	return int32(i64), nil
}

// FromInt64 converts a 64-bit integer to an IFC GUID.
// AutoCAD ObjectIDs are 64-bit integers.
func FromInt64(value int64) (string, error) {
	guid, err := int64ToUuid(value)
	if err != nil {
		return "", err
	}

	return FromUuid(guid)
}

// ToInt64 converts an IFC GUID to a 64-bit integer.
// AutoCAD ObjectIDs are 64-bit integers.
func ToInt64(ifcGuid string) (int64, error) {
	guid, err := ToUuid(ifcGuid)
	if err != nil {
		return 0, err
	}

	return uuidToInt64(guid)
}

// FromIntString converts a string representation of an integer to an IFC GUID.
// You can use this for string representations of AutoCAD ObjectIDs or Revit elementIDs.
func FromIntString(value string) (string, error) {
	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return "", err
	}

	return FromInt64(intVal)
}

// ToIntString converts an IFC GUID to a string representation of an integer.
// You can use this for string representations of AutoCAD ObjectIDs or Revit elementIDs.
func ToIntString(ifcGuid string) (string, error) {
	guid, err := ToUuid(ifcGuid)
	if err != nil {
		return "", err
	}

	return uuidToIntString(guid, "%d")
}

// ToUuid converts an IFC GUID to a UUID.
func ToUuid(ifcGuid string) (uuid.UUID, error) {

	if len(ifcGuid) != 22 {
		return uuid.Nil, fmt.Errorf("the ifcGuid must be 22 characters long")
	}

	lastBase64Num := ifcGuid[0]
	if (lastBase64Num - 48) > 3 {
		return uuid.Nil, fmt.Errorf("illegal GUID '%v' found, it is greater than 128 bits", ifcGuid)
	}

	num := make([]uint32, 6)
	digits := 2
	pos := 0
	for i := 0; i < 6; i++ {
		endPos := pos + digits
		num[i] = base64ToUInt32(ifcGuid[pos:endPos])
		pos += digits
		digits = 4
	}

	data1 := num[0]*16777216 + num[1]                // 16-13. bytes
	data2 := uint16(num[2] / 256)                    // 12-11. bytes
	data3 := uint16((num[2]%256)*256 + num[3]/65536) // 10-09. bytes
	d1Bytes := make([]byte, 4)
	d2Bytes := make([]byte, 2)
	d3Bytes := make([]byte, 2)
	binary.NativeEndian.PutUint32(d1Bytes, data1)
	binary.NativeEndian.PutUint16(d2Bytes, data2)
	binary.NativeEndian.PutUint16(d3Bytes, data3)

	guidBytes := make([]byte, 16)
	// write the int32 bytes in reverse order
	j := 3
	for i := 0; i < len(d1Bytes); i++ {
		guidBytes[j] = d1Bytes[i] // 16-13. bytes
		j--
	}
	// write 2 x int16 bytes in reverse order
	guidBytes[4] = d2Bytes[1]                  //    12. byte
	guidBytes[5] = d2Bytes[0]                  //    11. byte
	guidBytes[6] = d3Bytes[1]                  //    10. byte
	guidBytes[7] = d3Bytes[0]                  //    09. byte
	guidBytes[8] = byte((num[3] / 256) % 256)  //    08. byte
	guidBytes[9] = byte(num[3] % 256)          //    07. byte
	guidBytes[10] = byte(num[4] / 65536)       //    06. byte
	guidBytes[11] = byte((num[4] / 256) % 256) //    05. byte
	guidBytes[12] = byte(num[4] % 256)         //    04. byte
	guidBytes[13] = byte(num[5] / 65536)       //    03. byte
	guidBytes[14] = byte((num[5] / 256) % 256) //    02. byte
	guidBytes[15] = byte(num[5] % 256)         //    01. byte

	return uuid.FromBytes(guidBytes)
}

// FromUuid converts a UUID to an ifcGuid (= a 22 character length base 64 ifc compliant string).
func FromUuid(guid uuid.UUID) (string, error) {

	bytes, err := guid.MarshalBinary()
	if err != nil {
		return "", err
	}

	data1 := binary.NativeEndian.Uint32(reverseBytes(bytes, 0, 3))         // 4byte - int32
	data2 := uint32(binary.NativeEndian.Uint16(reverseBytes(bytes, 4, 5))) // 2byte - int16
	data3 := uint32(binary.NativeEndian.Uint16(reverseBytes(bytes, 6, 7))) // 2byte - int16

	num := make([]uint32, 6)

	num[0] = data1 / 16777216                                                    // 16. byte
	num[1] = data1 % 16777216                                                    // 15-13. bytes
	num[2] = data2*256 + data3/256                                               // 12-10. bytes
	num[3] = data3%256*65536 + uint32(bytes[8])*256 + uint32(bytes[9])           // 09-07. bytes
	num[4] = uint32(bytes[10])*65536 + uint32(bytes[11])*256 + uint32(bytes[12]) // 06-04. bytes
	num[5] = uint32(bytes[13])*65536 + uint32(bytes[14])*256 + uint32(bytes[15]) // 03-01. bytes

	//convert nums to base 64 characters
	digits := 2
	chars := strings.Builder{}
	for i := 0; i < 6; i++ {
		chars.WriteString(uint32ToBase64(num[i], digits))
		digits = 4
	}

	return chars.String(), nil
}

// base64ToUInt32 is a helper function used when converting a base64 string to a UUID.
func base64ToUInt32(base64String string) uint32 {
	var result uint32 = 0
	for _, charValue := range base64String {
		//c := string(charValue)
		//char := fmt.Sprintf("%c", charValue)
		index := uint32(strings.Index(_conversionTable, string(charValue)))
		result = (result * 64) + index
	}
	return result
}

// uint32ToBase64 is a helper function used when converting a UUID to a base64 string.
func uint32ToBase64(number uint32, digits int) string {
	result := make([]byte, digits)
	act := number
	for i := 0; i < digits; i++ {
		result[digits-i-1] = _conversionTable[int(act%64)]
		act = act / 64
	}
	return string(result)
}

// reverseBytes is a helper function used to reverse a byte slice.
func reverseBytes(bytes []byte, from, to int) []byte {
	arrayLength := to - from + 1
	reversedBytes := make([]byte, arrayLength)
	j := 0
	for i := to; i >= from; i-- {
		reversedBytes[j] = bytes[i]
		j++
	}
	return reversedBytes
}

// revitUniqueIdToUuid converts a Revit 'UniqueId' to a UUID.
func revitUniqueIdToUuid(uniqueId string) (uuid.UUID, error) {
	if !IsValidRevitUniqueId(uniqueId) {
		return uuid.Nil, fmt.Errorf(
			"the given string isn't a Revit uniqueId (length=%d != 45): %v", len(uniqueId), uniqueId,
		)
	}

	elementId, err := strconv.ParseInt(uniqueId[37:45], 16, 64)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error parsing Revit uniqueId: %v", uniqueId)
	}

	tempId, err := strconv.ParseInt(uniqueId[28:36], 16, 64)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error parsing Revit uniqueId: %v", uniqueId)
	}

	xorValue := tempId ^ elementId

	tmpGuidString := uniqueId[0:28] + fmt.Sprintf("%08x", xorValue)

	result, err := uuid.Parse(tmpGuidString)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error parsing Revit uniqueId: %v", uniqueId)
	}

	return result, nil
}

// IsValidRevitUniqueId checks if a string is a Revit 'uniqueId'.
func IsValidRevitUniqueId(uniqueId string) bool {
	/*
		The Revit element uniqueId is formatted in groups of 8-4-4-4-12-8 hexadecimal characters.
		It is similar to the standard GUID format, but has 8 additional characters at the end.
		These 8 additional hexadecimal characters are large enough to store 4 bytes or a 32-bit number,
		which is exactly the size of a Revit element id. 8-4-4-4-12-8 => 45 chars
	*/
	if len(uniqueId) != 45 {
		return false
	}

	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}-[a-fA-F0-9]{8}$")

	return r.MatchString(uniqueId)
}

// autoCadHandleToUuid converts an AutoCad handle to a UUID.
func autoCadHandleToUuid(handle string) (uuid.UUID, error) {
	return intStringToUuid(handle, 16)
}

// uuidToAutoCadHandle converts a UUID to an AutoCad handle.
func uuidToAutoCadHandle(guid uuid.UUID) (string, error) {
	return uuidToIntString(guid, "%x")
}

// intStringToUuid converts an integer string to a UUID.
func intStringToUuid(integerString string, base int) (uuid.UUID, error) {
	// convert the ID to an int64
	entityId, err := strconv.ParseInt(integerString, base, 64)
	if err != nil {
		return uuid.Nil, err
	}

	return int64ToUuid(entityId)
}

// uuidToIntString converts a UUID to an integer string using the given format.
func uuidToIntString(guid uuid.UUID, format string) (string, error) {
	number, err := uuidToInt64(guid)
	if err != nil {
		return "", err
	}

	// return the number in decimal format
	return fmt.Sprintf(format, number), nil
}

func int64ToUuid(i64 int64) (uuid.UUID, error) {
	// convert the int64 into a 16 character long string, left padded with zeros
	string16 := fmt.Sprintf("%016v", strconv.FormatInt(i64, _int64base))
	if len(string16) != 16 {
		return uuid.Nil, fmt.Errorf("expected 16 characters, got %d", len(string16))
	}

	// create a guid from the 16 character string (> 16 bytes)
	guid, err := uuid.FromBytes([]byte(string16))
	if err != nil {
		return uuid.Nil, err
	}

	return guid, nil
}

func uuidToInt64(guid uuid.UUID) (int64, error) {
	bytes, err := guid.MarshalBinary()
	if err != nil {
		return 0, err
	}

	// create the 16 character string
	string16 := string(bytes)
	if len(string16) != 16 {
		return 0, err
	}

	// parse the string to an int64
	number, err := strconv.ParseInt(string16, _int64base, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}
