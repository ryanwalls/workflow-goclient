// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CancelWorkflow Cancel a workflow
*/
func (a *Client) CancelWorkflow(params *CancelWorkflowParams, authInfo runtime.ClientAuthInfoWriter) (*CancelWorkflowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelWorkflowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cancelWorkflow",
		Method:             "PUT",
		PathPattern:        "/workflows/{id}/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CancelWorkflowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CancelWorkflowOK), nil

}

/*
GetWorkflow Get a workflow by id
*/
func (a *Client) GetWorkflow(params *GetWorkflowParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkflow",
		Method:             "GET",
		PathPattern:        "/workflows/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkflowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowOK), nil

}

/*
HeartbeatActivity Signal workflow that activity is still running
*/
func (a *Client) HeartbeatActivity(params *HeartbeatActivityParams, authInfo runtime.ClientAuthInfoWriter) (*HeartbeatActivityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHeartbeatActivityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "heartbeatActivity",
		Method:             "PUT",
		PathPattern:        "/workflows/{id}/activities/{activityId}/heartbeat",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HeartbeatActivityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HeartbeatActivityOK), nil

}

/*
SignalWorkflow Signal a workflow with the given id
*/
func (a *Client) SignalWorkflow(params *SignalWorkflowParams, authInfo runtime.ClientAuthInfoWriter) (*SignalWorkflowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSignalWorkflowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "signalWorkflow",
		Method:             "POST",
		PathPattern:        "/workflows/{id}/signals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SignalWorkflowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SignalWorkflowOK), nil

}

/*
UpdateActivity Create or update an activity
*/
func (a *Client) UpdateActivity(params *UpdateActivityParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateActivityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateActivityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateActivity",
		Method:             "PUT",
		PathPattern:        "/workflows/{id}/activities/{activityId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateActivityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateActivityOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}