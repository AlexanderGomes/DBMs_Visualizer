package buffer

import (
	"errors"
)

const MaxPoolSize = 50

type PageID int64
type Page struct {
	ID       PageID
	Data     []byte
	IsDirty  bool
	IsPinned bool
}

type FrameID int
type BufferPoolManager struct {
	pages     [MaxPoolSize]*Page
	freeList  []FrameID
	pageTable map[PageID]FrameID
}

func (bpm *BufferPoolManager) NewPage(pageID PageID, data []byte) *Page {
	return &Page{
		ID:       pageID,
		Data:     data,
		IsDirty:  false,
		IsPinned: false,
	}
}

func (bpm *BufferPoolManager) DeletePage(pageID PageID) error {
	if frameID, ok := bpm.pageTable[pageID]; ok {
		page := bpm.pages[frameID]
		if page.IsPinned {
			return errors.New("Page is pinned, cannot delete")
		}
		delete(bpm.pageTable, pageID)
		bpm.pages[frameID] = nil
		return nil
	}
	return errors.New("Page not found")
}

func (bpm *BufferPoolManager) FetchPage(pageID PageID) (*Page, error) {
	var page Page
	//# buffer hit
	if frameID, ok := bpm.pageTable[pageID]; ok {
		page = *bpm.pages[frameID]
		if page.IsPinned {
			return nil, errors.New("Page is pinned, cannot access")
		}
		bpm.Pin(pageID)
		return &page, nil
	}

	//#disk read
	return &page, nil
}

func (bpm *BufferPoolManager) Unpin(pageID PageID, isDirty bool) error {
	if FrameID, ok := bpm.pageTable[pageID]; ok {
		page := bpm.pages[FrameID]
		page.IsDirty = isDirty
		page.IsPinned = false
		return nil
	}

	return errors.New("Page Not Found")
}

func (bpm *BufferPoolManager) Pin(pageID PageID) error {
	if FrameID, ok := bpm.pageTable[pageID]; ok {
		page := bpm.pages[FrameID]
		page.IsPinned = true
		// #replacemet policy

		return nil
	}

	return errors.New("Page Not Found")
}

func NewBufferPoolManager() *BufferPoolManager {
	freeList := make([]FrameID, 0)
	pages := [MaxPoolSize]*Page{}
	for i := 0; i < MaxPoolSize; i++ {
		freeList = append(freeList, FrameID(i))
		pages[FrameID(i)] = nil
	}

	pageTable := make(map[PageID]FrameID)
	return &BufferPoolManager{pages, freeList, pageTable}
}

// FlushPage(page_id_t page_id) //check if it's pinned
// FlushAllPages() //check if it's pinned
