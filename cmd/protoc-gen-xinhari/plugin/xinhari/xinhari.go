package micro

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/ebelanja/micro/cmd/protoc-gen-micro/generator"
	"github.com/golang/protobuf/proto"
	pb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	options "google.golang.org/genproto/googleapis/api/annotations"
)

// Paths for packages used by code generated in this file,
// relative to the import_prefix of the generator.Generator.
const (
	apiPkgPath     = "github.com/ebelanja/go-micro/api"
	contextPkgPath = "context"
	clientPkgPath  = "github.com/ebelanja/go-micro/client"
	serverPkgPath  = "github.com/ebelanja/go-micro/server"
)

func init() {
	generator.RegisterPlugin(new(xinhari))
}

// xinhari is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for go-micro support.
type xinhari struct {
	gen *generator.Generator
}

// Name returns the name of this plugin, "xinhari".
func (g *xinhari) Name() string {
	return "xinhari"
}

// The names for packages imported in the generated code.
// They may vary from the final path component of the import path
// if the name is used by other packages.
var (
	apiPkg     string
	contextPkg string
	clientPkg  string
	serverPkg  string
	pkgImports map[generator.GoPackageName]bool
)

// Init initializes the plugin.
func (g *xinhari) Init(gen *generator.Generator) {
	g.gen = gen
	apiPkg = generator.RegisterUniquePackageName("api", nil)
	contextPkg = generator.RegisterUniquePackageName("context", nil)
	clientPkg = generator.RegisterUniquePackageName("client", nil)
	serverPkg = generator.RegisterUniquePackageName("server", nil)
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *xinhari) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (g *xinhari) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

// P forwards to g.gen.P.
func (g *xinhari) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the services in the given file.
func (g *xinhari) Generate(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}
	g.P("// Reference imports to suppress errors if they are not otherwise used.")
	g.P("var _ ", apiPkg, ".Endpoint")
	g.P("var _ ", contextPkg, ".Context")
	g.P("var _ ", clientPkg, ".Option")
	g.P("var _ ", serverPkg, ".Option")
	g.P()

	for i, service := range file.FileDescriptorProto.Service {
		g.generateService(file, service, i)
	}
}

// GenerateImports generates the import declaration for this file.
func (g *xinhari) GenerateImports(file *generator.FileDescriptor, imports map[generator.GoImportPath]generator.GoPackageName) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}
	g.P("import (")
	g.P(apiPkg, " ", strconv.Quote(path.Join(g.gen.ImportPrefix, apiPkgPath)))
	g.P(contextPkg, " ", strconv.Quote(path.Join(g.gen.ImportPrefix, contextPkgPath)))
	g.P(clientPkg, " ", strconv.Quote(path.Join(g.gen.ImportPrefix, clientPkgPath)))
	g.P(serverPkg, " ", strconv.Quote(path.Join(g.gen.ImportPrefix, serverPkgPath)))
	g.P(")")
	g.P()

	// We need to keep track of imported packages to make sure we don't produce
	// a name collision when generating types.
	pkgImports = make(map[generator.GoPackageName]bool)
	for _, name := range imports {
		pkgImports[name] = true
	}
}

// reservedClientName records whether a client name is reserved on the client side.
var reservedClientName = map[string]bool{
	// TODO: do we need any in go-micro?
}

func unexport(s string) string {
	if len(s) == 0 {
		return ""
	}
	name := strings.ToLower(s[:1]) + s[1:]
	if pkgImports[generator.GoPackageName(name)] {
		return name + "_"
	}
	return name
}

