package go_grafana_api

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	AlertStateNoData   = "no_data"
	AlertStatePaused   = "paused"
	AlertStateAlerting = "alerting"
	AlertStateOK       = "ok"
	AlertStatePending  = "pending"
)

const (
	AlertTypeAlertmanager   = "prometheus-alertmanager"
	AlertTypeDingDing       = "dingding"
	AlertTypeEmail          = "email"
	AlertTypeHipchat        = "hipchat"
	AlertTypeKafka          = "kafka"
	AlertTypeLine           = "line"
	AlertTypeOpsGenie       = "opsgenie"
	AlertTypePagerduty      = "pagerduty"
	AlertTypePushover       = "pushover"
	AlertTypeSensu          = "sensu"
	AlertTypeSlack          = "slack"
	AlertTypeMicrosoftTeams = "teams"
	AlertTypeTelegram       = "telegram"
	AlertTypeThreema        = "threema"
	AlertTypeVictorOps      = "victorops"
	AlertTypeWebhook        = "webhook"
)

type GetAlertsStatesForDashboardInput struct {
	DashboardId *int64 `url:"dashboardId,omitempty"`
}

type GetAlertsStatesForDashboardOutput []struct {
	Id           *int64     `json:"id,omitempty"`
	DashboardId  *int64     `json:"dashboardId,omitempty"`
	PanelId      *int64     `json:"panelId,omitempty"`
	State        *string    `json:"state,omitempty"`
	NewStateDate *time.Time `json:"newStateDate,omitempty"`
}

func (c *Client) GetAlertsStatesForDashboard(ctx context.Context, input *GetAlertsStatesForDashboardInput) (*GetAlertsStatesForDashboardOutput, error) {
	request, err := c.newRequest(ctx, http.MethodGet, "/alerts/states-for-dashboard", input)
	if err != nil {
		return nil, err
	}

	output := &GetAlertsStatesForDashboardOutput{}
	return output, c.send(request, output)
}

type GetAlertNotifiersInput struct {
}

type GetAlertNotifiersOutput []struct {
	Type            *string `json:"type,omitempty"`
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	OptionsTemplate *string `json:"optionsTemplate,omitempty"`
}

func (c *Client) GetAlertNotifiers(ctx context.Context, input *GetAlertNotifiersInput) (*GetAlertNotifiersOutput, error) {
	request, err := c.newRequest(ctx, http.MethodGet, "/alert-notifiers/", input)
	if err != nil {
		return nil, err
	}

	output := &GetAlertNotifiersOutput{}
	return output, c.send(request, output)
}

type GetAlertNotificationsInput struct {
}

