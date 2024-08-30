package handler

import (
	"portfolio/simple-Kanban/infrastructure/config"
	"portfolio/simple-Kanban/infrastructure/database"
	"portfolio/simple-Kanban/pkg/middlewares"
	"portfolio/simple-Kanban/repository/category_repository/category_repo"
	"portfolio/simple-Kanban/repository/task_repository/task_repo"
	"portfolio/simple-Kanban/repository/user_repository/user_repo"
	"portfolio/simple-Kanban/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	config.LoadAppConfig()
	db := database.GetDatabaseInstance()
	database.SeedAdmin(db)

	//Dependency Injection
	userRepo := user_repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	categoryRepo := category_repo.NewCategoryRepo(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	taskRepo := task_repo.NewTaskRepo(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := NewTaskHandler(taskService)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.CreateUser)
		userRoute.POST("/login", userHandler.Login)

		userRoute.Use(middlewares.Authentication())
		{
			userRoute.PUT("/update-account", userHandler.UpdateUser)
		}
	}

	categoriesRoute := route.Group("/categories")
	{
		categoriesRoute.Use(middlewares.Authentication())
		{
			categoriesRoute.GET("/", categoryHandler.GetCategories)

			categoriesRoute.Use(middlewares.AdminAuthorization())
			categoriesRoute.POST("/", categoryHandler.CreateCategory)
			categoriesRoute.PATCH("/:categoryId", categoryHandler.UpdateCategory)
			categoriesRoute.DELETE("/:categoryId", categoryHandler.DeleteCategory)
		}
	}

	taskRoute := route.Group("/tasks")
	{
		taskRoute.Use(middlewares.Authentication())
		{
			taskRoute.GET("/", taskHandler.GetTasks)
			taskRoute.POST("/", taskHandler.CreateTask)
			taskRoute.PUT("/:taskId", taskHandler.UpdateTask)

			taskRoute.Use(middlewares.Authentication())
			{
				taskRoute.PATCH("/update-status/:taskId", taskHandler.UpdateStatus)
				taskRoute.PATCH("/update-category/:taskId", taskHandler.UpdateTaskCategory)
			}
			taskRoute.DELETE("/:taskId", taskHandler.DeleteTask)

		}
	}

	route.Run(":" + config.GetAppConfig().Port)
}
