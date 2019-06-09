package model

import "database/sql"

type Course struct {
	AuthorsFullnames       sql.NullString `gorm:"column:AuthorsFullnames"`
	CachedOn               sql.NullString `gorm:"column:CachedOn"`
	DefaultImageURL        sql.NullString `gorm:"column:DefaultImageUrl"`
	Description            sql.NullString `gorm:"column:Description"`
	DurationInMilliseconds sql.NullInt64  `gorm:"column:DurationInMilliseconds"`
	HasTranscript          sql.NullInt64  `gorm:"column:HasTranscript"`
	ImageURL               sql.NullString `gorm:"column:ImageUrl"`
	IsStale                sql.NullInt64  `gorm:"column:IsStale"`
	Level                  sql.NullString `gorm:"column:Level"`
	Name                   sql.NullString `gorm:"column:Name"`
	ReleaseDate            sql.NullString `gorm:"column:ReleaseDate"`
	ShortDescription       sql.NullString `gorm:"column:ShortDescription"`
	Title                  sql.NullString `gorm:"column:Title"`
	UpdatedDate            sql.NullString `gorm:"column:UpdatedDate"`
	URLSlug                sql.NullString `gorm:"column:UrlSlug"`
}

// TableName sets the insert table name for this struct type
func (c *Course) TableName() string {
	return "Course"
}

type Module struct {
	AuthorHandle           sql.NullString `gorm:"column:AuthorHandle"`
	CourseName             sql.NullString `gorm:"column:CourseName"`
	Description            sql.NullString `gorm:"column:Description"`
	DurationInMilliseconds sql.NullInt64  `gorm:"column:DurationInMilliseconds"`
	ID                     sql.NullInt64  `gorm:"column:Id"`
	ModuleIndex            sql.NullInt64  `gorm:"column:ModuleIndex"`
	Name                   sql.NullString `gorm:"column:Name"`
	Title                  sql.NullString `gorm:"column:Title"`
}

// TableName sets the insert table name for this struct type
func (m *Module) TableName() string {
	return "Module"
}

type Clip struct {
	ClipIndex              sql.NullInt64  `gorm:"column:ClipIndex"`
	DurationInMilliseconds sql.NullInt64  `gorm:"column:DurationInMilliseconds"`
	ID                     sql.NullInt64  `gorm:"column:Id"`
	ModuleID               sql.NullInt64  `gorm:"column:ModuleId"`
	Name                   sql.NullString `gorm:"column:Name"`
	SupportsStandard       sql.NullInt64  `gorm:"column:SupportsStandard"`
	SupportsWidescreen     sql.NullInt64  `gorm:"column:SupportsWidescreen"`
	Title                  sql.NullString `gorm:"column:Title"`
}

// TableName sets the insert table name for this struct type
func (c *Clip) TableName() string {
	return "Clip"
}

type Cliptranscript struct {
	ClipID    sql.NullInt64  `gorm:"column:ClipId"`
	EndTime   sql.NullInt64  `gorm:"column:EndTime"`
	ID        sql.NullInt64  `gorm:"column:Id;primary_key"`
	StartTime sql.NullInt64  `gorm:"column:StartTime"`
	Text      sql.NullString `gorm:"column:Text"`
}

// TableName sets the insert table name for this struct type
func (c *Cliptranscript) TableName() string {
	return "ClipTranscript"
}

type User struct {
	DeviceID      sql.NullString `gorm:"column:DeviceId"`
	Jwt           sql.NullString `gorm:"column:Jwt"`
	JwtExpiration sql.NullString `gorm:"column:JwtExpiration"`
	RefreshToken  sql.NullString `gorm:"column:RefreshToken"`
	UserHandle    sql.NullString `gorm:"column:UserHandle"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "User"
}
