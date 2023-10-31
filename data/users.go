package data

type User struct {
	ID      int32
	Name    string
	City    string
	Phone   int64
	Height  float32
	Married bool
}

var users = []User{
	{1, "Steve", "LA", 1234567890, 5.8, true},
	{2, "Nick", "New York", 1234567899, 6.0, false},
	{3, "John", "New York", 1234567890, 5.9, true},
	{4, "Priya", "Mumbai", 9876543210, 5.5, true},
	{5, "Robert", "London", 7890123456, 6.1, false},
	{6, "Sunita", "Delhi", 8901234567, 5.3, false},
	{7, "Michael", "Los Angeles", 2345678901, 5.10, true},
	{8, "Neha", "Bangalore", 8765432109, 5.7, true},
	{9, "James", "Toronto", 7654321098, 6.0, false},
	{10, "Amit", "Kolkata", 6543210987, 5.4, false},
	{11, "Emily", "Sydney", 5432109876, 5.8, true},
	{12, "Raj", "Chennai", 4321098765, 5.6, true},
	{13, "Olivia", "Paris", 3210987654, 5.11, false},
	{14, "Suresh", "Mysore", 2109876543, 5.2, false},
	{15, "Daniel", "San Francisco", 1098765432, 5.7, true},
	{16, "Anjali", "Jaipur", 1098765431, 5.1, false},
	{17, "Sophia", "Beijing", 9876543210, 5.6, true},
	{18, "Manish", "Mumbai", 8765432109, 5.9, false},
	{19, "Liam", "Bengaluru", 7654321098, 6.0, true},
	{20, "Shivani", "Hyderabad", 6543210987, 5.5, false},
}

func GetUsers() []User {
	return users
}
