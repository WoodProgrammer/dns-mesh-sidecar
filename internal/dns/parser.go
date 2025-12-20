package dns

import "fmt"

func ParseQuery(query []byte) (string, string) {
	if len(query) < 12 {
		return "", ""
	}

	pos := 12
	domain := ""

	for pos < len(query) {
		length := int(query[pos])
		if length == 0 {
			pos++
			break
		}
		if length > 63 {
			return "", ""
		}
		if domain != "" {
			domain += "."
		}
		pos++
		if pos+length > len(query) {
			return "", ""
		}
		domain += string(query[pos : pos+length])
		pos += length
	}

	if pos+2 > len(query) {
		return domain, ""
	}
	qtype := uint16(query[pos])<<8 | uint16(query[pos+1])

	var qtypeStr string
	switch qtype {
	case 1:
		qtypeStr = "A"
	case 2:
		qtypeStr = "NS"
	case 5:
		qtypeStr = "CNAME"
	case 6:
		qtypeStr = "SOA"
	case 12:
		qtypeStr = "PTR"
	case 15:
		qtypeStr = "MX"
	case 16:
		qtypeStr = "TXT"
	case 28:
		qtypeStr = "AAAA"
	case 33:
		qtypeStr = "SRV"
	default:
		qtypeStr = fmt.Sprintf("TYPE%d", qtype)
	}

	return domain, qtypeStr
}
