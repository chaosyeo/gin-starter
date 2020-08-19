package requests

type CommonCreateRequest struct {}

type CommonRemoveRequest struct {
	Uri string ``
	Id  int64  ``
}

type CommonUpdateRequest struct {
	Uri string ``
	Id  int64  ``
}

type CommonCollectRequest struct {
	CommonPaginatorRequest
}

type CommonFetchRequest struct {
	Uri string ``
	Id  int64  ``
}

type CommonPurgeRequest struct {}