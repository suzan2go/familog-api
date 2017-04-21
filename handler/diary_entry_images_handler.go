package handler

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

// PatchDiaryEntryImage Create diary
func (h *Handler) PatchDiaryEntryImage(c echo.Context) error {
	diaryEntryImage, err := h.DB.FindMyDiaryEntryImage(
		c.Param("diary_entry_id"), c.Param("id"), h.CurrentUser)
	if err != nil {
		return err
	}
	file, _ := c.FormFile("image")
	if err := h.DB.UpdateDiaryEntryImage(file, diaryEntryImage); err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(diaryEntryImage)
	return c.JSONBlob(http.StatusOK, buf.Bytes())
}

// DeleteDiaryEntryImage Create diary
func (h *Handler) DeleteDiaryEntryImage(c echo.Context) error {
	diaryEntryImage, err := h.DB.FindMyDiaryEntryImage(
		c.Param("diary_entry_id"), c.Param("id"), h.CurrentUser)
	if err != nil {
		return err
	}
	if err := h.DB.DeleteDiaryEntryImage(diaryEntryImage); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
