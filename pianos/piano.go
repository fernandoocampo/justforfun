package pianos

const pianosKey = 53

func GuestColorKey(key int) int {
	if key < 1 {
		return -1
	}
	if key < pianosKey {
		return key
	}
	return (key % pianosKey)
}
