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
)

// New generates a new random IFC GUID.
func New() (string, error) {
	guid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return FromUuid(guid)
}

// FromRevitUniqueId converts a Revit 'unique identifier' to an IFC GUID.
func FromRevitUniqueId(uniqueId string) (string, error) {
	u, err := revitUniqueIdToUuid(uniqueId)
	if err != nil {
		return "", err
	}

	return FromUuid(u)
}

// FromAutoCadHandle converts an AutoCAD handle to an IFC GUID.
func FromAutoCadHandle(handle string) (string, error) {
	u, err := autoCadHandleToUuid(handle)
	if err != nil {
		return "", err
	}

	return FromUuid(u)
}

// ToAutoCadHandle converts an IFC GUID to an AutoCAD handle.
func ToAutoCadHandle(ifcGuid string) (string, error) {
	u, err := ToUuid(ifcGuid)
	if err != nil {
		return "", err
	}

	return uuidToAutoCadHandle(u)
}

// FromInt32 converts a 32-bit integer to an IFC GUID.
func FromInt32(value int32) (string, error) {
	return FromInt64(int64(value))
}

// ToInt32 converts an IFC GUID to a 32-bit integer.
func ToInt32(ifcGuid string) (int32, error) {
	i64, err := ToInt64(ifcGuid)
	if err != nil {
		return 0, err
	}

	return int32(i64), nil
}

// FromInt64 converts a 64-bit integer to an IFC GUID.
func FromInt64(value int64) (string, error) {
	u, err := int64ToUuid(value)
	if err != nil {
		return "", err
	}

	return FromUuid(u)
}

// ToInt64 converts an IFC GUID to a 64-bit integer.
func ToInt64(ifcGuid string) (int64, error) {
	u, err := ToUuid(ifcGuid)
	if err != nil {
		return 0, err
	}

	return uuidToInt64(u)
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
	u, err := ToUuid(ifcGuid)
	if err != nil {
		return "", err
	}

	return uuidToIntString(u, "%d")
}

// FromString calculates an IFC GUID from an arbitrary string.
// Note that this function uses the last 16bytes of the input string to create a UUID and thus the IFC GUID.
// The input will be truncated if it is longer than 16 bytes, or padded with zeros if necessary.
func FromString(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("the input string must not be empty")
	}

	// Convert string to runes
	runes := []rune(s)

	// Create a 16-byte slice
	bytes := make([]byte, 16)

	// Start from the end of the rune slice
	byteCount := 0
	runeIndex := len(runes) - 1

	for runeIndex >= 0 && byteCount < 16 {
		r := runes[runeIndex]
		runeBytes := []byte(string(r))
		runeBytesLen := len(runeBytes)

		// Check if adding this rune would exceed 16 bytes
		if byteCount+runeBytesLen > 16 {
			break
		}

		// Copy rune bytes to the end of our byte slice
		copy(bytes[16-byteCount-runeBytesLen:], runeBytes)

		byteCount += runeBytesLen
		runeIndex--
	}

	// If we haven't filled all 16 bytes, left-pad with zeros
	if byteCount < 16 {
		for i := 0; i < 16-byteCount; i++ {
			bytes[i] = 0
		}
	}

	// Create UUID from bytes
	u, err := uuid.FromBytes(bytes)
	if err != nil {
		return "", err
	}

	// Convert UUID to IFC GUID
	return FromUuid(u)
}

// ToString is the inverse of FromString and tries to convert an IFC GUID back to an arbitrary string.
//
// Note that this function will only return the original string if that string was exactly 16 characters long.
// If the original string was too long, it will have been truncated.
// If it was too short, it will have been padded with zeros.
func ToString(ifcGuid string) (string, error) {
	u, err := ToUuid(ifcGuid)
	if err != nil {
		return "", err
	}

	bytes, _ := u.MarshalBinary()
	return string(bytes), nil
}

// ToUuid converts an IFC GUID to a UUID.
func ToUuid(ifcGuid string) (uuid.UUID, error) {
	if len(ifcGuid) != 22 {
		return uuid.Nil, fmt.Errorf("the ifcGuid must be 22 characters long")
	}

	// Check for invalid characters
	validChars := regexp.MustCompile(`^[0-9A-Za-z_$]{22}$`)
	if !validChars.MatchString(ifcGuid) {
		return uuid.Nil, fmt.Errorf("invalid IFC GUID format: contains invalid characters")
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
		num[i] = b64ToU32(ifcGuid[pos:endPos])
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
func FromUuid(u uuid.UUID) (string, error) {
	if u == uuid.Nil {
		return "", fmt.Errorf("invalid UUID: nil UUID")
	}

	bytes, _ := u.MarshalBinary()

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
		chars.WriteString(u32ToB64(num[i], digits))
		digits = 4
	}

	return chars.String(), nil
}

// b64ToU32 converts a base64 string to an uint32.
func b64ToU32(s string) uint32 {
	var result uint32 = 0
	for _, charValue := range s {
		index := uint32(strings.Index(_conversionTable, string(charValue)))
		result = (result * 64) + index
	}
	return result
}

// u32ToB64 converts an uint32 to a base64 string.
func u32ToB64(v uint32, digits int) string {
	bytes := make([]byte, digits)
	for i := 0; i < digits; i++ {
		bytes[digits-i-1] = _conversionTable[int(v%64)]
		v = v / 64
	}
	return string(bytes)
}

// reverseBytes is a helper function used to reverse a byte slice.
func reverseBytes(b []byte, from, to int) []byte {
	l := to - from + 1
	rb := make([]byte, l)
	j := 0
	for i := to; i >= from; i-- {
		rb[j] = b[i]
		j++
	}
	return rb
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
	// Note the base 16! A handle is a hexadecimal string.
	return intStringToUuid(handle, 16)
}

// uuidToAutoCadHandle converts a UUID to an AutoCad handle.
func uuidToAutoCadHandle(u uuid.UUID) (string, error) {
	// Note the "%x" - format as hexadecimal string.
	return uuidToIntString(u, "%x")
}

// intStringToUuid converts an integer string to a UUID.
func intStringToUuid(s string, base int) (uuid.UUID, error) {
	// convert the ID to an int64
	entityId, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		return uuid.Nil, err
	}

	return int64ToUuid(entityId)
}

// uuidToIntString converts a UUID to an integer string using the given format.
func uuidToIntString(u uuid.UUID, format string) (string, error) {
	number, err := uuidToInt64(u)
	if err != nil {
		return "", err
	}

	// return the number in decimal format
	return fmt.Sprintf(format, number), nil
}

// int64ToUuid converts an int64 to a UUID.
func int64ToUuid(v int64) (uuid.UUID, error) {
	// Create a 16-byte array
	bytes := make([]byte, 16)

	// Convert int64 to bytes
	binary.BigEndian.PutUint64(bytes[8:], uint64(v))

	// Create UUID from bytes
	return uuid.FromBytes(bytes)
}

// uuidToInt64 converts a UUID to an int64.
func uuidToInt64(u uuid.UUID) (int64, error) {
	bytes, err := u.MarshalBinary()
	if err != nil {
		return 0, err
	}
	return int64(binary.BigEndian.Uint64(bytes[8:])), nil
}
