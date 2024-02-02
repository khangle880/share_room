package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/khangle880/share_room/graph/model"
	"github.com/khangle880/share_room/pg"
	"github.com/khangle880/share_room/utils"
)

// DeletedAt is the resolver for the deletedAt field.
func (r *budgetResolver) DeletedAt(ctx context.Context, obj *pg.Budget) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// Description is the resolver for the description field.
func (r *budgetResolver) Description(ctx context.Context, obj *pg.Budget) (*string, error) {
	panic(fmt.Errorf("not implemented: Description - description"))
}

// Amount is the resolver for the amount field.
func (r *budgetResolver) Amount(ctx context.Context, obj *pg.Budget) (float64, error) {
	panic(fmt.Errorf("not implemented: Amount - amount"))
}

// Icon is the resolver for the icon field.
func (r *budgetResolver) Icon(ctx context.Context, obj *pg.Budget) (*pg.Icon, error) {
	panic(fmt.Errorf("not implemented: Icon - icon"))
}

// Room is the resolver for the room field.
func (r *budgetResolver) Room(ctx context.Context, obj *pg.Budget) (*pg.Room, error) {
	panic(fmt.Errorf("not implemented: Room - room"))
}

// Period is the resolver for the period field.
func (r *budgetResolver) Period(ctx context.Context, obj *pg.Budget) (*pg.PeriodType, error) {
	panic(fmt.Errorf("not implemented: Period - period"))
}

// EndDate is the resolver for the end_date field.
func (r *budgetResolver) EndDate(ctx context.Context, obj *pg.Budget) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: EndDate - end_date"))
}

// Transactions is the resolver for the transactions field.
func (r *budgetResolver) Transactions(ctx context.Context, obj *pg.Budget) ([]pg.Transaction, error) {
	panic(fmt.Errorf("not implemented: Transactions - transactions"))
}

// Members is the resolver for the members field.
func (r *budgetResolver) Members(ctx context.Context, obj *pg.Budget) ([]pg.User, error) {
	panic(fmt.Errorf("not implemented: Members - members"))
}

// DeletedAt is the resolver for the deletedAt field.
func (r *categoryResolver) DeletedAt(ctx context.Context, obj *pg.Category) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// Icon is the resolver for the icon field.
func (r *categoryResolver) Icon(ctx context.Context, obj *pg.Category) (*pg.Icon, error) {
	panic(fmt.Errorf("not implemented: Icon - icon"))
}

// Parent is the resolver for the parent field.
func (r *categoryResolver) Parent(ctx context.Context, obj *pg.Category) (*pg.Category, error) {
	panic(fmt.Errorf("not implemented: Parent - parent"))
}

// Icon is the resolver for the icon field.
func (r *eventResolver) Icon(ctx context.Context, obj *model.Event) (*pg.Icon, error) {
	panic(fmt.Errorf("not implemented: Icon - icon"))
}

// DeletedAt is the resolver for the deletedAt field.
func (r *iconResolver) DeletedAt(ctx context.Context, obj *pg.Icon) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// Type is the resolver for the type field.
func (r *iconResolver) Type(ctx context.Context, obj *pg.Icon) (*string, error) {
	panic(fmt.Errorf("not implemented: Type - type"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.Token, error) {
	user, err := r.Repository.GetUserByEmail(ctx, pg.StringPtrToNullString(&email))
	if err != nil {
		return nil, errors.New("user not found")
	}
	if err := utils.ComparePasswords(password, user.HashedPassword); err != nil {
		return nil, err
	}
	token, err := utils.JwtGenerate(user.ID)
	if err != nil {
		return nil, err
	}
	return &model.Token{
		AccessToken: &token,
		User:        &user,
	}, nil
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.CreateUserInput) (*model.Token, error) {
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.Log.Err(err).Msg("error while hashing password")
		return nil, errors.New("something went wrong")
	}
	userParams := pg.CreateUserParams{
		Username:       input.Username,
		HashedPassword: hashedPassword,
		Email:          pg.StringPtrToNullString(input.Email),
		Phone:          pg.StringPtrToNullString(input.Phone),
	}
	profileParams := pg.CreateProfileParams{
		Firstname: pg.StringPtrToNullString(input.Firstname),
		Lastname:  pg.StringPtrToNullString(input.Lastname),
		Role:      *input.Role,
		Bio:       pg.StringPtrToNullString(input.Bio),
		Avatar:    pg.StringPtrToNullString(input.Avatar),
	}

	user, err := r.Repository.CreateUser(ctx, userParams, profileParams)
	if err != nil {
		utils.Log.Err(err).Msg("error while creating user")
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, fmt.Errorf("user already exists")
		}
		return nil, errors.New("someting went wrong")
	}

	token, err := utils.JwtGenerate(user.ID)
	if err != nil {
		utils.Log.Err(err).Msg("error while generate token")
		return nil, errors.New("something went wrong")
	}

	return &model.Token{
		AccessToken: &token,
		User:        user,
	}, nil
}

