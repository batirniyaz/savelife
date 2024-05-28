package database

// User represents a user in the database
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    // Add other user fields as needed
}

// Request represents an ambulance request in the database
type Request struct {
    ID         int    `json:"id"`
    UserID     int    `json:"user_id"`
    FullName   string `json:"full_name"`
    DateOfBirth string `json:"date_of_birth"`
    Illness    string `json:"illness"`
    Location   string `json:"location"`
    PhoneNumber string `json:"phone_number"`
}


