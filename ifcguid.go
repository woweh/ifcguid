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

// IsValid checks if the given string is a valid IFC GUID.
// Returns nil if the given string is a valid IFC GUID, otherwise returns an error.
func IsValid(ifcGuid string) error {
	if len(ifcGuid) != 22 {
		return fmt.Errorf("the IFC GUID must be 22 characters long")
	}
	validChars := regexp.MustCompile(`^[0-9A-Za-z_$]{22}$`)
	if !validChars.MatchString(ifcGuid) {
		return fmt.Errorf("invalid IFC GUID format: contains invalid characters")
	}
	lastBase64Num := ifcGuid[0]
	if (lastBase64Num - 48) > 3 {
		return fmt.Errorf("illegal IFC GUID: it is greater than 128 bits")
	}
	return nil
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
func FromIntString(value string) (string, error) {
	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return "", err
	}
	return FromInt64(intVal)
}

// ToIntString converts an IFC GUID to a string representation of an integer.
func ToIntString(ifcGuid string) (string, error) {
	u, err := ToUuid(ifcGuid)
	if err != nil {
		return "", err
	}
	return uuidToIntString(u, "%d")
}

// FromString calculates an IFC GUID from an arbitrary string.
// Note that this function uses the last 16 bytes of the input string to create a UUID and thus the IFC GUID.
// The input will be truncated if it is longer than 16 bytes, or padded with zeros if it shorter than 16 bytes.
func FromString(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("the input string must not be empty")
	}
	bytes := stringTo16Bytes(s)
	u, err := uuid.FromBytes(bytes)
	if err != nil {
		return "", err
	}
	return FromUuid(u)
}

// stringTo16Bytes converts a string to a byte slice of exactly 16 bytes.
// This takes 16 bytes from the end of the string, and if necessary, pads with zeros.
func stringTo16Bytes(s string) []byte {
	bytes := make([]byte, 16)
	// to handle unicode characters correctly, process the runes in reverse order
	runes := []rune(s)
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
	return bytes
}

// ToString is the inverse of FromString and tries to convert an IFC GUID back to an arbitrary string.
//
// Note that this function will only return the original string if that string occupied 16 bytes.
// See FromString for more details.
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
	err := IsValid(ifcGuid)
	if err != nil {
		return uuid.Nil, err
	}
	pos := 0
	digits := 2
	num := make([]uint32, 6)
	for i := 0; i < 6; i++ {
		endPos := pos + digits
		num[i] = b64ToU32(ifcGuid[pos:endPos])
		pos += digits
		digits = 4
	}
	data1 := num[0]*16777216 + num[1]                // 16-13. bytes
	data2 := uint16(num[2] / 256)                    // 12-11. bytes
	data3 := uint16((num[2]%256)*256 + num[3]/65536) // 10-09. bytes
	guidBytes := make([]byte, 16)
	// Note:
	// Microsoft GUIDs use a mixed-endian format where the first three components
	// are stored in little-endian order, while the remaining bytes are in big-endian order.
	// => Reverse the order of the bytes to be compatible with existing converters.
	// Write data1 (4 bytes) in little-endian order (reversed)
	guidBytes[0] = byte(data1 >> 24)
	guidBytes[1] = byte(data1 >> 16)
	guidBytes[2] = byte(data1 >> 8)
	guidBytes[3] = byte(data1)
	// Write data2 (2 bytes) in little-endian order (reversed)
	guidBytes[4] = byte(data2 >> 8)
	guidBytes[5] = byte(data2)
	// Write data3 (2 bytes) in little-endian order (reversed)
	guidBytes[6] = byte(data3 >> 8)
	guidBytes[7] = byte(data3)
	// Write remaining bytes directly
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
	// The UUID bytes are already in the correct order.
	// We just need to convert them to base 64.
	data1 := binary.BigEndian.Uint32(bytes[0:4])         // 4byte - int32
	data2 := uint32(binary.BigEndian.Uint16(bytes[4:6])) // 2byte - int16
	data3 := uint32(binary.BigEndian.Uint16(bytes[6:8])) // 2byte - int16
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

// FromUuidString converts a UUID string to an ifcGuid.
//
// See the uuid.Parse() documentation for supported UUID forms.
func FromUuidString(s string) (string, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return "", err
	}
	return FromUuid(u)
}

// ToUuidString converts an ifcGuid to a UUID string, is the form of `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`,
// or an empty string if the ifcGuid cannot be converted to a UUID.
func ToUuidString(ifcGuid string) (string, error) {
	u, err := ToUuid(ifcGuid)
	if err != nil {
		return "", err
	}
	return u.String(), nil
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
	return fmt.Sprintf(format, number), nil
}

// int64ToUuid converts an int64 to a UUID.
func int64ToUuid(v int64) (uuid.UUID, error) {
	bytes := make([]byte, 16)
	binary.BigEndian.PutUint64(bytes[8:], uint64(v))
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