// CreateIcon is the resolver for the createIcon field.
func (r *mutationResolver) CreateIcon(ctx context.Context, input model.CreateIconInput) (*pg.Icon, error) {
	panic(fmt.Errorf("not implemented: CreateIcon - createIcon"))
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input model.CreateEventInput) (*model.Event, error) {
	panic(fmt.Errorf("not implemented: CreateEvent - createEvent"))
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.CreateCategoryInput) (*pg.Category, error) {
	panic(fmt.Errorf("not implemented: CreateCategory - createCategory"))
}

// CreateBudget is the resolver for the createBudget field.
func (r *mutationResolver) CreateBudget(ctx context.Context, input model.CreateBudgetInput) (*pg.Budget, error) {
	panic(fmt.Errorf("not implemented: CreateBudget - createBudget"))
}

// CreateTransaction is the resolver for the createTransaction field.
func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.CreateTransInput) (*pg.Transaction, error) {
	panic(fmt.Errorf("not implemented: CreateTransaction - createTransaction"))
}

// UpdateBudget is the resolver for the updateBudget field.
func (r *mutationResolver) UpdateBudget(ctx context.Context, id uuid.UUID, input model.UpdateBudgetInput) (*pg.Budget, error) {
	panic(fmt.Errorf("not implemented: UpdateBudget - updateBudget"))
}

// DeleteBudget is the resolver for the deleteBudget field.
func (r *mutationResolver) DeleteBudget(ctx context.Context, id uuid.UUID) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteBudget - deleteBudget"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id uuid.UUID) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// DeletedAt is the resolver for the deletedAt field.
func (r *profileResolver) DeletedAt(ctx context.Context, obj *pg.Profile) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// Firstname is the resolver for the firstname field.
func (r *profileResolver) Firstname(ctx context.Context, obj *pg.Profile) (*string, error) {
	panic(fmt.Errorf("not implemented: Firstname - firstname"))
}

// Lastname is the resolver for the lastname field.
func (r *profileResolver) Lastname(ctx context.Context, obj *pg.Profile) (*string, error) {
	panic(fmt.Errorf("not implemented: Lastname - lastname"))
}

// Bio is the resolver for the bio field.
func (r *profileResolver) Bio(ctx context.Context, obj *pg.Profile) (*string, error) {
	panic(fmt.Errorf("not implemented: Bio - bio"))
}

