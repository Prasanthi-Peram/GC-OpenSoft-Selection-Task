package store
import(
	"context"
	"database/sql"
	"errors"
)

var(
	ErrNotFound= errors.New("Resource not found")
)
type Storage struct{
	Posts interface{
		GetByID(context.Context,int64) (*Post,error)
		Create(context.Context, *Post) error
		Delete(context.Context,int64) error
		Update(context.Context, *Post) error
	}

	Users interface{
		Create(context.Context, *User) error
	}

	Comments interface {
		//Create(context.Context, *Comment) error
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
	
}

func NewStorage(db *sql.DB) Storage{
	return Storage{
		Posts: &PostStore{db},
		Users: &UsersStore{db},
		Comments: &CommentStore{db},
	}
}