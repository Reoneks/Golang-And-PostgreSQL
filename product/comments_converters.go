package product

func FromCommentsDto(commentsDto CommentsDto) Comments {
	return Comments(commentsDto)
}

func FromCommentsDtos(commentsDtos []CommentsDto) (comments []Comments) {
	for _, dto := range commentsDtos {
		comments = append(comments, Comments(dto))
	}
	return
}

func ToCommentsDto(comments Comments) CommentsDto {
	return CommentsDto(comments)
}

func ToCommentsDtos(comments []Comments) (commentsDto []CommentsDto) {
	for _, dto := range comments {
		commentsDto = append(commentsDto, CommentsDto(dto))
	}
	return
}
