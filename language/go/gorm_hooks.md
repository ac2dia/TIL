# GORM Hooks

- Hooks 은 함수로써 생성/쿼리/업데이트/삭제 전/후에 불려집니다.
- 특정 Model 을 대상으로 해당 함수들을 정의하면 자동으로 불려지며, error 발생시 자동으로 해당 내용들을 롤백합니다.
- Hook 메서드 타입은 func(\*gorm.DB) error 로 반드시 선언되어야 합니다.

```plaintext
// begin transaction
BeforeSave
BeforeCreate
// save before associations
// insert into database
// save after associations
AfterCreate
AfterSave
// commit or rollback transaction
```

- 생성시 사용 가능한 Hooks
- 주로 새로운 row 추가시 uuid 정책에 사용될 수 있습니다.

```go
type Base struct {
	ID                uuid.UUID           `gorm:"size:36;primaryKey" json:"id`
	CreatedAt         time.Time           `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
    b.ID = uuid.NewV4()
	return
}
```

- 공통 Model 을 생성하여 해당 Model 을 포함하는 테이블의 경우 새로운 Row 생성시 BeforeCreate 로 인해 ID 컬럼에 자동으로 UUID 값이 들어갑니다.

```go
type User struct {
	Base              Base                `gorm:"embedded"`
	Username          string              `gorm:"size:20;not null;unique" json:"username"`
	Password          string              `gorm:"size:100;not null" json:"password"`
	Name              string              `gorm:"size:40;not null" json:"name"`
	Email             string              `gorm:"size:40;not null;unique" json:"email"`
	Enabled           int                 `gorm:"default:1" json:"enabled"`
	RoleName          string              `gorm:"size:40;not null" json:"role_name"`
	LoginFailCount    int                 `gorm:"default:0;size:1" json:"login_fail_count"`
	LastLoginAt       time.Time           `json:"last_login_at"`
	PasswordChangedAt time.Time           `json:"password_changed_at"`
}
```

- 특정 테이블에 공통 Model을 embedded 타입으로 선언하여 관리할 수 있습니다.

# Reference

[1] GORM Hooks, https://gorm.io/ko_KR/docs/hooks.html
