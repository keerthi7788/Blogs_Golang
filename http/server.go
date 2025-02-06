package http

import (
	"Blogs/config"
	"Blogs/http/handlers"
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Server struct {
	Logger         *zap.Logger // TODO
	Conf           *config.Config
	UserHandler    *handlers.UserHandler
	PostHandler    *handlers.PostHandler
	CommentHandler *handlers.CommentHandler
}

func NewServer(conf *config.Config, logger *zap.Logger, userHandler *handlers.UserHandler, postHandler *handlers.PostHandler, commentHandler *handlers.CommentHandler) *Server {
	return &Server{
		Conf:           conf,
		Logger:         logger,
		UserHandler:    userHandler,
		PostHandler:    postHandler,
		CommentHandler: commentHandler,
	}
}
// STRUCT METHOD
func (s *Server) Listen(ctx context.Context, addr string) error {
	// server := &http.Server{Addr: addr, Handler: r}
	r := chi.NewRouter()
	server := &http.Server{Addr: addr, Handler: r}

	r.Route("/blogs", func(r chi.Router) {
		/* Routes for user details */
		// r.Use("/middleware,"s.)
		r.Post("/users", s.UserHandler.CreateUser)
		r.Get("/users", s.UserHandler.GetAllUsers)
		r.Get("/users", s.UserHandler.GetUserByID)
		r.Delete("/users", s.UserHandler.DeleteAllUsers)
		r.Delete("/user", s.UserHandler.DeleteUserByID)
		r.Patch("/user", s.UserHandler.UpdateUserByID)
		r.Put("/users", s.UserHandler.UpdateUserDetails)

		/* Routes for post details */
		r.Post("/posts", s.PostHandler.CreatePost)
		r.Get("/posts", s.PostHandler.GetAllPosts)
		r.Get("/post", s.PostHandler.GetPostByID)
		r.Delete("/posts", s.PostHandler.DeleteAllPosts)
		r.Delete("/post", s.PostHandler.DeletePostById)
		r.Put("/posts", s.PostHandler.UpdateAllPosts)
		r.Patch("/post", s.PostHandler.UpdatePostById)

		/* Routes for comment details */
		r.Post("/comments", s.CommentHandler.CreateComment)
		r.Get("/comments", s.CommentHandler.GetAllComments)
		r.Get("/comment", s.CommentHandler.GetCommentByID)
		r.Delete("/comments", s.CommentHandler.DeleteAllComments)
		r.Delete("/comment", s.CommentHandler.DeleteCommentById)
		r.Patch("/comment", s.CommentHandler.UpdateCommentById)
		r.Put("/comment", s.CommentHandler.UpdateAllComments)
	})

	// server := &http.Server{Addr: addr, Handler: r}

	// Channel to capture errors
	errch := make(chan error, 1)

	go func() {
		s.Logger.Info("Starting server", zap.String("addr", addr))
		errch <- server.ListenAndServe()
	}()

	select {
	case err := <-errch:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return server.Shutdown(shutdownCtx)
	}
}
