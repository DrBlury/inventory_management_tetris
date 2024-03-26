package usecase

import (
	"context"
	domain "linuxcode/inventory_manager/pkg/domain/model"
	repo "linuxcode/inventory_manager/pkg/repo/generated"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// GetAllUsers returns all users that exist
func (a appLogicImpl) GetAllUsers(ctx context.Context) (*[]domain.User, error) {
	users, err := a.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	// map repo model to domain model
	var domainUsers = make([]domain.User, 0, len(users))
	for _, rUser := range users {
		user := domain.User{
			ID:       int(rUser.ID),
			Username: rUser.Username.String,
			Email:    rUser.Email.String,
		}
		domainUsers = append(domainUsers, user)
	}

	// get inventories for each user
	for i, user := range domainUsers {
		inventories, err := a.queries.ListInventoriesByUserID(ctx, pgtype.Int4{Int32: int32(user.ID), Valid: true})
		if err != nil {
			return nil, err
		}

		// map inventories to domain model
		for _, rInv := range inventories {
			inventory := domain.MapRepoInventoryToDomainInventoryMeta(&rInv)
			domainUsers[i].Inventories = append(domainUsers[i].Inventories, *inventory)
		}
	}

	return &domainUsers, nil
}

// GetUserById returns the user with the given id
func (a appLogicImpl) GetUserById(ctx context.Context, userId int) (*domain.User, error) {
	user, err := a.queries.GetUser(ctx, int32(userId))
	if err != nil {
		return nil, err
	}

	// map repo model to domain model
	domainUser := domain.User{
		ID:       int(user.ID),
		Username: user.Username.String,
		Email:    user.Email.String,
	}
	return &domainUser, nil
}

// GetUserByUsername returns the user with the given username
func (a appLogicImpl) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	user, err := a.queries.GetUserByUsername(ctx, pgtype.Text{String: username, Valid: true})
	if err != nil {
		return nil, err
	}

	// map repo model to domain model
	domainUser := &domain.User{
		ID:       int(user.ID),
		Username: user.Username.String,
		Email:    user.Email.String,
	}
	return domainUser, nil
}

// HashPassword hashes the given password using bcrypt and the given salt
func HashPassword(password, salt string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bcryptPassword), nil
}

// AddUser creates a new user
func (a appLogicImpl) AddUser(ctx context.Context, createUserParams domain.CreateUserParams) (*domain.User, error) {
	// hash password using bcrypt
	salt := "saltySalt-" + createUserParams.Username
	hashedPassword, err := HashPassword(createUserParams.Password, salt)
	if err != nil {
		return nil, err
	}

	createdUser, err := a.queries.CreateUser(ctx, repo.CreateUserParams{
		Username:     pgtype.Text{String: createUserParams.Username, Valid: true},
		Email:        pgtype.Text{String: createUserParams.Email, Valid: true},
		Salt:         pgtype.Text{String: salt, Valid: true},
		PasswordHash: pgtype.Text{String: hashedPassword, Valid: true},
		CreatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	// log what user was created
	a.log.Info("created user", zap.String("username", createUserParams.Username))

	// map user to domain model
	createdUserDomain := &domain.User{
		ID:       int(createdUser.ID),
		Username: createdUser.Username.String,
		Email:    createdUser.Email.String,
	}
	return createdUserDomain, nil
}

// DeleteUserById deletes the user with the given id
func (a appLogicImpl) DeleteUserById(ctx context.Context, userId int) error {
	repoUser, err := a.queries.DeleteUser(ctx, int32(userId))
	if err != nil {
		return err
	}

	// log what user was deleted
	a.log.Info("deleted user", zap.String("username", repoUser.Username.String))
	return nil
}

// UpdateUser updates the user with the given id
func (a appLogicImpl) UpdateUser(ctx context.Context, userId int, updateUserParams domain.CreateUserParams) (*domain.User, error) {
	// hash password using bcrypt
	salt := "saltySalt-" + updateUserParams.Username
	hashedPassword, err := HashPassword(updateUserParams.Password, salt)
	if err != nil {
		return nil, err
	}

	updatedUser, err := a.queries.UpdateUser(ctx, repo.UpdateUserParams{
		ID:           int32(userId),
		Username:     pgtype.Text{String: updateUserParams.Username, Valid: true},
		Email:        pgtype.Text{String: updateUserParams.Email, Valid: true},
		Salt:         pgtype.Text{String: salt, Valid: true},
		PasswordHash: pgtype.Text{String: hashedPassword, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	// map user to domain model
	updatedUserDomain := &domain.User{
		ID:       int(updatedUser.ID),
		Username: updatedUser.Username.String,
		Email:    updatedUser.Email.String,
	}

	// log what user was updated
	a.log.Info("updated user", zap.String("username", updateUserParams.Username))
	return updatedUserDomain, nil
}
