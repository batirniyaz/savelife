package database

import (
    "context"
    "database/sql"
)

// UserRepository handles database operations related to users
type UserRepository struct {
    db *sql.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db}
}

// CreateUser inserts a new user into the database
func (ur *UserRepository) CreateUser(ctx context.Context, user *User) (int64, error) {
    result, err := ur.db.ExecContext(ctx, "INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

// GetUserByUsername retrieves a user by username from the database
func (ur *UserRepository) GetUserByUsername(ctx context.Context, Username string) (*User, error) {
    var user User
    err := ur.db.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE username = ?", Username).Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, UserID int) (*User, error) {
    var user User
    err := ur.db.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE id = ?", UserID).Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// GetUserDetailsByID retrieves user details by user ID from the database
func (ur *UserRepository) GetUserDetailsByID(ctx context.Context, userID int) (*Request, error) {
    var userDetails Request
    err := ur.db.QueryRowContext(ctx, "SELECT id, full_name, date_of_birth, illness, location, phone_number FROM users WHERE id = ?", userID).Scan(&userDetails.ID, &userDetails.FullName, &userDetails.DateOfBirth, &userDetails.Illness, &userDetails.Location, &userDetails.PhoneNumber)
    if err != nil {
        return nil, err
    }
    return &userDetails, nil
}

// Add other CRUD operations as needed
