package main

type Model struct {
	Id        string  `db:"size=32" gorm:"primary_key" json:"id"`
	CreatedAt []uint8 `db:"created_at,null" json:"created_at"`
	UpdatedAt []uint8 `db:"updated_at,null" json:"updated_at"`
	DeletedAt []uint8 `db:"deleted_at,null" json:"deleted_at"`
}
type Comment struct {
	Model
	ViewpointId string `json:"viewpoint_id" db:"viewpoint_id"`
	CreatorId   string `json:"creator_id" db:"creator_id"`
	Content     string `json:"content" db:"content"`
}
type Viewpoint struct {
	Model
	ProjectId     string `json:"project_id" db:"project_id"`
	ProjectFileId string `json:"project_file_id" db:"project_file_id"`
	CreatorId     string `json:"creator_id" db:"creator_id"`
	Content       string `json:"content" db:"content"`
	Description   string `json:"description" db:"description"`
	FileType      int    `json:"file_type" db:"file_type"`
	DistrictType  string `json:"district_type" db:"district_type"`
}

type Attachment struct {
	Model
	Name       string `json:"name" db:"name"`
	ParentId   string `json:"parent_id" db:"parent_id"`
	UploaderId string `json:"uploader_id" db:"uploader_id"`
	Type       int    `json:"type" db:"type"`
	Size       int    `json:"size" db:"size"`
}
