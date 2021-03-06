package repository

import (
	"log"

	"github.com/GenkiHirano/layered-architecture/domain/model"
	"xorm.io/xorm"
)

// TaskRepository task repositoryのinterface
type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
	Get(id int) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	// Delete(task *model.Task) error
}

type taskRepository struct {
	Engine *xorm.Engine
}

// NewTaskRepository task repositoryのコンストラクタ
func NewTaskRepository(engine *xorm.Engine) TaskRepository {
	return &taskRepository{Engine: engine}
}

// Create taskの保存
func (tr *taskRepository) Create(task *model.Task) (*model.Task, error) {
	newTask := &model.Task{
		ID:      task.ID,
		Title:   task.Title,
		Content: task.Content,
	}

	_, err := tr.Engine.Table("task").Insert(newTask)
	if err != nil {
		log.Fatal(err)
	}

	return newTask, nil
}

// Get taskをIDで取得
func (tr *taskRepository) Get(id int) (*model.Task, error) {
	task := &model.Task{ID: id}

	_, err := tr.Engine.Where("id = ?", task.ID).Get(task)
	if err != nil {
		log.Fatal(err)
	}

	return task, nil

}

// Update taskの更新
func (tr *taskRepository) Update(task *model.Task) (*model.Task, error) {
	updateTask := &model.Task{
		ID:      task.ID,
		Title:   task.Title,
		Content: task.Content,
	}

	_, err := tr.Engine.Where("id = ?", updateTask.ID).Update(updateTask)
	if err != nil {
		log.Fatal(err)
	}

	return updateTask, nil
}

// // Delete taskの削除
// func (tr *taskRepository) Delete(task *model.Task) error {
// 	// FIXME: idのみ削除されないので対応する
// 	_, err := tr.Engine.Where("id = ?", task.ID).Delete(task)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