type GetAlertNotificationsOutput []struct {
	Id        *int64     `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Type      *string    `json:"type,omitempty"`
	IsDefault *bool      `json:"isDefault,omitempty"`
	Created   *time.Time `json:"created,omitempty"`
	Updated   *time.Time `json:"updated,omitempty"`
}

func (c *Client) GetAlertNotifications(ctx context.Context, input *GetAlertNotificationsInput) (*GetAlertNotificationsOutput, error) {
	request, err := c.newRequest(ctx, http.MethodGet, "/alert-notifications/", input)
	if err != nil {
		return nil, err
	}

	output := &GetAlertNotificationsOutput{}
	return output, c.send(request, output)
}

type GetAlertNotificationTestInput struct {
	Name     *string                `json:"name,omitempty"`
	Type     *string                `json:"type,omitempty"`
	Settings map[string]interface{} `json:"settings,omitempty"`
}

type GetAlertNotificationTestOutput struct {
	Message *string `json:"message"`
}

// @Role viewer
// @Role admin
func (c *Client) GetAlertNotificationTest(ctx context.Context, input *GetAlertNotificationTestInput) (*GetAlertNotificationTestOutput, error) {
	request, err := c.newRequest(ctx, http.MethodPost, "/alert-notifications/test", input)
	if err != nil {
		return nil, err
	}

	output := &GetAlertNotificationTestOutput{}
	return output, c.send(request, output)
}

type CreateAlertNotificationInput struct {
	Name      *string                `json:"name,omitempty"`
	Type      *string                `json:"type,omitempty"`
	IsDefault *bool                  `json:"isDefault,omitempty"`
	Settings  map[string]interface{} `json:"settings,omitempty"`
}

type CreateAlertNotificationOutput struct {
	Id        *int64                 `json:"id,omitempty"`
	Name      *string                `json:"name,omitempty"`
	Type      *string                `json:"type,omitempty"`
	IsDefault *bool                  `json:"isDefault,omitempty"`
	Settings  map[string]interface{} `json:"settings,omitempty"`
	Created   time.Time              `json:"created,omitempty"`
	Updated   time.Time              `json:"updated,omitempty"`
}

// @Role viewer
// @Role admin
func (c *Client) CreateAlertNotification(ctx context.Context, input *CreateAlertNotificationInput) (*CreateAlertNotificationOutput, error) {
	request, err := c.newRequest(ctx, http.MethodPost, "/alert-notifications/", input)
	if err != nil {
		return nil, err
	}

	output := &CreateAlertNotificationOutput{}
	return output, c.send(request, output)
}

type GetAlertNotificationInput struct {
	Id *int64 `json:"id,omitempty"`
}

type GetAlertNotificationOutput struct {
	Id        *int64                 `json:"id,omitempty"`
	Name      *string                `json:"name,omitempty"`
	Type      *string                `json:"type,omitempty"`
	IsDefault *bool                  `json:"isDefault,omitempty"`
	Settings  map[string]interface{} `json:"settings,omitempty"`
	Created   time.Time              `json:"created,omitempty"`
	Updated   time.Time              `json:"updated,omitempty"`
}

// @Role viewer
// @Role admin
func (c *Client) GetAlertNotification(ctx context.Context, input *GetAlertNotificationInput) (*GetAlertNotificationOutput, error) {
	id := Int64Value(input.Id)
	input.Id = nil

	request, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/alert-notifications/%d", id), input)
	if err != nil {
		return nil, err
	}

	output := &GetAlertNotificationOutput{}
	return output, c.send(request, output)
}

type UpdateAlertNotificationInput struct {
	Id        *int64                 `json:"id,omitempty"`
	Name      *string                `json:"name,omitempty"`
	Type      *string                `json:"type,omitempty"`
	IsDefault *bool                  `json:"isDefault,omitempty"`
	Settings  map[string]interface{} `json:"settings,omitempty"`
}

type UpdateAlertNotificationOutput struct {
	Id        *int64                 `json:"id,omitempty"`
	Name      *string                `json:"name,omitempty"`
	Type      *string                `json:"type,omitempty"`
	IsDefault *bool                  `json:"isDefault,omitempty"`
	Settings  map[string]interface{} `json:"settings,omitempty"`
	Created   time.Time              `json:"created,omitempty"`
	Updated   time.Time              `json:"updated,omitempty"`
}

// @Role viewer
// @Role admin
func (c *Client) UpdateAlertNotification(ctx context.Context, input *UpdateAlertNotificationInput) (*UpdateAlertNotificationOutput, error) {
	id := Int64Value(input.Id)

	request, err := c.newRequest(ctx, http.MethodPut, fmt.Sprintf("/alert-notifications/%d", id), input)
	if err != nil {
		return nil, err
	}

	output := &UpdateAlertNotificationOutput{}
	return output, c.send(request, output)
}

type DeleteAlertNotificationInput struct {
	Id *int64 `json:"id,omitempty"`
}

type DeleteAlertNotificationOutput struct {
	Message *string `json:"message,omitempty"`
}

// @Role viewer
// @Role admin
func (c *Client) DeleteAlertNotification(ctx context.Context, input *DeleteAlertNotificationInput) (*DeleteAlertNotificationOutput, error) {
	id := Int64Value(input.Id)
	input.Id = nil

	request, err := c.newRequest(ctx, http.MethodDelete, fmt.Sprintf("/alert-notifications/%d", id), input)
	if err != nil {
		return nil, err
	}

	output := &DeleteAlertNotificationOutput{}
	return output, c.send(request, output)
}
