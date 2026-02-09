package repository

import (
	"context"
	"my-personal-page/backend/internal/db/generated"
	"my-personal-page/backend/internal/domain"
)

type PostRepo struct {
	q *db.Queries
}

func NewPostRepo(q *db.Queries) *PostRepo{
	return &PostRepo{
		q: q,
	}
}

func (repo *PostRepo) Create(ctx context.Context, post *domain.Post) (domain.Post, error){
	row, err := repo.q.CreatePosts(ctx, db.CreatePostsParams{
		Title: post.Title,
		Content: post.Content,
		CreatedAt: post.CreatedAt,
	})
	if err != nil {
		return domain.Post{}, err
	}
	return domain.Post{
        ID:        row.ID,
        Title:     row.Title,
        Content:   row.Content,
        CreatedAt: row.CreatedAt,
    }, nil
}

func (repo *PostRepo) GetAll(ctx context.Context) ([]domain.Post, error){
	rows, err := repo.q.ListPosts(ctx)

	if err != nil {
		return []domain.Post{}, err
	}

	posts := make([]domain.Post, 0, len(rows))
	for _, row := range rows {
		posts = append(posts, domain.Post{
			ID:        row.ID,
        	Title:     row.Title,
        	Content:   row.Content,
        	CreatedAt: row.CreatedAt,
		})
	}

	return posts, nil
}

func (repo *PostRepo) Get(ctx context.Context, id int64) (domain.Post, error){
	row, err := repo.q.GetPosts(ctx, id)
	if err != nil {
		return domain.Post{}, err
	}

	return domain.Post{
		ID:        row.ID,
		Title:     row.Title,
		Content:   row.Content,
		CreatedAt: row.CreatedAt,
	}, nil
}

func (repo *PostRepo) Delete(ctx context.Context, id int64) (domain.Post, error){
	row, err := repo.q.DeletePosts(ctx, id)
	if err != nil {
		return domain.Post{}, err
	}

	return domain.Post{
		ID:        row.ID,
		Title:     row.Title,
		Content:   row.Content,
		CreatedAt: row.CreatedAt,
	}, nil
}

func (repo *PostRepo) Update(ctx context.Context, post *domain.Post) (domain.Post, error){
	row, err := repo.q.UpdatePosts(ctx, db.UpdatePostsParams{
		Title: post.Title,
		Content: post.Content,
		ID: post.ID,
	})
	
	if err != nil {
		return domain.Post{}, err
	}

	return domain.Post{
		ID:        row.ID,
		Title:     row.Title,
		Content:   row.Content,
		CreatedAt: row.CreatedAt,
	}, nil
}