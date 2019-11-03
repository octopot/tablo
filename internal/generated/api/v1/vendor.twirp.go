// Code generated by protoc-gen-twirp v5.8.0, DO NOT EDIT.
// source: v1/vendor.proto

package v1

import bytes "bytes"
import strings "strings"
import context "context"
import fmt "fmt"
import ioutil "io/ioutil"
import http "net/http"
import strconv "strconv"

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// =======================
// VendorService Interface
// =======================

type VendorService interface {
	Connect(context.Context, *Secret) (*Vendor, error)

	Connected(context.Context, *Void) (*VendorList, error)

	Disconnect(context.Context, *URI) (*Void, error)

	Sources(context.Context, *URI) (*SourceNode, error)
}

// =============================
// VendorService Protobuf Client
// =============================

type vendorServiceProtobufClient struct {
	client HTTPClient
	urls   [4]string
}

// NewVendorServiceProtobufClient creates a Protobuf client that implements the VendorService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewVendorServiceProtobufClient(addr string, client HTTPClient) VendorService {
	prefix := urlBase(addr) + VendorServicePathPrefix
	urls := [4]string{
		prefix + "Connect",
		prefix + "Connected",
		prefix + "Disconnect",
		prefix + "Sources",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &vendorServiceProtobufClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &vendorServiceProtobufClient{
		client: client,
		urls:   urls,
	}
}

