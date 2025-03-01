package savefilesresponse

import "time"

type PaginationEntry struct {
	CurrentPage  int             `json:"current_page"`
	Data         []EntryResponse `json:"data"`
	FirstPageURL string          `json:"first_page_url"`
	From         int             `json:"from"`
	LastPage     int             `json:"last_page"`
	LastPageURL  string          `json:"last_page_url"`
	Links        []LinkResponse  `json:"links"`
	NextPageURL  *string         `json:"next_page_url"`
	Path         string          `json:"path"`
	PerPage      int             `json:"per_page"`
	PrevPageURL  *string         `json:"prev_page_url"`
	To           int             `json:"to"`
	Total        int             `json:"total"`
}

type EntryResponse struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Description *string            `json:"description"`
	FileName    string             `json:"file_name"`
	Mime        *string            `json:"mime"`
	FileSize    int                `json:"file_size"`
	UserID      *int               `json:"user_id"`
	ParentID    int                `json:"parent_id"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	DeletedAt   *time.Time         `json:"deleted_at"`
	Path        string             `json:"path"`
	DiskPrefix  *string            `json:"disk_prefix"`
	Type        string             `json:"type"`
	Extension   *string            `json:"extension"`
	Public      bool               `json:"public"`
	Thumbnail   bool               `json:"thumbnail"`
	WorkspaceID *int               `json:"workspace_id"`
	OwnerID     int                `json:"owner_id"`
	Hash        string             `json:"hash"`
	URL         *string            `json:"url"`
	Users       []UserResponse     `json:"users"`
	Tags        []string           `json:"tags"`
	Permissions PermissionResponse `json:"permissions"`
}

type UserResponse struct {
	Email       string   `json:"email"`
	ID          int      `json:"id"`
	Avatar      string   `json:"avatar"`
	ModelType   string   `json:"model_type"`
	OwnsEntry   bool     `json:"owns_entry"`
	EntryPerms  []string `json:"entry_permissions"`
	DisplayName string   `json:"display_name"`
}

type PermissionResponse struct {
	Download bool `json:"files.download"`
	Update   bool `json:"files.update"`
	Delete   bool `json:"files.delete"`
}

type LinkResponse struct {
	URL    *string `json:"url"`
	Label  string  `json:"label"`
	Active bool    `json:"active"`
}
