package http

import (
	"Blogs/config"
	"Blogs/http/handlers"
	"Blogs/http/middleware"
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Server struct {
	Logger         *zap.Logger
	Conf           *config.Config
	UserHandler    *handlers.UserHandler
	PostHandler    *handlers.PostHandler
	CommentHandler *handlers.CommentHandler
	Authhandler    *handlers.Authhandler
}

func NewServer(conf *config.Config, logger *zap.Logger, userHandler *handlers.UserHandler, postHandler *handlers.PostHandler, commentHandler *handlers.CommentHandler, Authhandler *handlers.Authhandler) *Server {

	// conf.Database.Name = "Blogs"
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
	r := chi.NewRouter()
	// Global Middleware
	r.Use(middleware.RequestLogger)
	r.Post("/create", s.UserHandler.CreateUser) // Logs all request details

	r.Post("/login", s.Authhandler.LoginHandler)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthJWT) // Protecting these routes
		// r.Route("/blogs", func(r chi.Router) {
		// Protecting routes with JWT Authentication Middlewar

		/* Routes for user details */
		r.Route("/users", func(r chi.Router) {
			// r.Post("/", s.Authhandler.LoginHandler)
			// r.Post("/", s.UserHandler.CreateUser)
			r.Get("/", s.UserHandler.GetAllUsers)
			r.Get("/", s.UserHandler.GetUserByID)
			r.Put("/", s.UserHandler.UpdateUserDetails)
			r.Delete("/", s.UserHandler.DeleteAllUsers)
		})
		r.Route("/user", func(r chi.Router) {
			r.Delete("/", s.UserHandler.DeleteUserByID)
			r.Patch("/", s.UserHandler.UpdateUserByID)
		})
		r.Route("/posts", func(r chi.Router) {
			/* Routes for post details */
			r.Post("/posts", s.PostHandler.CreatePost)
			r.Get("/posts", s.PostHandler.GetAllPosts)
			r.Delete("/posts", s.PostHandler.DeleteAllPosts)
			r.Put("/posts", s.PostHandler.UpdateAllPosts)

		})
		r.Route("/post", func(r chi.Router) {
			r.Get("/post", s.PostHandler.GetPostByID)
			r.Delete("/post", s.PostHandler.DeletePostById)
			r.Patch("/post", s.PostHandler.UpdatePostById)
		})
		r.Route("/comments", func(r chi.Router) {
			r.Post("/comments", s.CommentHandler.CreateComment)
			r.Get("/comments", s.CommentHandler.GetAllComments)
			r.Delete("/comments", s.CommentHandler.DeleteAllComments)

		})
		r.Route("/comment", func(r chi.Router) {
			r.Get("/comment", s.CommentHandler.GetCommentByID)

			r.Delete("/comment", s.CommentHandler.DeleteCommentById)
			r.Patch("/comment", s.CommentHandler.UpdateCommentById)
			r.Put("/comment", s.CommentHandler.UpdateAllComments)
		})

	})
	// })

	server := &http.Server{Addr: addr, Handler: r}

	// Error channel
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