// Avatar is the resolver for the avatar field.
func (r *profileResolver) Avatar(ctx context.Context, obj *pg.Profile) (*string, error) {
	panic(fmt.Errorf("not implemented: Avatar - avatar"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, role pg.UserRole) ([]pg.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context, filter *model.BudgetFilter, limit *int, offset *int) ([]pg.Category, error) {
	panic(fmt.Errorf("not implemented: Categories - categories"))
}

// Budgets is the resolver for the budgets field.
func (r *queryResolver) Budgets(ctx context.Context) ([]pg.Budget, error) {
	panic(fmt.Errorf("not implemented: Budgets - budgets"))
}

// DeletedAt is the resolver for the deletedAt field.
func (r *roomResolver) DeletedAt(ctx context.Context, obj *pg.Room) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// Address is the resolver for the address field.
func (r *roomResolver) Address(ctx context.Context, obj *pg.Room) (*string, error) {
	panic(fmt.Errorf("not implemented: Address - address"))
}

// Admin is the resolver for the admin field.
func (r *roomResolver) Admin(ctx context.Context, obj *pg.Room) (*pg.User, error) {
	panic(fmt.Errorf("not implemented: Admin - admin"))
}

// Member is the resolver for the member field.
func (r *roomResolver) Member(ctx context.Context, obj *pg.Room) ([]pg.User, error) {
	panic(fmt.Errorf("not implemented: Member - member"))
}

// Avatar is the resolver for the avatar field.
func (r *roomResolver) Avatar(ctx context.Context, obj *pg.Room) (*string, error) {
	panic(fmt.Errorf("not implemented: Avatar - avatar"))
}

// Background is the resolver for the background field.
func (r *roomResolver) Background(ctx context.Context, obj *pg.Room) (*string, error) {
	panic(fmt.Errorf("not implemented: Background - background"))
}

// DeletedAt is the resolver for the deletedAt field.
func (r *transactionResolver) DeletedAt(ctx context.Context, obj *pg.Transaction) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// Category is the resolver for the category field.
func (r *transactionResolver) Category(ctx context.Context, obj *pg.Transaction) (*pg.Category, error) {
	panic(fmt.Errorf("not implemented: Category - category"))
}

// Budget is the resolver for the budget field.
func (r *transactionResolver) Budget(ctx context.Context, obj *pg.Transaction) (*pg.Budget, error) {
	panic(fmt.Errorf("not implemented: Budget - budget"))
}

// Event is the resolver for the event field.
func (r *transactionResolver) Event(ctx context.Context, obj *pg.Transaction) (*model.Event, error) {
	panic(fmt.Errorf("not implemented: Event - event"))
}

// Description is the resolver for the description field.
func (r *transactionResolver) Description(ctx context.Context, obj *pg.Transaction) (*string, error) {
	panic(fmt.Errorf("not implemented: Description - description"))
}

// Creators is the resolver for the creators field.
func (r *transactionResolver) Creators(ctx context.Context, obj *pg.Transaction) ([]pg.User, error) {
	panic(fmt.Errorf("not implemented: Creators - creators"))
}

// Partners is the resolver for the partners field.
func (r *transactionResolver) Partners(ctx context.Context, obj *pg.Transaction) ([]pg.User, error) {
	panic(fmt.Errorf("not implemented: Partners - partners"))
}

// Amount is the resolver for the amount field.
func (r *transactionResolver) Amount(ctx context.Context, obj *pg.Transaction) (float64, error) {
	panic(fmt.Errorf("not implemented: Amount - amount"))
}

// DeletedAt is the resolver for the deletedAt field.
func (r *userResolver) DeletedAt(ctx context.Context, obj *pg.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}

// LastLoginAt is the resolver for the lastLoginAt field.
func (r *userResolver) LastLoginAt(ctx context.Context, obj *pg.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: LastLoginAt - lastLoginAt"))
}

// Email is the resolver for the email field.
func (r *userResolver) Email(ctx context.Context, obj *pg.User) (*string, error) {
	panic(fmt.Errorf("not implemented: Email - email"))
}

// Phone is the resolver for the phone field.
func (r *userResolver) Phone(ctx context.Context, obj *pg.User) (*string, error) {
	panic(fmt.Errorf("not implemented: Phone - phone"))
}

// Profile is the resolver for the profile field.
func (r *userResolver) Profile(ctx context.Context, obj *pg.User) (*pg.Profile, error) {
	profile, err := r.Repository.GetProfileByUserID(ctx, obj.ID)
	// if err != nil {
	// 	return nil, err
	// }
	return &profile, err
}

// Budget returns BudgetResolver implementation.
func (r *Resolver) Budget() BudgetResolver { return &budgetResolver{r} }

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Event returns EventResolver implementation.
func (r *Resolver) Event() EventResolver { return &eventResolver{r} }

// Icon returns IconResolver implementation.
func (r *Resolver) Icon() IconResolver { return &iconResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Profile returns ProfileResolver implementation.
func (r *Resolver) Profile() ProfileResolver { return &profileResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Room returns RoomResolver implementation.
func (r *Resolver) Room() RoomResolver { return &roomResolver{r} }

// Transaction returns TransactionResolver implementation.
func (r *Resolver) Transaction() TransactionResolver { return &transactionResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type budgetResolver struct{ *Resolver }
type categoryResolver struct{ *Resolver }
type eventResolver struct{ *Resolver }
type iconResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type profileResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type roomResolver struct{ *Resolver }
type transactionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
