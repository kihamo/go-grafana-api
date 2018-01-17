package go_grafana_api

import (
	"context"
	"fmt"
	"net/http"
)

type FindAnnotationsInput struct {
	From        *int64    `url:"from,omitempty"`
	To          *int64    `url:"to,omitempty"`
	AlertId     *int64    `url:"alertId,omitempty"`
	DashboardId *int64    `url:"dashboardId,omitempty"`
	PanelId     *int64    `url:"panelId,omitempty"`
	Tags        []*string `url:"tags,omitempty"`
	Type        *string   `url:"type,omitempty"`
	Limit       *int64    `url:"limit,omitempty"`
}

type FindAnnotationsOutput []struct {
	Id          *int64       `json:"id,omitempty"`
	AlertId     *int64       `json:"alertId,omitempty"`
	AlertName   *string      `json:"alertName,omitempty"`
	DashboardId *int64       `json:"dashboardId,omitempty"`
	PanelId     *int64       `json:"panelId,omitempty"`
	UserId      *int64       `json:"userId,omitempty"`
	NewState    *string      `json:"newState,omitempty"`
	PrevState   *string      `json:"prevState,omitempty"`
	Time        *int64       `json:"time,omitempty"`
	Text        *string      `json:"text,omitempty"`
	RegionId    *int64       `json:"regionId,omitempty"`
	Tags        []*string    `json:"tags,omitempty"`
	Login       *string      `json:"login,omitempty"`
	Email       *string      `json:"email,omitempty"`
	AvatarUrl   *string      `json:"avatarUrl,omitempty"`
	Data        *interface{} `json:"data,omitempty"`
}

func (c *Client) FindAnnotations(ctx context.Context, input *FindAnnotationsInput) (*FindAnnotationsOutput, error) {
	request, err := c.newRequest(ctx, http.MethodGet, "/annotations/", input)
	if err != nil {
		return nil, err
	}

	output := &FindAnnotationsOutput{}
	return output, c.send(request, output)
}

type CreateAnnotationInput struct {
	DashboardId *int64      `json:"dashboardId,omitempty"`
	PanelId     *int64      `json:"panelId,omitempty"`
	Time        *int64      `json:"time,omitempty"`
	Text        *string     `json:"text"`
	Tags        []*string   `json:"tags,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	IsRegion    *bool       `json:"isRegion,omitempty"`
	TimeEnd     *int64      `json:"timeEnd,omitempty"`
}

type CreateAnnotationOutput struct {
	Id      *int64  `json:"id"`
	EndId   *int64  `json:"endId,omitempty"`
	Message *string `json:"message"`
}

func (c *Client) CreateAnnotation(ctx context.Context, input *CreateAnnotationInput) (*CreateAnnotationOutput, error) {
	request, err := c.newRequest(ctx, http.MethodPost, "/annotations/", input)
	if err != nil {
		return nil, err
	}

	output := &CreateAnnotationOutput{}
	return output, c.send(request, output)
}

type CreateAnnotationGraphiteInput struct {
	When *int64    `json:"when"`
	What *string   `json:"what"`
	Data *string   `json:"data"`
	Tags []*string `json:"tags,omitempty"`
}

type CreateAnnotationGraphiteOutput struct {
	Id      *int64  `json:"id"`
	Message *string `json:"message"`
}

func (c *Client) CreateAnnotationGraphite(ctx context.Context, input *CreateAnnotationGraphiteInput) (*CreateAnnotationGraphiteOutput, error) {
	request, err := c.newRequest(ctx, http.MethodPost, "/annotations/graphite", input)
	if err != nil {
		return nil, err
	}

	output := &CreateAnnotationGraphiteOutput{}
	return output, c.send(request, output)
}

type UpdateAnnotationInput struct {
	Id       *int64    `json:"id,omitempty"`
	Time     *int64    `json:"time,omitempty"`
	Text     *string   `json:"text"`
	Tags     []*string `json:"tags,omitempty"`
	IsRegion *bool     `json:"isRegion,omitempty"`
	TimeEnd  *int64    `json:"timeEnd,omitempty"`
}

type UpdateAnnotationOutput struct {
	Message *string `json:"message,omitempty"`
}

func (c *Client) UpdateAnnotation(ctx context.Context, input *UpdateAnnotationInput) (*UpdateAnnotationOutput, error) {
	id := Int64Value(input.Id)
	input.Id = nil

	request, err := c.newRequest(ctx, http.MethodPut, fmt.Sprintf("/annotations/%d", id), input)
	if err != nil {
		return nil, err
	}

	output := &UpdateAnnotationOutput{}
	return output, c.send(request, output)
}

type DeleteAnnotationsInput struct {
	AlertId     *int64 `json:"alertId,omitempty"`
	DashboardId *int64 `json:"dashboardId,omitempty"`
	PanelId     *int64 `json:"panelId,omitempty"`
}

type DeleteAnnotationsOutput struct {
	Message *string `json:"message,omitempty"`
}

// @Role admin
func (c *Client) DeleteAnnotations(ctx context.Context, input *DeleteAnnotationsInput) (*DeleteAnnotationsOutput, error) {
	request, err := c.newRequest(ctx, http.MethodPost, "/annotations/mass-delete", input)
	if err != nil {
		return nil, err
	}

	output := &DeleteAnnotationsOutput{}
	return output, c.send(request, output)
}

type DeleteAnnotationByIdInput struct {
	Id *int64 `json:"id,omitempty"`
}

type DeleteAnnotationByIdOutput struct {
	Message *string `json:"message,omitempty"`
}

func (c *Client) DeleteAnnotationById(ctx context.Context, input *DeleteAnnotationByIdInput) (*DeleteAnnotationByIdOutput, error) {
	id := Int64Value(input.Id)
	input.Id = nil

	request, err := c.newRequest(ctx, http.MethodDelete, fmt.Sprintf("/annotations/%d", id), input)
	if err != nil {
		return nil, err
	}

	output := &DeleteAnnotationByIdOutput{}
	return output, c.send(request, output)
}

type DeleteAnnotationRegionInput struct {
	RegionId *int64 `json:"regionId,omitempty"`
}

type DeleteAnnotationRegionOutput struct {
	Message *string `json:"message,omitempty"`
}

func (c *Client) DeleteAnnotationRegion(ctx context.Context, input *DeleteAnnotationRegionInput) (*DeleteAnnotationRegionOutput, error) {
	regionId := Int64Value(input.RegionId)
	input.RegionId = nil

	request, err := c.newRequest(ctx, http.MethodDelete, fmt.Sprintf("/annotations/region/%d", regionId), input)
	if err != nil {
		return nil, err
	}

	output := &DeleteAnnotationRegionOutput{}
	return output, c.send(request, output)
}
