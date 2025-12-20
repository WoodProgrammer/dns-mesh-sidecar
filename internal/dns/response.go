package dns

func CreateNXDomainResponse(query []byte) []byte {
	response := make([]byte, len(query))
	copy(response, query)

	response[2] = (query[2] & 0x01) | 0x84
	response[3] = 0x83

	response[6] = 0
	response[7] = 0

	response[8] = 0
	response[9] = 0

	response[10] = 0
	response[11] = 0

	return response
}
