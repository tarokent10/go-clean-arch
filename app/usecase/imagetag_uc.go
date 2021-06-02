package usecase

import (
	"context"
	"picture-go-app/domain/models"
	"picture-go-app/domain/repository"
)

type ImageTagUsecaseIF interface {
	FindAll(ctx context.Context) []*models.ImageTag
}

type ImageTagUsecase struct {
	rep repository.ImagetagRepository
}

func NewImagetagUsecase(r repository.ImagetagRepository) *ImageTagUsecase {
	return &ImageTagUsecase{
		rep: r,
	}
}

func (uc ImageTagUsecase) FindAll(ctx context.Context) []*models.ImageTag {
	return uc.rep.FindAllWithContext(ctx)
}

// // App comment
// type App struct {
// 	db *sql.DB
// }

// func A() {
// 	// TODO DB処理はifはドメイン、実装はインフラ層で分離。DIはmainで行える。
// 	sqlinfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "kentaro", "kentaro", "localhost", 3306, "picture")
// 	fmt.Println(sqlinfo)
// 	db, err := sql.Open("mysql", sqlinfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	if err := db.Ping(); err != nil {
// 		panic(err)
// 	}

// 	imagetag := []*entities.ImageTag{&entities.ImageTag{Name: "name1"}, &entities.ImageTag{Name: "name2"}}
// 	app := &App{db}
// 	if err := app.Store(imagetag); err != nil {
// 		panic(err)
// 	}
// }

// // Store comment
// func (a *App) Store(obj []*entities.ImageTag) error {
// 	for _, v := range obj {
// 		if err := v.Insert(context.Background(), a.db, boil.Infer()); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
