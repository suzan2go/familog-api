package model

// Migration migrate model schema
func (db *DB) Migration() {
	db.AutoMigrate(&Diary{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Device{})
	db.AutoMigrate(&SessionToken{})
	db.AutoMigrate(&DiarySubscriber{})
	db.AutoMigrate(&DiaryEntry{})
	db.AutoMigrate(&DiaryEntryImage{})
	db.AutoMigrate(&DiaryInvitation{})
	db.AutoMigrate(&DiaryEntryComment{})
}
