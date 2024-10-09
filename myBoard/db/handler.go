package db

import (
	"fmt"
	"myBoad/model"
)

// 게시글 생성
func (d *DBHandler) CreateBoard(board *model.Board) (*model.Board, error) {
	result := d.gDB.Create(board)

	// 오류가 있으면 에러를 래핑해서 반환
	if result.Error != nil {
		return nil, fmt.Errorf("db handler error: %w", result.Error)
	}

	return board, nil
}

// 게시글 리스트 조회
func (d *DBHandler) GetBoardList() ([]*model.Board, error) {
	boardList := []*model.Board{}
	result := d.gDB.Find(&boardList)

	if result.Error != nil {
		return nil, fmt.Errorf("db handler error: %w", result.Error)
	}

	return boardList, nil
}

// 게시글 ID로 조회
func (d *DBHandler) GetBoardByID(id uint) (*model.Board, error) {
	board := &model.Board{}
	result := d.gDB.First(board, id)

	if result.Error != nil {
		return nil, fmt.Errorf("db handler error: %w", result.Error)
	}

	return board, nil
}

// 게시글 수정
func (d *DBHandler) UpdateBoard(id uint, newBoard *model.Board) (*model.Board, error) {
	oldBoard := &model.Board{}
	result := d.gDB.Model(oldBoard).Where("id = ?", id).Updates(newBoard)

	if result.Error != nil {
		return nil, fmt.Errorf("db handler error: %w", result.Error)
	}

	return oldBoard, nil
}

// 게시글 삭제
func (d *DBHandler) DeleteBoardByID(id uint) error {
	result := d.gDB.Delete(&model.Board{}, id)

	if result.Error != nil {
		return fmt.Errorf("db handler error: %w", result.Error)
	}

	return nil
}