// generateService generates all the code for the named service.
func (g *xinhari) generateService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	path := fmt.Sprintf("6,%d", index) // 6 means service.

	origServName := service.GetName()
	serviceName := strings.ToLower(service.GetName())
	if pkg := file.GetPackage(); pkg != "" {
		serviceName = pkg
	}
	servName := generator.CamelCase(origServName)
	servAlias := servName + "Service"

	// strip suffix
	if strings.HasSuffix(servAlias, "ServiceService") {
		servAlias = strings.TrimSuffix(servAlias, "Service")
	}

	g.P()
	g.P("// Api Endpoints for ", servName, " service")
	g.P()

	g.P("func New", servName, "Endpoints () []*", apiPkg, ".Endpoint {")
	g.P("return []*", apiPkg, ".Endpoint{")
	for _, method := range service.Method {
		if method.Options != nil && proto.HasExtension(method.Options, options.E_Http) {
			g.P("&", apiPkg, ".Endpoint{")
			g.generateEndpoint(servName, method)
			g.P("},")
		}
	}
	g.P("}")
	g.P("}")
	g.P()

	g.P()
	g.P("// Client API for ", servName, " service")
	g.P()

	// Client interface.
	g.P("type ", servAlias, " interface {")
	for i, method := range service.Method {
		g.gen.PrintComments(fmt.Sprintf("%s,2,%d", path, i)) // 2 means method in a service.
		g.P(g.generateClientSignature(servName, method))
	}
	g.P("}")
	g.P()

	// Client structure.
	g.P("type ", unexport(servAlias), " struct {")
	g.P("c ", clientPkg, ".Client")
	g.P("name string")
	g.P("}")
	g.P()

	// NewClient factory.
	g.P("func New", servAlias, " (name string, c ", clientPkg, ".Client) ", servAlias, " {")
	/*
		g.P("if c == nil {")
		g.P("c = ", clientPkg, ".NewClient()")
		g.P("}")
		g.P("if len(name) == 0 {")
		g.P(`name = "`, serviceName, `"`)
		g.P("}")
	*/
	g.P("return &", unexport(servAlias), "{")
	g.P("c: c,")
	g.P("name: name,")
	g.P("}")
	g.P("}")
	g.P()
	var methodIndex, streamIndex int
	serviceDescVar := "_" + servName + "_serviceDesc"
	// Client method implementations.
	for _, method := range service.Method {
		var descExpr string
		if !method.GetServerStreaming() {
			// Unary RPC method
			descExpr = fmt.Sprintf("&%s.Methods[%d]", serviceDescVar, methodIndex)
			methodIndex++
		} else {
			// Streaming RPC method
			descExpr = fmt.Sprintf("&%s.Streams[%d]", serviceDescVar, streamIndex)
			streamIndex++
		}
		g.generateClientMethod(serviceName, servName, serviceDescVar, method, descExpr)
	}

	g.P("// Server API for ", servName, " service")
	g.P()

	// Server interface.
	serverType := servName + "Handler"
	g.P("type ", serverType, " interface {")
	for i, method := range service.Method {
		g.gen.PrintComments(fmt.Sprintf("%s,2,%d", path, i)) // 2 means method in a service.
		g.P(g.generateServerSignature(servName, method))
	}
	g.P("}")
	g.P()

	// Server registration.
	g.P("func Register", servName, "Handler(s ", serverPkg, ".Server, hdlr ", serverType, ", opts ...", serverPkg, ".HandlerOption) error {")
	g.P("type ", unexport(servName), " interface {")

	// generate interface methods
	for _, method := range service.Method {
		methName := generator.CamelCase(method.GetName())
		inType := g.typeName(method.GetInputType())
		outType := g.typeName(method.GetOutputType())

		if !method.GetServerStreaming() && !method.GetClientStreaming() {
			g.P(methName, "(ctx ", contextPkg, ".Context, in *", inType, ", out *", outType, ") error")
			continue
		}
		g.P(methName, "(ctx ", contextPkg, ".Context, stream server.Stream) error")
	}
	g.P("}")
	g.P("type ", servName, " struct {")
	g.P(unexport(servName))
	g.P("}")
	g.P("h := &", unexport(servName), "Handler{hdlr}")
	for _, method := range service.Method {
		if method.Options != nil && proto.HasExtension(method.Options, options.E_Http) {
			g.P("opts = append(opts, ", apiPkg, ".WithEndpoint(&", apiPkg, ".Endpoint{")
			g.generateEndpoint(servName, method)
			g.P("}))")
		}
	}
	g.P("return s.Handle(s.NewHandler(&", servName, "{h}, opts...))")
	g.P("}")
	g.P()

	g.P("type ", unexport(servName), "Handler struct {")
	g.P(serverType)
	g.P("}")

	// Server handler implementations.
	var handlerNames []string
	for _, method := range service.Method {
		hname := g.generateServerMethod(servName, method)
		handlerNames = append(handlerNames, hname)
	}
}

