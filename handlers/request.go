package handlers

import "github.com/chaosyeo/gin-starter/kernel/support/requests"

type LotteriesCreateRequest struct {
	requests.CommonCreateRequest
}

type LotteriesRemoveRequest struct {
	requests.CommonRemoveRequest
}

type LotteriesUpdateRequest struct {
	requests.CommonUpdateRequest
}

type LotteriesCollectRequest struct {
	requests.CommonCollectRequest
}

type LotteriesFetchRequest struct {
	requests.CommonFetchRequest
}