func (c *vendorServiceProtobufClient) Connect(ctx context.Context, in *Secret) (*Vendor, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Connect")
	out := new(Vendor)
	err := doProtobufRequest(ctx, c.client, c.urls[0], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceProtobufClient) Connected(ctx context.Context, in *Void) (*VendorList, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Connected")
	out := new(VendorList)
	err := doProtobufRequest(ctx, c.client, c.urls[1], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceProtobufClient) Disconnect(ctx context.Context, in *URI) (*Void, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Disconnect")
	out := new(Void)
	err := doProtobufRequest(ctx, c.client, c.urls[2], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceProtobufClient) Sources(ctx context.Context, in *URI) (*SourceNode, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Sources")
	out := new(SourceNode)
	err := doProtobufRequest(ctx, c.client, c.urls[3], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// =========================
// VendorService JSON Client
// =========================

type vendorServiceJSONClient struct {
	client HTTPClient
	urls   [4]string
}

// NewVendorServiceJSONClient creates a JSON client that implements the VendorService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewVendorServiceJSONClient(addr string, client HTTPClient) VendorService {
	prefix := urlBase(addr) + VendorServicePathPrefix
	urls := [4]string{
		prefix + "Connect",
		prefix + "Connected",
		prefix + "Disconnect",
		prefix + "Sources",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &vendorServiceJSONClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &vendorServiceJSONClient{
		client: client,
		urls:   urls,
	}
}

func (c *vendorServiceJSONClient) Connect(ctx context.Context, in *Secret) (*Vendor, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Connect")
	out := new(Vendor)
	err := doJSONRequest(ctx, c.client, c.urls[0], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceJSONClient) Connected(ctx context.Context, in *Void) (*VendorList, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Connected")
	out := new(VendorList)
	err := doJSONRequest(ctx, c.client, c.urls[1], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceJSONClient) Disconnect(ctx context.Context, in *URI) (*Void, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Disconnect")
	out := new(Void)
	err := doJSONRequest(ctx, c.client, c.urls[2], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceJSONClient) Sources(ctx context.Context, in *URI) (*SourceNode, error) {
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithMethodName(ctx, "Sources")
	out := new(SourceNode)
	err := doJSONRequest(ctx, c.client, c.urls[3], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ============================
// VendorService Server Handler
// ============================

type vendorServiceServer struct {
	VendorService
	hooks *twirp.ServerHooks
}

func NewVendorServiceServer(svc VendorService, hooks *twirp.ServerHooks) TwirpServer {
	return &vendorServiceServer{
		VendorService: svc,
		hooks:         hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *vendorServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// VendorServicePathPrefix is used for all URL paths on a twirp VendorService server.
// Requests are always: POST VendorServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const VendorServicePathPrefix = "/twirp/octolab.api.tablo.v1.VendorService/"

func (s *vendorServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "octolab.api.tablo.v1")
	ctx = ctxsetters.WithServiceName(ctx, "VendorService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/twirp/octolab.api.tablo.v1.VendorService/Connect":
		s.serveConnect(ctx, resp, req)
		return
	case "/twirp/octolab.api.tablo.v1.VendorService/Connected":
		s.serveConnected(ctx, resp, req)
		return
	case "/twirp/octolab.api.tablo.v1.VendorService/Disconnect":
		s.serveDisconnect(ctx, resp, req)
		return
	case "/twirp/octolab.api.tablo.v1.VendorService/Sources":
		s.serveSources(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *vendorServiceServer) serveConnect(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveConnectJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveConnectProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *vendorServiceServer) serveConnectJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Connect")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Secret)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Vendor
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Connect(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Vendor and nil error while calling Connect. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveConnectProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Connect")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(Secret)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Vendor
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Connect(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Vendor and nil error while calling Connect. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveConnected(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveConnectedJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveConnectedProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *vendorServiceServer) serveConnectedJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Connected")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Void)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *VendorList
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Connected(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *VendorList and nil error while calling Connected. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveConnectedProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Connected")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(Void)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *VendorList
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Connected(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *VendorList and nil error while calling Connected. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveDisconnect(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveDisconnectJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveDisconnectProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *vendorServiceServer) serveDisconnectJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Disconnect")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(URI)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Void
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Disconnect(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Void and nil error while calling Disconnect. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveDisconnectProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Disconnect")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(URI)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Void
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Disconnect(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Void and nil error while calling Disconnect. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveSources(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSourcesJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveSourcesProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *vendorServiceServer) serveSourcesJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Sources")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(URI)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *SourceNode
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Sources(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *SourceNode and nil error while calling Sources. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) serveSourcesProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Sources")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(URI)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *SourceNode
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.VendorService.Sources(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *SourceNode and nil error while calling Sources. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *vendorServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor1, 0
}

func (s *vendorServiceServer) ProtocGenTwirpVersion() string {
	return "v5.8.0"
}

func (s *vendorServiceServer) PathPrefix() string {
	return VendorServicePathPrefix
}

var twirpFileDescriptor1 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0x63, 0xd7, 0x75, 0xdb, 0x1b, 0x7d, 0x1f, 0xd5, 0xa8, 0x8b, 0x60, 0x21, 0x64, 0x45,
	0x42, 0x0a, 0x2c, 0x6c, 0x62, 0x36, 0xd0, 0x02, 0x8b, 0x94, 0x42, 0x22, 0x85, 0x3f, 0x72, 0x5a,
	0x16, 0xec, 0x26, 0x33, 0x57, 0x61, 0x84, 0x3d, 0x63, 0x8d, 0x07, 0x4b, 0x5d, 0xf3, 0x02, 0xbc,
	0x00, 0xaf, 0xc0, 0x33, 0x22, 0x8f, 0xdd, 0x08, 0x89, 0xc4, 0xa9, 0x90, 0xd8, 0x8d, 0xc7, 0xbf,
	0x7b, 0xe6, 0xdc, 0x33, 0xd7, 0x86, 0x3b, 0xd5, 0x38, 0xae, 0x50, 0x72, 0xa5, 0xa3, 0x42, 0x2b,
	0xa3, 0xc8, 0x89, 0x62, 0x46, 0x65, 0x74, 0x19, 0xd1, 0x42, 0x44, 0x86, 0x2e, 0x33, 0x15, 0x55,
	0xe3, 0xa0, 0xc6, 0x98, 0xca, 0x73, 0x25, 0x1b, 0x6c, 0xf8, 0xdd, 0x01, 0x7f, 0x81, 0x4c, 0xa3,
	0x21, 0x27, 0xb0, 0x6f, 0xd4, 0x17, 0x94, 0x03, 0x27, 0x74, 0x46, 0x47, 0x69, 0xf3, 0x40, 0x08,
	0x78, 0x92, 0xe6, 0x38, 0x70, 0xed, 0xa6, 0x5d, 0x93, 0x10, 0xfa, 0x1c, 0x4b, 0xa6, 0x45, 0x61,
	0x84, 0x92, 0x83, 0x3d, 0xfb, 0xea, 0xf7, 0x2d, 0x72, 0x0a, 0x87, 0x85, 0x56, 0x95, 0xe0, 0xa8,
	0x07, 0x5e, 0xe8, 0x8c, 0xfe, 0x4f, 0xee, 0x47, 0x9b, 0x0c, 0x45, 0x1f, 0x5a, 0x2a, 0x5d, 0xf3,
	0xc3, 0x1f, 0x0e, 0xf8, 0x1f, 0x6d, 0x2b, 0xe4, 0x21, 0xb8, 0x82, 0x5b, 0x3f, 0xfd, 0xe4, 0xee,
	0x66, 0x81, 0xab, 0x74, 0x96, 0xba, 0x82, 0xff, 0xa5, 0xcf, 0x04, 0x3c, 0x73, 0x5d, 0xe0, 0x2d,
	0x3d, 0x5a, 0x76, 0xf8, 0x12, 0xa0, 0xb1, 0x37, 0x17, 0xa5, 0x21, 0x8f, 0xc1, 0xcb, 0x44, 0x69,
	0x06, 0x4e, 0xb8, 0x37, 0xea, 0x27, 0xf7, 0x36, 0x2b, 0x34, 0x7c, 0x6a, 0xc9, 0xa1, 0x00, 0x7f,
	0xa1, 0xbe, 0x6a, 0x86, 0xff, 0xbc, 0xbd, 0xda, 0x6a, 0x73, 0xd4, 0xed, 0xad, 0x36, 0x7c, 0x6b,
	0xf5, 0x9b, 0x7b, 0x23, 0xf0, 0x4e, 0x71, 0x24, 0xcf, 0xc0, 0xcb, 0xd1, 0xd0, 0xd6, 0xf1, 0x83,
	0x2e, 0x81, 0x9a, 0x8f, 0xde, 0xa2, 0xa1, 0xa9, 0x2d, 0x21, 0x4f, 0x61, 0x9f, 0x7d, 0x16, 0x19,
	0xb7, 0x0d, 0xf4, 0x93, 0x70, 0x57, 0xed, 0xb4, 0x97, 0x36, 0x05, 0xe4, 0x14, 0xfc, 0x0c, 0x69,
	0x85, 0xa5, 0x6d, 0x70, 0x47, 0x69, 0xdd, 0xe7, 0xb4, 0x97, 0xb6, 0x15, 0xc1, 0x73, 0xf0, 0x6a,
	0x0f, 0xeb, 0xf4, 0x9c, 0xed, 0xe9, 0xb9, 0x7f, 0xa4, 0x37, 0xf1, 0xc1, 0xe3, 0xd4, 0xd0, 0x47,
	0x43, 0x38, 0xbc, 0x19, 0x01, 0x02, 0xe0, 0xbf, 0x99, 0x5d, 0x4e, 0xaf, 0x26, 0xc7, 0xbd, 0x7a,
	0x7d, 0x99, 0x5e, 0xcc, 0xe7, 0xef, 0x8f, 0x9d, 0xe4, 0xa7, 0x0b, 0xff, 0x35, 0xb7, 0xbc, 0x40,
	0x5d, 0x09, 0x86, 0xe4, 0x02, 0x0e, 0xce, 0x95, 0x94, 0xc8, 0x0c, 0xd9, 0x16, 0xb5, 0xfd, 0xee,
	0x82, 0xce, 0x99, 0x21, 0x33, 0x38, 0x6a, 0x65, 0x90, 0x93, 0x60, 0x0b, 0xaa, 0x04, 0x0f, 0xc2,
	0x2e, 0x19, 0x7b, 0xff, 0xe7, 0x00, 0xaf, 0x44, 0xc9, 0x5a, 0x53, 0xdb, 0x07, 0x2e, 0xe8, 0x38,
	0x86, 0xbc, 0x86, 0x83, 0x26, 0xea, 0xb2, 0x4b, 0x61, 0xe7, 0xfd, 0x4e, 0x5e, 0x7c, 0x3a, 0x5b,
	0xa9, 0x35, 0xa5, 0xf4, 0x2a, 0x46, 0xa6, 0xca, 0xeb, 0xd2, 0x60, 0x1e, 0x5b, 0x3e, 0x16, 0xd2,
	0xa0, 0x96, 0x34, 0x8b, 0x57, 0x28, 0x51, 0x53, 0x83, 0x3c, 0xa6, 0x85, 0x88, 0xab, 0xf1, 0x59,
	0x35, 0x5e, 0xfa, 0xf6, 0xf7, 0xf5, 0xe4, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x93, 0x7e,
	0x91, 0xf8, 0x04, 0x00, 0x00,
}