// generateEndpoint creates the api endpoint
func (g *xinhari) generateEndpoint(servName string, method *pb.MethodDescriptorProto) {
	if method.Options == nil || !proto.HasExtension(method.Options, options.E_Http) {
		return
	}
	// http rules
	r, err := proto.GetExtension(method.Options, options.E_Http)
	if err != nil {
		return
	}
	rule := r.(*options.HttpRule)
	var meth string
	var path string
	switch {
	case len(rule.GetDelete()) > 0:
		meth = "DELETE"
		path = rule.GetDelete()
	case len(rule.GetGet()) > 0:
		meth = "GET"
		path = rule.GetGet()
	case len(rule.GetPatch()) > 0:
		meth = "PATCH"
		path = rule.GetPatch()
	case len(rule.GetPost()) > 0:
		meth = "POST"
		path = rule.GetPost()
	case len(rule.GetPut()) > 0:
		meth = "PUT"
		path = rule.GetPut()
	}
	if len(meth) == 0 || len(path) == 0 {
		return
	}
	// TODO: process additional bindings
	g.P("Name:", fmt.Sprintf(`"%s.%s",`, servName, method.GetName()))
	g.P("Path:", fmt.Sprintf(`[]string{"%s"},`, path))
	g.P("Method:", fmt.Sprintf(`[]string{"%s"},`, meth))
	if len(rule.GetGet()) == 0 {
		g.P("Body:", fmt.Sprintf(`"%s",`, rule.GetBody()))
	}
	if method.GetServerStreaming() || method.GetClientStreaming() {
		g.P("Stream: true,")
	}
	g.P(`Handler: "rpc",`)
}

// generateClientSignature returns the client-side signature for a method.
func (g *xinhari) generateClientSignature(servName string, method *pb.MethodDescriptorProto) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	if reservedClientName[methName] {
		methName += "_"
	}
	reqArg := ", in *" + g.typeName(method.GetInputType())
	if method.GetClientStreaming() {
		reqArg = ""
	}
	respName := "*" + g.typeName(method.GetOutputType())
	if method.GetServerStreaming() || method.GetClientStreaming() {
		respName = servName + "_" + generator.CamelCase(origMethName) + "Service"
	}

	return fmt.Sprintf("%s(ctx %s.Context%s, opts ...%s.CallOption) (%s, error)", methName, contextPkg, reqArg, clientPkg, respName)
}

func (g *xinhari) generateClientMethod(reqServ, servName, serviceDescVar string, method *pb.MethodDescriptorProto, descExpr string) {
	reqMethod := fmt.Sprintf("%s.%s", servName, method.GetName())
	methName := generator.CamelCase(method.GetName())
	inType := g.typeName(method.GetInputType())
	outType := g.typeName(method.GetOutputType())

	servAlias := servName + "Service"

	// strip suffix
	if strings.HasSuffix(servAlias, "ServiceService") {
		servAlias = strings.TrimSuffix(servAlias, "Service")
	}

	g.P("func (c *", unexport(servAlias), ") ", g.generateClientSignature(servName, method), "{")
	if !method.GetServerStreaming() && !method.GetClientStreaming() {
		g.P(`req := c.c.NewRequest(c.name, "`, reqMethod, `", in)`)
		g.P("out := new(", outType, ")")
		// TODO: Pass descExpr to Invoke.
		g.P("err := ", `c.c.Call(ctx, req, out, opts...)`)
		g.P("if err != nil { return nil, err }")
		g.P("return out, nil")
		g.P("}")
		g.P()
		return
	}
	streamType := unexport(servAlias) + methName
	g.P(`req := c.c.NewRequest(c.name, "`, reqMethod, `", &`, inType, `{})`)
	g.P("stream, err := c.c.Stream(ctx, req, opts...)")
	g.P("if err != nil { return nil, err }")

	if !method.GetClientStreaming() {
		g.P("if err := stream.Send(in); err != nil { return nil, err }")
	}

	g.P("return &", streamType, "{stream}, nil")
	g.P("}")
	g.P()

	genSend := method.GetClientStreaming()
	genRecv := method.GetServerStreaming()

	// Stream auxiliary types and methods.
	g.P("type ", servName, "_", methName, "Service interface {")
	g.P("Context() context.Context")
	g.P("SendMsg(interface{}) error")
	g.P("RecvMsg(interface{}) error")
	g.P("Close() error")

	if genSend {
		g.P("Send(*", inType, ") error")
	}
	if genRecv {
		g.P("Recv() (*", outType, ", error)")
	}
	g.P("}")
	g.P()

	g.P("type ", streamType, " struct {")
	g.P("stream ", clientPkg, ".Stream")
	g.P("}")
	g.P()

	g.P("func (x *", streamType, ") Close() error {")
	g.P("return x.stream.Close()")
	g.P("}")
	g.P()

	g.P("func (x *", streamType, ") Context() context.Context {")
	g.P("return x.stream.Context()")
	g.P("}")
	g.P()

	g.P("func (x *", streamType, ") SendMsg(m interface{}) error {")
	g.P("return x.stream.Send(m)")
	g.P("}")
	g.P()

	g.P("func (x *", streamType, ") RecvMsg(m interface{}) error {")
	g.P("return x.stream.Recv(m)")
	g.P("}")
	g.P()

	if genSend {
		g.P("func (x *", streamType, ") Send(m *", inType, ") error {")
		g.P("return x.stream.Send(m)")
		g.P("}")
		g.P()

	}

	if genRecv {
		g.P("func (x *", streamType, ") Recv() (*", outType, ", error) {")
		g.P("m := new(", outType, ")")
		g.P("err := x.stream.Recv(m)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("return m, nil")
		g.P("}")
		g.P()
	}
}

// generateServerSignature returns the server-side signature for a method.
func (g *xinhari) generateServerSignature(servName string, method *pb.MethodDescriptorProto) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	if reservedClientName[methName] {
		methName += "_"
	}

	var reqArgs []string
	ret := "error"
	reqArgs = append(reqArgs, contextPkg+".Context")

	if !method.GetClientStreaming() {
		reqArgs = append(reqArgs, "*"+g.typeName(method.GetInputType()))
	}
	if method.GetServerStreaming() || method.GetClientStreaming() {
		reqArgs = append(reqArgs, servName+"_"+generator.CamelCase(origMethName)+"Stream")
	}
	if !method.GetClientStreaming() && !method.GetServerStreaming() {
		reqArgs = append(reqArgs, "*"+g.typeName(method.GetOutputType()))
	}
	return methName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}

