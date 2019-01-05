package goprettyhash

func PrettyHash(buf []byte) string {
	var bufString string
	bufString = string(buf[:])
	if len(bufString) > 8 {
		return bufString[0:6] + ".." + bufString[len(bufString)-2:]
	}
	return bufString
}