func (g *xinhari) generateServerMethod(servName string, method *pb.MethodDescriptorProto) string {
	methName := generator.CamelCase(method.GetName())
	hname := fmt.Sprintf("_%s_%s_Handler", servName, methName)
	serveType := servName + "Handler"
	inType := g.typeName(method.GetInputType())
	outType := g.typeName(method.GetOutputType())

	if !method.GetServerStreaming() && !method.GetClientStreaming() {
		g.P("func (h *", unexport(servName), "Handler) ", methName, "(ctx ", contextPkg, ".Context, in *", inType, ", out *", outType, ") error {")
		g.P("return h.", serveType, ".", methName, "(ctx, in, out)")
		g.P("}")
		g.P()
		return hname
	}
	streamType := unexport(servName) + methName + "Stream"
	g.P("func (h *", unexport(servName), "Handler) ", methName, "(ctx ", contextPkg, ".Context, stream server.Stream) error {")
	if !method.GetClientStreaming() {
		g.P("m := new(", inType, ")")
		g.P("if err := stream.Recv(m); err != nil { return err }")
		g.P("return h.", serveType, ".", methName, "(ctx, m, &", streamType, "{stream})")
	} else {
		g.P("return h.", serveType, ".", methName, "(ctx, &", streamType, "{stream})")
	}
	g.P("}")
	g.P()

	genSend := method.GetServerStreaming()
	genRecv := method.GetClientStreaming()

	// Stream auxiliary types and methods.
	g.P("type ", servName, "_", methName, "Stream interface {")
	g.P("Context() context.Context")
	g.P("SendMsg(interface{}) error")
	g.P("RecvMsg(interface{}) error")
	g.P("Close() error")

	if genSend {
		g.P("Send(*", outType, ") error")
	}

	if genRecv {
		g.P("Recv() (*", inType, ", error)")
	}

	g.P("}")
	g.P()

	g.P("type ", streamType, " struct {")
	g.P("stream ", serverPkg, ".Stream")
	g.P("}")
	g.P()

	g.P("func (x *", streamType, ") Close() error {")
	g.P("return x.stream.Close()")
	g.P("}")
	g.P()

	g.P("func (x *", streamType, ") Context() context.Context {")
	g.P("return x.stream.Context()")
	g.P("}")
	g.P()

	g.P("func (x *", streamType, ") SendMsg(m interface{}) error {")
	g.P("return x.stream.Send(m)")
	g.P("}")
	g.P()

	g.P("func (x *", streamType, ") RecvMsg(m interface{}) error {")
	g.P("return x.stream.Recv(m)")
	g.P("}")
	g.P()

	if genSend {
		g.P("func (x *", streamType, ") Send(m *", outType, ") error {")
		g.P("return x.stream.Send(m)")
		g.P("}")
		g.P()
	}

	if genRecv {
		g.P("func (x *", streamType, ") Recv() (*", inType, ", error) {")
		g.P("m := new(", inType, ")")
		g.P("if err := x.stream.Recv(m); err != nil { return nil, err }")
		g.P("return m, nil")
		g.P("}")
		g.P()
	}

	return hname
}